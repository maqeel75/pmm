// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	cmdTimeout          = time.Minute
	resyncTimeout       = 5 * time.Minute
	statusCheckInterval = 5 * time.Second
	maxRestoreChecks    = 10
)

type pbmSeverity int

type describeInfo struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	pbmFatalSeverity pbmSeverity = iota
	pbmErrorSeverity
	pbmWarningSeverity
	pbmInfoSeverity
	pbmDebugSeverity
)

func (s pbmSeverity) String() string {
	switch s {
	case pbmFatalSeverity:
		return "F"
	case pbmErrorSeverity:
		return "E"
	case pbmWarningSeverity:
		return "W"
	case pbmInfoSeverity:
		return "I"
	case pbmDebugSeverity:
		return "D"
	default:
		return ""
	}
}

type pbmLogEntry struct {
	TS         int64 `json:"ts"`
	pbmLogKeys `json:",inline"`
	Msg        string `json:"msg"`
}

func (e pbmLogEntry) String() string {
	return fmt.Sprintf("%s %s [%s/%s] [%s/%s] %s",
		time.Unix(e.TS, 0).Format(time.RFC3339), e.Severity, e.RS, e.Node, e.Event, e.ObjName, e.Msg)
}

type pbmLogKeys struct {
	Severity pbmSeverity `json:"s"`
	RS       string      `json:"rs"`
	Node     string      `json:"node"`
	Event    string      `json:"e"`
	ObjName  string      `json:"eobj"`
	OPID     string      `json:"opid,omitempty"`
}

type pbmBackup struct {
	Name    string `json:"name"`
	Storage string `json:"storage"`
}

type pbmRestore struct {
	StartedAt time.Time
	Name      string `json:"name"`
	Snapshot  string `json:"snapshot"`
	PITR      string `json:"point-in-time"`
}

type pbmSnapshot struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	RestoreTo  int64  `json:"restoreTo"`
	PbmVersion string `json:"pbmVersion"`
	Type       string `json:"type"`
}

type pbmList struct {
	Snapshots []pbmSnapshot `json:"snapshots"`
	Pitr      struct {
		On     bool        `json:"on"`
		Ranges interface{} `json:"ranges"`
	} `json:"pitr"`
}

type pbmListRestore struct {
	Start    int    `json:"start"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Snapshot string `json:"snapshot"`
	PITR     int64  `json:"point-in-time"`
	Name     string `json:"name"`
	Error    string `json:"error"`
}

type pbmStatus struct {
	Backups struct {
		Type       string        `json:"type"`
		Path       string        `json:"path"`
		Region     string        `json:"region"`
		Snapshot   []pbmSnapshot `json:"snapshot"`
		PitrChunks struct {
			Size int `json:"size"`
		} `json:"pitrChunks"`
	} `json:"backups"`
	Cluster []struct {
		Rs    string `json:"rs"`
		Nodes []struct {
			Host  string `json:"host"`
			Agent string `json:"agent"`
			Role  string `json:"role"`
			Ok    bool   `json:"ok"`
		} `json:"nodes"`
	} `json:"cluster"`
	Pitr struct {
		Conf bool `json:"conf"`
		Run  bool `json:"run"`
	} `json:"pitr"`
	Running struct {
		Type    string `json:"type"`
		Name    string `json:"name"`
		StartTS int    `json:"startTS"`
		Status  string `json:"status"`
		OpID    string `json:"opID"`
	} `json:"running"`
}

type pbmError struct {
	Error string `json:"Error"`
}

func execPBMCommand(ctx context.Context, dbURL *url.URL, to interface{}, args ...string) error {
	nCtx, cancel := context.WithTimeout(ctx, cmdTimeout)
	defer cancel()

	args = append(args, "--out=json", "--mongodb-uri="+dbURL.String())
	cmd := exec.CommandContext(nCtx, pbmBin, args...) // #nosec G204

	b, err := cmd.Output()
	if err != nil {
		// try to parse pbm error message
		if len(b) != 0 {
			var pbmErr pbmError
			if e := json.Unmarshal(b, &pbmErr); e == nil {
				return errors.New(pbmErr.Error)
			}
		}
		return err
	}

	return json.Unmarshal(b, to)
}

func retrieveLogs(ctx context.Context, dbURL *url.URL, event string) ([]pbmLogEntry, error) {
	var logs []pbmLogEntry

	if err := execPBMCommand(ctx, dbURL, &logs, "logs", "--event="+event, "--tail=0"); err != nil {
		return nil, err
	}

	return logs, nil
}

func waitForPBMNoRunningOperations(ctx context.Context, l logrus.FieldLogger, dbURL *url.URL) error {
	l.Info("Waiting for no running pbm operations.")

	ticker := time.NewTicker(statusCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			var status pbmStatus
			if err := execPBMCommand(ctx, dbURL, &status, "status"); err != nil {
				return errors.Wrapf(err, "pbm status error")
			}
			if status.Running.Type == "" {
				return nil
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func waitForPBMBackup(ctx context.Context, l logrus.FieldLogger, dbURL *url.URL, name string) error {
	l.Infof("waiting for pbm backup: %s", name)
	ticker := time.NewTicker(statusCheckInterval)
	defer ticker.Stop()

	retryCount := 50

	for {
		select {
		case <-ticker.C:
			var info describeInfo
			err := execPBMCommand(ctx, dbURL, &info, "describe-backup", name)
			if err != nil {
				// for the first couple of seconds after backup process starts describe-backup command may return this error
				if (strings.HasSuffix(err.Error(), "no such file") || strings.HasSuffix(err.Error(), "file is empty")) && retryCount > 0 {
					retryCount--
					continue
				}

				return errors.Wrap(err, "failed to get backup status")
			}

			switch info.Status {
			case "done":
				return nil
			case "canceled":
				return errors.New("backup was canceled")
			case "error":
				return errors.New(info.Error)
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func findPITRRestore(list []pbmListRestore, restoreInfoPITRTime int64, startedAt time.Time) *pbmListRestore {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i].Type != "pitr" {
			continue
		}
		// list[i].Name is a string which represents time the restore was started.
		restoreStartedAt, err := time.Parse("2006-01-02T15:04:05.000000000Z", list[i].Name)
		if err != nil {
			continue
		}
		// Because of https://jira.percona.com/browse/PBM-723 to find our restore record in the list of all records we're checking:
		// 1. We received PITR field as a response on starting process
		// 2. There is a record with the same PITR field in the list of restoring records
		// 3. Start time of this record is not before the time we asked for restoring.
		if list[i].PITR == restoreInfoPITRTime && !restoreStartedAt.Before(startedAt) {
			return &list[i]
		}
	}
	return nil
}

func findPITRRestoreName(ctx context.Context, dbURL *url.URL, restoreInfo *pbmRestore) (string, error) {
	restoreInfoPITRTime, err := time.Parse("2006-01-02T15:04:05", restoreInfo.PITR)
	if err != nil {
		return "", err
	}

	ticker := time.NewTicker(statusCheckInterval)
	defer ticker.Stop()

	checks := 0
	for {
		select {
		case <-ticker.C:
			checks++
			var list []pbmListRestore
			if err := execPBMCommand(ctx, dbURL, &list, "list", "--restore"); err != nil {
				return "", errors.Wrapf(err, "pbm status error")
			}
			entry := findPITRRestore(list, restoreInfoPITRTime.Unix(), restoreInfo.StartedAt)
			if entry == nil {
				if checks > maxRestoreChecks {
					return "", errors.Errorf("failed to start restore")
				}
				continue
			} else {
				return entry.Name, nil
			}
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

func waitForPBMRestore(ctx context.Context, l logrus.FieldLogger, dbURL *url.URL, restoreInfo *pbmRestore, backupType, confFile string) error {
	l.Infof("waiting for pbm restore")
	var name string
	var err error

	// @TODO Do like this until https://jira.percona.com/browse/PBM-723 is not done.
	if restoreInfo.PITR != "" { // TODO add more checks of PBM responses.
		name, err = findPITRRestoreName(ctx, dbURL, restoreInfo)
		if err != nil {
			return err
		}
	} else {
		name = restoreInfo.Name
	}

	ticker := time.NewTicker(statusCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			var info describeInfo
			if backupType == "physical" {
				err = execPBMCommand(ctx, dbURL, &info, "describe-restore", "--config="+confFile, name)
			} else {
				err = execPBMCommand(ctx, dbURL, &info, "describe-restore", name)
			}
			if err != nil {
				return errors.Wrap(err, "failed to get restore status")
			}

			switch info.Status {
			case "done":
				return nil
			case "canceled":
				return errors.New("restore was canceled")
			case "error":
				return errors.New(info.Error)
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func pbmConfigure(ctx context.Context, l logrus.FieldLogger, dbURL *url.URL, confFile string) error {
	l.Info("Configuring S3 location.")
	nCtx, cancel := context.WithTimeout(ctx, cmdTimeout)
	defer cancel()

	output, err := exec.CommandContext( //nolint:gosec
		nCtx,
		pbmBin,
		"config",
		"--mongodb-uri="+dbURL.String(),
		"--file="+confFile).CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "pbm config error: %s", string(output))
	}

	return nil
}

func writePBMConfigFile(conf *PBMConfig) (string, error) {
	tmp, err := os.CreateTemp("", "pbm-config-*.yml")
	if err != nil {
		return "", errors.Wrap(err, "failed to create pbm configuration file")
	}

	bytes, err := yaml.Marshal(&conf)
	if err != nil {
		tmp.Close() //nolint:errcheck
		return "", errors.Wrap(err, "failed to marshall pbm configuration")
	}

	if _, err := tmp.Write(bytes); err != nil {
		tmp.Close() //nolint:errcheck
		return "", errors.Wrap(err, "failed to write pbm configuration file")
	}

	return tmp.Name(), tmp.Close()
}

// Serialization helpers

// Storage represents target storage parameters.
type Storage struct {
	Type       string     `yaml:"type"`
	S3         S3         `yaml:"s3"`
	FileSystem FileSystem `yaml:"filesystem"`
}

// S3 represents S3 storage parameters.
type S3 struct {
	Region      string      `yaml:"region"`
	Bucket      string      `yaml:"bucket"`
	Prefix      string      `yaml:"prefix"`
	EndpointURL string      `yaml:"endpointUrl"`
	Credentials Credentials `yaml:"credentials"`
}

// FileSystem  represents local storage parameters.
type FileSystem struct {
	Path string `yaml:"path"`
}

// Credentials contains S3 credentials.
type Credentials struct {
	AccessKeyID     string `yaml:"access-key-id"`
	SecretAccessKey string `yaml:"secret-access-key"`
}

// PITR contains Point-in-Time recovery reature related parameters.
type PITR struct {
	Enabled bool `yaml:"enabled"`
}

// PBMConfig represents pbm configuration file.
type PBMConfig struct {
	Storage Storage `yaml:"storage"`
	PITR    PITR    `yaml:"pitr"`
}

// createPBMConfig returns object that is ready to be serialized into YAML.
func createPBMConfig(locationConfig *BackupLocationConfig, prefix string, pitr bool) (*PBMConfig, error) {
	conf := &PBMConfig{
		PITR: PITR{
			Enabled: pitr,
		},
	}

	switch locationConfig.Type {
	case S3BackupLocationType:
		conf.Storage = Storage{
			Type: "s3",
			S3: S3{
				EndpointURL: locationConfig.S3Config.Endpoint,
				Region:      locationConfig.S3Config.BucketRegion,
				Bucket:      locationConfig.S3Config.BucketName,
				Prefix:      prefix,
				Credentials: Credentials{
					AccessKeyID:     locationConfig.S3Config.AccessKey,
					SecretAccessKey: locationConfig.S3Config.SecretKey,
				},
			},
		}
	case PMMClientBackupLocationType:
		conf.Storage = Storage{
			Type: "filesystem",
			FileSystem: FileSystem{
				Path: path.Join(locationConfig.LocalStorageConfig.Path, prefix),
			},
		}
	default:
		return nil, errors.New("unknown location config")
	}
	return conf, nil
}

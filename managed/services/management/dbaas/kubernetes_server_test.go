// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package dbaas

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	goversion "github.com/hashicorp/go-version"
	controllerv1beta1 "github.com/percona-platform/dbaas-api/gen/controller"
	dbaasv1 "github.com/percona/dbaas-operator/api/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	dbaasv1beta1 "github.com/percona/pmm/api/managementpb/dbaas"
	"github.com/percona/pmm/managed/models"
	"github.com/percona/pmm/managed/services/dbaas/olm"
	"github.com/percona/pmm/managed/utils/logger"
	"github.com/percona/pmm/managed/utils/testdb"
	"github.com/percona/pmm/managed/utils/tests"
	pmmversion "github.com/percona/pmm/version"
)

func TestKubernetesServer(t *testing.T) {
	setup := func(t *testing.T) (ctx context.Context, ks dbaasv1beta1.KubernetesServer, dbaasClient *mockDbaasClient,
		kubernetesClient *mockKubernetesClient, olms *olm.MockOperatorServiceManager, grafanaClient *mockGrafanaClient,
		versionService *mockVersionService, teardown func(t *testing.T),
	) {
		t.Helper()

		ctx = logger.Set(context.Background(), t.Name())
		uuid.SetRand(&tests.IDReader{})

		sqlDB := testdb.Open(t, models.SetupFixtures, nil)
		// To enable verbose queries output use:
		// db = reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(t.Logf))
		db := reform.NewDB(sqlDB, postgresql.Dialect, nil)
		dbaasClient = &mockDbaasClient{}
		kubernetesClient = &mockKubernetesClient{}
		grafanaClient = &mockGrafanaClient{}
		olms = &olm.MockOperatorServiceManager{}

		teardown = func(t *testing.T) {
			uuid.SetRand(nil)
			dbaasClient.AssertExpectations(t)
			require.NoError(t, sqlDB.Close())
		}
		// versionService = NewVersionServiceClient("https://check-dev.percona.com/versions/v1")
		versionService = &mockVersionService{}
		ks = NewKubernetesServer(db, dbaasClient, kubernetesClient, versionService, grafanaClient, olms)
		return
	}

	// This is for local testing. When running local tests, if pmmversion.PMMVersion is empty
	// these lines in kubernetes_server.go will throw an error and tests won't finish.
	//
	//     pmmVersion, err := goversion.NewVersion(pmmversion.PMMVersion)
	//     if err != nil {
	//     	return nil, status.Error(codes.Internal, err.Error())
	//     }
	//
	if pmmversion.PMMVersion == "" {
		pmmversion.PMMVersion = version230
	}

	t.Run("Basic", func(t *testing.T) {
		ctx, ks, dbaasClient, kubernetesClient, olms, grafanaClient, versionService, teardown := setup(t)
		kubernetesClient.On("SetKubeconfig", mock.Anything).Return(nil)
		kubernetesClient.On("SetKubeconfig", mock.Anything).Return(nil)
		defer teardown(t)

		v1120, _ := goversion.NewVersion("1.12.0")
		versionService.On("LatestOperatorVersion", mock.Anything, mock.Anything).Return(v1120, v1120, nil)

		clusters, err := ks.ListKubernetesClusters(ctx, &dbaasv1beta1.ListKubernetesClustersRequest{})
		require.NoError(t, err)
		require.Empty(t, clusters.KubernetesClusters)

		kubeconfig := "preferences: {}\n"
		dbaasClient.On("CheckKubernetesClusterConnection", mock.Anything, mock.Anything).Return(&controllerv1beta1.CheckKubernetesClusterConnectionResponse{
			Operators: &controllerv1beta1.Operators{
				PxcOperatorVersion:   "",
				PsmdbOperatorVersion: onePointEight,
			},
			Status: controllerv1beta1.KubernetesClusterStatus_KUBERNETES_CLUSTER_STATUS_OK,
		}, nil)

		olms.On("SetKubeConfig", mock.Anything).WaitUntil(time.After(time.Second)).Return(nil)
		grafanaClient.On("CreateAdminAPIKey", mock.Anything, mock.Anything).Return(int64(123456), "api-key", nil)
		olms.On("InstallOLMOperator", mock.Anything, mock.Anything).Return(nil)
		olms.On("InstallOperator", mock.Anything, mock.Anything).Return(nil)
		dbaasClient.On("StartMonitoring", mock.Anything, mock.Anything).WaitUntil(time.After(time.Second)).Return(&controllerv1beta1.StartMonitoringResponse{}, nil)

		kubernetesClusterName := "test-cluster"
		registerKubernetesClusterResponse, err := ks.RegisterKubernetesCluster(ctx, &dbaasv1beta1.RegisterKubernetesClusterRequest{
			KubernetesClusterName: kubernetesClusterName,
			KubeAuth:              &dbaasv1beta1.KubeAuth{Kubeconfig: kubeconfig},
		})
		require.NoError(t, err)
		assert.NotNil(t, registerKubernetesClusterResponse)

		getClusterResponse, err := ks.GetKubernetesCluster(ctx, &dbaasv1beta1.GetKubernetesClusterRequest{
			KubernetesClusterName: kubernetesClusterName,
		})
		require.NoError(t, err)
		assert.NotNil(t, getClusterResponse)
		assert.NotNil(t, getClusterResponse.KubeAuth)
		assert.Equal(t, kubeconfig, getClusterResponse.KubeAuth.Kubeconfig)

		supportedOperatorVersions := map[string][]string{
			"pxc-operator": {
				"1.10.0",
				"1.11.0",
				"1.9.0",
			},
			"psmdb-operator": {
				"1.11.0",
				"1.12.0",
				"1.10.0",
			},
		}
		versionService.On("SupportedOperatorVersionsList", mock.Anything, mock.Anything).Return(supportedOperatorVersions, nil)

		clusters, err = ks.ListKubernetesClusters(ctx, &dbaasv1beta1.ListKubernetesClustersRequest{})
		assert.NoError(t, err)
		assert.Equal(t, 1, len(clusters.KubernetesClusters))
		expected := []*dbaasv1beta1.ListKubernetesClustersResponse_Cluster{
			{
				KubernetesClusterName: kubernetesClusterName,
				Operators: &dbaasv1beta1.Operators{
					Pxc:   &dbaasv1beta1.Operator{Status: dbaasv1beta1.OperatorsStatus_OPERATORS_STATUS_NOT_INSTALLED},
					Psmdb: &dbaasv1beta1.Operator{Version: onePointEight, Status: dbaasv1beta1.OperatorsStatus_OPERATORS_STATUS_UNSUPPORTED},
					Dbaas: &dbaasv1beta1.Operator{Version: "", Status: dbaasv1beta1.OperatorsStatus_OPERATORS_STATUS_INVALID},
				},
				Status: dbaasv1beta1.KubernetesClusterStatus_KUBERNETES_CLUSTER_STATUS_OK,
			},
		}

		assert.Equal(t, expected[0].Operators, clusters.KubernetesClusters[0].Operators)
		assert.Equal(t, expected[0].KubernetesClusterName, clusters.KubernetesClusters[0].KubernetesClusterName)
		assert.True(
			t,
			clusters.KubernetesClusters[0].Status == dbaasv1beta1.KubernetesClusterStatus_KUBERNETES_CLUSTER_STATUS_OK ||
				clusters.KubernetesClusters[0].Status == dbaasv1beta1.KubernetesClusterStatus_KUBERNETES_CLUSTER_STATUS_PROVISIONING,
		)
		mockK8sResp := []dbaasv1.DatabaseCluster{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "first-pxc-test",
				},
				Spec: dbaasv1.DatabaseSpec{
					Database:      "pxc",
					DatabaseImage: "percona/percona-xtradb-cluster:8.0.27-18.1",
					ClusterSize:   5,
					DBInstance: dbaasv1.DBInstanceSpec{
						CPU:      resource.MustParse("3m"),
						Memory:   resource.MustParse("256"),
						DiskSize: resource.MustParse("1073741824"),
					},
					LoadBalancer: dbaasv1.LoadBalancerSpec{
						Type: "proxysql",
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceCPU:    resource.MustParse("2m"),
								corev1.ResourceMemory: resource.MustParse("124"),
							},
						},
					},
				},
				Status: dbaasv1.DatabaseClusterStatus{
					Ready: 15,
					Size:  15,
				},
			},
		}

		dbaasClient.On("StopMonitoring", mock.Anything, mock.Anything).Return(&controllerv1beta1.StopMonitoringResponse{}, nil)
		listDatabaseMock := kubernetesClient.On("ListDatabaseClusters", ctx)
		listDatabaseMock.Return(&dbaasv1.DatabaseClusterList{Items: mockK8sResp}, nil)

		_, err = ks.UnregisterKubernetesCluster(ctx, &dbaasv1beta1.UnregisterKubernetesClusterRequest{
			KubernetesClusterName: kubernetesClusterName,
		})
		tests.AssertGRPCError(t, status.Newf(codes.FailedPrecondition, "Kubernetes cluster %s has database clusters", kubernetesClusterName), err)

		mockK8sResp = []dbaasv1.DatabaseCluster{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "first-psmdb-test",
				},
				Spec: dbaasv1.DatabaseSpec{
					Database:      "psmdb",
					DatabaseImage: "percona/percona-server-mongodb:4.4.5-7",
					ClusterSize:   5,
					DBInstance: dbaasv1.DBInstanceSpec{
						CPU:      resource.MustParse("3m"),
						Memory:   resource.MustParse("256"),
						DiskSize: resource.MustParse("1073741824"),
					},
					LoadBalancer: dbaasv1.LoadBalancerSpec{
						Type: "mongos",
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceCPU:    resource.MustParse("2m"),
								corev1.ResourceMemory: resource.MustParse("124"),
							},
						},
					},
				},
				Status: dbaasv1.DatabaseClusterStatus{
					Ready: 10,
					Size:  10,
				},
			},
		}
		listDatabaseMock.Return(&dbaasv1.DatabaseClusterList{Items: mockK8sResp}, nil)
		tests.AssertGRPCError(t, status.Newf(codes.FailedPrecondition, "Kubernetes cluster %s has database clusters", kubernetesClusterName), err)
		mockK8sResp = []dbaasv1.DatabaseCluster{}

		listDatabaseMock.Return(&dbaasv1.DatabaseClusterList{Items: mockK8sResp}, nil)
		unregisterKubernetesClusterResponse, err := ks.UnregisterKubernetesCluster(ctx, &dbaasv1beta1.UnregisterKubernetesClusterRequest{
			KubernetesClusterName: kubernetesClusterName,
		})
		require.NoError(t, err)
		assert.NotNil(t, unregisterKubernetesClusterResponse)

		clusters, err = ks.ListKubernetesClusters(ctx, &dbaasv1beta1.ListKubernetesClustersRequest{})
		assert.NoError(t, err)
		assert.Empty(t, clusters.KubernetesClusters)
	})
}

func TestGetFlagValue(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		args          []string
		flagName      string
		expectedValue string
	}{
		{
			args:          []string{"token", "--foo", "bar"},
			flagName:      "--foo",
			expectedValue: "bar",
		},
		{
			args:          []string{"token", "--foo", "bar"},
			flagName:      "--bar",
			expectedValue: "",
		},
		{
			args:          []string{"token", "--foo"},
			flagName:      "--foo",
			expectedValue: "",
		},
	}
	for _, tt := range testCases {
		value := getFlagValue(tt.args, tt.flagName)
		assert.Equal(t, tt.expectedValue, value)
	}
}

const awsIAMAuthenticatorKubeconfig = `kind: Config
apiVersion: v1
current-context: arn:aws:eks:zone-2:123465545:cluster/cluster
clusters:
    - cluster:
        certificate-authority-data: base64data
        name: arn:aws:eks:zone-2:123465545:cluster/cluster
        server: https://DDDDD.bla.zone-2.eks.amazonaws.com
contexts:
    - context:
        cluster: arn:aws:eks:zone-2:123465545:cluster/cluster
        name: arn:aws:eks:zone-2:123465545:cluster/cluster
        user: arn:aws:eks:zone-2:123465545:cluster/cluster
preferences: {}
users:
    - name: arn:aws:eks:zone-2:123465545:cluster/cluster
      user:
        exec:
            apiVersion: client.authentication.k8s.io/v1alpha1
            args:
                - token
                - -i
                - test-cluster1
                - --region
                - zone-2
            command: aws-iam-authenticator
            env:
                - name: AWS_STS_REGIONAL_ENDPOINTS
                  value: regional
            provideClusterInfo: false
`

const awsIAMAuthenticatorKubeconfigTransformed = `kind: Config
apiVersion: v1
current-context: arn:aws:eks:zone-2:123465545:cluster/cluster
clusters:
    - cluster:
        certificate-authority-data: base64data
        name: arn:aws:eks:zone-2:123465545:cluster/cluster
        server: https://DDDDD.bla.zone-2.eks.amazonaws.com
contexts:
    - context:
        cluster: arn:aws:eks:zone-2:123465545:cluster/cluster
        name: arn:aws:eks:zone-2:123465545:cluster/cluster
        user: arn:aws:eks:zone-2:123465545:cluster/cluster
preferences: {}
users:
    - name: arn:aws:eks:zone-2:123465545:cluster/cluster
      user:
        exec:
            apiVersion: client.authentication.k8s.io/v1alpha1
            args:
                - token
                - -i
                - test-cluster1
                - --region
                - zone-2
            command: aws-iam-authenticator
            env:
                - name: AWS_STS_REGIONAL_ENDPOINTS
                  value: regional
                - name: AWS_ACCESS_KEY_ID
                  value: keyID
                - name: AWS_SECRET_ACCESS_KEY
                  value: key
            provideClusterInfo: false
`

const awsKubeconfig = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: base64data
    name: arn:aws:eks:zone-2:123465545:cluster/cluster
    server: https://DDDDD.bla.zone-2.eks.amazonaws.com
contexts:
- context:
    cluster: arn:aws:eks:zone-2:123465545:cluster/cluster
    name: arn:aws:eks:zone-2:123465545:cluster/cluster
    user: arn:aws:eks:zone-2:123465545:cluster/cluster
current-context: "arn:aws:eks:zone-2:123465545:cluster/cluster"
kind: Config
preferences: {}
users:
- name: arn:aws:eks:zone-2:123465545:cluster/cluster
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1alpha1
      args:
      - eks
      - get-token
      - --cluster-name
      - test-cluster1
      - --region
      - zone-2
      command: aws
      env:
      - name: AWS_STS_REGIONAL_ENDPOINTS
        value: regional
      provideClusterInfo: false
`

const awsKubeconfigWithKeys = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: base64data
    name: arn:aws:eks:zone-2:123465545:cluster/cluster
    server: https://DDDDD.bla.zone-2.eks.amazonaws.com
contexts:
- context:
    cluster: arn:aws:eks:zone-2:123465545:cluster/cluster
    name: arn:aws:eks:zone-2:123465545:cluster/cluster
    user: arn:aws:eks:zone-2:123465545:cluster/cluster
current-context: "arn:aws:eks:zone-2:123465545:cluster/cluster"
kind: Config
preferences: {}
users:
- name: arn:aws:eks:zone-2:123465545:cluster/cluster
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1alpha1
      args:
      - eks
      - get-token
      - --cluster-name
      - test-cluster1
      - --region
      - zone-2
      command: aws
      env:
      - name: AWS_STS_REGIONAL_ENDPOINTS
        value: regional
      - name: AWS_ACCESS_KEY_ID
        value: keyID
      - name: AWS_SECRET_ACCESS_KEY
        value: key
      provideClusterInfo: false
`

func TestUseIAMAuthenticator(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name              string
		kubeconfig        string
		expectedError     error
		expectedTransform string
		keyID             string
		key               string
	}{
		{
			name:              "transform aws to aws-iam-authenticator with keys",
			kubeconfig:        awsKubeconfig,
			expectedTransform: awsIAMAuthenticatorKubeconfigTransformed,
			expectedError:     nil,
			keyID:             "keyID",
			key:               "key",
		},
		{
			name:              "transform aws with keys to aws-iam-authenticator",
			kubeconfig:        awsKubeconfigWithKeys,
			expectedTransform: awsIAMAuthenticatorKubeconfigTransformed,
			expectedError:     nil,
		},
		{
			name:              "transform aws to aws-iam-authenticator without keys",
			kubeconfig:        awsKubeconfig,
			expectedTransform: awsIAMAuthenticatorKubeconfig,
			expectedError:     nil,
		},
		{
			name:              "add environment variables to aws-iam-authenticator",
			kubeconfig:        awsIAMAuthenticatorKubeconfig,
			expectedTransform: awsIAMAuthenticatorKubeconfigTransformed,
			expectedError:     nil,
			keyID:             "keyID",
			key:               "key",
		},
		{
			name:              "return error if kubeconfig is empty",
			kubeconfig:        "     ",
			expectedTransform: "",
			expectedError:     errKubeconfigIsEmpty,
		},
		{
			name:              "don't transform aws-iam-authenticator with keys",
			kubeconfig:        awsIAMAuthenticatorKubeconfigTransformed,
			expectedTransform: awsIAMAuthenticatorKubeconfigTransformed,
			expectedError:     nil,
		},
	}
	for i, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			value, err := replaceAWSAuthIfPresent(tt.kubeconfig, tt.keyID, tt.key)
			assert.ErrorIsf(t, err, tt.expectedError, "Errors don't match in the test case number %d.", i)
			assert.Equalf(t, tt.expectedTransform, value, "Given and expected kubeconfigs don't match in the test case number %d.", i)
		})
	}
}

func TestGetOperatorConfigs(t *testing.T) {
	data := `---
operators:
  psmdb:
    name: percona-server-mongodb-operator 
    catalog: percona-dbaas-catalog
    channel: stable-v79014
    auto-install-on-registration: true
  pxc:
    name: percona-xtradb-cluster-operator
    catalog: percona-dbaas-catalog
    channel: stable-v79015
    auto-install-on-registration: true
  dbaas:
    name: dbaas-operator
    catalog: percona-dbaas-catalog
    channel: stable-v79016
    auto-install-on-registration: true`

	f, err := os.CreateTemp("", "test-olm-operator-config.*.yaml")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte(data))
	assert.NoError(t, err)

	t.Setenv("PMM_OLM_OPERATORS_CONFIG", f.Name())

	cfg, err := getOperatorConfigs()
	assert.NoError(t, err)

	assert.Equal(t, 3, len(cfg.Operators))
	assert.Equal(t, "stable-v79015", cfg.Operators["pxc"].Channel)
}

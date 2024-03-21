#!/bin/bash -e

# Update submodules and PR branches
/bin/bash $(dirname $0)/update.sh

get_branch_name() {
  local module="${1:-}"
  local branch_name
  local path

  path=$(git config -f .gitmodules submodule.${module}.path)
  cd "$path" || exit 1
  branch_name=$(git branch --show-current)
  cd - > /dev/null
  echo $branch_name
}

# Define global variables
pmm_commit=$(git submodule status | grep 'sources/pmm/src' | awk -F ' ' '{print $1}')
echo $pmm_commit > apiCommitSha
pmm_branch=$(get_branch_name pmm)
echo $pmm_branch > apiBranch
pmm_url=$(git config -f .gitmodules submodule.pmm.url)
echo $pmm_url > apiURL
pmm_qa_branch=$(get_branch_name pmm-qa)
echo $pmm_qa_branch > pmmQABranch
pmm_qa_commit=$(git submodule status | grep 'pmm-qa' | awk -F ' ' '{print $1}')
echo $pmm_qa_commit > pmmQACommitSha
pmm_ui_tests_branch=$(get_branch_name pmm-ui-tests)
echo $pmm_ui_tests_branch > pmmUITestBranch
pmm_ui_tests_commit=$(git submodule status | grep 'pmm-ui-tests' | awk -F ' ' '{print $1}')
echo $pmm_ui_tests_commit > pmmUITestsCommitSha
fb_commit_sha=$(git rev-parse HEAD)
echo $fb_commit_sha > fbCommitSha

# We use a special docker image to build various PMM artifacts - `perconalab/rpmbuild:ol9`.
# Important: the docker container's user need to be able to write to these directories.
# The docker container's user is `builder` with uid 1000 and gid 1000. You need to make sure
# that the directories we create on the host are owned by a user with same uid and gid.

# Create cache directories. 
test -d "${root_dir}/go-path" || mkdir -p "go-path"
test -d "${root_dir}/go-build" || mkdir -p "go-build"

PATH_TO_SCRIPTS="sources/pmm/src/github.com/percona/pmm/build/scripts"
export RPMBUILD_DOCKER_IMAGE=perconalab/rpmbuild:ol9 

# Local reference test environment
# CPU: 4 cores
# RAM: 16GB
# OS: Ubuntu 22.04.1 LTS

# Build client source: 4m39s from scratch, 0m27s using cache
"$PATH_TO_SCRIPTS/build-client-source"

# Build client binary: ??? from scratch, 0m20s using cache
"$PATH_TO_SCRIPTS/build-client-binary"

# Building client source rpm takes 13s (caching does not apply)
"$PATH_TO_SCRIPTS/build-client-srpm"

# Building client rpm takes 1m40s
"$PATH_TO_SCRIPTS/build-client-rpm"

# Building client docker image takes 17s
GIT_COMMIT=$(git rev-parse HEAD | head -c 8)
export DOCKER_CLIENT_TAG=local/pmm-client:${GIT_COMMIT}
"$PATH_TO_SCRIPTS/build-client-docker"

# Building PMM CLient locally (non-CI, i.e. non-Jenkins)
# total time: 6m26s - build from scratch, no initial cache
# total time: 2m49s - subsequent build, using cache from prior builds


# Building PMM CLient in a CI environment, i.e. Jenkins running on AWS
# total time: 8m45s - build from scratch, no initial cache
# total time: ??? - subsequent build, using cache from prior builds

export RPM_EPOCH=1
export RPMBUILD_DIST="el9"
"$PATH_TO_SCRIPTS/build-server-rpm" percona-dashboards grafana-dashboards
"$PATH_TO_SCRIPTS/build-server-rpm" pmm-managed pmm
"$PATH_TO_SCRIPTS/build-server-rpm" percona-qan-api2 pmm
"$PATH_TO_SCRIPTS/build-server-rpm" pmm-update pmm
"$PATH_TO_SCRIPTS/build-server-rpm" pmm-dump
"$PATH_TO_SCRIPTS/build-server-rpm" grafana-db-migrator
"$PATH_TO_SCRIPTS/build-server-rpm" vmproxy pmm

# 3rd-party
"$PATH_TO_SCRIPTS/build-server-rpm" victoriametrics
"$PATH_TO_SCRIPTS/build-server-rpm" grafana

export DOCKER_TAG=local/pmm-server:${GIT_COMMIT}
export RPMBUILD_DOCKER_IMAGE=perconalab/rpmbuild:ol9
export RPMBUILD_DIST=el9
export DOCKERFILE=Dockerfile.el9.local
${PATH_TO_SCRIPTS}/build-server-docker

# Clean up temporary files
rm -f apiBranch \
  apiCommitSha \
	apiURL \
	fbCommitSha \
	pmmQABranch \
	pmmQACommitSha \
	pmmUITestBranch \
	pmmUITestsCommitSha
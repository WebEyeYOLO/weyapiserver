#!/bin/bash
set -o errexit

# define package name
PROGRAM="WebEyeYOLO"
WORK_DIR=/go/src/weyapiserver
BASE_NAME=weyapiserver
releasedir=./.release
distdir=${releasedir}/dist
GO_VERSION=1.12
BASEDIR="$( cd "$( dirname "$0"  )" && pwd  )"
VERSION=$(git branch | grep '^*' | cut -d ' ' -f 2 | awk -F'V' '{print $2}')
buildTime=$(date +%F-%H)
git_commit=$(git log -n 1 --pretty --format=%h)
if [ -z "$VERSION" ];then
    VERSION=cloud
fi
release_desc=${VERSION}-${git_commit}-${buildTime}

function prepare() {
	rm -rf $releasedir
    mkdir -pv $releasedir/{tmp,dist}
    path=$PWD
    [ ! -d "$distdir/usr/local/" ] && mkdir -p $distdir/usr/local/bin
}

function build() {
	echo "---> Build Binary"
	echo "version:$release_desc"

	echo "build...."
    docker run --rm -v `pwd`:${WORK_DIR} -w ${WORK_DIR} -it golang:${GO_VERSION} go build -ldflags "-w -s -X version.version=${release_desc}"  -o $releasedir/dist/usr/local/bin/weyapiserver .

	cd $releasedir/dist/usr/local/
	tar zcf pkg.tgz `find . -maxdepth 1|sed 1d`

	cat >Dockerfile <<EOF
FROM alpine:3.6
COPY pkg.tgz /
EOF
	docker build -t weyapiserver:v$VERSION .
}

function localbuild() {
		go build -ldflags "-w -s -X version.version=${release_desc}"  -o _output/${VERSION}/weyapiserver .
}


function build::image() {
	echo "---> Build Image"
	DOCKER_PATH=./hack/docker
        cd $BASEDIR/../
	HOME=`pwd`

	docker run --rm -v `pwd`:${WORK_DIR} -w ${WORK_DIR} -it golang:${GO_VERSION} go build -ldflags "-w -s -X verion.version=${release_desc}"  -o ${DOCKER_PATH}/${BASE_NAME} .

	cd  ${DOCKER_PATH}
	sed "s/__RELEASE_DESC__/${release_desc}/" Dockerfile > Dockerfile.release
	docker build -t weyapiserver:${VERSION} -f Dockerfile.release .
	rm -f ./Dockerfile.release
	rm -f ./${BASE_NAME}
	cd $HOME
}

case $1 in
	build)
		prepare
		build
	;;
	buildimage)
		build::image
	;;
	localbuild)
		prepare
		localbuild
	;;
esac
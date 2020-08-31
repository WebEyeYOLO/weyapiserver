FROM alpine:3.4
MAINTAINER zhichao.yan@outlook.com

ADD weyapiserver /run/weyapiserver
ADD entrypoint.sh /run/entrypoint.sh
WORKDIR /run
ENV CUR_NET=midonet
ENV EX_DOMAIN=
ENV RELEASE_DESC=__RELEASE_DESC__

ENTRYPOINT ["/run/entrypoint.sh"]
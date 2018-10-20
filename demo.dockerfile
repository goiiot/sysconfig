FROM golang:alpine AS build

LABEL "com.github.actions.name"="goiiot/sysconfig"
LABEL "com.github.actions.description"="Sysconfig demo app"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="purple"

LABEL "repository"="http://github.com/goiiot/sysconfig"
LABEL "homepage"="http://github.com/goiiot/sysconfig"
LABEL "maintainer"="JeffreyStoke <jeffctor@gmail.com>"

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

# build project
ENV GOPATH=/gopath
ENV BUILD_DIR=/build
ENV CGO_ENABLED=0
ENV PATH=${PATH}:${GOPATH}/bin

COPY . ${BUILD_DIR}

WORKDIR /

RUN apk add --no-cache --update --virtual .build_deps \
        upx git nodejs nodejs-npm make musl-dev dep curl \
    \
    && mkdir -p ${GOPATH} \
    \
    && go get github.com/rakyll/statik \
    && go get -d github.com/goreleaser/goreleaser \
    && cd ${GOPATH}/src/github.com/goreleaser/goreleaser \
    && dep ensure -vendor-only \
    && make setup build \
    && go install \
    \
    && cd ${BUILD_DIR} \
    && ./x-install-deps.sh \
    && ./x-build.sh

# build actual image
FROM alpine:latest

COPY --from=build /build/dist/linux_amd64/sysconfig /app/sysconfig

COPY --from=build /build/config.example.yaml /path/to/config.yaml
COPY --from=build /build/testdata/tls_cert.pem /build/testdata/tls_key.pem /path/to/
COPY --from=build /build/testdata/test_conf/* /path/to/

COPY --from=build /build/scripts/templates/t-bus-helper.sh /path/to/bus-helper.sh
COPY --from=build /build/scripts/templates/t-cell-helper.sh /path/to/cell-helper.sh
COPY --from=build /build/scripts/templates/t-iface-helper.sh /path/to/iface-helper.sh
COPY --from=build /build/scripts/templates/t-lora-gw-helper.sh /path/to/lora-gw-helper.sh
COPY --from=build /build/scripts/templates/t-lora-pf-helper.sh /path/to/lora-pf-helper.sh
COPY --from=build /build/scripts/templates/t-periph-helper.sh /path/to/periph-helper.sh
COPY --from=build /build/scripts/templates/t-wifi-helper.sh /path/to/wifi-helper.sh

EXPOSE 8080 8443

CMD ["/app/sysconfig", "-c", "/path/to/config.yaml"]

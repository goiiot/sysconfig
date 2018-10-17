FROM golang:alpine

# build app
COPY . /build

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

ENV GOPATH=/gopath
ENV CGO_ENABLED=0

RUN \
    apk add --no-cache --virtual .build_deps \
     upx git nodejs make musl-dev \
    && mkdir -p ${GOPATH}

RUN go get github.com/rakyll/statik \
    \
    && go get -d github.com/goreleaser/goreleaser \
    && cd ${GOPATH}/src/github.com/goreleaser/goreleaser \
    && dep ensure -vendor-only \
    && make setup build \
    && cd - \
    \
    && cd /build \
    && go mod download \
    && ./x-build.sh \
    && go clean -modcache -cache \
    \
    && rm -rf ${GOPATH} \
    && apk del .build_deps \
    \
    && mkdir -p /app /path/to \
    && mv dist/linux_amd64/sysconfig /app/sysconfig \
    && mv config.example.yaml /path/to/config.yaml \
    && mv testdata/tls_cert.pem testdata/tls_cert.pem /path/to/ \
    && mv testdata/test_conf/* /path/to/ \
    && mv scripts/templates/t-bus-helper.sh /path/to/bus-helper.sh \
    && mv scripts/templates/t-cell-helper.sh /path/to/cell-helper.sh \
    && mv scripts/templates/t-iface-helper.sh /path/to/iface-helper.sh \
    && mv scripts/templates/t-lora-gw-helper.sh /path/to/lora-gw-helper.sh \
    && mv scripts/templates/t-lora-pf-helper.sh /path/to/lora-pf-helper.sh \
    && mv scripts/templates/t-periph-helper.sh /path/to/periph-helper.sh \
    && mv scripts/templates/t-wifi-helper.sh /path/to/wifi-helper.sh \
    && chmod -R +x /path/to/*.sh /app/sysconfig

CMD ["/app/sysconfig", "-c", "/path/to/config.yaml"]

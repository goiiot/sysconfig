FROM alpine:latest

COPY dist/linux_amd64/sysconfig /app/sysconfig

COPY config.example.yaml /path/to/config.yaml

COPY testdata/tls_cert.pem testdata/tls_cert.pem /path/to/
COPY testdata/test_conf/* /path/to/

# copy helper scripts
COPY scripts/templates/t-bus-helper.sh /path/to/bus-helper.sh
COPY scripts/templates/t-cell-helper.sh /path/to/cell-helper.sh
COPY scripts/templates/t-iface-helper.sh /path/to/iface-helper.sh
COPY scripts/templates/t-lora-gw-helper.sh /path/to/lora-gw-helper.sh
COPY scripts/templates/t-lora-pf-helper.sh /path/to/lora-pf-helper.sh
COPY scripts/templates/t-periph-helper.sh /path/to/periph-helper.sh
COPY scripts/templates/t-wifi-helper.sh /path/to/wifi-helper.sh

RUN chmod -R +x /path/to/*.sh /app/sysconfig

CMD ["/app/sysconfig", "-c", "/path/to/config.yaml"]

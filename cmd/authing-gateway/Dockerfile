FROM jrei/systemd-centos:7

COPY AuthingGateway.viper.yaml /etc/AuthingGateway/AuthingGateway.viper.yaml
COPY authing-gateway /usr/local/bin/authing-gateway

CMD ["authing-gateway", "run"]

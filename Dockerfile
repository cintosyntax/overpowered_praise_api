FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/overpowered_praise_api /opt/bin/overpowered_praise_api

CMD ["/opt/bin/overpowered_praise_api"]

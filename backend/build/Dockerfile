FROM golang:1.27.0
WORKDIR /go/src/backend

SHELL ["/bin/bash", "-c"]

COPY tool-versions.env /tmp/tool-versions.env
RUN source /tmp/tool-versions.env \
    && go install github.com/air-verse/air@"${AIR_VERSION}" \
    && rm /tmp/tool-versions.env
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# change timezone
RUN \
    echo "Asia/Tokyo" > /etc/timezone \
    && rm /etc/localtime \
    && ln -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && dpkg-reconfigure --frontend noninteractive tzdata

COPY config/.bashrc /root/.bashrc

ENTRYPOINT ["./script/entrypoint.sh"]


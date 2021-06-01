FROM golang:1.14.0
WORKDIR /go/src/app

SHELL ["/bin/bash", "-c"]

RUN go get -u github.com/cosmtrek/air

# change timezone
RUN \
    echo "Asia/Tokyo" > /etc/timezone \
    && rm /etc/localtime \
    && ln -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && dpkg-reconfigure --frontend noninteractive tzdata

COPY config/.bashrc /root/.bashrc

ENTRYPOINT ["./scripts/entrypoint.sh"]



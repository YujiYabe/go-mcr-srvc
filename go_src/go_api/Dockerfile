FROM golang:1.13.0
# RUN mkdir -p /go/src/go_api
WORKDIR /go/src/go_api

SHELL ["/bin/bash", "-c"]

# change timezone
RUN \
    echo "Asia/Tokyo" > /etc/timezone \
    && rm /etc/localtime \
    && ln -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && dpkg-reconfigure --frontend noninteractive tzdata

# install gin
RUN go get github.com/codegangsta/gin
RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure

COPY config/.bashrc /root/.bashrc

ENTRYPOINT ["./entrypoint.sh"]



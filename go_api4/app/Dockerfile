FROM golang:1.13.0
# RUN mkdir -p /go/src/go_api
WORKDIR /go/src/app

SHELL ["/bin/bash", "-c"]

# change timezone
RUN \
    echo "Asia/Tokyo" > /etc/timezone \
    && rm /etc/localtime \
    && ln -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && dpkg-reconfigure --frontend noninteractive tzdata

# install gin

RUN go get -u github.com/golang/dep/cmd/dep \
    && go get github.com/mdempsky/gocode \
    && go get github.com/uudashr/gopkgs/cmd/gopkgs \
    && go get github.com/ramya-rao-a/go-outline \
    && go get github.com/stamblerre/gocode \
    && go get github.com/rogpeppe/godef \
    && go get golang.org/x/lint/golint


COPY config/.bashrc /root/.bashrc

ENTRYPOINT ["./entrypoint.sh"]



FROM golang:1.23
WORKDIR /go/src/backend

SHELL ["/bin/bash", "-c"]

# RUN go get -u github.com/cosmtrek/air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# change timezone
RUN \
    echo "Asia/Tokyo" > /etc/timezone \
    && rm /etc/localtime \
    && ln -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && dpkg-reconfigure --frontend noninteractive tzdata

COPY config/.bashrc /root/.bashrc

ENTRYPOINT ["./script/entrypoint.sh"]




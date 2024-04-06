FROM golang:1.19

RUN apt-get update \
        && apt-get install -y \
            bc \
            dos2unix \
            diffutils \
            coreutils;


RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN ln -s /go/bin/linux_amd64/migrate /usr/local/bin/migrate

RUN PATH="/usr/local/opt/coreutils/libexec/gnubin:$PATH"

RUN go install github.com/mitranim/gow@latest

WORKDIR /app
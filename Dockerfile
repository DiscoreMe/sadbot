FROM golang:1.14.9

WORKDIR /go/src/github.com/discoreme/sadbot
COPY bot bot
COPY cache cache
COPY cmd cmd
COPY config config
COPY weather weather
COPY dict dict
COPY go.mod go.mod
COPY go.sum go.sum

RUN go build -o /usr/bin/sadbot  /go/src/github.com/discoreme/sadbot/cmd/main.go

WORKDIR /usr/bin

ENTRYPOINT ["./sadbot"]
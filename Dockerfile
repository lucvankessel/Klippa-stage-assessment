FROM golang:1.16-alpine

ADD . /go/src/intern-assessment
WORKDIR /go/src/intern-assessment

RUN go build cmd/klippa-assessment/main.go

ENTRYPOINT [ "./main" ]
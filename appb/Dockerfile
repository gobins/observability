FROM golang:1.17 as builder
WORKDIR /work
COPY . .
RUN go build . -o appb
ENTRYPOINT [ "/work/appb" ]

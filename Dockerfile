FROM golang:1.17 as builder
WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
ENTRYPOINT [ "/work/app" ]

FROM scratch
COPY --from=builder /work/app /work/app
ENTRYPOINT [ "/work/app" ]

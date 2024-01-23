FROM quay.io/projectquay/golang:1.20  as builder
ARG TARGET_OS=linux
RUN echo "TARGET_OS: $TARGET_OS"
WORKDIR /go/src/app
COPY . .
RUN go version
RUN make $TARGET_OS


FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/kbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./kbot", "start"]

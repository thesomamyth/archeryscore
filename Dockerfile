FROM golang:1.19-alpine as builder
WORKDIR /build
ADD . .
RUN CGO_ENABLED=0 go build -buildvcs=false -o archeryscore main.go

FROM scratch
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=0:0 --from=builder /build/* /
USER 65535
EXPOSE 8080:8080
ENTRYPOINT [ "/archeryscore" ]
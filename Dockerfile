FROM golang:1.17-alpine as builder

WORKDIR /build
COPY shs.go /build/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o shs shs.go

FROM scratch as runner
COPY --from=builder /build/shs /shs
COPY passwd.noroot /etc/passwd
COPY www /www/

EXPOSE 8080
USER nobody
ENTRYPOINT ["/shs", "-l", ":8080", "-s", "/www"]


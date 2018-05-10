FROM alpine:latest
MAINTAINER Sean Hagen <sean.hagen@gmail.com>
RUN apk --no-cache add ca-certificates
ADD server server
ADD config/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ADD config/zoneinfo.tar.gz /
ENV PORT 80
EXPOSE 80
ENTRYPOINT ["/server"]

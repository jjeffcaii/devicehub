FROM golang:alpine AS builder
WORKDIR /go/src/devicehub
COPY . .
RUN go build

FROM alpine
LABEL maintainer "wcai@maxleap.com"
COPY db/ data/
COPY --from=builder /go/src/devicehub/devicehub /usr/local/bin/devicehub
EXPOSE 3000
VOLUME ["/data"]
ENTRYPOINT ["devicehub"]
CMD ["-d=/data"]
FROM golang:1.18-alpine as builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata && apk add upx

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go get -d -v
RUN go build -ldflags="-s -w" -o /app/ucabot .
RUN upx /app/ucabot

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Europe/Paris /usr/share/zoneinfo/Europe/Paris
ENV TZ Europe/Paris

WORKDIR /app

COPY --from=builder /app/ucabot /app/ucabot
COPY --from=builder /build/.env /app/.env

CMD ["/app/ucabot"]
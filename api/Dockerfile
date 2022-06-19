ARG GO_VERSION=1.18
FROM golang:${GO_VERSION}-alpine AS builder

# gcc binary fix
RUN apk add build-base

WORKDIR /app

COPY . .

RUN go get
RUN go build -o ./ducks-life-api

FROM alpine:3.15

RUN apk update && apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/ducks-life-api .

CMD [ "./ducks-life-api" ]
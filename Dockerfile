FROM golang:1.23-alpine AS builder
LABEL authors="ewik2k"

WORKDIR /usr/local/src

RUN apk --no-cache add bash make git gcc musl-dev

#dependecies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#build
COPY . ./
RUN go build -o ./bin/grpc_user ./cmd/user_server/main.go

FROM alpine AS runner


COPY --from=builder /usr/local/src/bin/grpc_user /
COPY --from=builder /usr/local/src/.env /

CMD ["/grpc_user"]
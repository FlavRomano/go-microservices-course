FROM golang:1.23.0-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
# ADD EXECUTION OF RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto
RUN CGO_ENABLED=0 go build -o loggerService ./cmd/api
RUN chmod +x /app/loggerService

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/loggerService /app
CMD ["/app/loggerService"]
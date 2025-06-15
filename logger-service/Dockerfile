FROM golang:1.18-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o loggerService ./cmd/api
RUN chmod +x /app/loggerService

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/loggerService /app
CMD ["/app/loggerService"]
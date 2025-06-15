FROM golang:1.18-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o mailService ./cmd/api
RUN chmod +x /app/mailService

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/mailService /app
COPY templates /templates
CMD ["/app/mailService"]
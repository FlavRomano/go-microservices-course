FROM golang:1.18-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o authenticationService ./cmd/api
RUN chmod +x /app/authenticationService

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/authenticationService /app
CMD ["/app/authenticationService"]
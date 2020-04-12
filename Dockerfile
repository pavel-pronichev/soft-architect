FROM golang:1.14.2-alpine3.11 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /main .
RUN chmod 711 main


FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]

FROM golang
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./grpc /app/grpc
COPY ./UserService /app/UserService
COPY ./OrderService /app/OrderService
COPY ./AuthorizationService /app/AuthorizationService
WORKDIR /app/AuthorizationService
RUN go mod download
RUN go build -o authserver ./cmd/web*
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/AuthorizationService .
RUN chmod +x ./authserver
CMD ["./authserver"]
EXPOSE 8081
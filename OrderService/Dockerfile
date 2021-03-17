FROM golang
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./grpc /app/grpc
COPY ./FoodService /app/FoodService
COPY ./AuthorizationService /app/AuthorizationService
COPY ./OrderService /app/OrderService
WORKDIR /app/OrderService
RUN go mod download
RUN go build -o orderserver ./cmd/web4*
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/OrderService .
RUN chmod +x ./orderserver
CMD ["./orderserver"]
EXPOSE 4004
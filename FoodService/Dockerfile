FROM golang
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./grpc /app/grpc
COPY ./OrderService /app/OrderService
COPY ./FoodService /app/FoodService
COPY ./AuthorizationService /app/AuthorizationService
WORKDIR /app/FoodService
RUN go mod download
RUN go build -o foodserver ./cmd/web3*
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/FoodService .
RUN chmod +x ./authserver
CMD ["./foodserver"]
EXPOSE 4040
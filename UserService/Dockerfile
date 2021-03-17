FROM golang
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY ./grpc /app/grpc
COPY ./OrderService /app/OrderService
COPY ./FoodService /app/FoodService
COPY ./AuthorizationService /app/AuthorizationService
COPY ./UserService /app/UserService
WORKDIR /app/UserService
RUN go mod download
RUN go build -o userserver ./cmd/web2*
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/UserService .
RUN chmod +x ./userserver
CMD ["./userserver"]
EXPOSE 4005
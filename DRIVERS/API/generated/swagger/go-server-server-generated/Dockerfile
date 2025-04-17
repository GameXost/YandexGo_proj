
FROM golang:1.23.4 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR /app/API/generated/swagger/drivers/go-server-server-generated
# Собираем статически слинкованный бинарник
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o swagger .

FROM alpine:latest
WORKDIR /root/
# Копируем только скомпилированный бинарный файл
COPY --from=build /app/API/generated/swagger/drivers/go-server-server-generated/swagger .
EXPOSE 8081
ENTRYPOINT ["./swagger"]

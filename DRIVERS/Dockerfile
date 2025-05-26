FROM golang:1.23 AS build
WORKDIR /app

COPY . .



ENV CGO_ENABLED=0
RUN go build -o /app/bin/server ./cmd/main.go

FROM gcr.io/distroless/static-debian12
ENV GIN_MODE=release
COPY --from=build /app/bin/server /server
COPY --from=build /app/config /DRIVERS/config
EXPOSE 8080
ENTRYPOINT ["/server"]

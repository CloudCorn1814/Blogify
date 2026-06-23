FROM golang:1.26 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o main /cmd/main.go

FROM alpine:latest
COPY --from=build /app ./

EXPOSE 8080
ENTRYPOINT ["./main"]
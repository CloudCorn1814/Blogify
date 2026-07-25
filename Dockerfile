FROM golang:1.26 AS build
WORKDIR /app
COPY . .
ARG SERVICE
RUN CGO_ENABLED=0 go build -o main ./cmd/${SERVICE}/main.go

FROM alpine:latest
COPY --from=build /app/main ./

ENTRYPOINT ["./main"]
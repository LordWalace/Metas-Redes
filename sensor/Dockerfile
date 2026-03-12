# build stage
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN go build -o /sensor .

# runtime stage
FROM alpine:latest
COPY --from=build /sensor /sensor
EXPOSE 9000
ENTRYPOINT ["/sensor"]

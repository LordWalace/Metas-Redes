# build stage
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN go build -o /consumer .

# runtime stage
FROM alpine:latest
COPY --from=build /consumer /consumer
ENTRYPOINT ["/consumer"]

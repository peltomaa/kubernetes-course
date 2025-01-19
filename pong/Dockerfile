# Stage 1: Build stage
FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final stage
FROM alpine:edge

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]

EXPOSE 3000

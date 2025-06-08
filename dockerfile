FROM golang:1.24.3 AS builder

WORKDIR /app

# Copy go mod files
COPY absolutecinema/go.mod absolutecinema/go.sum ./absolutecinema/
WORKDIR /app/absolutecinema
RUN go mod download

# Copy the entire backend source
COPY absolutecinema/ /app/absolutecinema/

# Copy the resources folder explicitly
COPY absolutecinema/resources /app/absolutecinema/resources

# Copy .env
COPY absolutecinema/.env /app/absolutecinema/.env

# Build the application
WORKDIR /app/absolutecinema/cmd
RUN go build -o /main main.go

# Stage 2: Runtime
FROM debian:bookworm-slim

# Set working directory to where the app and resources live
WORKDIR /app/absolutecinema

# Copy the built binary from the builder stage
COPY --from=builder /main .

# Copy the .env file as well
COPY --from=builder /app/absolutecinema/.env .

EXPOSE 8080

CMD ["./main"]

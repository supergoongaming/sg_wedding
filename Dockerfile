FROM golang:1.21 AS build-stage
# Make the /app folder and cd into it
WORKDIR /app
# Get our go modules so that we can install them
COPY go.mod ./
# Download them
RUN go mod download
COPY src/ /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /wedding

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage
# Use this for debugging as it has sh and bash
# FROM debian:buster-slim AS build-release-stage
WORKDIR /app
COPY --from=build-stage /wedding /app
COPY templates /app/templates
COPY static /app/static
EXPOSE 8080
ENTRYPOINT ["/app/wedding"]
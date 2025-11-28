# Fetch
FROM golang:latest AS fetch-stage
COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Build
FROM golang:latest AS build-stage
COPY --from=generate-stage /app /app
WORKDIR /app
# RUN apt-get update && apt-get install -y nodejs npm && npm install -g @unocss/cli
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Deploy
FROM gcr.io/distroless/base-debian12 AS deploy-stage
WORKDIR /
COPY --from=build-stage /main /main
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/main"]

# ---------- Build (leve e rápido)
FROM golang:1.22-alpine AS build
WORKDIR /app
RUN apk add --no-cache git ca-certificates
COPY go.mod ./
RUN go mod download
COPY . .
# build estático
RUN CGO_ENABLED=0 GOOS=linux go build -o incident-assistant

# ---------- Runtime (seguro, sem shell, non-root)
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=build /app/incident-assistant .
EXPOSE 8080
USER nonroot:nonroot
CMD ["./incident-assistant"]

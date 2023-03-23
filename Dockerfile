# BUILDER
FROM golang:1.19 as builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o drone-synochat

# PRODUCTION IMAGE
FROM plugins/base:multiarch
COPY --from=builder /app/drone-synochat /bin/drone-synochat
ENTRYPOINT ["/bin/drone-synochat"]

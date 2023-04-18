# This Dockerfile allows grabbing private dependencies in two ways:

FROM golang:1.20 as builder
WORKDIR /app
COPY . .

RUN make _bindata
RUN mkdir -p internal/bindata/migrations/ && go get github.com/go-bindata/go-bindata/go-bindata && go run github.com/go-bindata/go-bindata/go-bindata -nocompress -prefix "./migrations/" -pkg "migrations" -o "internal/bindata/migrations/migrations.go" "migrations"
RUN CGO_ENABLED=0 GOOS=linux go build -tags bindata -a -installsuffix cgo -o mtl ./cmd/cli/

FROM gcr.io/distroless/static-debian10

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder --chown=nonroot:nonroot /app/mtl .

ARG VERSION
ENV APP__VERSION="${VERSION}"
USER nonroot

# Command to run the executable
CMD ["./mtl", "server", "start"]

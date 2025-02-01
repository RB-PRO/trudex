FROM docker.io/rbproxy/trud_base:1.0 AS build

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY . .
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/trud-distributor-cli trud_distributor/main.go


FROM alpine:3.21.2

WORKDIR /app
COPY --from=build /build/trud-distributor-cli /app/trud-distributor-cli
RUN chmod +x /app/trud-distributor-cli
RUN ls -l /app/trud-distributor-cli

CMD ["/app/trud-distributor-cli"]
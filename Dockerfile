FROM rust as builder
WORKDIR /app
COPY . .
RUN cargo install --path .

FROM debian:12-slim
COPY --from=builder /usr/local/cargo/bin/example-service /usr/local/bin/example-service
ENTRYPOINT ["/usr/local/bin/example-service"]

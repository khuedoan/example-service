FROM rust as builder
WORKDIR /app
COPY . .
RUN cargo install --path .

FROM debian:12-slim
# This project is for testing, there's no need to optimize the image size
RUN apt-get update && apt-get install -y iputils-ping curl
COPY --from=builder /usr/local/cargo/bin/example-service /usr/local/bin/example-service
ENTRYPOINT ["/usr/local/bin/example-service"]

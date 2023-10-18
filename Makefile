.POSIX:

.PHONY: default
default: build

.PHONY: dev
dev:
	cargo run

.PHONY: build
build:
	cargo clippy
	cargo build --release

.PHONY: fmt
fmt:
	cargo fmt

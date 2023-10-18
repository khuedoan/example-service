use axum::{extract::Path, http::StatusCode, response::Json, routing::get, Router};
use serde_json::{json, Value};
use std::env;

async fn healthz() -> StatusCode {
    StatusCode::OK
}

async fn ping() -> &'static str {
    "pong"
}

async fn info() -> Json<Value> {
    Json(json!({
        "hostname": hostname::get().unwrap().to_string_lossy()
    }))
}

async fn env(Path(name): Path<String>) -> String {
    match env::var(name) {
        Ok(value) => value,
        Err(_) => "".to_string(),
    }
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/", get(info))
        .route("/healthz", get(healthz))
        .route("/ping", get(ping))
        .route("/env/:name", get(env));

    axum::Server::bind(&"0.0.0.0:3000".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}

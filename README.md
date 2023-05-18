# micro service template

## this repository is template for easier establishment of micro service

### technology

- Go

  - Gorilla Mux (go get github.com/gorilla/mux)
  - Gorm (go get gorm.io/gorm)
  - go-migrate (go get github.com/golang-migrate/migrate/v4/cmd/migrate)

- Rust

  - warp
    tokio = { version = "1", features = ["full"] }
    warp = "0.3"
  - hyper (hyper = "version")
  - sqlx (sqlx = "version")

- database

  - MySQL
  - Redis

- AWS

  - localstack
    - S3
    - EC2

- SSL/TLS

  - Let's encrypt
  - Nginx

- CI/CD
  - circle CI

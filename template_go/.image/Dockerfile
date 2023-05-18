FROM golang:latest AS build_stage
WORKDIR /app

COPY ./src/ .
RUN go install
RUN go build -o main

# ビルドステージ2: 最終的なイメージを構築する
FROM base_image AS final_stage
WORKDIR /app

# ビルドステージ1から必要な成果物をコピーする
COPY --from=build_stage /app/build ./build

# アプリケーションを実行するコマンドを指定する
CMD ["main"]

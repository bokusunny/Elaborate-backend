FROM golang:1.11.11

WORKDIR /go/src/Elaborate-backend
COPY . .
ENV GO111MODULE=on
ENV GO_ENV=prod

# バイナリの実行ファイルを生成する
RUN go build -o main /go/src/Elaborate-backend

CMD ["./main"]

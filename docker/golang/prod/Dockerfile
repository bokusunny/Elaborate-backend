FROM golang:1.11.11

WORKDIR /go/src/Elaborate-backend
COPY . /go/src/Elaborate-backend
ENV GO111MODULE=on
ENV GO_ENV=prod

# バイナリの実行ファイルを生成する
RUN go build -o main .

CMD ["./main"]

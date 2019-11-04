FROM golang:1.11.11

WORKDIR /go/src/Elaborate-backend
COPY . .
ENV GO111MODULE=on
ENV port=3000
ENV GO_ENV=prod

RUN go get github.com/pilu/fresh
CMD ["fresh"]

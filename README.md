# Elaborate-backend
Elaborate backend with Golang.

## Set up
### building
```sh
# https://hub.docker.com/search/?offering=community&type=edition などで予めDockerのインストールをしておく
$ git clone https://github.com/bokusunny/Elaborate-backend.git

$ docker-compose build # 初回およびdocker関連のファイルを書き変えた時のみ
$ docker-compose up -d

# セットアップが上手くいっているか確認
$ docker exec -it elaborate-backend go version # go version go1.11.11 linux/amd64と表示されたらOK
$ docker exec -it elaborate-mysql mysql -u root -p # Enter password: と出てくるのでpasswordと打ち込んでmysqlに入れたらOK
```
### FYI
```sh
$ docker ps # 起動中のコンテナを確認
$ docker exec -it elaborate-backend bash # backendコンテナに入ってコマンドを打ちたい時
$ docker exec -it elaborate-mysql bash # mysqlコンテナに入ってコマンドを打ちたい時
$ docker-compose stop # コンテナを停止する
```

# Elaborate-backend
Elaborate backend with Golang.

## Set up
### firebase secret key
Firebase secret keys like below are required.
```json
{
    "type": "service_account",
    "project_id": "...",
    "private_key_id": "...",
    "private_key": "-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----\n",
    "client_email": "firebase-adminsdk-awott@progate-mafia-tmp.iam.gserviceaccount.com",
    "client_id": "...",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-awott%40progate-mafia-tmp.iam.gserviceaccount.com"
}
```
Please contact the author to get the keys and paste them after creating a file like this.
```sh
# At root directory of Elaborate-backend
touch service_account_key.json
```
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

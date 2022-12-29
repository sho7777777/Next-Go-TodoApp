# Next-Go-TodoApp

#### 1. go-appディレクトリに移動し、下記コマンドを実行
- docker compose up

#### 2. 下記コマンドを実行し、DBを作成
- docker exec -it <コンテナ名> /bin/bash
- mysql -u root -p
- go-app/database.sql の中身を実行

#### 3. go-appディレクトリに移動し、下記コマンドを実行
- DB_NAME=todo DB_USER=root DB_PASSWD=password DB_ADDR=localhost:3306 DB_NET=tcp go run main.go

#### 4. next-appディレクトリに移動し、下記コマンドを実行
- npm upgrade
- npm run dev


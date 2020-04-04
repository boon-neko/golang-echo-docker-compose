# golang-echo-docker-compose

## 概要
- golang-version 1.10.3
- Flamework Echo <https://echo.labstack.com/>  
ローカル上でdep ensureしないと動かない  
ローカルの好きな場所にプロジェクトをクローン  
ローカルのGOPATH/src配下にプロジェクトへのシンボリックリンクを作成  
例　/Users/tanakaamoru/.go/src/nxsw  
上記パスのディレクトリへ移動し、dep ensureすることで依存pkgをインストールしてくれる  
docker-compose up -d で起動、docker-compose downで停止  
docker-exec -it コンテナの名前 bash でコンテナの中へ  
cd src/nxsw/nxsw-golang-docker/src/  
go run main.goで動く  
http://locallhost:8081 にアクセスして表示されれば動いてる  
curl -X POST "http://localhost:8081/routing_check"  -H 'Content-Type: application/json' -d '{"name":"next tarou","email":"tarou@nxsw.co.jp"}'  
で同じ内容のjsonが帰って来ればPOST成功  

syncの速度を解決するためにdocker-syncを導入したいけど動かないから後ほど  

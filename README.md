
# Beego 学习项目1

go get -u github.com/go-sql-driver/mysql

go get -u github.com/lib/pq

bee migrate -driver=mysql -conn="root:******@tcp(127.0.0.1:8889)/study-dev"

## 请确保 mysql 支持 inodb 引擎和 utf8mb4 编码
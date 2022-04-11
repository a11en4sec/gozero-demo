## 1 install
```
go install github.com/Mikaelemmmm/sql2pb@latest
```

## 2 在pb目录下,使用
```
$ sql2pb -go_package ./pb -host localhost -package pb -password root -port 3306 -schema usercenter -service_name usersrv -user root > usersrv.proto
```
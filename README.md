A project build with Go language.

Libraries used in the project:
* gin: used router, middleware

``` go
go get -u github.com/gin-gonic/gin
```

* zap: used to save logs
```go
go get -u go.uber.org/zap
```

* lumberjack: its primary purpose is to facilitate logging activities by providing functionality specifically designed for managing log files.

```go
go get gopkg.in/natefinch/lumberjack.v2
```

* viper: used to config manager

```go
go get github.com/spf13/viper
```

* Run project: 
```go
go run .\cmd\server\main.go 
```






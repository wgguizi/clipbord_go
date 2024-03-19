#### 使用到的第三方库
1. 配置文件 viper，go get -u github.com/spf13/viper
2. 日志记录库 zap，go get -u go.uber.org/zap
3. 日志分割库 lumberjack，go get -u github.com/natefinch/lumberjack
4. 使用web路由框架 httprouter, go get -u github.com/julienschmidt/httprouter
5. 使用mysql驱动，go get -u github.com/go-sql-driver/mysql
6. 使用sqlx驱动，go get -u github.com/jmoiron/sqlx

#### 配置定时清除日志脚本
* * * * * cd /data/wwwroot/clipboard_go && ./main -task removeExpire -expireSeconds 600 > /dev/null 2>&1
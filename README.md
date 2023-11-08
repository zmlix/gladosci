<h1 align="center">Glados Checkin</h1>

# 项目介绍

Glados自动签到工具，签到后会向微信发送签到结果。

# 使用方法

## 编译
```shell
go build .
```

## 使用
```shell
# 每天8点自动签到
./gladosci -cron "0 0 8 * * *"

# 添加签到用户
./gladosci -add

# 指定用户进行签到
# 'all' 全部签到
# uId 指定uId用户签到
./gladosci -checkin all
./gladosci -checkin xxxxxxxxxxx
```
## Gin使用教程

#### 下载httpie
```shell
sudo apt-get install httpie
```

#### 测试使用
##### 登入
```
http -v --json POST localhost:8000/login username=admin password=admin

#### 实际请求
POST /login HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 42
Content-Type: application/json
Host: localhost:8000
User-Agent: HTTPie/0.9.8

{
    "password": "admin",
    "username": "admin"
}

# 返回内容
HTTP/1.1 200 OK
Content-Length: 213
Content-Type: application/json; charset=utf-8
Date: Mon, 25 Nov 2019 06:08:22 GMT

{
    "code": 200,
    "expire": "2019-11-25T15:08:22+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ2NjU3MDIsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU3NDY2MjEwMn0.QolKvUF3a9oJ-UoTofq7uO3cuNTzgfS7GeJVpixE1vY"
}

```

##### 请求
```
http -v -f GET localhost:8000/auth/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ2NjU3NjAsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU3NDY2MjE2MH0.n78ochCTgI36sWovw053awy2an-asLND9FAxozOcEqA"  "Content-Type: application/json"

# 请求
GET /auth/hello HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ2NjU3NjAsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU3NDY2MjE2MH0.n78ochCTgI36sWovw053awy2an-asLND9FAxozOcEqA
Connection: keep-alive
Content-Type: application/json
Host: localhost:8000
User-Agent: HTTPie/0.9.8


# 回复
HTTP/1.1 200 OK
Content-Length: 60
Content-Type: application/json; charset=utf-8
Date: Mon, 25 Nov 2019 06:14:11 GMT

{
    "text": "Hello World.",
    "userID": "admin",
    "userName": "admin"
}
http -v -f GET localhost:8000/auth/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ2NjU3NjAsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU3NDY2MjE2MH0.n78ochCTgI36sWovw053awy2an-asLND9FAxozOcEqA"  "Content-Type: application/json"
GET /auth/hello HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ2NjU3NjAsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU3NDY2MjE2MH0.n78ochCTgI36sWovw053awy2an-asLND9FAxozOcEqA
Connection: keep-alive
Content-Type: application/json
Host: localhost:8000
User-Agent: HTTPie/0.9.8



HTTP/1.1 200 OK
Content-Length: 60
Content-Type: application/json; charset=utf-8
Date: Mon, 25 Nov 2019 06:14:11 GMT

{
    "text": "Hello World.",
    "userID": "admin",
    "userName": "admin"
}

```

#### 参考资料
+ [《10分钟了解JSON Web令牌（JWT）》](https://baijiahao.baidu.com/s?id=1608021814182894637&wfr=spider&for=pc)
+ [《JWT 也不是万能的呀，入坑需谨慎！》](https://cloud.tencent.com/developer/article/1495531)
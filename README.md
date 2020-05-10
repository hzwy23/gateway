# API 网关

## 获取代码
```shell
go get -u https://github.com/wisrc/gateway
```

## 配置介绍
```yml
server:
    port: 8790
    contextPath: /
    timeout: 10
    host:
registerCenter:
    refreshFrequency: 30
    eureka:
        serviceUrls: [ http://localhost:8761 ]
router:
    gateway:
        ignoredPatterns: [ /gateway, /js, /css ]
        sensitiveHeaders: [Cookie]
        routers:
            user:
                path: /map/**
                serviceId: map
                stripPrefix: true
                timeout: 30
            ai:
                path: /gitchat/columns/category/5d8b7c3786194a1921979123?page=1
                url: https://gitbook.cn
                stripPrefix: false
                timeout: 5

```

## 测试地址：
```shell
 http://localhost:8790/gitchat/columns/category/5d8b7c3786194a1921979123?page=1
```
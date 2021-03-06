# Xilixili
A demo website for video watching and upload.  
你可以在xilixili上登录，上传，观看视频。 Go语言编写，服务端使用了Gin+Gorm框架，使用MySQL和Redis储存。  
<https://www.xilixili.site>

## 关键词
Go+Gin+Gorm  
MySQL+Redis  
Docker+Kubernetes  
Vue.js  前端请参照 <https://github.com/whjdarking/xilixili_view>

## 文件树
```
│  .gitignore
│  Dockerfile_example  //打包到docker
│  env_example  //环境变量设置同上，本地开发测试时方便设置
│  go.mod
│  go.sum
│  main.go  //主入口
│  README.md
│  Refactor_xilixili  //build后得到的linux可执行文件
│  
├─cache  //redis连接等
│      redis.go
│      
├─config
│      config.go  //初始化
│      cron.go  //计时任务，用来每周清空排行榜
│      
├─model  //mysql连接，orm的model等
│      init.go
│      user.go
│      video.go
│      
├─serializer  //序列化，将struct变成json
│      common.go
│      user.go
│      video.go
│      
├─server  //router等网络后端相关
│      apiMethods.go
│      cors.go
│      router.go
│      session.go
│      
└─service  //反序列化json,并实际处理后端任务
       create_video_service.go
       delete_video_service.go
       list_video_service.go
       ......
```

## 数据库
项目需要预先配置好MySQL和Redis，否则会直接panic。  
视频和封面保存在阿里云OSS，MySQL保存用户数据和视频等内容的地址。  
Redis专门保存视频点击量，同时用作weekly排行榜。

## 其它
* cookie+sessionID模式  
* 类MVC架构，restful设计，前后端分离  
* 本项目运行在Centos7+Kubernetes内，实现了https化，部分yaml文件可参照[other_resources](other_resources)文件夹。  
* 提供简单的postman接口测试，了解各url接口设计的功能，参照文件夹。

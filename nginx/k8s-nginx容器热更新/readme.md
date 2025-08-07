## 说明
Nginx广泛被作为web服务器使用，自身支持热更新操作。如我们在宿主机上通过yum/apt/编译安装/二进制安装nginx，在我们修改配置文件之后，执行nginx -s reload命令可以不停服务重新加载配置。

对于使用docker部署的nginx，我们也可以使用docker exec -it nginx-container service nginx reload 来完成修改后的配置文件重新加载，但无论对于docker/k8s这样过多的人工干预还是很痛苦的。

本文将使用sidecar(边车)方案来完成nginx的自动热更新，同时通过演示k8s部署nginx下载服务器来更好的帮助大家理解。
### 部署过程
部署过程主要分两部分，一是nginx热更新镜像build，二是k8s部署nginx下载服务器。当然，镜像也可直接使用我分享的~

nginx热更新镜像build
Dockerfile
Sidecar容器，nginx-reloader镜像的Dockerfile如下：
```
# cat Dockerfile

FROM golang:1.12.0 as build
RUN go get /fsnotify/fsnotify
RUN go get /shirou/gopsutil/process
RUN mkdir -p /go/src/app
ADD main.go /go/src/app/
WORKDIR /go/src/app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o nginx-reloader .
# main image
FROM nginx:1.14.2-alpine
COPY --from=build /go/src/app/nginx-reloader /
CMD ["/nginx-reloader"]

```

nginx热更新功能是用一个叫main.go的go脚本实现的,见 k8s-nginx-reload.go

### 构建镜像
基于该Dockerfile构建nginx-reloader镜像：
``` 
docker build -t /danqing/nginx-reloader:20250806 .
```

### k8s部署nginx
见deploy目录文件


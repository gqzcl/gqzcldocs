## 详细文档参阅

- [docker中文文档](http://www.dockerinfo.net/)
- [docker使用指南](http://shouce.jb51.net/docker_practice/content.html)
- [docker常用命令](https://zhuanlan.zhihu.com/p/245641060)

## 环境：

```5.8.18-1-MANJARO```

## 开机自动启动

```sh
$ sudo systemctl enable docker
```

启动docker

```sh
$ sudo systemctl start docker
```

重启doker

```sh
$ sudo systemctl restart docker
```
检查是否安装成功：

```sh
$ sudo docker run hello-world
```

如果安装成功应该会看到Hello from Docker。

## 镜像加速：

参考：https://www.runoob.com/docker/docker-mirror-acceleration.html
推荐使用阿里云镜像加速器：https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors

如果没有docker文件夹则新建一个

```sh
$ sudo mkdir -p /etc/docker
```

写入：

```sh
$ sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://<id>.mirror.aliyuncs.com"]
}
EOF
```

保存后重启服务：

```sh
$ sudo systemctl daemon-reload
$ sudo systemctl restart docker
```

## 下载镜像：

```sh
$ sudo docker pull ubuntu
```

也可以自己下别的，具体去 [docker hub](https://hub.docker.com/) 看

## 运行镜像：

```sh
$ sudo docker run -t -i ubuntu /bin/bash
```

查看镜像信息：

```sh
$ cat /proc/version
```

退出： exit

列出本地镜像：

```sh
$ sudo docker images
```

创建Docker用户组：

```sh
$ sudo usermod -aG docker <username>
$ restart
```

查看info 

```sh
$ docker info
```

## 启动一个bash终端

可以使用下面的exec命令

```
$ sudo docker run -t -i ubuntu:14.04 /bin/bash
```

其中，-t 选项让Docker分配一个伪终端（pseudo-tty）并绑定到容器的标准输入上， -i 则让容器的标准输入保持打开

## docker exec

在容器中执行命令

```
$ docker exec id command
```

绑定容器到bash

```sh
docker exec -it id bash
```

## 守护态运行

更多的时候，需要让 Docker 容器在后台以守护态（Daemonized）形式运行。此时，可以通过添加 -d 参数来实现。

例如下面的命令会在后台运行容器。

```sh
$ sudo docker run -d ubuntu:14.04 /bin/bash -c "while true; do echo hello world; sleep 1; done"
```

注： 容器是否会长久运行，是和docker run指定的命令有关，和 -d 参数无关。

要获取容器的输出信息，可以通过 docker logs 命令。

```sh
$ sudo docker logs id
```

## docker查看容器IP地址

```sh
$ docker inspect id
```

## docker 为镜像添加标签

```sh
$ docker tag [REPOSITORY] [name]
```

## 使用 attach 卡住问题

大概就是因为启动容器时并不是一个bash输出或者进程并不支持写入

改用

```sh
$ docker exec -it +容器ID bash
```

退出使用ctrl + 大写P + 大写Q
实测 ctrl + 大写Q




## docker stop 没有作用

遇到docker无法删除，或者kill相应的容器，要么是运行完docker stop xxx后发现xxx仍然存在，要么就根本无法删除，或者发现会报错，提示 Error response from daemon: Conflict, cannot remove the default name of the container

可能是因为容器意外退出，但是并没有**优雅的**终止容器，剩下的文件阻止你重新生成旧名称的新容器或者占用网络端口

1、强制删除容器

```sh
$ docker rm -f name/id
```

2、清理此容器的网络占用

```sh
docker network disconnect --force bridge name/id
```

##  防火墙设置允许访问的端口

### ufw

对于启用了ufw的主机（基于Debian的发行版）, 你可以通过ufw命令允许指定端口上的所有流量连接. 通过如下命令允许访问端口9000

```sh
ufw allow 9000
```

如下命令允许端口9000-9010上的所有传入流量

```sh
ufw allow 9000:9010/tcp

```

### firewall-cmd

对于启用了firewall-cmd的主机（CentOS）, 你可以通过firewall-cmd命令允许指定端口上的所有流量连接。 通过如下命令允许访问端口9000

```sh
Copyfirewall-cmd --get-active-zones
```
这个命令获取当前正在使用的区域。 现在，就可以为以上返回的区域应用端口规则了。 假如返回的区域是 public, 使用如下命令

```sh
Copyfirewall-cmd --zone=public --add-port=9000/tcp --permanent
```

这里的permanent参数表示持久化存储规则，可用于防火墙启动、重启和重新加载。 最后，需要防火墙重新加载，让我们刚刚的修改生效。

```sh
Copyfirewall-cmd --reload
```

### iptables

对于启用了iptables的主机（RHEL, CentOS, etc）, 你可以通过iptables命令允许指定端口上的所有流量连接。 通过如下命令允许访问端口9000

```sh
Copyiptables -A INPUT -p tcp --dport 9000 -j ACCEPT
service iptables restart
```

如下命令允许端口9000-9010上的所有传入流量。

```sh
Copyiptables -A INPUT -p tcp --dport 9000:9010 -j ACCEPT
service iptables restart
```
## 删除容器

可以使用 docker rm 来删除一个处于终止状态的容器。 例如

```sh
$sudo docker rm  trusting_newton
trusting_newton
```
如果要删除一个运行中的容器，可以添加 -f 参数。Docker 会发送 SIGKILL 信号给容器。

## 删除本地镜像

如果要移除本地的镜像，可以使用 docker rmi 命令。注意 docker rm 命令是移除容器。

```sh
$ sudo docker rmi training/sinatra
Untagged: training/sinatra:latest
Deleted: 5bc342fa0b91cabf65246837015197eecfa24b2213ed6a51a8974ae250fedd8d
Deleted: ed0fffdcdae5eb2c3a55549857a8be7fc8bc4241fb19ad714364cbfd7a56b22f
Deleted: 5c58979d73ae448df5af1d8142436d81116187a7633082650549c52c3a2418f0
```

-> *注意：在删除镜像之前要先用 docker rm 删掉依赖于这个镜像的所有容器。

## 存出镜像
如果要导出镜像到本地文件，可以使用 docker save 命令。

```sh
$ sudo docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
ubuntu              14.04               c4ff7513909d        5 weeks ago         225.4 MB
...
$sudo docker save -o ubuntu_14.04.tar ubuntu:14.04
```

## 载入镜像
可以使用 docker load 从导出的本地文件中再导入到本地镜像库，例如

```sh
$ sudo docker load --input ubuntu_14.04.tar
```
或

```sh
$ sudo docker load < ubuntu_14.04.tar
```

这将导入镜像以及其相关的元数据信息（包括标签等）。

minios
```sh
docker run -t -i -d -p 9000:9000 \
  -e "MINIO_ROOT_USER=gqzcl" \
  -e "MINIO_ROOT_PASSWORD=systemlogin" \
  minio/minio server /data /bin/bash
```

## 如果使用docker作为开发环境进行开发

docker run -p 82(本地端口):80(docker端口) -v filepath id

-v 指的是把我们本地的代码动态检测并拷贝到容器里面的位置上。
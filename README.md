这个是一个tcp流量转发的代理工具，写他的目的是为了绕过跳板机器ssh登录服务器。

# Usage
```shell
make build
// 将本地7013的流量代理到hellojukay.cn:8080
./tcp_proxy --listen=7013 --target=hellojukay.cn:8080
```

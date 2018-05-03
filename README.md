这是一个ssh端口代理程序，主要是用来绕过堡垒机端口和ip限制进行ssh登录！

# Usage
```shell
go build ssh_proxy 
// 在一个开放7013(或者其他端口)的端口机器上执行
./ssh_proxy -ssh="server:port"  -port=7013
```
本机登录ssh, 免密码登录同理
```shell
ssh -p 7013 user@targethost
```

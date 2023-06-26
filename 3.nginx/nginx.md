# nginx

## 正向代理和反向代理

## 负载均衡
负载均衡策略：
* 内置策略  轮询 加权轮询 Ip hash
* 扩展策略   

## 常用命令
```shell
# 启动
nginx

# 停止 
nginx -s stop
# 安全退出
nginx -s quit
# 重新加载配置文件
nginx -s reload
# 查看nginx进程
ps aux|grep nginx
```

## linux相关命令
```shell
# 开启防火墙
service firewalld start

# 重启防火墙
service firewalld restart
# 关闭防火墙
service firewalld stop
# 产看防火墙规则
firewall-cmd --list-all
# 查看端口是否开放
firewall-cmd --query-port=8000/tcp
# 开放80端口
firewall-cmd --permanent --add-port=80/tcp
# 移除端口
firewall-cmd -permanent --remove-port=80/tcp

#参数解释
firewall-cmd 
--permanent 设置为持久
--add-port 添加端口
--remove-port 移除的端口

```
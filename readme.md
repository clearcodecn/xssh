### xssh 

一款自用多服务器 ssh 登录管理工具. 

安装: 
```azure
go get -u github.com/clearcodecn/xssh 
```

功能如下: 
```azure
# 新增一个服务器配置 s
# -a 设置别名备注, -c 配置 
xssh add -a baidu -c root@baidu.com
    
# 列表
xssh list 

# 删除 
xssh del 
    
# 登录 
xssh ssh 
```

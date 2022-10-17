# ovpn-manager

## 功能列表

* 新增openvpn用户;
* 使用角色管理，用户绑定角色，获取对应的权限
* 撤销/恢复用户证书;
* 生成可供用户使用的配置文件;
* 为 Prometheus 提供指标，包括证书到期日期、（已连接/总）用户数、已连接用户信;
* (optionally) Specifying CCD (`client-config-dir`) for each user;
* (optionally) Operating in a master/slave mode (syncing certs & CCD with other server);
* (optionally) Specifying/changing password for additional authorization in OpenVPN;
* (optionally) Specifying the Kubernetes LoadBalancer if it's used in front of the Open

### er图

![ovpn-admin UI](img/openvpn.png)

### 展示

Managing users in ovpn-admin:
![ovpn-admin UI](img/ovpn-admin-users.png)

An example of dashboard made using ovpn-admin metrics:
![ovpn-admin metrics](img/ovpn-admin-metrics.png)

## 本地启动说明
### 生成前端 js
```shell
cd h5/openvpn-ui
npm install
```
### 运行go项目增加环境变了
```shell
本地启动需要的参数
OVPN_INDEX_PATH=./easyrsa_master/pki/index.txt;OVPN_CCD=True;OVPN_CCD_PATH=./ccd_master;OVPN_AUTH=true;EASYRSA_PATH=./easyrsa_master;DB_PATH=./easyrsa_master/openvpn.db
```

###编译 Dockerfile.openvpn 命令
```shell
docker build -t openvpn:v1 -f Dockerfile.openvpn .
docker run -t -i openvpn:v1 /bin/bash
```

### 启动
```shell
docker-compose up
```

### 用户名密码初始化
打开 http://127.0.0.1:8080
admin/123456

### 结束
```shell
docker-compose down
```

### 借鉴工程
感谢

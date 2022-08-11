## 完成内容
后端逻辑直接修改

    1.用户登录功能
    2.用户拦截功能
    3.用户角色功能
    4.用户角色配置功能
    5.角色路由配置功能
前端逻辑

    最后修改


## 本地启动说明
### 生成前端 js
```shell
cd frontend 
npm install
npm run build
```
### 运行go项目增加环境变了
```shell
OVPN_INDEX_PATH=./easyrsa_master/pki/index.txt;OVPN_CCD=True;OVPN_CCD_PATH=./ccd_master;OVPN_AUTH=true;OVPN_AUTH_DB_PATH=./easyrsa_master/pki/users.db;EASYRSA_PATH=./easyrsa_master;DB_PATH=./easyrsa_master/openvpn.db
```

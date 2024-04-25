## 基于始皇 oaifree 的GPT私人拼车实现


实现了网站密码，聊天隔离，限定访问用户，自动刷新access_token等

### 使用方法

#### docker compose 运行
```yaml
version: '3.8'
services:
  oai-proxy:
    image: neccen/oai-proxy
    ports:
      - "自定义端口:48881"
    container_name: oai-proxy
    volumes:
      - /文件config目录/config.yml:/home/oai-proxy/config.yml
    restart: unless-stopped
```

#### docker 运行
 - `docker run -itd -p 自定义端口:48881 -v /config文件目录/config.yml:/home/oai-proxy/config.yml --name oai-proxy --restart always neccen/oai-proxy:latest`

#### 本地运行
 - 在[Releases · ycvk/oai-proxy](https://github.com/ycvk/oai-proxy/releases) 下载最新的对应版本压缩文件
 - 解压并进入目录，在执行程序目录下，新建`config.yml`文件，并修改配置文件
 - 运行程序

#### docker compose 编译运行
- git pull下来源码后，修改`config.yml`配置文件
- 自定义修改`docker-compose.yml`中的 端口 和映射的`config.yml`文件目录
- `docker compose up -d`

### `config.yml`配置文件说明
```yaml
access_token: xxx #OpenAI的access_token 
refresh_token: xxx #可以为空，会自动根据access_token获取并保存
site_password: abc123 #网站的密码
allow_users: #允许登录的用户
   - 
   - abc
   - e#F8!@
   - "666" #纯数字需要加引号
```
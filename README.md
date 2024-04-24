## 基于始皇 oaifree 的GPT私人拼车实现


实现了网站密码，聊天隔离，自动刷新access_token等

### 使用方法

1. docker compose 编译运行
 - git pull下来源码后，修改`config.yml`配置文件
 - 自定义修改`docker-compose.yml`中的 端口 和映射的`config.yml`文件目录
 - `docker compose up -d`


2. docker 运行
 - `docker run -itd -p 自定义端口:48881 --name oai-proxy --restart always -v /config文件目录/config.yml:/home/oai-proxy/config.yml neccen/oai-proxy:latest`
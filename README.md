## 基于 oaifree 的GPT私人拼车实现


实现了网站密码，聊天隔离，限定访问用户，自动刷新access_token等
使用转OAuth的方式登陆，提高安全性，避免share_token泄露

## 页面展示

### 登录
![ebd7fd45550498b1806aaebeac83e260.png](https://i3.mjj.rip/2024/06/27/ebd7fd45550498b1806aaebeac83e260.png)

![7a5b2258dbb2e6bd23097ec195e55f65.png](https://i3.mjj.rip/2024/06/27/7a5b2258dbb2e6bd23097ec195e55f65.png)

### 错误信息

![d4727db474c8de07b4e2d4c1828de86d.png](https://i3.mjj.rip/2024/06/27/d4727db474c8de07b4e2d4c1828de86d.png)

![44c0d1f36c3a4cec8b4b82adb29b8db7.png](https://i3.mjj.rip/2024/06/27/44c0d1f36c3a4cec8b4b82adb29b8db7.png)

### 登录成功

自动跳转gpt页面

![f7dbaeb2142697313232b67e0b72c25b.png](https://i3.mjj.rip/2024/06/27/f7dbaeb2142697313232b67e0b72c25b.png)


## 使用方法

#### docker compose 运行
```yaml
version: '3.8'
services:
  oai-proxy:
    image: neccen/oai-proxy
    ports:
      - "自定义端口:8080"
    container_name: oai-proxy
    volumes:
      - /文件config目录/config.yml:/home/oai-proxy/config.yml
    restart: unless-stopped
```

#### docker 运行
 - `docker run -itd -p 自定义端口:8080 -v /config文件目录/config.yml:/home/oai-proxy/config.yml --restart always neccen/oai-proxy:latest`

#### 本地运行
 - 在[Releases · ycvk/oai-proxy](https://github.com/ycvk/oai-proxy/releases) 下载最新的对应版本压缩文件
 - 解压并进入目录，在执行程序目录下，新建`config.yml`文件，并修改配置文件
 - 运行程序

## `config.yml`配置文件说明
```yaml
access_token: xxx #access_token，可以为空，会自动根据refresh_token获取并保存
refresh_token: xxx #OpenAI的refresh_token, 必填
site_password: abc123 #网站的密码
allow_users: #允许登录的用户
   - abc
   - e#F8!@
   - "666" #纯数字需要加引号
```
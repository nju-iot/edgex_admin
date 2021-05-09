# Edgex-admin

### 环境
go1.15

mysql8.0.21

redis:6.2.2

### 本地启动
- 本地配置mysql，使用./dal/edgex.sql生成数据库表
- 本地配置redis
##### 运行命令
```./build.sh && ./output/bootstrap.sh ```

#### 镜像启动

```docker-compose build```

```docker-compose up -d```
##### 验证
```http://localhost:6789/ping```

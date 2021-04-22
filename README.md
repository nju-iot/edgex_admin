# Edgex-admin

### 环境
go1.15

### 运行命令
```./build.sh && ./output/bootstrap.sh ```

#### 镜像启动

```docker build -t edgex_admin:latest --rm .```

```docker run -it --rm -p 6789:6789 -d edgex_admin:latest```
##### 验证
```http://localhost:6789/ping```

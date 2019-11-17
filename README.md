# 使用说明

### 功能简介

> 该demo使用fasthttp做网关和api接口，使用go-micro做server微服务
>
> 使用etcd或yaml文件做配置
>
> 功能简单明了
>

### 提供了2个脚本

```shell
./proto.sh srv.user.proto
./local.sh srv.user --registry=etcd --registry_address=127.0.0.1:2379 --etcd_addr=127.0.0.1:2379
```


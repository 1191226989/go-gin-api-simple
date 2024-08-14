## 关于

`go-gin-api-simple` 是基于 [go-gin-api](https://github.com/xinliangnote/go-gin-api) 改编设计的简化版 API 框架，省略了后台管理系统功能，致力于进行快速的用户端接口业务研发。

供参考学习，线上使用请谨慎！

## 代码生成命令
```sh
# 生成数据表 CURD
# 在根目录下执行脚本：`./scripts/gormgen.sh addr user pass name tables`；
# addr：数据库地址，例如：127.0.0.1:3306
# user：账号，例如：root
# pass：密码，例如：root
# name：数据库名称，例如：go_gin_api_simple
# tables：表名，默认为 *，多个表名可用英文逗号分割，例如：user_demo
./scripts/gormgen.sh 127.0.0.1:3306 root 123456 go_gin_api_simple prize

# 生成控制器方法
# 1. 在 ./internal/api 目录中，创建 prize 目录；
# 2. 在 prize 目录中，创建 handler.go 文件；
# 3. 在 handler.go 文件中定义需要实现的接口，具体可参考其他 handler.go 文件
./scripts/handlergen.sh prize

# 更新swagger文档
./scripts/swagger.sh
```

## 原版文档索引

- 中文文档：[go-gin-api - 语雀](https://www.yuque.com/xinliangnote/go-gin-api/ngc3x5)
- English Document：[en.md](https://go-gin-api/blob/master/en.md)

Penny Hardaway

#### 目录结构

* conf：用于存储配置文件
* middleware：应用中间件
* models：应用数据库模型
* pkg：第三方包
* routers 路由逻辑处理
* runtime 应用运行时数据

#### 授权

* 获取授权码 http://127.0.0.1:8080/auth?username==__UNAME__&password=__PWD__ 
* 访问时加上授权码 http://127.0.0.1:8080/api/v1/articles?token=eyJhbGci...
* token 三个小时过期

#### swagger doc

* localhost:8080/swagger/index.html
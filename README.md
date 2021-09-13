## 启动
~~~ bash
make
./webconsole
~~~

## TODO
- 加入消息队列以处理缓存与数据库之间的一致性问题(现在使用缓存ttl作为保底，更新数据库前先清除缓存，清除不成功只能等待缓存失效) 计划实现一个定制版的nsq[https://github.com/impact-eintr/esq]
- 加入链路追踪
- 将服务拆分，用微服务思想重构


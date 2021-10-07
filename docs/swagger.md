# 交通一张图后端系统
交通一张图

## Version: 1.0.0

### /api/v1/cache/hit/{infotype}/{year}/{level}

#### GET
##### Summary

检查缓存命中接口

##### Description

检查缓存中是否有请求的值 有就返回没有将请求转发

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | No | string |
| infotype | path | 查询类型 | Yes | string |
| year | path | 查询年份 格式: 202X  | Yes | string |
| level | path | 查询等级 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v1/data/info/{infotype}/{year}/{level}

#### GET
##### Summary

更新缓存接口

##### Description

获取数据库原始数据接口 访问后会更新缓存

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |
| infotype | path | 查询类型 : road(路)  bridge(桥) tunnel(隧道) service(服务区) portal(收费门架) toll(收费站) | Yes | string |
| year | path | 查询年份 格式: 202X  | Yes | string |
| level | path | 查询等级 : 0(高速) 1(一级) 2(二级) 3(三级) 4(四级) 5(等外) | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

#### POST
##### Summary

删除指定数据

##### Description

删除数据库原始数据接口 访问后会删除缓存

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |
| infotype | path | 查询类型 : road(路)  bridge(桥) tunnel(隧道) service(服务区) portal(收费门架) toll(收费站) | Yes | string |
| year | path | 查询年份 格式: 202X  | Yes | string |
| level | path | 查询等级 : 0(高速) 1(一级) 2(二级) 3(三级) 4(四级) 5(等外) | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v1/data/info/{infotype}/{year}/{level}/query

#### POST
##### Summary

获取查询数据

##### Description

查询表数据（年报）数据接口

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |
| infotype | path | 查询类型 : road(路)  bridge(桥) tunnel(隧道) service(服务区) portal(收费门架) toll(收费站) | Yes | string |
| year | path | 查询年份 格式: 202X  | Yes | string |
| level | path | 查询等级 : 0(高速) 1(一级) 2(二级) 3(三级) 4(四级) 5(等外) | Yes | integer |
| 查询信息 | body | 道路信息（示例） | Yes | [model.L21](#modell21) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v1/data/menus

#### GET
##### Summary

获取数据导航栏

##### Description

用于前端渲染数据管理侧边栏

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v1/home/menus

#### GET
##### Summary

获取导航栏数据

##### Description

用于前端渲染用户家目录侧边栏

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v1/home/roles

#### GET
##### Summary

获取权限数据接口

##### Description

请求后可以拿到权限数据

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

#### POST
##### Summary

更新权限数据接口

##### Description

请求后可以跟新权限数据 构造一个新的权限json 作为 body

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v1/home/users

#### GET
##### Summary

获取用户数据接口

##### Description

请求后可以拿到用户数据

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

#### POST
##### Summary

更新用户数据接口

##### Description

请求后可以将post请求body中提供的用户新数据替换原来的用户数据

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |
| 用户信息 | body | 用户信息 | Yes | [model.RespUser](#modelrespuser) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v1/home/users/{id}

#### DELETE
##### Summary

删除用户数据接口

##### Description

请求后可以删除用户数据（直接从数据库删除）

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |
| id | path | 用户id | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v1/home/users/query

#### GET
##### Summary

查询用户数据接口

##### Description

请求后可以获取指定的用户数据

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |
| column | query | 属性 | Yes | string |
| value | query | 值 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v1/login

#### POST
##### Summary

处理登录请求

##### Description

首页登录

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| 登录信息 | body | 用户登录信息 | Yes | [model.ParamLogin](#modelparamlogin) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v1/signup

#### GET
##### Summary

申请一个注册码

##### Description

申请一个注册码 重复请求上一个就会失效 重启后端服务也会失效

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

#### POST
##### Summary

处理注册请求

##### Description

首页注册

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| 注册信息 | body | 用户注册信息 | Yes | [model.ParamSignUp](#modelparamsignup) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [respcode.ResponseData](#respcoderesponsedata) & object |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### /api/v2/iserver/services/{service}/rest/{service}

#### POST
##### Summary

获取IServer数据接口

##### Description

用于转发IServer数据，并做权限检验

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Authorization | header | Bearer 用户令牌 | Yes | string |
| 查询信息 | body | 查询组建构建信息 | Yes | [model.IServerReq](#modeliserverreq) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [model.IServerResp](#modeliserverresp) |

##### Security

| Security Schema | Scopes |
| --- | --- |
| ApiKeyAuth | |

### Models

#### model.IServerReq

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| datasetNames | [ string ] |  | No |
| getFeatureMode | string |  | No |
| queryParameter | object |  | No |

#### model.IServerResp

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| model.IServerResp | object |  |  |

#### model.L21

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |
| 修建年度 | string |  | No |
| 养护里程 | string |  | No |
| 可绿化里程 | string |  | No |
| 国道调整前路线编号 | string |  | No |
| 地貌代码 | string |  | No |
| 地貌汉字 | string |  | No |
| 备注 | string |  | No |
| 已绿化里程 | string |  | No |
| 所在行政区划代码 | string |  | No |
| 技术等级 | string |  | No |
| 技术等级代码 | string |  | No |
| 改建年度 | string |  | No |
| 断链类型 | string |  | No |
| 是否一幅高速 | string |  | No |
| 是否城管路段 | string |  | No |
| 是否按干线公路管理接养 | string |  | No |
| 是否断头路段 | string |  | No |
| 最近一次修复养护年度 | string |  | No |
| 止点名称 | string |  | No |
| 止点桩号 | string |  | No |
| 涵洞数量(个) | string |  | No |
| 省际出入口 | string |  | No |
| 管养单位名称 | string |  | No |
| 设计时速(公里/小时) | string |  | No |
| 起点名称 | string |  | No |
| 起点桩号 | string |  | No |
| 路基宽度(米) | string |  | No |
| 路段收费性质 | string |  | No |
| 路线名称 | string |  | No |
| 路线编号 | string |  | No |
| 路面宽度(米) | string |  | No |
| 车道数量(个) | string |  | No |
| 里程(公里) | string |  | No |
| 重复路段线路编号 | string |  | No |
| 重复路段终点桩号 | string |  | No |
| 重复路段起点桩号 | string |  | No |
| 面层厚度(厘米) | string |  | No |
| 面层类型 | string |  | No |
| 面层类型代码 | string |  | No |

#### model.ParamLogin

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| password | string |  | Yes |
| username | string |  | Yes |

#### model.ParamSignUp

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| password | string |  | Yes |
| repassword | string |  | Yes |
| unit | string |  | Yes |
| username | string |  | Yes |
| verifycode | string |  | Yes |

#### model.RespUser

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| role | string |  | No |
| unit | string |  | No |
| user_id | integer |  | No |
| user_id_str | string |  | No |
| username | string |  | No |

#### respcode.ResponseData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| code | integer |  | No |
| data | object |  | No |
| msg | object |  | No |

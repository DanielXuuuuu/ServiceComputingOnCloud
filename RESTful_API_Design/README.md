# RESTful API  设计

### RESTful架构简介

[参考资料]( http://www.ruanyifeng.com/blog/2011/09/restful.html )

REST，即Representational State Transfer的缩写，由Roy Thomas Fielding在2000年提出。 如果一个架构符合REST原则，就称它为RESTful架构。 要理解REST，首先要先明白以下几个概念：

+ **资源**： 所谓"资源"，就是网络上的一个实体，或者说是网络上的一个具体信息。 我们可以通过一个URI（统一资源标识符 ）指向它，每种资源对应一个特定的URI。 
+ **表现层**： "资源"是一种信息实体，它可以有多种外在表现形式。我们把"资源"具体呈现出来的形式，叫做它的"表现层"（Representation）。 URI只代表资源的实体，不代表它的形式。 

+ **状态转化** ： 互联网通信协议HTTP协议，是一个无状态协议。这意味着，所有的状态都保存在服务器端。因此，如果客户端想要操作服务器，必须通过某种手段，让服务器端发生"状态转化"（State Transfer）。而这种转化是建立在表现层之上的，所以就是"表现层状态转化"。 而这里所说的手段，只能是HTTP协议。具体来说，就是HTTP协议里面，四个表示操作方式的动词：GET、POST、PUT、DELETE。 

综上所述，可以对于RESTful架构得到如下定义：

+ 每一个URI代表一种资源；
+ 客户端和服务器之间，传递这种资源的某种表现层；
+ 客户端通过四个HTTP动词，对服务器端资源进行操作，实现"表现层状态转化"。

RESTful架构，就是目前最流行的一种互联网软件架构与API设计规范，用于Web数据接口的设计。它结构清晰、符合标准、易于理解、扩展方便，所以正得到越来越多网站的采用。 

### RESTful API设计实践

> 模仿 [Github]( https://developer.github.com/v3/ )，设计一个博客网站的 API 。
>
> [设计参考]( http://www.ruanyifeng.com/blog/2018/10/restful-api-best-practices.html )

#### 文章模块

|          描述          |                        url                         |  方法  |
| :--------------------: | :------------------------------------------------: | :----: |
|    获取文章分类列表    |                  /api/categories                   |  GET   |
| 获取全站下某分类的文章 |       /api/articles?categoryId={categoryId}        |  GET   |
|  获取某用户的全部文章  |             /api/users/:owner/articles             |  GET   |
| 获取某用户的某一类文章 | /api/users/:owner/articles?categoryId={categoryId} |  GET   |
|  获取某用户的指定文章  |  /api/users/:owner/articles?articleId={articleId}  |  GET   |
|      用户创建文章      |                 /api/user/articles                 |  POST  |
|      用户修改文章      |          /api/articles/:owner/:articleId           |  PUT   |
|      用户删除文章      |           /api/articles/:owner/:article            | DELETE |
|   获取当前用户的文章   |                 /api/user/articles                 |  GET   |

#### 用户模块

|        描述        |              url              | 方法 |
| :----------------: | :---------------------------: | :--: |
|      用户注册      |          /api/users           | POST |
|      用户登录      |       /api/users/login        | POST |
|  获取当前用户信息  |       /api/user/profile       | GET  |
|  当前用户更新信息  |       /api/user/profile       | PUT  |
| 获取当前用户的粉丝 |      /api/user/followers      | GET  |
| 获取当前用户的关注 |      /api/user/following      | GET  |
| 获取当前用户收藏夹 | /api/user/favorite/folderList | GET  |

#### 搜索模块

|     描述     |              url               | 方法 |
| :----------: | :----------------------------: | :--: |
| 搜索指定用户 |  /api/search/users?q={query}   | GET  |
| 搜索指定文章 | /api/search/articles?q={query} | GET  |




## 项目地址
* [https://github.com/monomania/foot](https://github.com/monomania/foot)

## 项目介绍
```
热衷于足球多年,之余也会去研究一下,时间久了,都会有自己的心得.
但不可能每次都那么费劲的自己人工去看盘分析,
所以结合所学,就有这个项目.
```

```
foot-parent 是一个集足球数据采集器,简单分析,同步到微信及其他发布平台一体化的项目.
程序采用go语言开发,项目结构清晰完整,非常容易入手并进行二次开发分析.
```

## 项目为什么开源?
~~~
让有编程能力的朋友,同样爱好足球的朋友,也可以参与进来.

大家购彩多年,都有自己购彩的心得.
如果大家将这些心得,集中起来,形成一个模型习惯算法库.进行持续优化.
再交给机器来对之前的大数据量的比赛进行复盘学习,学习过程中对各种模型习惯进行加权重加评分.
是否是一个比较有价值的预测东西,,,或只是个人妄想而已.哈哈 
~~~

##如该项目对你有帮助,请给一个star,谢谢!

## 公众号演示
 * 公众号: AI球探(ai00268)
 >> ![](https://oscimg.oschina.net/oscnet/up-e1c184e44f8f98c962274667d01f9670639.JPEG "go mod")
## 战绩截图
<img src="https://oscimg.oschina.net/oscnet/up-a2c999d4924ad795a582a8514f49fabe420.png" width="180px">
<img src="https://oscimg.oschina.net/oscnet/up-c71f54f3bf588fc4ffc6b6edc94919b7671.png" width="180px">

##目前程序已经完成了对很多足球相关数据荐的收集,包括且不仅限于:
* 所有的联赛信息,
* 球队信息,
* 今日比赛列表,
* 自动更新比赛结果,
* 所有亚指的数据,
* 所有的亚指的变化数据,
* 部分欧指数据,(可配置)
* 部分欧指数据的变化过程(可配置),
* 对阵双方的积分榜收集,
* 对阵双方的对战历史,
* 对阵双方的近30场战绩,
* 对阵双方的未来三场赛事.

##项目中同时也提供一些分析示例,都是个人及一些网友的忙得,且在持续更新中:
###模型说明
* A1模型  
  >> A1模型分析算法,以亚指指数为校验基准,为亚欧盘联动变化过程不符合模型设定值时,产生.
* A2模型  
  >> A2模型分析算法,以亚指指数为校验基准,比较复杂,开发实现中...
* C1模型  
  >> C1模型分析算法,以亚指指数为校验基准,纯基本面分析.
  >> 通过基本面计算出让球(BF让球),与盘口让球对比是否在合理范畴内
* E1模型  
  >> E1模型分析算法,以亚指指数为校验基准,比较复杂,暂无说明.
* E2模型  
  >> E2模型分析算法,以胜平负为校验基准,由于是单边防平双选,因此结果非对即错.
  >> 为公众号主推模型,望大家长期关注,对比历史查看.
* Q1模型  
  >> Q1模型分析算法,以亚指指数为校验基准,由网友(强)提供,逻辑简单有效,经长测表现稳定.
  >> 分析算法逻辑,对比计算对竞彩官方与***波菜的即时盘赔率得出预测结果.
  >> 目前只可在待选池中可查看,未加入到推荐列表中.


## 技术选型
* [xorm](https://github.com/go-xorm/xorm)
* [go_spider](https://github.com/hu17889/go_spider)
* [beego](https://github.com/astaxie/beego)
* [walk](https://github.com/lxn/walk)
* [go版wechat sdk](https://github.com/chanxuehong/wechat)





## 项目结构
~~~
|-- assets 素材文件夹
|   |-- common
|   |   `-- template
|   |       `-- analycontent 主要用于生成推荐文字说明内容
|   |-- leisu
|   |   `-- html 
|   `-- wechat
|       |-- html  发布公众号使用到的素材html
|       `-- img   发布公众号使用到的素材图片
|-- bin     一键打包的存放目录
|   |-- assets
|   |   |-- common
|   |   |   `-- template
|   |   |       `-- analycontent
|   |   |-- leisu
|   |   |   `-- html
|   |   `-- wechat
|   |       |-- html
|   |       `-- img
|   `-- conf
|-- conf    配置文件夹
|-- foot-api  实体类项目,用于存放模块的载体类
|   |-- common
|   |   `-- base
|   |       `-- pojo
|   `-- module
|       |-- analy
|       |   |-- pojo
|       |   `-- vo
|       |-- core
|       |   `-- pojo
|       |-- elem
|       |   `-- pojo
|       |-- match
|       |   `-- pojo
|       |-- odds
|       |   `-- pojo
|       `-- suggest
|           |-- enums
|           |-- pojo
|           `-- vo
|-- foot-core  核心库,用于提供数据库CRUD的功能,及对接第三方网络的功能
|   |-- common 通用库
|   |   |-- base 
|   |   |   |-- controller
|   |   |   `-- service
|   |   |       `-- mysql
|   |   |-- fliters
|   |   |-- routers
|   |   `-- utils
|   |-- launch
|   |-- module
|   |   |-- analy  分析模型模块
|   |   |   |-- constants
|   |   |   `-- service
|   |   |-- check
|   |   |   `-- sql
|   |   |-- core
|   |   |   `-- service
|   |   |-- elem
|   |   |   `-- service
|   |   |-- index
|   |   |   `-- controller
|   |   |-- leisu
|   |   |   |-- constants
|   |   |   |-- controller
|   |   |   |-- service
|   |   |   |-- utils
|   |   |   `-- vo
|   |   |-- match 提供赛事的相关数据库操作
|   |   |   |-- controller
|   |   |   `-- service
|   |   |-- odds  提供指数的相关数据库操作
|   |   |   `-- service
|   |   |-- spider
|   |   |   `-- constants
|   |   |-- suggest  用于获取推荐的比赛列表
|   |   |   `-- service
|   |   |-- tesou  无用
|   |   |   |-- constants
|   |   |   |-- service
|   |   |   |-- utils
|   |   |   `-- vo
|   |   `-- wechat  微信发布相关
|   |       |-- constants
|   |       |-- controller
|   |       `-- service
|   `-- test
|       `-- bson
|-- foot-gui
|   `-- handler
|-- foot-robot
|   `-- helper
|-- foot-spider 足球相关赛事指数数据的爬虫模块
|   |-- common
|   |   `-- base
|   |       `-- down
|   |-- launch
|   `-- module
|       `-- win007
|           |-- down
|           |-- proc
|           `-- vo
|-- foot-web
`-- logs

~~~
## 模块依赖
  
  | 模块名    |  依赖模块     |
  | --------    | :----:   |
  | foot-api      |无|
  | foot-core    |foot-api|
  | foot-gui     |foot-core|
  | foot-spider  |foot-core|
  | foot-web     |暂无|
   


## 实现功能
    ```
    本项目仅作娱乐研究参考所用,
    ```
## 后台数据截图
> 1. <img src="https://oscimg.oschina.net/oscnet/up-fb352eee77e897424c365a77b07269388ca.png" width="180px">
> 2. <img src="https://oscimg.oschina.net/oscnet/up-7da97167e12e1d89e455a342c0e17bbe21d.png" width="180px">
> 3. <img src="https://oscimg.oschina.net/oscnet/up-e1dc8255364a999bcc473489b163e1aa98c.png" width="180px">

## 使用教程

* 配置环境
  * 安装 go
    * 配置GOPATH
  * 环境变量
    
  | 变量名称=值    |  说明     |
  | --------    | :----:   |
  | GO111MODULE=on  |开启go mod模块支持|
  | GOPROXY=https://goproxy.cn,direct     |依赖包下载代理地址|
  | GOSUMDB=sum.golang.google.cn     |包的哈希值校验地址|
  
* 导入项目到[JetBrains GoLand](https://www.jetbrains.com/go/)并启用go mod
  >> ![](https://oscimg.oschina.net/oscnet/265bf76794ead3bac4c19a38dc4dbbe8bbb.png "go mod")
* 下载资源包
    ```
      cd ./foot-api && go mod tidy
      cd ../foot-core && go mod tidy
      cd ../foot-gui && go mod tidy
      cd ../foot-spider && go mod tidy
      cd ../foot-web && go mod tidy
    ```
* 手动创建数据库
  
  数据库名为: foot 
* 配置数据库连接
  * conf文件修改配置
    * ./conf/app.ini
     >>>> 详情看配置文件内的说明
* 同步数据库表
     * FC001DBInitApplication.go 
      
## 主要入口
  *  build_linux.bat            一键打包linux发布程序
  *  build_windows.bat          一键打包windows发布程序
  *  FC000Application.go        运行beego
  *  FC001DBInitApplication.go  数据库表同步初始化
  *  FC002AnalyApplication.go   运行结果分析    
  *  FOOT000.go                 linux入口（主要使用）
  *  FOOT000CmdApplication.go   windows入口（主要使用）
  *  FS000Application.go        运行数据爬虫

## 本地调试运行
~~~
1. 创建数据库foot
2. FC001DBInitApplication.go  同步数据库表
3. FS000Application.go 同步数据库表
4. FC002AnalyApplication 分析得出推荐列表
~~~
## 打包部署
~~~
1.运行build_linux.bat 进行打包
2.FOOT000 auto        启动
~~~
 
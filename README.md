## ‍🚀 AI足球大数据爬虫分析预测一体化项目(golang)

## 📝 项目地址
* [https://gitee.com/aoe5188/foot](https://gitee.com/aoe5188/foot)
***

![](https://img.shields.io/badge/build-go1.13.4-yeoll?style=for-the-badge&logo=appveyor)
![](https://img.shields.io/badge/build-python3.7-yellow?style=for-the-badge&logo=appveyor)
![](https://img.shields.io/badge/database-mysql5.7.28-red?style=for-the-badge&logo=appveyor)
![](https://img.shields.io/badge/ide-goland2018-blue?style=for-the-badge&logo=appveyor)
![](https://img.shields.io/badge/status-dev-green?style=for-the-badge&logo=appveyor)
***

## 🎉 项目简介 
> + ✂ 💉 ⚙ 🔨 📐 ☁ 📊 📦 📚 🌐 📈 📞
> + 👊 foot-parent 是一个集足球数据采集器,简单分析的项目.
> + 👍 程序采用golang开发,项目模块化结构清晰完整,非常容易入手并进行二次开发分析.
> + ⚙ AI球探为程序全自动处理,全程无人为参与干预足球分析预测程序.
> + ⚡️ 避免了人为分析的主观性及不稳定因素.
> + ✨ 程序根据各大指数多维度数据,结合作者多年足球分析经验,精雕细琢,
    集天地之灵气,汲日月之精华,历时七七四十九天,经Bug九九八十一个,编码而成.
> + 🎯 程序执行流程包括且不仅限于(数据自动获取-->分析学习-->自动推送发布).  
> + 😎 经近三个月的实验准确率一直能维持在一个较高的水平.
> + ☕ 同时也是一个学习golang的一个入门级项目.
***

## 🌰 项目来由
```
热衷于足球多年,之余也会去研究一下,时间久了,都会有自己的心得.
但不可能每次都那么费劲的自己人工去看盘分析,
所以结合所学,就有这个项目.
```
***

## 😎 如该项目对你有帮助,请给一个👉 star,谢谢!
## 😎 如该项目对你有帮助,请给一个👉 star,谢谢!
## 😎 如该项目对你有帮助,请给一个👉 star,谢谢!
***

## 🙋‍ 公众号演示
 * 公众号: AI球探(ai00268)
 ![](https://oscimg.oschina.net/oscnet/up-e1c184e44f8f98c962274667d01f9670639.JPEG "go mod")
***

## 👏 目前程序已经完成了对很多足球相关数据的收集,包括且不仅限于:
* 所有的联赛信息,
* 球队信息,
* 今日比赛列表,
* 自动更新比赛结果,
* 所有亚指的数据,
* 所有的亚指的变化数据,
* 所有的欧指数据,(可配置)
* 所有的欧指数据的变化过程(可配置),
* 对阵双方的积分榜收集,
* 对阵双方的对战历史,
* 对阵双方的近30场战绩,
* 对阵双方的未来三场赛事.
* 必发交易量
* 大小球指数数据
* 大小球指数数据的变化数据
***

## 🌰 技术选型
* [xorm](https://github.com/go-xorm/xorm)
* [go_spider](https://github.com/hu17889/go_spider)
* [beego](https://github.com/astaxie/beego)
* [walk](https://github.com/lxn/walk)
* [go版wechat sdk](https://github.com/chanxuehong/wechat)
***


## 🌰 模块依赖
  
  | 模块名    |  依赖模块     |  说明     |
  |:----:    | :----:   |:----   |
  | foot-api      |无| 存放各载体struct|
  | foot-core    |foot-api| 提供CRUD能力逻辑处理|
  | foot-gui     |foot-core| windows桌面控制| 
  | foot-spider  |foot-core|爬虫数据源|
  | foot-web     |暂无| 可能会用于提供API|
***

## 🌰 使用教程

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
  ![](https://oscimg.oschina.net/oscnet/265bf76794ead3bac4c19a38dc4dbbe8bbb.png "go mod")
  * 或可手动下载资源包
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
     ~~~
     详情看配置文件内的说明
     ~~~
* 同步数据库表
     * FC001DBInit.go 
      
## 🌰 主要入口
  *  build_linux.bat            一键打包linux发布程序
  *  build_windows.bat          一键打包windows发布程序
  *  FC000.go                   运行beego
  *  FC001DBInit.go             数据库表同步初始化
  *  FC002Analy.go              运行结果分析    
  *  FOOT000.go                 linux入口（主要使用）
  *  FOOT000Cmd.go              windows入口（主要使用）
  *  FS000.go                   运行数据爬虫

## 🌰 本地调试运行
~~~
(有变动需要自行查看源码)
1. 创建数据库foot
2. FC001DBInit.go  同步数据库表
3. FS000.go 运行数据爬虫
4. FC002Analy.go 分析得出推荐列表
~~~
## 🌰打包部署
* 编译打包

* 压缩编译后的文件,并上传

* 上传到服务器后,Terminal Record

***

## 🌰 项目结构
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
## 项目地址
* [https://gitee.com/aoe5188/foot](https://gitee.com/aoe5188/foot)

## 公众号演示
 * 公众号: ai00268
 >> ![](https://oscimg.oschina.net/oscnet/up-e1c184e44f8f98c962274667d01f9670639.JPEG "go mod")
## 战绩截图
<img src="https://oscimg.oschina.net/oscnet/up-a2c999d4924ad795a582a8514f49fabe420.png" width="180px">
<img src="https://oscimg.oschina.net/oscnet/up-c71f54f3bf588fc4ffc6b6edc94919b7671.png" width="180px">


## 项目己部分闭源

## 项目介绍
   >> 1. 究极足球爱好者,平时也会去体彩店支持一下国足,或是自己喜欢的球队.
   >> 2. 入了门道,就想结合一下所学所专,尝试着分析预测一下足球比赛.最近命中率感觉还可以阿.
   >> 3. 娱乐娱乐!!!如果该项目对您有帮助,请您给一个star.

## 技术选型
* [xorm](https://github.com/go-xorm/xorm)
* [go_spider](https://github.com/hu17889/go_spider)
* [beego](https://github.com/astaxie/beego)
* [walk](https://github.com/lxn/walk)
* [go版wechat sdk](https://github.com/chanxuehong/wechat)

## 项目结构
~~~
foot-parent 
├─assets    资源素材
│  ├─img
│  ├─leisu
│  └─wechat
│      └─html   公众号素材发布模板
├─conf      配置文件
├─foot-api  实体类项目
│  ├─common
│  └─module
│      ├─analy
│      ├─core
│      ├─elem
│      ├─match
│      ├─odds
│      └─suggest
├─foot-core    核心项目模块
│  ├─common
│  │  ├─base
│  │  ├─fliters beego的过滤器设置
│  │  ├─routers beego的路由设置
│  │  └─utils
│  ├─launch
│  ├─module
│  │  ├─analy   数据分析(大家可自行扩展)
│  │  ├─check   数据检查
│  │  ├─core
│  │  ├─elem    联赛指数公司模块
│  │  ├─index   http入口控制器
│  │  ├─leisu   雷速发布推荐相关       
│  │  ├─match   比赛数据模块
│  │  ├─odds    指数数据模块
│  │  ├─suggest 推荐比赛模块
│  │  ├─tesou   无用
│  │  └─wechat  微信发布推荐相关
│  └─test
├─foot-gui  未完成的gui界面
│  ├─conf
│  └─handler
├─foot-spider   数据爬虫模块
│  ├─common
│  ├─launch
│  └─module
└─foot-web  未用
    └─launch
~~~
## 模块依赖
  
  | 模块名    |  依赖模块     |
  | --------    | :----:   |
  | foot-api  |无|
  | foot-core  |foot-api|
  | foot-gui  |foot-core|
  | foot-spider  |foot-core|
  | foot-web  |暂无|
   


## 实现功能
* 当前及历史比赛数据爬取
* 当前及历史亚赔数据爬取
* 当前及历史欧赔数据爬取
* 欧亚赔的分析预测
* 数据推送到雷速发布
* 数据推送到微信公众号发布
* 定时更新微信公众号发布素材

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
      
## 运行入口
  *  build_linux.bat            一键打包linux发布程序
  *  build_windows.bat          一键打包windows发布程序
  *  FC000Application.go        运行beego
  *  FC001DBInitApplication.go  数据库表同步初始化
  *  FC002AnalyApplication.go   运行结果分析    
  *  FC003PubApplication.go     发布数据到互联网平台
  *  FOOT000.go                 linux入口（主要使用）
  *  FOOT000CmdApplication.go   windows入口（主要使用）
  *  FOOT000TestApplication.go  无用
  *  FS000Application.go        运行数据爬虫
  *  FS001AsiaModifyApplication.go  亚指数据遗漏检测,重新尝试
  *  FS001EuroIncompleteApplication.go  欧指数据遗漏检测,重新尝试
  *  FS001EuroModifyApplication.go  欧指数据单独获取



 
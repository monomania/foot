
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
     >>>> 详情看配置文件内的说明(主要修改数据库连接配置)
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
3. FS000Application.go 运行数据爬虫
4. FC002AnalyApplication 分析得出推荐列表
~~~
## 打包部署
~~~
1.运行build_linux.bat 进行打包
2.FOOT000 auto        启动
~~~

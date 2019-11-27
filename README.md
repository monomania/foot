## 项目介绍
    究极足球爱好者,平时也会去体彩店支持一下国足,或是自己喜欢的球队.
    入了门道,就想结合一下所学所专,尝试着分析预测一下足球比赛.最近命中率感觉还可以阿.
    娱乐娱乐!!!

## 技术选型
* [xorm](https://github.com/go-xorm/xorm)
* [go_spider](https://github.com/hu17889/go_spider)
* [beego](https://github.com/astaxie/beego)

## 项目结构
~~~
foot-parent
|-- foot-api            实体类模块
|   |-- common          公共工具
|   |   `-- base        基础工具
|   `-- module          业务模块
|       `-- core
|-- foot-core           后台核心模块
|   |-- common          公共工具
|   |   |-- base        基础工具
|   |   |-- log         日志工具
|   |   `-- pinyin      拼音转换工具
|   |-- conf            配置文件
|   |-- module          业务模块            
|   |   `-- core
|   `-- test
|       `-- bson
|-- foot-spider         爬虫模块
|   |-- common          公共工具              
|   |   `-- base        基础工具
|   |-- conf            配置文件
|   |-- launch          爬虫启动类
|   `-- module          业务模块
|       `-- gushiwen
`-- foot-web            http服务模块
|    |-- common         公共工具
|    |   |-- base       基础工具
|    |   |-- fliters    过滤器        
|    |   `-- routers    路由配置
|    |-- conf           配置文件
|    |-- module         业务模块
|    |   |-- core
|    |   |-- index
|    |   `-- spider
|    `-- test
~~~
## 模块依赖
  
  | 模块名    |  依赖模块     |
  | --------    | :----:   |
  | foot-api  |无|
  | foot-core  |foot-api|
  | foot-spider  |foot-api oem-core|
  | foot-web  |foot-api foot-core foot-spider|
   


## 实现功能
* 当前比赛的获取
* 当前亚赔的获取
* 当前欧赔的获取
* 历史欧赔的获取
* 欧赔的分析预测
* 亚欧赔率的分析预测


    ```
    本项目仅作娱乐研究参考所用,请勿使用到非法途径,
    ```

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
    ![](https://oscimg.oschina.net/oscnet/265bf76794ead3bac4c19a38dc4dbbe8bbb.png "go mod")
* 下载资源包
    ```
      cd ./foot-api
      go mod tidy
      cd ../foot-core
      go mod tidy
      cd ../foot-spider
      go mod tidy
      cd ../foot-web
      go mod tidy
    ```
* 手动创建数据库
  
  数据库名为: foot 
* 配置数据库连接
  
  * 各模块的conf下的 mysql.ini文件修改配置
    * ./foot-core/conf/mysql.ini
    * ./foot-spider/conf/mysql.ini
    * ./foot-web/conf/mysql.ini
* 同步数据库表

  运行入口: ./foot-core/FC000Application.go
  ```
    注意运行时: working directory需为 ****/foot-parent/foot-core 下
  ```  
  ![](https://oscimg.oschina.net/oscnet/6aeea26d87faf8cc37c7a8de61d29f6c1e5.png "working directory")
* 执行爬取数据

   运行入口: ./foot-spider/FS000Application.go
   
* 启动http服务

   运行入口: ./foot-web/FW000Application.go
  ![](https://oscimg.oschina.net/oscnet/b87398056bd5ffc0e7680f748c160bc7608.png "api")
  
* 联系作者
<table>
  <tr>
    <td>    </td>
    <td><img src="https://oscimg.oschina.net/oscnet/917bee8edddbf16a7645a56d085e887a59f.jpg"/></td> 
    <td><img src="https://oscimg.oschina.net/oscnet/aaf253aa4757b62af61036493f6fba683c2.jpg"/></td> 
    <td>    </td>
  </tr>
</table>
 
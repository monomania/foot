#go-xorm封装公共CURD类

##$ 前言
* go-xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便。

##$ 配置目标
* 封装一个公共的父级操作类,
* 通过继承可让子类带有简单的操作CRUD分页查询数据库的能力
* 可减少很多冗余,重复相似性很高的查询的代码编写
* 提高代码整洁性,可维护性

## app.ini配置
~~~
[mysql]
url=root:abc.123@tcp(localhost:3306)/foot?charset=utf8
maxIdle=10
maxConn=50
~~~

## utils工具类源码(用于读取配置信息)
~~~
package utils

import (
	"gopkg.in/ini.v1"
	"tesou.io/platform/foot-parent/foot-api/common/base"
)

var (
	//配置信息
	iniFile *ini.File
)

func init() {
	file, e := ini.Load("conf/app.ini")
	if e != nil {
		base.Log.Info("Fail to load conf/app.ini" + e.Error())
		return
	}
	iniFile = file
}

func GetSection(sectionName string) *ini.Section {
	section, e := iniFile.GetSection(sectionName)
	if e != nil {
		base.Log.Info("未找到对应的配置信息:" + sectionName + e.Error())
		return nil
	}
	return section
}

func GetSectionMap(sectionName string) map[string]string {
	section, e := iniFile.GetSection(sectionName)
	if e != nil {
		base.Log.Info("未找到对应的配置信息:" + sectionName + e.Error())
		return nil
	}
	section_map := make(map[string]string, 0)
	for _, e := range section.Keys() {
		section_map[e.Name()] = e.Value()
	}
	return section_map
}

func GetVal(sectionName string, key string) string {
	var temp_val string
	section := GetSection(sectionName)
	if nil != section {
		temp_val = section.Key(key).Value()
	}
	return temp_val;
}
~~~

## Page源码(分页结构体)
~~~
package pojo

type Page struct {
	//记录总数
	Counts int64
	//每页显示记录数
	PageSize int64
	//总页数
	TotalPage int64
	//当前页
	CurPage int64
	//页面显示开始记录数
	FirstResult int64
	//页面显示最后记录数
	LastResult int64
	//排序类型
	OrderType string
	//排序名称
	OrderName string
}

func (this *Page) Build(counts int64, pageSize int64) {
	this.Counts = counts
	this.PageSize = pageSize
	if (counts%pageSize == 0) {
		this.TotalPage = this.Counts / this.PageSize
	} else {
		this.TotalPage = this.Counts/this.PageSize + 1
	}
}

func (this *Page) GetCounts() int64 {
	return this.Counts
}

/**
 *  Counts
 *            the Counts to set
 */
func (this *Page) SetCounts(counts int64) {
	// 计算所有的页面数
	this.Counts = counts
	// this.TotalPage = (int)Math.ceil((this.Counts + this.perPageSize - 1)
	// / this.perPageSize)
	if (counts%this.PageSize == 0) {
		this.TotalPage = this.Counts / this.PageSize
	} else {
		this.TotalPage = this.Counts/this.PageSize + 1
	}
}

func (this *Page) GetPageSize() int64 {
	return this.PageSize
}

func (this *Page) SetPageSize(pageSize int64) {
	this.PageSize = pageSize
}

/**
 *  the TotalPage
 */
func (this *Page) GetTotalPage() int64 {
	if this.TotalPage < 1 {
		return 1
	}
	return this.TotalPage
}

/**
 *  TotalPage
 *            the TotalPage to set
 */
func (this *Page) SetTotalPage(totalPage int64) {
	this.TotalPage = totalPage
}

func (this *Page) GetCurPage() int64 {
	return this.CurPage
}

func (this *Page) SetCurPage(curPage int64) {
	this.CurPage = curPage
}

/**
 *  the FirstResult
 */
func (this *Page) GetFirstResult() int64 {
	temp := this.CurPage - 1
	if (temp <= 0) {
		return 0
	}
	this.FirstResult = (this.CurPage - 1) * this.PageSize
	return this.FirstResult
}

/**
 *  FirstResult
 *            the FirstResult to set
 */
func (this *Page) SetFirstResult(firstResult int64) {
	this.FirstResult = firstResult
}

/**
 *  the LastResult
 */
func (this *Page) GetLastResult() int64 {
	this.LastResult = this.FirstResult + this.PageSize
	return this.LastResult
}

/**
 *  LastResult
 *            the LastResult to set
 */
func (this *Page) SetLastResult(lastResult int64) {
	this.LastResult = lastResult
}

/**
 *  the OrderName
 */
func (this *Page) GetOrderName() string {
	return this.OrderName
}

/**
 *  OrderName
 *            the OrderName to set
 */
func (this *Page) SetOrderName(orderName string) {
	this.OrderName = orderName
}

/**
 *  the orderBy
 */
func (this *Page) getOrderType() string {
	return this.OrderType
}

/**
 *  orderBy
 *            the orderBy to set
 */
func (this *Page) SetOrderType(orderType string) {
	this.OrderType = orderType
}

/**
 *  the orderBy
 */
func (this *Page) GetOrderBy() string {
	if len(this.GetOrderName()) <= 0 {
		return ""
	}
	orderBy := " order by " + this.GetOrderName() + " " + this.getOrderType()
	return orderBy
}

~~~

## BaseService源码(公共的CRUD类)
~~~
package mysql

import (
	"container/list"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/utils"
)

type BaseService struct {
}

var (
	engine *xorm.Engine
)

func GetEngine() *xorm.Engine {
	if nil == engine {
		setEngine()
	}
	return engine
}

func ShowSQL(show bool) {
	engine := GetEngine()
	engine.ShowSQL(show)
	engine.ShowExecTime(show)

}

func setEngine() *xorm.Engine {
	url := utils.GetVal("mysql", "url")
	maxIdle, _ := strconv.Atoi(utils.GetVal("mysql", "maxIdle"))
	maxConn, _ := strconv.Atoi(utils.GetVal("mysql", "maxConn"))

	var err error
	engine, err = xorm.NewEngine("mysql", url)
	if nil != err {
		base.Log.Error("init" + err.Error())
	}

	//engine.ShowExecTime(true)
	//则会在控制台打印出生成的SQL语句
	//则会在控制台打印调试及以上的信息
	engine.ShowSQL(true)
	//engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxIdleConns(maxIdle)
	engine.SetMaxOpenConns(maxConn)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	engine.SetTableMapper(tbMapper)
	engine.SetColumnMapper(core.SameMapper{})
	/**
	当使用了Distinct,Having,GroupBy方法将不会使用缓存
	在Get或者Find时使用了Cols,Omit方法，则在开启缓存后此方法无效，系统仍旧会取出这个表中的所有字段。
	在使用Exec方法执行了方法之后，可能会导致缓存与数据库不一致的地方。因此如果启用缓存，尽量避免使用Exec。
	如果必须使用，则需要在使用了Exec之后调用ClearCache手动做缓存清除的工作。比如：
	engine.Exec("update user set name = ? where id = ?", "xlw", 1)
	engine.ClearCache(new(User))
	*/
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 999)
	//engine.SetDefaultCacher(cacher)

	return engine
}

func init() {
	//设置初始化数据库引擎
	setEngine()
}

func beforeModify(entity interface{}) {
	//当前时间
	//current_date := time.Now().Format("2006-01-02 15:04:05")
	//默认更新者
	default_user := "100000"
	//对象操作
	entity_value := reflect.ValueOf(entity).Elem()

	/*//设置更新时间
	field_ModifyDate := entity_value.FieldByName("ModifyTime")
	if field_ModifyDate.String() == "" {
		field_ModifyDate.SetString(current_date)
	}*/
	//设置更新者
	field_ModifyUser := entity_value.FieldByName("ModifyUser")
	if field_ModifyUser.String() == "" {
		field_ModifyUser.SetString(default_user)
	}

}

func beforeDelete(entity interface{}) {
	//当前时间
	//current_date := time.Now().Format("2006-01-02 15:04:05")
	//默认删除者
	default_user := "100000"
	//对象操作
	entity_value := reflect.ValueOf(entity).Elem()

	/*//设置更新时间
	field_ModifyDate := entity_value.FieldByName("ModifyTime")
	if field_ModifyDate.String() == "" {
		field_ModifyDate.SetString(current_date)
	}*/
	//设置删除者
	field_DeleteUser := entity_value.FieldByName("DeleteUser")
	if field_DeleteUser.String() == "" {
		field_DeleteUser.SetString(default_user)
	}

}

func beforeSave(entity interface{}) interface{} {
	//当前时间
	//current_date := time.Now().Format("2006-01-02 15:04:05")
	//默认创建者
	default_user := "100000"
	//对象操作
	entity_value := reflect.ValueOf(entity).Elem()

	/*	//设置创建时间  通过`xorm:"created"`配置
		field_CreateDate := entity_value.FieldByName("CreateTime")
		if field_CreateDate.String() == "" {
			field_CreateDate.SetString(current_date)
		}*/
	//设置创建者
	field_CreateUser := entity_value.FieldByName("CreateUser")
	if field_CreateUser.String() == "" {
		field_CreateUser.SetString(default_user)
	}

	beforeModify(entity)

	var id interface{}
	//设置id
	field_Id := entity_value.FieldByName("Id")
	if field_Id.String() == "" {
		//使用bson.NewObject作为主键
		id = bson.NewObjectId().Hex()
		field_Id.Set(reflect.ValueOf(id))
	}
	return id
}

func (this *BaseService) SaveOrModify(entity interface{}) {
	b, err := engine.Exist(entity)
	if nil != err {
		base.Log.Info("SaveOrModify:" + err.Error())
	}
	if b {
		this.Modify(entity)
	} else {
		this.Save(entity)
	}
}
func (this *BaseService) Save(entity interface{}) interface{} {
	id := beforeSave(entity)
	_, err := engine.InsertOne(entity)
	if nil != err {
		base.Log.Info("Save:" + err.Error())
	}
	return id
}

func (this *BaseService) SaveList(entitys []interface{}) *list.List {
	if len(entitys) <= 0 {
		return nil
	}
	list_ids := list.New()
	for _, v := range entitys {
		id := beforeSave(v)
		list_ids.PushBack(id)
	}

	_, err := engine.Insert(entitys...)
	if nil != err {
		base.Log.Info("SaveList:" + err.Error())
	}
	return list_ids
}

func (this *BaseService) Del(entity interface{}) int64 {
	beforeDelete(entity)
	entity_value := reflect.ValueOf(entity).Elem()
	id_field := entity_value.FieldByName("Id")
	i, err := engine.Id(id_field.Interface()).Delete(entity)
	if err != nil {
		base.Log.Info("Del:", err)
	}
	return i
}

func (this *BaseService) Modify(entity interface{}) int64 {
	beforeModify(entity)

	entity_value := reflect.ValueOf(entity).Elem()
	id_field := entity_value.FieldByName("Id")
	i, err := engine.Id(id_field.Interface()).AllCols().Update(entity)
	if err != nil {
		base.Log.Info("Modify:", err)
	}
	return i
}

func (this *BaseService) ModifyList(entitys []interface{}) int64 {
	if len(entitys) <= 0 {
		return 0
	}
	//i, err := engine.In("id",ids).Update(entitys)
	for _, v := range entitys {
		//entity_value := reflect.ValueOf(v).Elem()
		//id_field := entity_value.FieldByName("Id")
		this.Modify(v)
	}
	return 1
}

func (this *BaseService) Exist(entity interface{}) bool {
	//对象操作
	exist, err := engine.Exist(entity)
	if nil != err {
		base.Log.Info("ExistByName:" + err.Error())
	}
	return exist
}

func (this *BaseService) Find(entity interface{}) {
	engine.Find(entity)
}

func (this *BaseService) FindBySQL(sql string, entity interface{}) {
	engine.SQL(sql).Find(entity)
}

func (this *BaseService) FindAll(entity interface{}) {
	err := engine.Find(entity)
	if nil != err {
		base.Log.Info("FindAll: " + err.Error())
	}
}

/**
分页查询
*/
func (this *BaseService) Page(v interface{}, page *pojo.Page, dataList interface{}) error {
	tableName := engine.TableName(v)
	sql := "select t.* from " + tableName + " t where 1=1 "

	return this.PageSql(sql, page, dataList)
}

/**
分页查询
*/
func (this *BaseService) PageSql(sql string, page *pojo.Page, dataList interface{}) error {
	//声明结果变量
	var err error
	var counts int64
	//获取总记录数处理
	countSql := " select count(1) from (" + sql + ") t"
	counts, err = engine.SQL(countSql).Count()
	if nil != err {
		return err
	} else {
		page.SetCounts(counts)
	}
	//排序处理
	orderBy := page.GetOrderBy()
	if len(orderBy) > 0 {
		sql += orderBy
	}
	sql += " limit " + strconv.FormatInt(page.GetFirstResult(), 10) + "," + strconv.FormatInt(page.GetPageSize(), 10)
	err = engine.SQL(sql).Find(dataList)
	return err
}


~~~

## 使用示例
~~~
package service

import (
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/module/match/pojo"
	"tesou.io/platform/foot-parent/foot-core/common/base/service/mysql"
)

type MatchHisService struct {
    //引入公共类
	mysql.BaseService
}

func (this *MatchHisService) Exist(v *pojo.MatchHis) bool {
	has, err := mysql.GetEngine().Table("`t_match_his`").Where(" `Id` = ?  ", v.Id).Exist()
	if err != nil {
		base.Log.Error("Exist", err)
	}
	return has
}

func (this *MatchHisService) FindAll() []*pojo.MatchHis {
	dataList := make([]*pojo.MatchHis, 0)
	mysql.GetEngine().OrderBy("MatchDate").Find(&dataList)
	return dataList
}

func (this *MatchHisService) FindById(matchId string) *pojo.MatchHis {
	data := new(pojo.MatchHis)
	data.Id = matchId
	_, err := mysql.GetEngine().Get(data)
	if err != nil {
		base.Log.Error("FindById:", err)
	}
	return data
}

func (this *MatchHisService) FindBySeason(season string) []*pojo.MatchLast {
	sql_build := `
SELECT 
  la.* 
FROM
  foot.t_match_his la
WHERE 1=1
	`
	sql_build = sql_build + " AND la.MatchDate >= '" + season + "-01-01 00:00:00' AND la.MatchDate <= '" + season + "-12-31 23:59:59'"

	//结果值
	dataList := make([]*pojo.MatchLast, 0)
	//执行查询
	this.FindBySQL(sql_build, &dataList)
	return dataList
}

~~~

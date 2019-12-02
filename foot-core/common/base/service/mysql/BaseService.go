package mysql

import (
	"container/list"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strconv"
	"tesou.io/platform/foot-parent/foot-api/common/base"
	"tesou.io/platform/foot-parent/foot-api/common/base/pojo"
)

type BaseService struct {
}

var (
	mysql_conf map[string]string
	engine     *xorm.Engine
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
	url := mysql_conf["url"]
	maxIdle, _ := strconv.Atoi(mysql_conf["maxIdle"])
	maxConn, _ := strconv.Atoi(mysql_conf["maxConn"])

	var err error
	engine, err = xorm.NewEngine("mysql", url)
	if nil != err {
		base.Log.Error("init" + err.Error())
	}

	//engine.ShowExecTime(true)
	//则会在控制台打印出生成的SQL语句
	//则会在控制台打印调试及以上的信息
	//engine.ShowSQL(true)
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
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 999)
	engine.SetDefaultCacher(cacher)

	return engine
}

func init() {
	//加载配置
	loadConfig()

	//设置初始化数据库引擎
	setEngine()
}

func loadConfig() {
	configer, e := new(config.IniConfig).Parse("conf/mysql.ini")
	if e != nil {
		base.Log.Info("loadConfig加载配置文件失败:", e.Error())
		return
	}
	section, e := configer.GetSection("mysql")
	if e != nil {
		base.Log.Info("loadConfig加载配置文件失败:", e.Error())
		return
	}
	mysql_conf = section
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
	i, err := engine.Delete(entity)
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
		base.Log.Info("Exist:" + err.Error())
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
	sql += " limit " + strconv.FormatInt(page.GetFirstResult(), 10) + "," + strconv.FormatInt(page.GetLastResult(), 10)
	err = engine.SQL(sql).Find(dataList)
	return err
}

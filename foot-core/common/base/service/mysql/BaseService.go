package mysql

import (
	"container/list"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
	"strconv"
	"time"
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

func setEngine() *xorm.Engine {
	url := mysql_conf["url"]
	maxIdle, _ := strconv.Atoi(mysql_conf["maxIdle"])
	maxConn, _ := strconv.Atoi(mysql_conf["maxConn"])

	var err error
	engine, err = xorm.NewEngine("mysql", url)
	if nil != err {
		log.Println("init" + err.Error())
	}
	engine.ShowExecTime(true)
	//则会在控制台打印出生成的SQL语句
	//则会在控制台打印调试及以上的信息
	//engine.ShowSQL(true)
	//engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxIdleConns(maxIdle)
	engine.SetMaxOpenConns(maxConn)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	engine.SetTableMapper(tbMapper)
	engine.SetColumnMapper(core.SameMapper{})

	//设置缓存
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 9999)
	//engine.SetDefaultCacher(cacher)

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
		log.Println("加载配置文件失败:", e.Error())
		return
	}
	section, e := configer.GetSection("mysql")
	if e != nil {
		log.Println("加载配置文件失败:", e.Error())
		return
	}
	mysql_conf = section
}

func BeforeModify(entity interface{}) {
	//当前时间
	current_date := time.Now().Format("2006-01-02 15:04:05")
	//默认创建者
	default_user := "100000"
	//对象操作
	entity_value := reflect.ValueOf(entity).Elem()

	//设置更新时间
	field_ModifyDate := entity_value.FieldByName("ModifyTime")
	if field_ModifyDate.String() == "" {
		field_ModifyDate.SetString(current_date)
	}
	//设置更新者
	field_ModifyUser := entity_value.FieldByName("ModifyUser")
	if field_ModifyUser.String() == "" {
		field_ModifyUser.SetString(default_user)
	}

}

func BeforeSave(entity interface{}) string {
	//当前时间
	current_date := time.Now().Format("2006-01-02 15:04:05")
	//默认创建者
	default_user := "100000"
	//对象操作
	entity_value := reflect.ValueOf(entity).Elem()

	//设置创建时间
	field_CreateDate := entity_value.FieldByName("CreateTime")
	if field_CreateDate.String() == "" {
		field_CreateDate.SetString(current_date)
	}
	//设置创建者
	field_CreateUser := entity_value.FieldByName("CreateUser")
	if field_CreateUser.String() == "" {
		field_CreateUser.SetString(default_user)
	}

	BeforeModify(entity)

	var id string
	//设置id
	field_Id := entity_value.FieldByName("Id")
	if field_Id.String() == "" {
		//使用bson.NewObject作为主键
		id = bson.NewObjectId().Hex()
		field_Id.Set(reflect.ValueOf(id))
	} else {
		id = field_Id.String()
	}
	return id
}

func Save(entity interface{}) string {
	id := BeforeSave(entity)
	i, err := engine.Insert(entity)
	println(i)
	if nil != err {
		log.Println("Save" + err.Error())
	}
	return id
}

func SaveList(entitys []interface{}) *list.List {
	if len(entitys) <= 0 {
		return nil
	}
	list_ids := list.New()
	for _, v := range entitys {
		id := BeforeSave(v)
		list_ids.PushBack(id)
	}

	i, err := engine.Insert(entitys...)
	println(i)
	if nil != err {
		log.Println("SaveList" + err.Error())
	}
	return list_ids
}


func Del(entity interface{}) int64 {

	i, err := engine.Delete(entity)
	log.Println(i)
	if err != nil {
		log.Println("错误:", err)
	}
	return i
}

func Modify(id string, entity interface{}) int64 {
	BeforeModify(entity)
	i, err := engine.Id(id).Update(entity)
	log.Println(i)
	if err != nil {
		log.Println("错误:", err)
	}
	return i
}

func ModifyList(entitys []interface{}) int64 {
	if len(entitys) <= 0 {
		return 0
	}
	//i, err := engine.In("id",ids).Update(entitys)
	for i, v := range entitys {
		entity_value := reflect.ValueOf(v).Elem()
		id_field := entity_value.FieldByName("Id")
		Modify(id_field.String(), entitys[i])
	}
	return 1
	/*	log.Println(i)
		if err != nil {
			log.Println("错误:", err)
		}
		return i*/
}

func Get(entity interface{}) {
	//对象操作
	entity_value := reflect.ValueOf(entity).Elem()
	id_field := entity_value.FieldByName("Id")
	if id_field.String() == "" {
		return
	}
	_, err := engine.Get(entity)
	if nil != err {
		log.Println("Get" + err.Error())
	}
}

func FindAll(entity interface{}) {
	err := engine.Find(entity)
	if nil != err {
		log.Println("FindAllIds" + err.Error())
	}
}

func Find(entity interface{},sql string){
	engine.SQL(sql).Find(entity)
}



package mongo

import (
	"container/list"
	"github.com/astaxie/beego/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
	"time"
)


type BaseService struct {

}

var (
	mongo_conf map[string]string
	session    *mgo.Session
)

func init() {
	//加载配置
	loadConfig()
	//初始化session
	getSession(mongo_conf["hosts"])
}

func loadConfig() {
	configer, e := new(config.IniConfig).Parse("foot-core/conf/mongo.ini")
	if e != nil {
		log.Println("加载配置文件失败:", e.Error())
		return
	}
	section, e := configer.GetSection("mongo")
	if e != nil {
		log.Println("加载配置文件失败:", e.Error())
		return
	}
	mongo_conf = section
}

func getSession(hosts string) *mgo.Session {
	if session == nil {
		session, err := mgo.Dial(hosts)
		if err != nil {
			log.Println("初始化数据库Session失败:", err.Error())
		}
		session.SetMode(mgo.Monotonic, true)
		return session
	}
	return session
}

func getCollection(hosts string, dbname string, collectionName string, op func(*mgo.Collection)) {
	if hosts == "" || dbname == "" {
		if mongo_conf == nil {
			loadConfig()
		}
		hosts = mongo_conf["hosts"]
		dbname = mongo_conf["dbname"]
	}
	session := getSession(hosts)
	defer session.Close()
	database := session.DB(dbname)
	collection := database.C(collectionName)
	op(collection)
}

func getCollectionName(face interface{}) string {
	this_type := reflect.TypeOf(face)
	kind := this_type.Kind()
	if kind == reflect.Ptr {
		this_type = reflect.ValueOf(face).Elem().Type()
	}
	name := this_type.Name()
	return name
}
func Exec(face interface{}, op func(*mgo.Collection)) {
	name := getCollectionName(face)
	getCollection("", "", name, op)
}

type BaseMongoService struct {
}

func BeforeProperties(entity interface{}) bson.ObjectId {
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

	var id bson.ObjectId
	//设置id
	field_Id := entity_value.FieldByName("Id")
	if field_Id.String() == "" {
		id = bson.NewObjectId()
		field_Id.Set(reflect.ValueOf(id))
	} else {
		id = field_Id.Interface().(bson.ObjectId)
	}
	return id
}

func Save(entity interface{}) string {
	id := BeforeProperties(entity)
	Exec(entity, func(collection *mgo.Collection) {
		err := collection.Insert(entity)
		if err != nil {
			log.Println("错误:", err)
		}
	})

	return id.Hex()
}

func SaveList(entitys []interface{}) *list.List {
	if len(entitys) <= 0 {
		return nil
	}
	list_ids := list.New()
	for _, v := range entitys {
		id := BeforeProperties(v)
		list_ids.PushBack(id)
	}
	Exec(entitys[0], func(collection *mgo.Collection) {
		err := collection.Insert(entitys...)
		if err != nil {
			log.Println("错误:", err)
		}

	})
	return list_ids
}

func FindById(id string, entity interface{}) interface{} {
	if id == "" {
		return nil
	}

	objectId := bson.ObjectIdHex(id)
	Exec(entity, func(collection *mgo.Collection) {
		err := collection.FindId(objectId).One(entity)
		if err != nil {
			log.Println("错误:", err)
		}
	})

	return entity
}

func FindAll(entity interface{}) interface{} {
	entitys := make([]interface{}, 0)
	Exec(entity, func(collection *mgo.Collection) {
		err := collection.Find(nil).All(&entitys)
		if err != nil {
			log.Println("错误:", err)
		}
	})
	return entitys
}

func Del(entity interface{}) int {
	entity_value := reflect.ValueOf(entity).Elem()
	id := entity_value.FieldByName("Id").Interface()
	if nil == id {
		return 0
	}

	Exec(entity, func(collection *mgo.Collection) {
		err := collection.RemoveId(id)
		if err != nil {
			log.Println("错误:", err)
		}
	})
	return 1
}

func DelList(entitys []interface{}) int {
	if len(entitys) <= 0 {
		return 0
	}

	Exec(entitys[0], func(collection *mgo.Collection) {
		info, err := collection.RemoveAll(entitys)
		println(info)
		if err != nil {
			log.Println("错误:", err)
		}
	})
	return 1
}

func DelByIds(entity interface{}, ids []bson.ObjectId) int {
	entitys := make([]interface{}, 0)
	for _, v := range ids {
		entity_type := reflect.New(reflect.TypeOf(entity))
		elem := entity_type.Elem()
		name := elem.FieldByName("id")
		name.Set(reflect.ValueOf(v))
		entitys = append(entitys, entity_type)
	}
	return DelList(entitys)
}

func Modify(entity interface{}) int {
	entity_value := reflect.ValueOf(entity).Elem()
	id := entity_value.FieldByName("Id").Interface()
	if nil == id {
		return 0
	}

	Exec(entity, func(collection *mgo.Collection) {
		err := collection.UpdateId(bson.M{"_id": id.(bson.ObjectId)}, entity)
		if err != nil {
			log.Println("错误:", err)
		}
	})
	return 1
}

func ModifyList(entitys []interface{}) int {
	if len(entitys) <= 0 {
		return 0
	}

	ids := [100]bson.ObjectId{}
	for i, v := range entitys {
		entity_value := reflect.ValueOf(v).Elem()
		id := entity_value.FieldByName("Id").Interface()
		ids[i] = id.(bson.ObjectId)
	}
	entity := entitys[0]
	Exec(entity, func(collection *mgo.Collection) {
		info, err := collection.UpdateAll(bson.M{"_id": bson.M{"$in": ids}}, entitys)
		println(info)
		if err != nil {
			log.Println("错误:", err)
		}
	})
	return 1
}

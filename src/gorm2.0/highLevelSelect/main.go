package main

import (
	"database/sql"
	"fmt"
	. "goProject/gorm2.0/common"
)

var db = CreateDbConn()

func main() {
	FirstOrInit()
}

// SelectField 自动选择字段,其实就是定义了一个小的结构体gorm内部反射进行名称匹配赋值
func SelectField() {
	type UserModel struct {
		Id          int    `gorm:"column:Id"`
		Name        string `gorm:"column:Name"`
		PhoneNumber string `gorm:"column:PhoneNumber"`
	}
	var vm []UserModel
	db.Model(new(User)).Debug().Find(&vm)
	// SELECT "User"."Id","User"."Name","User"."PhoneNumber" FROM "User"
	fmt.Printf("%#v", vm)
}

// SubQuery 子查询
func SubQuery() {
	var users []User
	db.Where("age > (?)", db.Table("User").Select("AVG(age)")).Debug().Find(&users)
	// SELECT "User"."Id","User"."Name","User"."Age","User"."PhoneNumber","User"."Address","User"."CreateTime" FROM "User" WHERE age > (SELECT AVG(age) FROM "User")
	fmt.Println(users)

	subQuery := db.Select("AVG(age)").Where("name LIKE ?", "%wang%").Table("User")
	db.Select("AVG(age) as avgAge").Group("name").Having("AVG(age) > (?)", subQuery).Debug().Find(&users)
	// SELECT AVG(age) as avgAge FROM "User" GROUP BY "name" HAVING AVG(age) > (SELECT AVG(age) FROM "User" WHERE name LIKE '%wang%')
}

// FromSubQuery Form子查询,基于子查询结果查询
func FromSubQuery() {
	db.Table("(?) as u",
		db.Model(&User{}).Select("name", "age")).Where("age = ?", 18).Debug().Find(&User{})
	// SELECT * FROM (SELECT name,age FROM "User") as u WHERE age = 18

	subQuery1 := db.Model(&User{}).Select("name")
	subQuery2 := db.Model(&User{}).Select("age")
	db.Table("(?) as name, (?) as age", subQuery1, subQuery2).Debug().Find(&User{})
	// SELECT * FROM (SELECT name FROM "User") as name, (SELECT age FROM "User") as age
}

// WhereGroup 使用条件分组编写复杂sql
func WhereGroup() {
	db.Where(
		db.Where("address = ?", "shanghai").Where(db.Where("age = ?", 19).Or("age = ?", 20)),
	).Or(
		db.Where("address = ?", "beijing").Where("age = ?", 18),
	).Debug().Find(&User{})
	//  SELECT * FROM "User" WHERE (address = 'shanghai' AND (age = 19 OR age = 20)) OR (address = 'beijing' AND age = 18)
}

// MultiseriateIn 多列的in操作,注意：多列In在SqlServer不支持
func MultiseriateIn() {
	db.Where("(name, age) IN ?", [][]interface{}{{"w1", 18}, {"w2", 19}}).
		Debug().Find(new(User))
	// SELECT * FROM "User" WHERE (name, age) IN (('w1',18),('w2',19))
}

// NamedParameter 使用命名参数,GORM 支持 sql.NamedArg 和 map[string]interface{}{} 形式的命名参数
func NamedParameter() {
	db.Where("name = @name OR address = @address",
		sql.Named("name", "w1"),
		sql.Named("address", "shanghai")).
		Debug().Find(new([]User))
	//  SELECT * FROM "User" WHERE name = 'w1' OR address = 'shanghai'

	db.Where("name = @name OR address = @address",
		map[string]interface{}{"name": "w1", "address": "shanghai"}).
		Debug().Find(new([]User))
	// SELECT * FROM "User" WHERE name = 'w1' OR address = 'shanghai'
}

// ResultToMap 扫描结果到Map
func ResultToMap() {
	result := map[string]interface{}{}
	db.Model(&User{}).First(&result, "id = ?", 1)

	var results []map[string]interface{}
	db.Table("users").Find(&results)
}

// FirstOrInit Attrs/Assign获取第一条匹配的记录，或者根据给定的条件初始化一个实例（仅支持 struct 和 map 条件）
func FirstOrInit() {
	var user User
	// 未找到 user，则根据给定的条件初始化一条记录.注意：FirstOrInit只会赋值给结构体，而不会写入表
	db.FirstOrInit(&user, User{Name: "w7"})
	fmt.Println(user)

	// 找到name=w1的user
	var user1 User
	db.Where(User{Name: "w1"}).FirstOrInit(&user1)
	fmt.Println(user1)

	// 找到name=w1的user
	var user2 User
	db.FirstOrInit(&user2, map[string]interface{}{"name": "w1"})
	fmt.Println(user2)

	// 如果没有找到记录，可以使用包含更多的属性的结构体初始化 user，Attrs 不会被用于生成查询 SQL
	// 未找到 user，则根据给定的条件以及 Attrs 初始化 user
	db.Where(User{Name: "w7"}).Attrs(User{Age: 20}).FirstOrInit(&user)
	// user -> User{Name: "w7", Age: 20}

	// 未找到 user，则根据给定的条件以及 Attrs 初始化 user
	db.Where(User{Name: "w7"}).Attrs("age", 20).FirstOrInit(&user)
	// user -> User{Name: "w7", Age: 20}

	// 找到name=w1的user，则忽略 Attrs
	db.Where(User{Name: "w1"}).Attrs(User{Age: 20}).FirstOrInit(&user)

	// 不管是否找到记录，Assign 都会将属性赋值给 struct，但这些属性不会被用于生成查询 SQL，也不会被保存到数据库
	// 未找到 user，根据条件和 Assign 属性初始化 struct
	db.Where(User{Name: "w7"}).Assign(User{Age: 20}).FirstOrInit(&user)
	// user -> User{Name: "w1", Age: 20}

	// 找到 name=w1 的记录，依然会更新 Assign 相关的属性
	db.Where(User{Name: "w1"}).Assign(User{Age: 20}).FirstOrInit(&user)
}

// FirstOrCreate 获取第一条匹配的记录，或者根据给定的条件创建一条新纪录（仅支持 struct 和 map 条件）
func FirstOrCreate() {

}

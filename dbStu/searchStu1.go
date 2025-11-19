package dbStu

import (
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	// var users []User
	// var user User
	// db.Debug().First(&user)
	// db.Debug().Take(&user)
	// db.Debug().Last(&user)
	// db.Debug().First(&user, 2)

	// db.Debug().Find(&user,[]int{2,3})
	// res := db.Find(&user)
	// fmt.Println(user)
	// fmt.Println(res.RowsAffected, res.Error)

	// 查询的是字符串
	// db.Debug().Find(&user, "name=?", "wangwu")
	// fmt.Println(user)

	// var user2 = User{ID:2}
	// db.Debug().First(&user2)
	// fmt.Println(user2)

	// var result User
	// db.Model(User{ID: 3}).First(&result)
	// fmt.Println(result)

	// var users []User
	// res:=db.Debug().Find(&users)
	// fmt.Println(users)
	// fmt.Println(res.RowsAffected, res.Error)

	// // 不等于
	// db.Where("name <> ?", "wangwu").Find(&user)
	// fmt.Println(user)

	// IN
	// db.Debug().Where("name IN ?", []string{"wangwu", "zhaoliu"}).Find(&users)
	// fmt.Println(users)

	//LIKE
	// db.Debug().Where("name like ?", "%zhao%").Find(&users)
	// fmt.Println(users)

	//AND
	// db.Debug().Where("name like ? OR age > ?", "%zhao%", "40").Find(&users)
	// fmt.Println(users)

	// Time
	// db.Where("birthday > ?", "2025-11-17 23:07:19").Find(&users)
	// fmt.Println(users)

	// between
	// db.Where("age between ? AND ?", "18", "25").Find(&users)
	// fmt.Println(users)

	// 结构体
	// db.Debug().Where(&User{Name: "lisi", Age: 23}).Find(&users)
	// fmt.Println(users)

	// Map映射,只有使用此方法才能查询字段值为空0、空字符串''或false其他零值
	// db.Debug().Where(map[string]interface{}{"name":"lisi", "age":"23"}).Find(&users)
	// fmt.Println(users)

	// 主键切片
	// db.Where([]int{2,4}).Find(&users)
	// fmt.Println(users)

	// db.First(&users, "id = ?", "2")
	// fmt.Println(users)

	// db.Table("users").Select("name", "age").Where("name=?", "lisi").Scan(&users)
	// fmt.Println(users)

	// 更新字段
	// db.First(&user)
	// user.Name = "lisi2"
	// user.Age = 30
	// db.Debug().Save(&user)

	// 删除数据，这种方式只能删除id，其他的条件不行
	// user.ID = 1
	// db.Debug().Delete(&user)
	// 其他条件的删除
	// db.Debug().Where("name=?", "zhaoba").Delete(&user)
	// fmt.Println(user)

	// 测试生成的sql
	// stmt := db.Session(&gorm.Session{DryRun: true}).First(&user).Statement
	// fmt.Println(stmt.Vars...)
	// fmt.Println(stmt.SQL.String())

	// 获取搜索行的结果
	// var name string
	// var age int
	// rows, err := db.Model(&User{}).Where("name=?", "zhaosi").Select("name,age").Rows()

	// if err == nil {
	// 	defer rows.Close()
	// 	for rows.Next() {
	// 		rows.Scan(&name, &age)
	// 		fmt.Println(name)
	// 		fmt.Println(age)
	// 	}
	// }

}

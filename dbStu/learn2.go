package dbstu

import (
	"fmt"

	"gorm.io/gorm"
)

// 钩子函数，创建之前的动作
// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	if u.Age > 18 {
// 		return errors.New("invalid age!")
// 	}

// 	return nil
// }

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User2 struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

func Run(db *gorm.DB) {

	// 一次创建单个
	// now := time.Now()
	// user := User{Name: "lisi", Age: 23, Birthday: &now}
	// res := db.Debug().Create(&user)

	// fmt.Println(res.Error, res.RowsAffected, user.ID)

	// 一次创建多个
	// users := []*User{
	// 	{Name: "wangwu", Age: 22, Birthday: &now},
	// 	{Name: "zhaoliu", Age: 11, Birthday: &now},
	// }
	// res2 := db.Debug().Create(users)

	// fmt.Println(res2.Error, res2.RowsAffected)

	// 用指定字段创建记录
	// res := db.Debug().Select("Name", "Age", "Birthday").Create(&user)
	// fmt.Println(res.Error, res.RowsAffected, user.ID)
	// 创建记录并忽略所选的值
	// res := db.Debug().Omit("Name").Create(&user)
	// fmt.Println(res.Error, res.RowsAffected)

	// 钩子函数，跳过hooks
	// res := db.Session(&gorm.Session{SkipHooks: true}).Create(&user)
	// fmt.Println(res.Error, res.RowsAffected, user.ID)

	err := db.AutoMigrate(&User2{}, &CreditCard{})
	if err != nil{
		fmt.Println(err)
	}

}

package dbStu

// 钩子函数，创建之前的动作
// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	if u.Age > 18 {
// 		return errors.New("invalid age!")
// 	}
// 	return nil
// }

// type CreditCard struct {
// 	gorm.Model
// 	Number string
// 	UserID uint
// }

// type User2 struct {
// 	gorm.Model
// 	Name       string
// 	CreditCard CreditCard `gorm:"foreignKey:UserID"`
// }

// func Run(db *gorm.DB) {
// 	db.AutoMigrate(&User{})

// 	// 一次创建单个
// 	now := time.Now()
// 	user := User{Name: "lisi", Age: 23, Birthday: &now}
// 	res := db.Debug().Create(&user)
// 	fmt.Println(res.Error, res.RowsAffected, user.ID)

// 	// 一次创建多个
// 	users := []*User{
// 		{Name: "wangwu", Age: 22, Birthday: &now},
// 		{Name: "zhaoliu", Age: 11, Birthday: &now},
// 	}
// 	res2 := db.Debug().Create(users)
// 	fmt.Println(res2.Error, res2.RowsAffected)

// 	// 用指定字段创建记录
// 	// res := db.Debug().Select("Name", "Age", "Birthday").Create(&user)
// 	// fmt.Println(res.Error, res.RowsAffected, user.ID)
// 	// 创建记录并忽略所选的值
// 	// res := db.Debug().Omit("Name").Create(&user)
// 	// fmt.Println(res.Error, res.RowsAffected)

// 	// 钩子函数，跳过hooks
// 	// res := db.Session(&gorm.Session{SkipHooks: true}).Create(&user)
// 	// fmt.Println(res.Error, res.RowsAffected, user.ID)

// 	// 创建关联表
// 	// err := db.AutoMigrate(&User2{}, &CreditCard{})
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// db.Create(&User2{
// 	// 	Name:       "jinzhu",
// 	// 	CreditCard: CreditCard{Number: "411111111111"},
// 	// })
// 	// 跳过关联更新
// 	// user2 := User2{
// 	// 	Name:       "yyyyyyy",
// 	// 	CreditCard: CreditCard{Number: "6666666666"},
// 	// }
// 	// db.Omit("CreditCard").Create(&user2)		// 方式一
// 	// db.Omit(clause.Associations).Create(&user2)	// 方式二

// }

package associationstu

// User 拥有并属于多种 language，`user_languages` 是连接表
// type User struct {
// 	ID        uint
// 	Name      string
// 	Age       int
// 	Languages []Language `gorm:"many2many:user_languages;"`
// }

// type Language struct {
// 	ID   uint
// 	Name string
// }

// func Run(db *gorm.DB) {

// db.Debug().AutoMigrate(&User{}, &Language{})
// user := []User{
// 	{
// 		Name: "zhangsan",
// 		Languages: []Language{
// 			{Name: "ZH"},
// 			{Name: "EN"},
// 		},
// 	},
// 	{
// 		Name: "zhaoliu",
// 		Languages: []Language{
// 			{Name: "FR"},
// 			{Name: "JS"},
// 		},
// 	},
// }
// user := User{
// 	Name: "lisi",
// }
// db.Create(&user)
// user := User{}
// user.ID = 2
// // db.Omit("age").Create(&user)

// 关联删除，Languages为表名，不加s会报错
// db.Debug().Select("Languages").Delete(&User{ID: 1})

// user := User{ID: 4}
// var languages []Language
// db.Debug().Model(&user).Association("Languages").Find(&languages)
// fmt.Println(user)
// fmt.Println(languages)

// 关联计数
// count := db.Model(&user).Association("Languages").Count()
// fmt.Println(count)

// 	db.Preload("Languages").Find(&user)
// 	fmt.Println(user)

// }

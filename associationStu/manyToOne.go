package associationstu

// `User` 属于 `Company`，`CompanyID` 是外键
// type User struct {
// 	gorm.Model
// 	Name      string
// 	CompanyID int // 默认外键
// 	Company   Company
// }

// type Company struct {
// 	ID   int
// 	Name string
// }

// func Run(db *gorm.DB) {
// 	// db.AutoMigrate(&Company{}, &User{})
// 	var users []User
// 	db.Model(&User{}).Preload("Company").Find(&users, "name=?", "zhangsan")
// 	fmt.Println(users)
// }

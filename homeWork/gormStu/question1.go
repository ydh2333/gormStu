package gormstu

type User struct {
	ID        uint
	Name      string
	Age       uint
	Post      []Post
	Comment   []Comment
	PostCount uint
}

type Post struct {
	ID         uint
	Title      string
	Content    string
	Comment    []Comment
	UserID     uint
	CommCount  uint
	CommStatus string
}

type Comment struct {
	ID      uint
	Content string
	PostID  uint
	UserID  uint
}

// func Run(db *gorm.DB) {
// 题目一：
// db.Debug().AutoMigrate(&User{}, &Post{},&Comment{})
// users := []User{
// 	{
// 		Name: "张三",
// 		Age: 20,
// 	},
// 	{
// 		Name: "李四",
// 		Age: 22,
// 	},
// 	{
// 		Name: "王五",
// 		Age: 24,
// 	},
// 	{
// 		Name: "赵六",
// 		Age: 30,
// 	},
// 	{
// 		Name: "汤姆",
// 		Age: 42,
// 	},
// }
// db.Debug().Create(&users)
// posts := []Post{
// 	{
// 		Title:   "文章一",
// 		Content: "文章内容文章一，哈哈哈哈哈哈哈哈哈哈",
// 		UserID:  1,
// 	},
// 	{
// 		Title:   "文章七",
// 		Content: "文章内容文章七，哈哈哈哈哈哈哈哈哈哈",
// 		UserID:  1,
// 	},
// 	{
// 		Title:   "文章二",
// 		Content: "文章内容文章二，哈哈哈哈哈哈哈哈哈哈",
// 		UserID:  1,
// 	},
// 	{
// 		Title:   "文章三",
// 		Content: "文章内容文章三，哈哈哈哈哈哈哈哈哈哈",
// 		UserID:  1,
// 	},
// 	{
// 		Title:   "文章四",
// 		Content: "文章内容文章四，哈哈哈哈哈哈哈哈哈哈",
// 		UserID:  2,
// 	},
// 	{
// 		Title:   "文章五",
// 		Content: "文章内容文章五，哈哈哈哈哈哈哈哈哈哈",
// 		UserID:  3,
// 	},
// 	{
// 		Title:   "文章六",
// 		Content: "文章内容文章六，哈哈哈哈哈哈哈哈哈哈",
// 		UserID:  2,
// 	},
// }
// db.Debug().Create(&posts)
// comms := []Comment{
// 	{
// 		Content: "评论内容一",
// 		UserID:  1,
// 		PostID:  8,
// 	},
// 	{
// 		Content: "评论内容二",
// 		UserID:  1,
// 		PostID:  8,
// 	},
// 	{
// 		Content: "评论内容三",
// 		UserID:  1,
// 		PostID:  9,
// 	},
// 	{
// 		Content: "评论内容四",
// 		UserID:  2,
// 		PostID:  8,
// 	},
// 	{
// 		Content: "评论内容五",
// 		UserID:  2,
// 		PostID:  9,
// 	},
// 	{
// 		Content: "评论内容六",
// 		UserID:  1,
// 		PostID:  10,
// 	},
// 	{
// 		Content: "评论内容七",
// 		UserID:  1,
// 		PostID:  10,
// 	},
// 	{
// 		Content: "评论内容八",
// 		UserID:  3,
// 		PostID:  8,
// 	},
// 	{
// 		Content: "评论内容九",
// 		UserID:  4,
// 		PostID:  8,
// 	},
// 	{
// 		Content: "评论内容十",
// 		UserID:  3,
// 		PostID:  9,
// 	},
// 	{
// 		Content: "评论内容十一",
// 		UserID:  4,
// 		PostID:  10,
// 	},
// }
// db.Debug().Create(&comms)
// }

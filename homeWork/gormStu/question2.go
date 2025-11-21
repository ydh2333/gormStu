package gormstu

type PostCount struct {
	PostID    uint
	CommCount int
}

// func Run(db *gorm.DB) {
// 问题二
// 1、查询张三发布的所有文章：
// var user User
// db.Debug().Preload("Post").Preload("Comment").Where("name=?", "张三").Find(&user)
// db.Debug().Model(&User{Name: "张三"}).Preload("Post").Preload("Comment").Find(&user)
// fmt.Println(user)

// 2、查询评论数量最多的文章信息：
// PostCount存储评论数量最多的文章id和数量
// var postCount PostCount
// db.Debug().
// 	Model(&Comment{}).
// 	Select("post_id, count(*) as comm_count").
// 	Group("post_id").Order("comm_count DESC").
// 	Limit(1).Scan(&postCount)
// fmt.Println(postCount)

// var post1 Post
// 通过id获取文章信息
// db.Debug().
// 	Model(&Post{}).
// 	Where("ID=?", postCount.PostID).
// 	Preload("Comment").Find(&post1)

// fmt.Println("文章信息：", post1)
// fmt.Println("评论数量：", postCount.CommCount)
// }

package gormstu

import (
	"fmt"

	"gorm.io/gorm"
)

/*
问题1：为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
*/
func (p *Post) AfterCreate(db *gorm.DB) error {
	var coun int64
	// 获取当前user的文章数量
	if err := db.Debug().Model(&Post{}).Where("user_id=?", p.UserID).Count(&coun).Error; err != nil {
		return err
	}

	// 将文章数量更新到user表中
	res := db.Model(&User{}).Where("id=?", p.UserID).Update("post_count", coun)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("该用户更新文章数量统计失败！")
	}

	return nil
}

/*
问题2：为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func (c *Comment) AfterDelete(db *gorm.DB) error {
	var coun int64
	if err := db.Debug().Model(&Comment{}).Where("post_id=?", c.PostID).Count(&coun).Error; err != nil {
		return err
	}
	fmt.Println("-----------count:", coun)
	if coun == 0 {
		res := db.Debug().Model(&Post{}).Where("id=?", c.PostID).Update("comm_status", "无评论")
		if res.Error != nil {
			fmt.Println(res.Error)
			return res.Error
		}
		if res.RowsAffected == 0 {
			return fmt.Errorf("文章的评论状态未设置为 无评论！")
		}
	}

	return nil
}

func Run(db *gorm.DB) {
	//问题1：为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	// po := Post{
	// 	Title:   "更新用户文章数量统计标题2",
	// 	Content: "更新用户文章数量统计内容2",
	// 	UserID:  1,
	// }
	// db.Debug().Create(&po)

	// 问题2：为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	comm := Comment{ID: 10}
	db.Find(&comm)
	db.Debug().Delete(&comm)


}

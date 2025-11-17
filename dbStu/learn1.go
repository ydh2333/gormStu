package dbstu

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// 入门案例, 普通模型定义
type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

// 使用model///////////////////////////////////////////////////////////////////////////////
type Member struct {
	gorm.Model
	// 指定列名
	Name string `gorm:"column:membername"`
	Age  uint8
}

// 使用嵌入结构体（匿名字段）//////////////////////////////////////////////////////////////////
// type Author struct {
//   Name  string
//   Email string
// }

//	type Blog struct {
//	  Author
//	  ID      int
//	  Upvotes int32
//	}
//
// 使用嵌入结构体（结构体字段）/////////////////////////////////////////////////////////////////
type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
	Upvotes int32
}

// func Run(db *gorm.DB) {
// 	// db.Debug().AutoMigrate(&User{})

// 	// db.Debug().AutoMigrate(&Member{})

// 	db.Debug().AutoMigrate(&Blog{})

// }

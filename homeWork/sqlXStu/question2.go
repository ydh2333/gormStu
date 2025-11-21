package sqlxstu

// import (
// 	"fmt"

// 	"github.com/jmoiron/sqlx"
// )

// type Books struct {
// 	Title  string  `db:"title"`
// 	Author string  `db:"author"`
// 	Price  float64 `db:"price"`
// }

// func Run(db *sqlx.DB) {
	// 1、新建表
	// createTableSQL := `
	// 	CREATE TABLE books(
	// 		id int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
	// 		title VARCHAR(255),
	// 		author VARCHAR(255),
	// 		price DOUBLE
	// 	) COMMENT '';
	// `
	// _, err := db.Exec(createTableSQL)
	// if err != nil {
	// 	panic(err)
	// 	return
	// }

	// 2、造数据
	// books := []Books{
	// 	{
	// 		Title:  "西游记",
	// 		Author: "张三",
	// 		Price:  33,
	// 	},
	// 	{
	// 		Title:  "红楼梦",
	// 		Author: "李四",
	// 		Price:  50,
	// 	},
	// 	{
	// 		Title:  "三国演义",
	// 		Author: "王五",
	// 		Price:  60,
	// 	},
	// 	{
	// 		Title:  "水浒传",
	// 		Author: "赵六",
	// 		Price:  51,
	// 	},
	// }
	// for _, book := range books {
	// 	_, err := db.NamedExec("insert into books (title, author, price) values(:title, :author, :price)", &book)
	// 	if err != nil {
	// 		fmt.Printf("插入失败：%v\n", err)
	// 		return
	// 	}

	// }

	// 3、查数据
// 	var books []Books
// 	db.Select(&books, "select title, author, price from books where price > ?", 50)
// 	fmt.Println(books)

// }

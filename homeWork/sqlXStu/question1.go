package sqlxstu

type Employees struct {
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// func Run(db *sqlx.DB) {
// 1、新建表
// createTableSQL := `
// 	CREATE TABLE employees(
// 		id int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
// 		name VARCHAR(255),
// 		department VARCHAR(255),
// 		salary DOUBLE
// 	) COMMENT '';
// `
// _, err := db.Exec(createTableSQL)
// if err != nil {
// 	panic(err)
// 	return
// }

// 2、造数据
// emss := []Employees{
// 	{
// 		Name:       "张三",
// 		Department: "技术部",
// 		Salary:     6000,
// 	},
// 	{
// 		Name:       "李四",
// 		Department: "技术部",
// 		Salary:     5000,
// 	},
// 	{
// 		Name:       "王五",
// 		Department: "人事部",
// 		Salary:     3000,
// 	},
// 	{
// 		Name:       "赵六",
// 		Department: "技术部",
// 		Salary:     7000,
// 	},
// 	{
// 		Name:       "王二",
// 		Department: "人事部",
// 		Salary:     4000,
// 	},
// 	{
// 		Name:       "赵八",
// 		Department: "销售部",
// 		Salary:     10000,
// 	},
// }

// for _, ems := range emss {
// 	fmt.Println(ems)
// 	_, err := db.NamedExec("INSERT INTO employees (name,department,salary) VALUES (:name,:department,:salary)", &ems)
// 	if err != nil {
// 		fmt.Printf("插入失败：%v\n", err)
// 		return
// 	}
// }

// 要求一：
// var emss []Employees
// err := db.Select(&emss, "select name, department, salary from employees where department=?", "技术部")
// if err != nil {
// 	fmt.Printf("插入失败：%v\n", err)
// 	return
// }
// fmt.Println(ems)

// 要求二：
// var ems Employees
// 方式一：
// db.Get(&ems, "select name, department, salary from employees order by salary desc limit 1")
// 方式二：
// db.Get(&ems, "select name, department, salary from employees where salary=(select max(salary) from employees)")
// fmt.Println(ems)

// }

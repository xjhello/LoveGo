package main

// 定义一个全局对象db
//var db *sql.DB
//type user struct {
//	id   int
//	age  int
//	name string
//}
func main() {
	//initDB()
	// DSN:Data Source Name
	//dsn := "root:123456@tcp(127.0.0.1:3306)/dbname"
	//db, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close() // 注意这行代码要写在上面err判断的下面

	//queryRowDemo()

}

// 查询单条数据示例
//func queryRowDemo() {
//	sqlStr := "select id, name, age from user where id=?"
//	var u user
//	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
//	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
//	if err != nil {
//		fmt.Printf("scan failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
//}
//
//func initDB() (err error) {
//	// DSN:Data Source Name
//	dsn := "root:123@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
//	// 不会校验账号密码是否正确
//	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
//	db, err = sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	// 尝试与数据库建立连接（校验dsn是否正确）
//	err = db.Ping()
//	if err != nil {
//		return err
//	}
//	return nil
//}

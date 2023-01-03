package sqldemo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var (
	dbhost     = "119.91.115.202:port"
	dbusername = "root"
	dbpassword = "password"
	dbname     = "crawler"
)

// 注意: 一定要匿名导包.
// go get github.com/go-sql-driver/mysql

func TestSql1() {

	db, err := InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, _ := db.Query("select id,company_name from job_info limit 10")
	var ID int
	var NAME string
	for rows.Next() {
		rows.Scan(&ID, &NAME)
		fmt.Println(ID, "---", NAME)
	}
}

func TestSql2() {

	db, err := InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := db.QueryRow("select * from job_info where id = 10")
	var id int
	var company_name string
	rows.Scan(&id, &company_name)
	fmt.Println(id, "---", company_name)
}

func InitDB() (*sql.DB, error) {
	dbinfo := strings.Join([]string{dbusername, ":", dbpassword, "@tcp(", dbhost, ")/", dbname, "?charset=utf8"}, "")
	fmt.Println(dbinfo)
	dbins, err := sql.Open("mysql", dbinfo)
	if err != nil {
		fmt.Println("Open DB Error: ", err)
		return nil, err
	}
	dbins.SetConnMaxLifetime(20)
	dbins.SetConnMaxIdleTime(5)
	if err = dbins.Ping(); nil != err {
		fmt.Println("Open DB Fail, Error: ", err)
		return nil, err
	}
	fmt.Println("Connect Success!")
	return dbins, nil
}

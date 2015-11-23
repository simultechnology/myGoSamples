package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {

	fmt.Println("start");

	yamlData, err := ioutil.ReadFile("/home/charing/deploysh/config.yaml")

	if err != nil {
		panic(err)
		return
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yamlData, &m)
	if err != nil {
		panic(err)
	}

	dbInfo := m["production"].
	(map[interface {}]interface {})["db"].
	(map[interface {}]interface {})["mysql"].
	(map[interface {}]interface {})["default"]

	host := dbInfo.(map[interface {}]interface {})["host"]
	port := dbInfo.(map[interface {}]interface {})["port"]
	dbName := dbInfo.(map[interface {}]interface {})["db"]
	userName := dbInfo.(map[interface {}]interface {})["username"]
	password := dbInfo.(map[interface {}]interface {})["password"]
	//fmt.Printf("%d\n", m["b"].(map[interface {}]interface {})["c"])

	connText := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userName, password, host, port,  dbName)
	db, err := sql.Open("mysql", connText)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	res, err := db.Exec("update items set trade_status_code = 'END' where limit_date < CURDATE() and trade_status_code <> 'END'")
	count, err := res.RowsAffected()
	fmt.Println(count)
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM items where limit_date < CURDATE() and trade_status_code <> 'END'")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}


}

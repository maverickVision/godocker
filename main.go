package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"log"

	_ "github.com/godror/godror"
)

type table1 struct {
	f1 string
	f2 string
}

func main() {
	db, err := sql.Open("godror", "Admin/Marcelo@testdb#2020@testdb_high")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	var res string
	_, fnErrr := db.Exec("BEGIN :1 := pck_Test.fn_test(:2); END;", sql.Out{Dest: &res}, "é um param")

	// rows, err := db.Query("select f1, f2 from table1")
	if fnErrr != nil {
		fmt.Println("Error running query")
		fmt.Println(fnErrr)
		return
	}

	fmt.Println(res)
	// defer rows.Close()

	// for rows.Next() {
	// 	var t table1
	// 	rows.Scan(&t.f1)
	// 	fmt.Println(t)
	// }

	var refCursor driver.Rows
	_, curErr := db.Exec("BEGIN pck_Test.p_Cursor(:1); END;", sql.Out{Dest: &refCursor})
	defer refCursor.Close()

	if curErr != nil {
		fmt.Println("Error running cursor query")
		fmt.Println(curErr)
		return
	}

	values := make([]driver.Value, len(refCursor.Columns()))

	for {
		curErr = refCursor.Next(values)
		if curErr == io.EOF {
			break
		} else if curErr != nil {
			log.Fatal(curErr)
		} else {
			fmt.Println(values)
		}
	}

}

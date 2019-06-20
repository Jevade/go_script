package main

import (
	"fmt"
    "strings"
	"github.com/xuri/excelize"
)

func removeSpace(orj interface{})(interface{}){
    return strings.Replace(orj.(string)," ","",-1)
}

func Map(lens int,f func(int)interface{})([]interface{}){
    s := make([]interface{}, lens)
    for i:=0;i<lens;i++{
        s[i] = f(i)
    }
    return s
}

func main() {
	xlsx, err := excelize.OpenFile("contact.xlsx")
	if err != nil {
	    fmt.Println(err)
	    return
	}
	rows, _ := xlsx.GetRows("Sheet1")
    for _, row := range rows[1:] {
        ro := Map(len(row),func(i int)interface{}{
            return removeSpace(row[i])
        })
        fmt.Printf("insert into \"company_employee\" (\"name\",\"telephone\",\"mobilephone\",\"office\",\"department\",\"other_duty\") values(\"%v\",\"%v\",\"%v\",\"%v\",\"%v\",\"%v\"); ",ro[0],ro[1],ro[2],ro[3],ro[4],ro[5])
        fmt.Println()

	}
}

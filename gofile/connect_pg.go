package main
import (
	"database/sql"
    "fmt"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=liu dbname=scrum sslmode=verify-full"
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	    fmt.Println("failed")
        return
	}
	fmt.Println("success")


}

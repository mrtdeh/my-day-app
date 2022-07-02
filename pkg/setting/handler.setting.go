package setting

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:my_secret_password@tcp(db:3306)/app_db")
	if err != nil {
		log.Fatal(err)
	}
}

func Create(c *gin.Context) {
	userId := c.PostForm("user_id")
	data := c.PostForm("data")

	sql := fmt.Sprintf("INSERT INTO settings(user_id,data) VALUES ('%s','%s')", userId, data)
	fmt.Println(sql)
	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)

	c.JSON(http.StatusOK, gin.H{
		"id": fmt.Sprint(lastId),
	})

}

func Update(c *gin.Context) {
	userId := c.PostForm("user_id")
	data := c.PostForm("data")

	res, err := db.Exec("UPDATE settings SET data = ? WHERE user_id = ?", data, userId)

	if err != nil {
		panic(err)
	}

	id, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": fmt.Sprint(id),
	})

}

func GetById(c *gin.Context) {
	id := c.Param("id")

	fmt.Println("id : ", id)

	sql := fmt.Sprintf("SELECT * FROM settings WHERE id='%s'", id)
	res, err := db.Query(sql)

	if err != nil {
		panic(err.Error())
	}
	var set Setting
	res.Next()
	// for each row, scan the result into our tag composite object
	err = res.Scan(&set.Id, &set.UserId, &set.Data, &set.Timestamp)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// and then print out the tag's Name attribute
	// return set
	fmt.Println(set)
	c.JSON(http.StatusOK, gin.H{
		"id":        fmt.Sprint(set.Id),
		"user_id":   fmt.Sprint(set.UserId),
		"data":      fmt.Sprint(set.Data),
		"timestamp": fmt.Sprint(set.Timestamp),
	})

}

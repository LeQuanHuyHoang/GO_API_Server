package controller

import (
	database "GO_API_Server/mysql"
	"GO_API_Server/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(c *gin.Context) {
	db := database.DBcon()

	newUser := user.User{}
	newUser.ID = uuid.New()

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sql := `INSERT INTO users (Id,Name,Age,FullAddress) VALUES("` + newUser.ID.String() + `","` + newUser.Name + `","` + fmt.Sprintf("%v", newUser.Age) + `","` + newUser.FullAddress + `");`
	if _, err := db.Query(sql); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"data": newUser,
		})
	}

	defer db.Close()
}

func GetUser(c *gin.Context) {
	db := database.DBcon()

	n := c.Param("name")

	sql := `SELECT * FROM users WHERE Name = "` + n + `";`
	if rows, err := db.Query(sql); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		user := user.User{}
		for rows.Next() {
			err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.FullAddress)
			if err != nil {
				panic(err.Error())
			}
		}
		c.JSON(200, gin.H{
			"data": user,
		})

	}
	defer db.Close()
}

func UpdateUser(c *gin.Context) {
	db := database.DBcon()

	user := user.User{}
	n := c.Param("name")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sql := `UPDATE users SET Age = "` + fmt.Sprintf("%v", user.Age) + `",FullAddress = "` + user.FullAddress + `" WHERE Name = "` + n + `";`
	if _, err := db.Query(sql); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
			"data":    user,
		})

	}
	defer db.Close()
}

func DeleteUser(c *gin.Context) {
	db := database.DBcon()

	n := c.Param("name")

	sql := `DELETE FROM users WHERE name = "` + n + `"`
	if _, err := db.Query(sql); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
		})

	}
	defer db.Close()
}

package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    ID   uint   `json:"id" gorm:"primaryKey"`
    FirstName string `json:"name" gorm:"size:100;not null"`
    lastName string `json:"name" gorm:"size:100;not null"`
    Age  uint8  `json:"age"`
}

func (u User) FullName() string {
    return u.FirstName + " " + u.LastName
}

func UserDTO struct {
    ID        uint   `json:"id"`
    FullName  string `json:"full_name"`
    Age       uint8  `json:"age"`
}

func main() {
    dsn := os.Getenv("DB_DSN")
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    if err := db.AutoMigrate(&User{}); err != nil {
        log.Fatal(err)
    }

    r := gin.Default()

    r.GET("/users", func(c *gin.Context) {
        var users []User
        if err := db.Find(&users).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, users)
    })

    r.POST("/users", func(c *gin.Context) {
        var u User
        if err := c.ShouldBindJSON(&u); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := db.Create(&u).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, u)
    })

    log.Println("listening on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}


package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var items = []Item{
    {ID: "1", Name: "Item 1"},
}

func main() {
    router := gin.Default()

    router.GET("/items", GetItems)
    router.GET("/items/:id", GetItem)
    router.POST("/items", CreateItem)
    router.PUT("/items/:id", UpdateItem)
    router.DELETE("/items/:id", DeleteItem)

    router.Run(":8081")
}

func GetItems(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
    id := c.Param("id")
    for _, a := range items {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func CreateItem(c *gin.Context) {
    var newItem Item
    if err := c.BindJSON(&newItem); err != nil {
        return
    }
    items = append(items, newItem)
    c.IndentedJSON(http.StatusCreated, newItem)
}

func UpdateItem(c *gin.Context) {
    id := c.Param("id")
    var updatedItem Item
    if err := c.BindJSON(&updatedItem); err != nil {
        return
    }
    for i, a := range items {
        if a.ID == id {
            items[i] = updatedItem
            c.IndentedJSON(http.StatusOK, updatedItem)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func DeleteItem(c *gin.Context) {
    id := c.Param("id")
    for i, a := range items {
        if a.ID == id {
            items = append(items[:i], items[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "item deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
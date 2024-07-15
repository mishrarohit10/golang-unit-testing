package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/items", GetItems)
	router.GET("/items/:id", GetItem)
	router.POST("/items", CreateItem)
	router.PUT("/items/:id", UpdateItem)
	router.DELETE("/items/:id", DeleteItem)

	return router
}


func TestGetItemsWithGoConvey(t *testing.T) {
    convey.Convey("Given a running server", t, func() {
        router := setupRouter()

        convey.Convey("When GET /items is called", func() {
            req, _ := http.NewRequest("GET", "/items", nil)
            resp := httptest.NewRecorder()
            router.ServeHTTP(resp, req)

            convey.Convey("Then the status code should be 200 OK", func() {
                convey.So(resp.Code, convey.ShouldEqual, http.StatusOK)

                convey.Convey("And the response should be a list", func() {
                    var items []interface{} 
                    err := json.Unmarshal(resp.Body.Bytes(), &items)
                    convey.So(err, convey.ShouldBeNil)
                    convey.So(len(items), convey.ShouldBeGreaterThan, 0)
                })
            })
        })
    })
}

func TestGetItemWithGoConvey(t *testing.T) {
    convey.Convey("Given a running server", t, func() {
        router := setupRouter()

        convey.Convey("When GET /items/1 is called", func() {
            req, _ := http.NewRequest("GET", "/items/1", nil)
            resp := httptest.NewRecorder()
            router.ServeHTTP(resp, req)

            convey.Convey("Then the status code should be 200 OK", func() {
                convey.So(resp.Code, convey.ShouldEqual, http.StatusOK)

                convey.Convey("And the response should be the item 1", func() {
                    var item Item
                    err := json.Unmarshal(resp.Body.Bytes(), &item)
                    convey.So(err, convey.ShouldBeNil)
                    convey.So(item.ID, convey.ShouldEqual, "1")
                    convey.So(item.Name, convey.ShouldEqual, "Item 1")
                })
            })
        })

        convey.Convey("When GET /items/999 is called (non-existent ID)", func() {
            req, _ := http.NewRequest("GET", "/items/999", nil)
            resp := httptest.NewRecorder()
            router.ServeHTTP(resp, req)

            convey.Convey("Then the status code should be 404 Not Found", func() {
                convey.So(resp.Code, convey.ShouldEqual, http.StatusNotFound)

                convey.Convey("And the response should contain an error message", func() {
                    var errorResponse map[string]string
                    err := json.Unmarshal(resp.Body.Bytes(), &errorResponse)
                    convey.So(err, convey.ShouldBeNil)
                    convey.So(errorResponse["message"], convey.ShouldEqual, "Item not found")
                })
            })
        })
    })
}

func TestCreateItemWithGoConvey(t *testing.T) {
    convey.Convey("Given a running server", t, func() {
        router := setupRouter()

        convey.Convey("When POST /items with a new item is called", func() {
            newItem := Item{ID: "2", Name: "Item 2"}
            jsonValue, _ := json.Marshal(newItem)
            req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonValue))
            req.Header.Set("Content-Type", "application/json")
            resp := httptest.NewRecorder()
            router.ServeHTTP(resp, req)

            convey.Convey("Then the status code should be 201 Created", func() {
                convey.So(resp.Code, convey.ShouldEqual, http.StatusCreated)

                convey.Convey("And the response should be the new item", func() {
                    var item Item
                    err := json.Unmarshal(resp.Body.Bytes(), &item)
                    convey.So(err, convey.ShouldBeNil)
                    convey.So(item.ID, convey.ShouldEqual, newItem.ID)
                    convey.So(item.Name, convey.ShouldEqual, newItem.Name)
                })
            })
        })

        convey.Convey("When POST /items with invalid data is called", func() {
            invalidItem := map[string]string{"invalid": "data"}
            jsonValue, _ := json.Marshal(invalidItem)
            req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonValue))
            req.Header.Set("Content-Type", "application/json")
            resp := httptest.NewRecorder()
            router.ServeHTTP(resp, req)

            convey.Convey("Then the status code should be 400 Bad Request", func() {
                convey.So(resp.Code, convey.ShouldEqual, http.StatusBadRequest)

                convey.Convey("And the response should contain an error message", func() {
                    var errorResponse map[string]string
                    err := json.Unmarshal(resp.Body.Bytes(), &errorResponse)
                    convey.So(err, convey.ShouldBeNil)
                    convey.So(errorResponse["message"], convey.ShouldEqual, "Could not create new item")
                })
            })
        })
    })
}

func TestUpdateItemWithGoConvey(t *testing.T) {
	convey.Convey("Given a running server", t, func() {
		router := setupRouter()

		convey.Convey("When PUT /items/1 with an updated item is called", func() {
			updatedItem := Item{ID: "1", Name: "Item 1 Updated"}
			jsonValue, _ := json.Marshal(updatedItem)
			req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(jsonValue))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			convey.Convey("Then the status code should be 200 OK", func() {
				convey.So(resp.Code, convey.ShouldEqual, http.StatusOK)

				convey.Convey("And the response should be the updated item", func() {
					var item Item
					err := json.Unmarshal(resp.Body.Bytes(), &item)
					convey.So(err, convey.ShouldBeNil)
					convey.So(item.ID, convey.ShouldEqual, updatedItem.ID)
					convey.So(item.Name, convey.ShouldEqual, updatedItem.Name)
				})
			})
		})
	})
}

func TestDeleteItemWithGoConvey(t *testing.T) {
	convey.Convey("Given a running server", t, func() {
		router := setupRouter()

		convey.Convey("When DELETE /items/1 is called", func() {
			req, _ := http.NewRequest("DELETE", "/items/1", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			convey.Convey("Then the status code should be 200 OK", func() {
				convey.So(resp.Code, convey.ShouldEqual, http.StatusOK)
			})
		})
	})
}
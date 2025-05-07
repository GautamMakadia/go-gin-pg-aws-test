package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"botmg.com/go-server/config"
	"github.com/gin-contrib/cache"
	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Phone    int    `json:"phone"`
	Role     string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserRoute(router *gin.Engine) {
	userRoute := router.Group("/user")

	userRoute.GET("", getAllUsers)
	userRoute.GET("/:id", getUserById)
	userRoute.GET("/cache", cache.CachePage(config.CacheStore, time.Second*30, getAllUsers))

	userRoute.POST("", saveUser)
}

func getAllUsers(ctx *gin.Context) {
	result, err := config.DbPool.Query(context.Background(), "select * from users;")

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"name":  "sql error",
			"route": "/users",
		})
		return
	}

	defer result.Close()

	users, err := pgx.CollectRows(result, pgx.RowToStructByName[User])

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"msg": "error fetching row from databse",
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func getUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusNotFound)
		return
	}

	result, err := config.DbPool.Query(context.Background(), "select * from users where id = $1", id)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"name":  "sql error",
			"route": fmt.Sprintf("/user/%d", id),
		})
		return
	}

	defer result.Close()

	user, err := pgx.CollectOneRow(result, pgx.RowToStructByName[User])

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"msg": "error fetching row from databse",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func saveUser(ctx *gin.Context) {

	var userInput LoginInput

	if err := ctx.ShouldBindBodyWithJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

}

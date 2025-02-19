package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"botmg.com/go-server/config"
	"github.com/gin-contrib/cache"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Session struct {
	Id        int              `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Punch_In  pgtype.Timestamp `json:"punch_in"`
	Punch_Out pgtype.Timestamp `json:"punch_out"`
	Duration  pgtype.Interval  `json:"duration"`
	Category  string           `json:"category"`
}

func SessionRoute(router *gin.Engine) {
	sessionRouter := router.Group("/session")

	sessionRouter.GET("", cache.CachePage(config.CacheStore, time.Minute*20, getAllSession))
	// sessionRouter.GET("", getAllSession)
}


func getAllSession(ctx *gin.Context) {
	dbCtx := context.Background()
	result, err := config.DbPool.Query(dbCtx, "select * from test_session limit 50;")

	if err != nil { // Db query error
		fmt.Fprintf(os.Stderr, "Query failed to execute: %v\n", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"name": "sql error",
		})
	}

	sessions, err := pgx.CollectRows(result, pgx.RowToStructByName[Session])
	result.Conn().Close(dbCtx)

	if err != nil { // query parsing error
		fmt.Fprintf(os.Stderr, "Query failed to execute: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"name":    "parsing error",
			"message": "can't parse result to `struct Session`",
		})
	}

	ctx.JSON(http.StatusOK, sessions)

}

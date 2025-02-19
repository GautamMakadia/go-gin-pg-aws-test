package config

import (
	"time"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
	CacheStore = persistence.NewInMemoryStore(time.Hour)
)
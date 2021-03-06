package main

import (
	"github.com/gin-gonic/gin"
	session "github.com/loop-xxx/gin-session"
	"github.com/loop-xxx/gin-session/dao"
	"net/http"
)

type User struct{
	FirstName string
	SecondName string
}

func main() {
	if keeper, err := dao.DefaultRedis("192.168.1.245:6379", "", 0); err == nil {
		engine := gin.Default()
		engine.Use(session.DefaultGinSessionManager(keeper, "localhost"))

		engine.GET("/login", func(ctx *gin.Context){
			if s, exist := session.GetSession(ctx); exist {
				s.Set("name", "loop")
				_= s.SetStruct("user", User{"li", "loop"})
			}
			ctx.String(http.StatusOK, "ok")
		})


		engine.GET("/show", func(ctx *gin.Context){
			if s, ok := session.GetSession(ctx); ok {
				if name , ok := s.Get("name"); ok{
					var u User
					if err := s.GetStruct("user",&u); err == nil{
						ctx.JSON(http.StatusOK, gin.H{"name": name,  "user":u})
					}
				}
			}
		})
		_ = engine.Run(":2333")
	}
}

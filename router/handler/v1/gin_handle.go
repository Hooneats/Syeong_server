package v1

import "github.com/gin-gonic/gin"

func GinHandle(gr *gin.Engine) {
	v1 := gr.Group("app/v1")
	{
		i := v1.Group("/info")
		{
			Info(i)
		}

		u := v1.Group("/users")
		{
			Users(u)
		}
	}
}

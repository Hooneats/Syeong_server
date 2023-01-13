package middleware

import (
	"github.com/Hooneats/Syeong_server/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AccessToken 검증 및 재발급 하는 미들웨어
func ValidateJWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Request.Cookie("access-token")
		// 쿠키에서 토큰이 잘 가져와졌는지 확인
		if err != nil {
			c.JSON(401, gin.H{
				"msg": "get Access Cookie failed",
			})
			c.Abort()
			return
		}
		atValue := accessToken.Value
		// 토큰에 value가 잘 들어있는지 확인
		if atValue == "" {
			c.JSON(401, gin.H{
				"msg": "accessToken is None",
			})
			c.Abort()
			return
		}

		// 토큰 파싱
		claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(atValue, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTConfig.Salt), nil
		})
		if err != nil {
			c.JSON(401, gin.H{
				"msg": "accessToken is None",
			})
			c.Abort()
		}
		// AccessToken이 검증불가한 경우 RefreshToken을 확인하여 재발급 절차를 거친다.
		// if err != nil {
		// 	var ac *controller.AuthController
		// 	ac.VerifyToken(c)
		// }

		c.Set("userid", claims["userid"])
		c.Next()
	}
}

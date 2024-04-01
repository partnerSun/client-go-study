package middlewares

import (
	"client-go-study/utils/jwtutils"
	"client-go-study/utils/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthCheck(c *gin.Context) {
	//除了login和logout之外的所有接口，都要验证请求是否携带token，并且token是否合法
	requestUrl := c.FullPath()
	logs.Info(map[string]interface{}{"请求路径": requestUrl}, "打印路径")
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" {
		c.Next()
		return

	} else {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{
				"status":  401,
				"message": "请求未携带token，请登录后尝试",
			})
			c.Abort()
			return
		}
		//token不为空，要去验证token是否合法
		claims, err := jwtutils.ParseJwtToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  401,
				"message": "token不合法",
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}

}

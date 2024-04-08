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

func CORSMiddleware(c *gin.Context) {

	// 允许指定的跨域来源
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 允许指定的请求方法
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 允许指定的请求头部信息
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// 处理 OPTIONS 请求，返回允许的请求方法
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	// 继续处理请求
	c.Next()

}

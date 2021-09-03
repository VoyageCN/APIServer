package router

import (
	_ "APISERVER/docs"
	"APISERVER/handler/printer"
	"APISERVER/handler/user"

	"APISERVER/handler/sd"
	"APISERVER/router/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	//404 handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// Swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	g.POST("/v1/login", user.Login)       // 登陆用户
	g.POST("/v1/register", user.Register) // 注册用户

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.DELETE("/:id", user.Delete)        // 删除用户
		u.PUT("/:id", user.Update)           // 更新用户
		u.GET("", user.Get)                  // 获取登陆用户的详细信息
		u.GET("/printers", user.GetPrinters) // 获取登陆用户的绑定打印机列表
		u.PUT("/bind", user.BindPrinter)     // 绑定打印机
		u.DELETE("/unbind", user.Unbind)     // 取消绑定打印机
	}

	p := g.Group("/v1/printer")
	p.GET("", printer.Get)
	p.POST("/register", printer.Register) // 注册打印机
	p.PUT("", printer.Update)             // 更新打印机
	p.DELETE("/:id", printer.Delete)      //删除打印机
	p.POST("/connect", printer.Connect)   //连接打印机

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}

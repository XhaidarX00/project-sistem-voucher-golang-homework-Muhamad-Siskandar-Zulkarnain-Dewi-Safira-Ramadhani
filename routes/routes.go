package routes

import (
	"project-voucher-team3/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	// r.POST("/users", ctx.Ctl.User.CreateUser)
	// r.GET("/users/:id", ctx.Ctl.User.GetUser)
	// r.PUT("/users/:id", ctx.Ctl.User.UpdateUser)
	// r.DELETE("/users/:id", ctx.Ctl.User.DeleteUser)

	reedemRoutes(r, ctx)
	return r
}

func reedemRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	reedemGroup := r.Group("/reedem")

	reedemGroup.GET("/:vourcher-type", ctx.Ctl.Redeem.GetUserRedeemVoucherController)
}

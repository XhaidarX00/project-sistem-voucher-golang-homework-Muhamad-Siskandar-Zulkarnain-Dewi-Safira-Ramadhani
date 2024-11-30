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

	redeemRoutes(r, ctx)
	vourcherRouter(r, ctx)
	return r
}

func redeemRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	redeemGroup := r.Group("/redeem")

	redeemGroup.GET("/:vourcher-type", ctx.Ctl.Redeem.GetUserRedeemVoucherController)
}

func vourcherRouter(r *gin.Engine, ctx infra.ServiceContext) {
	voucherGroup := r.Group("/voucher")

	voucherGroup.GET("/validate", ctx.Ctl.Voucher.ValidateVoucherController)
}

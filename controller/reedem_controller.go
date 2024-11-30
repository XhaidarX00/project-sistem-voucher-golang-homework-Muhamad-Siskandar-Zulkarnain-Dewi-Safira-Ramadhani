package controller

import (
	"net/http"
	"project-voucher-team3/models"
	"project-voucher-team3/service"
	"project-voucher-team3/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RedeemController struct {
	service service.RedeemService
	logger  *zap.Logger
}

func NewRedeemController(service service.RedeemService, logger *zap.Logger) *RedeemController {
	return &RedeemController{service, logger}
}

func (ctrl *RedeemController) GetUserRedeemVoucherController(c *gin.Context) {
	userID := 1

	voucherType := c.Param("vourcher-type")
	if voucherType == "" {
		ctrl.logger.Error("voucher type is empty")
		utils.ResponseError(c, "EMPTY_PARAM", "voucher type is empty", http.StatusBadRequest)
		return
	}
	voucherFilter := models.Voucher{
		VoucherType: voucherType,
	}
	userRedeem, err := ctrl.service.GetAllUserRedeems(userID, voucherFilter)
	if err != nil {
		ctrl.logger.Error("Failed to get user redeem vouchers", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", err.Error(), http.StatusInternalServerError)
		return
	}

	if len(userRedeem) == 0 {
		ctrl.logger.Info("User has no redeem vouchers")
		utils.ResponseOK(c, userRedeem, "user has no redeem voucher")
		return
	}
	utils.ResponseOK(c, userRedeem, "user redeem successfully retrieved")
}

package controller

import (
	"net/http"
	"project-voucher-team3/models"
	"project-voucher-team3/service"
	"project-voucher-team3/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VoucherController struct {
	service service.VoucherService
	logger  *zap.Logger
}

func NewVoucherController(service service.VoucherService, logger *zap.Logger) *VoucherController {
	return &VoucherController{service, logger}
}

func (ctrl *VoucherController) ValidateVoucherController(c *gin.Context) {
	var voucherInput models.VoucherDTO
	if err := c.ShouldBindJSON(&voucherInput); err != nil {
		ctrl.logger.Error("Failed to bind voucher data", zap.Error(err))
		utils.ResponseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	validateResult, err := ctrl.service.ValidateVoucher(voucherInput)
	if err != nil {
		ctrl.logger.Error("Invalid voucher data", zap.Error(err))
		utils.ResponseError(c, "INVALID_DATA", err.Error(), http.StatusBadRequest)
		return
	}

	ctrl.logger.Info("Voucher data validated successfully")
	utils.ResponseOK(c, validateResult, "Voucher data validated successfully")
}

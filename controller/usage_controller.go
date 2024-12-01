package controller

import (
	"errors"
	"net/http"
	"project-voucher-team3/models"
	"project-voucher-team3/service"
	"project-voucher-team3/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UsageController struct {
	service service.UsageService
	logger  *zap.Logger
}

func NewUsageController(service service.UsageService, logger *zap.Logger) *UsageController {
	return &UsageController{service, logger}
}

func (ctrl *UsageController) CreateUsageController(c *gin.Context) {
	var usageInput models.UsageDTO

	// Bind JSON input to usageInput
	if err := c.ShouldBindJSON(&usageInput); err != nil {
		ctrl.logger.Error("Failed to bind usage input", zap.Error(err))
		utils.ResponseError(c, "BIND_ERROR", "Invalid input data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service to create usage
	if err := ctrl.service.CreateUsage(usageInput.UserID, usageInput.VoucherInput); err != nil {
		ctrl.logger.Error("Failed to create usage", zap.Error(err))

		// Distinguish between different error types if possible
		statusCode := http.StatusInternalServerError
		if errors.Is(err, errors.New("invalid voucher")) {
			statusCode = http.StatusBadRequest
		}

		utils.ResponseError(c, "CREATE_ERROR", "Could not process the usage: "+err.Error(), statusCode)
		return
	}

	// Log success and send response
	ctrl.logger.Info("Usage created successfully", zap.Int("UserID", usageInput.UserID))
	utils.ResponseOK(c, nil, "Usage created successfully")
}

func (ctrl *UsageController) GetUsageVoucherByUserIDController(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		ctrl.logger.Error("Invalid user ID", zap.Error(err))
		utils.ResponseError(c, "INVALID_ID", "Invalid user ID", http.StatusBadRequest)
		return
	}

	usageVoucher, err := ctrl.service.GetUsageVoucherByUserID(userID)
	if err != nil {
		ctrl.logger.Error("Failed to get usage voucher", zap.Error(err))
		utils.ResponseError(c, "INTERNAL_SERVER_ERROR", "Failed to get usage voucher", http.StatusInternalServerError)
		return
	}

	if len(usageVoucher) == 0 {
		ctrl.logger.Info("User has no usage vouchers")
		utils.ResponseOK(c, usageVoucher, "user has no usage voucher")
		return
	}
	utils.ResponseOK(c, usageVoucher, "user usage voucher successfully retrieved")
}

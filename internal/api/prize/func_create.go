package prize

import (
	"net/http"

	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/pkg/validation"
	"go-gin-api-simple/internal/services/prize"
)

type createRequest struct {
	Name    string  `form:"name" binding:"required"`    // 奖品名称
	Image   string  `form:"image" binding:"required"`   // 奖品图片
	Worth   float64 `form:"worth" binding:"required"`   // 奖品价值
	Content string  `form:"content" binding:"required"` // 奖品描述
}

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Create 新增奖品
// @Summary 新增奖品
// @Description 新增奖品
// @Tags API.prize
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name formData string true "奖品名称"
// @Param image formData string true "奖品图片"
// @Param worth formData number true "奖品价值"
// @Param content formData string true "奖品描述"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/prize [post]
// @Security LoginToken
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		createData := new(prize.CreatePrizeData)
		createData.Name = req.Name
		createData.Image = req.Image
		createData.Worth = req.Worth
		createData.Content = req.Content

		id, err := h.prizeService.Create(ctx, createData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.PrizeCreateError,
				code.Text(code.PrizeCreateError)).WithError(err),
			)
			return
		}

		res.Id = id
		ctx.Payload(res)
	}
}

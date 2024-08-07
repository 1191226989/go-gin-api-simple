package prize

import (
	"net/http"

	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/pkg/validation"
	"go-gin-api-simple/internal/services/prize"
)

type modifyRequest struct {
	Id      string  `form:"id" binding:"required"`      // 主键ID
	Name    string  `form:"name" binding:"required"`    // 奖品名称
	Image   string  `form:"image" binding:"required"`   // 奖品图片
	Worth   float64 `form:"worth" binding:"required"`   // 奖品价值
	Content string  `form:"content" binding:"required"` // 奖品描述
	IsUsed  int32   `form:"is_used" binding:"required"` // 是否启用
}

type modifyResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Modify 编辑奖品
// @Summary 编辑奖品
// @Description 编辑奖品
// @Tags API.prize
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "奖品id"
// @Param name formData string true "奖品名称"
// @Param image formData string true "奖品图片"
// @Param worth formData number true "奖品价值"
// @Param content formData string true "奖品描述"
// @Param is_used formData number true "是否启用"
// @Success 200 {object} modifyResponse
// @Failure 400 {object} code.Failure
// @Router /api/prize/{id} [post]
// @Security LoginToken
func (h *handler) Modify() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(modifyRequest)
		res := new(modifyResponse)

		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		id := int32(ids[0])

		modifyData := new(prize.ModifyPrizeData)
		modifyData.Name = req.Name
		modifyData.Image = req.Image
		modifyData.Worth = req.Worth
		modifyData.Content = req.Content
		modifyData.IsUsed = req.IsUsed

		err = h.prizeService.Modify(ctx, id, modifyData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.PrizeModifyError,
				code.Text(code.PrizeModifyError)).WithError(err),
			)
			return
		}

		res.Id = id
		ctx.Payload(res)
	}
}

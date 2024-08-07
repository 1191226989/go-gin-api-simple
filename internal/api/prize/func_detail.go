package prize

import (
	"errors"
	"net/http"

	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/pkg/validation"
	"go-gin-api-simple/internal/services/prize"
	"go-gin-api-simple/pkg/timeutil"

	"github.com/spf13/cast"
)

type detailRequest struct {
	Id string `uri:"id" binding:"required"` // 主键ID
}

type detailResponse struct {
	Id        int     `json:"id"`         // ID
	HashID    string  `json:"hashid"`     // hashid
	Name      string  `json:"name"`       // 奖品名称
	Image     string  `json:"image"`      // 奖品图片
	Worth     float64 `json:"worth"`      // 奖品价值
	Content   string  `json:"content"`    // 奖品描述
	IsUsed    int     `json:"is_used"`    // 是否启用 1:是 0:否
	CreatedAt string  `json:"created_at"` // 创建时间
	UpdatedAt string  `json:"updated_at"` // 更新时间
}

// Detail 奖品详情
// @Summary 奖品详情
// @Description 奖品详情
// @Tags API.prize
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/prize/{id} [get]
// @Security LoginToken
func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)

		if err := ctx.ShouldBindURI(req); err != nil {
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

		detailData := new(prize.SearchOneData)
		detailData.Id = id

		prizeData, err := h.prizeService.Detail(ctx, detailData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.PrizeDetailError,
				code.Text(code.PrizeDetailError)).WithError(err),
			)
			return
		}

		if prizeData == nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.PrizeDetailError,
				code.Text(code.PrizeDetailError)).WithError(errors.New("该奖品数据未找到")),
			)
			return
		}

		res.Id = cast.ToInt(id)
		res.HashID = req.Id
		res.Name = prizeData.Name
		res.Image = prizeData.Image
		res.Worth = prizeData.Worth
		res.Content = prizeData.Content
		res.IsUsed = cast.ToInt(prizeData.IsUsed)
		res.CreatedAt = prizeData.CreatedAt.Format(timeutil.CSTLayout)
		res.UpdatedAt = prizeData.UpdatedAt.Format(timeutil.CSTLayout)

		ctx.Payload(res)
	}
}

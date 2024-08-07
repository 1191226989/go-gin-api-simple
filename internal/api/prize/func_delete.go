package prize

import (
	"net/http"

	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
)

type deleteRequest struct {
	Id string `uri:"id"` // HashID
}

type deleteResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Delete 删除奖品
// @Summary 删除奖品
// @Description 删除奖品
// @Tags API.prize
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /api/prize/{id} [delete]
// @Security LoginToken
func (h *handler) Delete() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(deleteRequest)
		res := new(deleteResponse)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
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

		err = h.prizeService.Delete(ctx, id)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.PrizeDeleteError,
				code.Text(code.PrizeDeleteError)).WithError(err),
			)
			return
		}

		res.Id = id
		ctx.Payload(res)
	}
}

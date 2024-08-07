package user

import (
	"net/http"

	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/services/user"
)

type detailResponse struct {
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Mobile   string `json:"mobile"`   // 手机号
}

// Detail 用户详情
// @Summary 用户详情
// @Description 用户详情
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/info [get]
// @Security LoginToken
func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {
		res := new(detailResponse)

		searchOneData := new(user.SearchOneData)
		searchOneData.Id = ctx.SessionUserInfo().UserID
		searchOneData.IsUsed = 1

		info, err := h.userService.Detail(ctx, searchOneData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserDetailError,
				code.Text(code.UserDetailError)).WithError(err),
			)
			return
		}

		res.Username = info.Username
		res.Nickname = info.Nickname
		res.Mobile = info.Mobile
		ctx.Payload(res)
	}
}

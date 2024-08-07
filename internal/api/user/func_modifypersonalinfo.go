package user

import (
	"net/http"

	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/services/user"
)

type modifyPersonalInfoRequest struct {
	Nickname string `form:"nickname"` // 昵称
	Mobile   string `form:"mobile"`   // 手机号
}

type modifyPersonalInfoResponse struct {
	Username string `json:"username"` // 用户账号
}

// ModifyPersonalInfo 修改个人信息
// @Summary 修改个人信息
// @Description 修改个人信息
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param nickname formData string true "昵称"
// @Param mobile formData string true "手机号"
// @Success 200 {object} modifyPersonalInfoResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/modify_personal_info [patch]
// @Security LoginToken
func (h *handler) ModifyPersonalInfo() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(modifyPersonalInfoRequest)
		res := new(modifyPersonalInfoResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		modifyData := new(user.ModifyData)
		modifyData.Nickname = req.Nickname
		modifyData.Mobile = req.Mobile

		if err := h.userService.ModifyPersonalInfo(ctx, ctx.SessionUserInfo().UserID, modifyData); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserModifyPersonalInfoError,
				code.Text(code.UserModifyPersonalInfoError)).WithError(err),
			)
			return
		}

		res.Username = ctx.SessionUserInfo().UserName
		ctx.Payload(res)
	}
}

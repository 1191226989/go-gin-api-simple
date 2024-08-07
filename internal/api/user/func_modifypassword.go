package user

import (
	"net/http"

	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/pkg/password"
	"go-gin-api-simple/internal/services/user"
)

type modifyPasswordRequest struct {
	OldPassword string `form:"old_password"` // 旧密码
	NewPassword string `form:"new_password"` // 新密码
}

type modifyPasswordResponse struct {
	Username string `json:"username"` // 用户账号
}

// ModifyPassword 修改密码
// @Summary 修改密码
// @Description 修改密码
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param old_password formData string true "旧密码（md5）"
// @Param new_password formData string true "新密码（md5）"
// @Success 200 {object} modifyPasswordResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/modify_password [patch]
// @Security LoginToken
func (h *handler) ModifyPassword() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(modifyPasswordRequest)
		res := new(modifyPasswordResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		searchOneData := new(user.SearchOneData)
		searchOneData.Id = ctx.SessionUserInfo().UserID
		searchOneData.Password = password.GeneratePassword(req.OldPassword)
		searchOneData.IsUsed = 1

		info, err := h.userService.Detail(ctx, searchOneData)
		if err != nil || info == nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserOldPasswordError,
				code.Text(code.UserOldPasswordError)).WithError(err),
			)
			return
		}

		if err := h.userService.ModifyPassword(ctx, ctx.SessionUserInfo().UserID, req.NewPassword); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserModifyPasswordError,
				code.Text(code.UserModifyPasswordError)).WithError(err),
			)
			return
		}

		res.Username = ctx.SessionUserInfo().UserName
		ctx.Payload(res)
	}
}

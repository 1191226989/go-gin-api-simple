package user

import (
	"net/http"

	"go-gin-api-simple/configs"
	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/repository/redis"
	"go-gin-api-simple/pkg/errors"
)

type logoutResponse struct {
	Username string `json:"username"` // 用户账号
}

// Logout 用户登出
// @Summary 用户登出
// @Description 用户登出
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} logoutResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/logout [post]
// @Security LoginToken
func (h *handler) Logout() core.HandlerFunc {
	return func(c core.Context) {
		res := new(logoutResponse)
		res.Username = c.SessionUserInfo().UserName

		if !h.cache.Del(configs.RedisKeyPrefixLoginUser+c.GetHeader(configs.HeaderLoginToken), redis.WithTrace(c.Trace())) {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserLogOutError,
				code.Text(code.UserLogOutError)).WithError(errors.New("cache del err")),
			)
			return
		}

		c.Payload(res)
	}
}

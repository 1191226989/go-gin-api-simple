package user

import (
	"net/http"

	"go-gin-api-simple/configs"
	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/pkg/password"
	"go-gin-api-simple/internal/proposal"
	"go-gin-api-simple/internal/repository/redis"
	"go-gin-api-simple/internal/services/user"
	"go-gin-api-simple/pkg/errors"
)

type loginRequest struct {
	Username string `form:"username"` // 用户名
	Password string `form:"password"` // 密码
}

type loginResponse struct {
	Token string `json:"token"` // 用户身份标识
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "MD5后的密码"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /api/login [post]
func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		req := new(loginRequest)
		res := new(loginResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		searchOneData := new(user.SearchOneData)
		searchOneData.Username = req.Username
		searchOneData.Password = password.GeneratePassword(req.Password)
		searchOneData.IsUsed = 1

		info, err := h.userService.Detail(c, searchOneData)
		if err != nil || info == nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserLoginError,
				code.Text(code.UserLoginError)).WithError(err),
			)
			return
		}

		if info == nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserLoginError,
				code.Text(code.UserLoginError)).WithError(errors.New("未查询出符合条件的用户")),
			)
			return
		}

		token := password.GenerateLoginToken(info.Id)

		// 用户信息
		sessionUserInfo := &proposal.SessionUserInfo{
			UserID:   info.Id,
			UserName: info.Username,
		}

		// 将用户信息记录到 Redis 中
		err = h.cache.Set(configs.RedisKeyPrefixLoginUser+token, string(sessionUserInfo.Marshal()), configs.LoginSessionTTL, redis.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserLoginError,
				code.Text(code.UserLoginError)).WithError(err),
			)
			return
		}

		res.Token = token
		c.Payload(res)
	}
}

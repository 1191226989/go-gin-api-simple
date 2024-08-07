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

type signupRequest struct {
	Username string `form:"username"` // 用户名
	Password string `form:"password"` // 密码
}

type signupResponse struct {
	Token string `json:"token"` // 用户身份标识
}

// Signup 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "MD5后的密码"
// @Success 200 {object} signupResponse
// @Failure 400 {object} code.Failure
// @Router /api/signup [post]
func (h *handler) Signup() core.HandlerFunc {
	return func(c core.Context) {
		req := new(signupRequest)
		res := new(signupResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		// 用户名是否已经注册
		searchOneData := new(user.SearchOneData)
		searchOneData.Username = req.Username
		searchOneData.IsUsed = 1

		info, err := h.userService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserSignupError,
				code.Text(code.UserSignupError)).WithError(err),
			)
			return
		}

		if info != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserSignupError,
				code.Text(code.UserSignupError)).WithError(errors.New("该用户名已经注册")),
			)
			return
		}

		// 用户保存数据库
		createData := new(user.CreateUserData)
		createData.Username = req.Username
		createData.Password = password.GeneratePassword(req.Password)
		createData.Nickname = ""
		createData.Mobile = ""
		createData.IsUsed = 1
		createData.IsDeleted = -1

		id, err := h.userService.Create(c, createData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UserSignupError,
				code.Text(code.UserSignupError)).WithError(err),
			)
			return
		}

		// 注册成功则登录
		token := password.GenerateLoginToken(id)

		// 用户信息
		sessionUserInfo := &proposal.SessionUserInfo{
			UserID:   id,
			UserName: req.Username,
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

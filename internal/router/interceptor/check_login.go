package interceptor

import (
	"encoding/json"
	"net/http"

	"go-gin-api-simple/configs"
	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/proposal"
	"go-gin-api-simple/internal/repository/redis"
	"go-gin-api-simple/pkg/errors"
)

func (i *interceptor) CheckLogin(ctx core.Context) (sessionUserInfo proposal.SessionUserInfo, err core.BusinessError) {
	token := ctx.GetHeader(configs.HeaderLoginToken)
	if token == "" {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(errors.New("Header 中缺少 Token 参数"))

		return
	}

	if !i.cache.Exists(configs.RedisKeyPrefixLoginUser + token) {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(errors.New("请先登录"))

		return
	}

	cacheData, cacheErr := i.cache.Get(configs.RedisKeyPrefixLoginUser+token, redis.WithTrace(ctx.Trace()))
	if cacheErr != nil {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(cacheErr)

		return
	}

	jsonErr := json.Unmarshal([]byte(cacheData), &sessionUserInfo)
	if jsonErr != nil {
		core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(jsonErr)

		return
	}

	return
}

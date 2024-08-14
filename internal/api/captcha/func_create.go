package captcha

import (
	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/captcha"
	"go-gin-api-simple/internal/pkg/core"
	"net/http"

	"github.com/mojocn/base64Captcha"
)

type createRequest struct{}

type createResponse struct {
	CaptchaId    string `json:"captcha_id"`   // 验证码id
	Base64String string `json:"base64string"` // 验证码图片base64字符串
}

// Create 生成验证码id和图片
// @Summary 生成验证码id和图片
// @Description 生成验证码id和图片
// @Tags API.captcha
// @Produce json
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/captcha [get]
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		store := captcha.NewStoreRedis(h.cache)
		captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, store)
		captchaId, base64, _, err := captcha.Generate()
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CaptchaGenerateError,
				code.Text(code.CaptchaGenerateError)).WithError(err),
			)
			return
		}

		res := new(createResponse)
		res.CaptchaId = captchaId
		res.Base64String = base64

		ctx.Payload(res)
	}
}

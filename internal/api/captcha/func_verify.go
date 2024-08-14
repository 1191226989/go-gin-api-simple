package captcha

import (
	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/captcha"
	"go-gin-api-simple/internal/pkg/core"
	"net/http"

	"github.com/mojocn/base64Captcha"
)

type verifyRequest struct {
	CaptchaId     string `form:"captcha_id"`     // 验证码id
	CaptchaAnswer string `form:"captcha_answer"` // 验证码答案
}

type verifyResponse struct {
	VerifyResult bool `json:"verify_result"` // 验证结果
}

// Verify 验证码校验
// @Summary 验证码校验
// @Description 验证码校验
// @Tags API.captcha
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param captcha_id formData string true "验证码id"
// @Param captcha_answer formData string true "验证码答案"
// @Success 200 {object} verifyResponse
// @Failure 400 {object} code.Failure
// @Router /api/captcha [post]
func (h *handler) Verify() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(verifyRequest)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		store := captcha.NewStoreRedis(h.cache)
		captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, store)
		match := captcha.Verify(req.CaptchaId, req.CaptchaAnswer, true)

		res := new(verifyResponse)
		res.VerifyResult = match

		ctx.Payload(res)
	}
}

package captcha

import (
	"net/http"

	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/captcha"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/pkg/validation"

	"github.com/mojocn/base64Captcha"
)

type verifyRequest struct {
	CaptchaId     string `json:"captcha_id" binding:"required" msg:"验证码id不能为空"`     // 验证码id
	CaptchaAnswer string `json:"captcha_answer" binding:"required" msg:"验证码答案不能为空"` // 验证码答案
}

type verifyResponse struct {
	VerifyResult bool `json:"verify_result"` // 验证结果
}

// Verify 验证码校验
// @Summary 验证码校验
// @Description 验证码校验
// @Tags API.captcha
// @Accept json
// @Produce json
// @Param Request body verifyRequest true "请求参数"
// @Success 200 {object} verifyResponse
// @Failure 400 {object} code.Failure
// @Router /api/captcha [post]
func (h *handler) Verify() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(verifyRequest)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.CustomErrorMessage(err, req)).WithError(err),
			// code.Text(code.ParamBindError)).WithError(err),
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

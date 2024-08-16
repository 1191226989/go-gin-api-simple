package captcha

import (
	"fmt"
	"net/http"

	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/captcha"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/pkg/validation"

	"github.com/mojocn/base64Captcha"
)

type createRequest struct {
	Height int32 `uri:"height" binding:"required,max=1000,min=10" label:"高度"` // 验证码高度
	Width  int32 `uri:"width" binding:"required,max=1000,min=10" label:"宽度"`  // 验证码宽度
	Length int32 `uri:"length" binding:"required,max=10,min=1" label:"长度"`    // 验证码长度
}

type createResponse struct {
	CaptchaId    string `json:"captcha_id"`   // 验证码id
	Base64String string `json:"base64string"` // 验证码图片base64字符串
}

// Create 生成验证码id和图片
// @Summary 生成验证码id和图片
// @Description 生成验证码id和图片
// @Tags API.captcha
// @Accept json
// @Produce json
// @Param height path integer true "验证码高度"
// @Param width path integer true "验证码宽度"
// @Param length path integer true "验证码长度"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/captcha/{height}/{width}/{length} [get]
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(createRequest)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				fmt.Sprintf("%s: %s", code.Text(code.ParamBindError), validation.Error(err))).WithError(err),
			// code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		driver := base64Captcha.DefaultDriverDigit
		if req.Height > 0 {
			driver.Height = int(req.Height)
		}
		if req.Width > 0 {
			driver.Width = int(req.Width)
		}
		if req.Length > 0 {
			driver.Length = int(req.Length)
		}

		store := captcha.NewStoreRedis(h.cache)

		captcha := base64Captcha.NewCaptcha(driver, store)
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

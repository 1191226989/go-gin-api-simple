package alert

import (
	"go-gin-api-simple/configs"
	"go-gin-api-simple/internal/proposal"
	"go-gin-api-simple/pkg/errors"
	"go-gin-api-simple/pkg/mail"

	"go.uber.org/zap"
)

// NotifyHandler 告警通知
func NotifyHandler(logger *zap.Logger) func(msg *proposal.AlertMessage) {
	if logger == nil {
		panic("logger required")
	}

	return func(msg *proposal.AlertMessage) {
		cfg := configs.Get().Mail
		if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Pass == "" || cfg.To == "" {
			logger.Error("Mail config error")
			return
		}

		subject, body, err := newHTMLEmail(
			msg.Method,
			msg.HOST,
			msg.URI,
			msg.TraceID,
			msg.ErrorMessage,
			msg.ErrorStack,
		)
		if err != nil {
			logger.Error("email template error", zap.Error(err))
			return
		}

		options := &mail.Options{
			MailHost: cfg.Host,
			MailPort: cfg.Port,
			MailUser: cfg.User,
			MailPass: cfg.Pass,
			MailTo:   cfg.To,
			Subject:  subject,
			Body:     body,
		}
		if err := mail.Send(options); err != nil {
			logger.Error("发送告警通知邮件失败", zap.Error(errors.WithStack(err)))
		}

		return
	}
}

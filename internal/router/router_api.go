package router

import (
	"go-gin-api-simple/internal/api/captcha"
	"go-gin-api-simple/internal/api/config"
	"go-gin-api-simple/internal/api/helper"
	"go-gin-api-simple/internal/api/prize"
	"go-gin-api-simple/internal/api/tool"
	"go-gin-api-simple/internal/api/user"
	"go-gin-api-simple/internal/pkg/core"
)

func setApiRouter(r *resource) {
	// helper
	helperHandler := helper.New(r.logger, r.db, r.cache)

	helpers := r.mux.Group("/helper")
	{
		helpers.GET("/md5/:str", helperHandler.Md5())
	}

	// user
	userHandler := user.New(r.logger, r.db, r.cache)

	// 无需登录验证
	nologin := r.mux.Group("/api")
	{
		nologin.POST("/signup", userHandler.Signup())
		nologin.POST("/login", userHandler.Login())

		// captcha
		captchaHandler := captcha.New(r.logger, r.cache)
		nologin.POST("/captcha", captchaHandler.Verify())
		nologin.GET("/captcha", captchaHandler.Create())
	}

	// 需要登录验证
	api := r.mux.Group("/api", core.WrapAuthHandler(r.interceptors.CheckLogin))
	{
		// tool
		toolHandler := tool.New(r.logger, r.db, r.cache)
		api.GET("/tool/hashids/encode/:id", core.AliasForRecordMetrics("/api/tool/hashids/encode"), toolHandler.HashIdsEncode())
		api.GET("/tool/hashids/decode/:id", core.AliasForRecordMetrics("/api/tool/hashids/decode"), toolHandler.HashIdsDecode())
		api.POST("/tool/cache/search", toolHandler.SearchCache())
		api.PATCH("/tool/cache/clear", toolHandler.ClearCache())
		api.GET("/tool/data/dbs", toolHandler.Dbs())
		api.POST("/tool/data/tables", toolHandler.Tables())
		api.POST("/tool/data/mysql", toolHandler.SearchMySQL())
		api.POST("/tool/send_message", toolHandler.SendMessage())

		// config
		configHandler := config.New(r.logger, r.db, r.cache)
		api.PATCH("/config/email", configHandler.Email())

		// user
		api.POST("/user/logout", userHandler.Logout())
		api.PATCH("/user/modify_password", userHandler.ModifyPassword())
		api.GET("/user/info", userHandler.Detail())
		api.PATCH("/user/modify_personal_info", userHandler.ModifyPersonalInfo())

		// prize
		prizeHandler := prize.New(r.logger, r.db, r.cache)
		api.POST("/prize", prizeHandler.Create())
		api.GET("/prize", prizeHandler.List())
		api.GET("/prize/:id", core.AliasForRecordMetrics("/api/prize/detail"), prizeHandler.Detail())
		api.POST("/prize/:id", core.AliasForRecordMetrics("/api/prize/modify"), prizeHandler.Modify())
		api.PATCH("/prize/used", prizeHandler.UpdateUsed())
		api.DELETE("/prize/:id", core.AliasForRecordMetrics("/api/prize"), prizeHandler.Delete())

	}
}

package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"self-wechat/config"
	_ "self-wechat/dao"
	"self-wechat/handler"
)

func Router(r *gin.Engine) {
	wxServerCheck := r.Group("/api/v2")
	{
		wxServerCheck.GET("/check", handler.HandlerSignatureHandler)
		wxServerCheck.POST("/check", handler.HandMessagesHandler)

	}

	cc := r.Group("/api/v1")

	// 微信相关
	wechatGroup := cc.Group("/wechat")
	{
		wechatGroup.GET("/accesstokne_by_code", handler.WechatAuthHandler)                    // code 换access_token或openid
		wechatGroup.GET("/download_media_by_audioid", handler.WechatDownloadMediaDataHandler) // 根据media_id下载音频数据
		wechatGroup.POST("/send_template_info", handler.WechatSendTemplateInfoHandler)        // 发送模板消息
		wechatGroup.GET("/jsconfig", handler.WechatGetJSConfigHandler)                        // 获取jsconfig
		wechatGroup.POST("/receive_msg", handler.WechatReceiveMsgHandler)
	}
	err := r.Run(fmt.Sprintf(":%d", config.Config.Cfg.Port))
	if err != nil {
		panic(err)
	}
}

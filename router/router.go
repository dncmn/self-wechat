package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"self-wechat/config"
	_ "self-wechat/dao"
	"self-wechat/handler"
)

func Router(r *gin.Engine) {
	//r.StaticFS("/compoments/runtime/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	wxServerCheck := r.Group("/api/v2")
	{
		wxServerCheck.GET("/", handler.HandlerSignatureHandler)

	}

	cc := r.Group("/api/v1", handler.VerifyToken)

	// 微信相关
	wechatGroup := cc.Group("/wechat")
	{
		wechatGroup.GET("/accesstokne_by_code", handler.WechatAuthHandler)                    // code 换access_token或openid
		wechatGroup.GET("/download_media_by_audioid", handler.WechatDownloadMediaDataHandler) // 根据media_id下载音频数据
		wechatGroup.POST("/send_template_info", handler.WechatSendTemplateInfoHandler)        // 发送模板消息
		wechatGroup.GET("/jsconfig", handler.WechatGetJSConfigHandler)                        // 获取jsconfig
	}
	err := r.Run(fmt.Sprintf(":%d", config.Config.Cfg.Port))
	if err != nil {
		panic(err)
	}
}

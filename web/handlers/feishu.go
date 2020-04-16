package handlers

import (
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller/impl"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FeishuCallback struct {
	wechatMessageController controller.WechatMessageController
}

func NewFeishuCallback() *FeishuCallback {
	wechatMessageController := impl.NewWechatMessageController()
	return &FeishuCallback{wechatMessageController: wechatMessageController}
}


func (p *FeishuCallback) Callback(c *gin.Context) {
	data := make(map[string]interface{})
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("error: %+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if msgType, ok := data["type"]; ok {
		if msgType == "url_verification" {
			if receivedToken, ok := data["token"]; ok {
				if receivedToken != configs.DefaultConfig.LarkBot.Token {
					c.JSON(200, gin.H{})
					c.Abort()
					return
				} else {
					if challenge, ok := data["challenge"]; ok {
						c.JSON(http.StatusOK, gin.H{"challenge": challenge})
						c.Abort()
						return
					}
				}
			}
		}
	}
	log.Printf("%+v\n", data)
	log.Printf("%+v\n", data["action"])
	if action, ok := data["action"]; ok {
		if valueMap, ok := action.(map[string]interface{}); ok {
			wxid, _ := valueMap["wx_id"]
			roomID, _ := valueMap["room_id"]
			createdAt, _ := valueMap["timestamp"]
			direction, _ := valueMap["direction"]

			strWxid, _ := wxid.(string)
			strRoomID, _ := roomID.(string)
			strCreatedAt, _ := createdAt.(string)
			strDirection, _ := direction.(string)
			if len(strWxid) > 0 && len(strRoomID) > 0 && len(strCreatedAt) > 0 && len(strDirection) > 0 {
				p.wechatMessageController.GetRecentMessages(strWxid, strRoomID, strCreatedAt, strDirection)
			}
		}
	}
	c.JSON(http.StatusCreated, gin.H{})
	return
}
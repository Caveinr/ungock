package CQHTTP

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strings"
	"ungock/GORM"
	"ungock/wiringPi"
)

func Handle(ch CQHTTP, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var groupMsg groupMsg

		if err := c.ShouldBind(&groupMsg); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}

		if strings.TrimSpace(groupMsg.Message) != ch.KeyWord {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		insert := GORM.AuthLog{
			REQUESTER: fmt.Sprintf("%d", groupMsg.UserID),
			METHOD:    "GO_CQHTTP",
		}
		for _, groupId := range ch.GroupID {
			if groupMsg.GroupID == groupId {
				wiringPi.Pass()
				c.JSON(http.StatusOK, gin.H{
					"reply": ch.FastReplayOk,
				})
				insert.STATUS = "PASSED"
				db.Omit("ID", "TIME").Create(&insert)
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"reply": ch.FastReplayNo,
		})
		insert.STATUS = "FAILED"
		db.Omit("ID", "TIME").Create(&insert)
		return
	}

}

func heartbeatHandle(ch CQHTTP, hb heartbeat, c *gin.Context) {
}

func groupMsgHandle(ch CQHTTP, msg groupMsg, c *gin.Context) {
}

func Auth(ch CQHTTP) gin.HandlerFunc {
	return func(c *gin.Context) {
		hSignature := c.GetHeader("X-Signature")
		if hSignature != "" {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			// gin框架中request body只能读取一次，需要复写 request body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			// 使用密钥计算request body的 hmac码
			mac := hmac.New(sha1.New, []byte(ch.Secret))
			if _, err := mac.Write(body); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			// 校验hmac签名
			if "sha1="+hex.EncodeToString(mac.Sum(nil)) != hSignature {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}

}

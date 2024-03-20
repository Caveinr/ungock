package MFRCScan

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
	"ungock/GORM"
	"ungock/wiringPi"
)

func Handle(config MFRCScan, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg msg
		if err := c.ShouldBind(&msg); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		//if config.Tolerance > 0 {
		//	if math.Abs(msg.Timestamp-time.Now().Unix()) < config.Tolerance {
		//		c.AbortWithStatus(http.StatusForbidden)
		//		return
		//	}
		//}

		var query []GORM.Auth_uid
		insert := GORM.AuthLog{
			REQUESTER: fmt.Sprintf("%s%d", msg.UID, msg.SAK),
			METHOD:    "MFRC522",
		}
		db.Table("auth_uid").Where("UID = ? AND SAK = ?", msg.UID, msg.SAK).Find(&query)
		if len(query) == 1 {
			wiringPi.Pass()
			c.JSON(http.StatusOK, gin.H{})

			insert.STATUS = "PASSED"
			db.Omit("ID", "TIME").Create(&insert)
			return
		} else {
			insert.STATUS = "FAILED"
			db.Omit("ID", "TIME").Create(&insert)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}

func Auth(mfrcs MFRCScan) gin.HandlerFunc {
	return func(c *gin.Context) {
		hSignature := c.GetHeader("X-Signature")
		if hSignature != "" {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			mac := hmac.New(sha1.New, []byte(mfrcs.Secret))
			if _, err := mac.Write(body); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			if hex.EncodeToString(mac.Sum(nil)) != hSignature {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}

}

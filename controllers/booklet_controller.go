package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	booklet "github.com/leon0399/cardyo-pdf/services/booklet"
)

// GenerateA5Booklet generate
// @Failure 500 {object} string "internal server error"
// @Router /booklet/a5 [get]
func GenerateA5Booklet(c *gin.Context) {
	theme := c.DefaultQuery("theme", "white")
	url := c.Query("url")
	_, download := c.GetQuery("download")

	pdf, err := booklet.GenerateBookletA5(theme, url)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var disposition string

	if download {
		disposition = "attachment"
	} else {
		disposition = "inline"
	}

	c.Status(http.StatusOK)
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", disposition+"; filename=\"Cardyo-Booklet-A5.pdf\"")

	if err := pdf.Write(c.Writer); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

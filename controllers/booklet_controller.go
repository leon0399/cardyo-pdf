package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	booklets "github.com/leon0399/cardyo-pdf/services/booklet"
)

type booklet struct {
	Theme    string `form:"theme,default=white" binding:"oneof=black white"`
	URL      string `form:"url" binding:"required"`
	Download bool   `form:"download,default=false"`
}

// GenerateBookletA5Api generate
// @Failure 500 {object} string "internal server error"
// @Router /booklet/a5 [get]
func GenerateBookletA5Api(c *gin.Context) {
	var b booklet

	if err := c.ShouldBind(&b); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	pdf, err := booklets.GenerateBookletA5(b.Theme, b.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var disposition string

	if b.Download {
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

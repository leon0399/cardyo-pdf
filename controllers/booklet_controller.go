package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	booklet "github.com/leon0399/cardyo-pdf/services/booklet"
	"github.com/signintech/gopdf"
)

// GenerateA5Booklet generate
// @Failure 500 {object} string "internal server error"
// @Router /booklet/a5 [get]
func GenerateA5Booklet(c *gin.Context) {
	theme := c.DefaultQuery("theme", "white")
	url := c.Query("url")

	pdf := gopdf.GoPdf{}
	defer pdf.Close()

	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA5})

	{
		pdf.AddPage()

		if err := booklet.GenerateBookletA5(&pdf, theme, url); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.Status(http.StatusOK)
	c.Header("Content-Type", "application/pdf")

	if err := pdf.Write(c.Writer); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

package services

import (
	"github.com/signintech/gopdf"
)

func GenerateBookletA5(pdf *gopdf.GoPdf, theme string, url string) error {
	tpl := pdf.ImportPage("./assets/templates/a5/"+theme+".pdf", 1, "/MediaBox")

	{
		pdf.UseImportedTemplate(tpl, 0, 0, 430, 595)
	}

	return nil
}

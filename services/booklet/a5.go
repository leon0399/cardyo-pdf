package services

import (
	"github.com/signintech/gopdf"
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateBookletA5(theme string, url string) (*gopdf.GoPdf, error) {
	pdf := gopdf.GoPdf{}
	defer pdf.Close()

	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA5})
	pdf.AddPage()

	{
		tpl := pdf.ImportPage("./assets/templates/a5/"+theme+".pdf", 1, "/MediaBox")
		pdf.UseImportedTemplate(tpl, 0, 0, 420, 595)
	}

	{
		qr, err := qrcode.New(url, qrcode.Medium)
		if err != nil {
			return nil, err
		}

		qr.DisableBorder = true

		png, err := qr.PNG(256)
		if err != nil {
			return nil, err
		}

		holder, err := gopdf.ImageHolderByBytes(png)
		if err != nil {
			return nil, err
		}

		pdf.ImageByHolderWithOptions(holder, gopdf.ImageOptions{X: 40, Y: 477, Rect: &gopdf.Rect{W: 78, H: 78}})
	}

	return &pdf, nil
}

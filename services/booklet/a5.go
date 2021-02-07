package services

import (
	"github.com/signintech/gopdf"
)

func GenerateBookletA5(theme string, url string) (*gopdf.GoPdf, error) {
	pdf := gopdf.GoPdf{}
	defer pdf.Close()

	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA5})
	pdf.AddPage()

	if err := importPageAsBackground(&pdf, "./assets/templates/a5/"+theme+".pdf", gopdf.PageSizeA5); err != nil {
		return nil, err
	}

	if err := addQRCode(&pdf, url, &gopdf.ImageOptions{X: 40, Y: 477, Rect: &gopdf.Rect{W: 78, H: 78}}); err != nil {
		return nil, err
	}

	return &pdf, nil
}

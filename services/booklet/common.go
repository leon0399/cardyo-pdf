package services

import (
	"github.com/signintech/gopdf"
	qrcode "github.com/skip2/go-qrcode"
)

func importPageAsBackground(pdf *gopdf.GoPdf, path string, size *gopdf.Rect) error {
	tpl := pdf.ImportPage(path, 1, "/MediaBox")
	pdf.UseImportedTemplate(tpl, 0, 0, size.W, size.H)

	return nil
}

func addQRCode(pdf *gopdf.GoPdf, content string, opts *gopdf.ImageOptions) error {
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return err
	}

	qr.DisableBorder = true

	png, err := qr.PNG(256)
	if err != nil {
		return err
	}

	holder, err := gopdf.ImageHolderByBytes(png)
	if err != nil {
		return err
	}

	pdf.ImageByHolderWithOptions(holder, *opts)

	return nil
}

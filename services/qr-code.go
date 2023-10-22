package services

import (
	"fmt"
	. "web_practicum/models"

	"strconv"

	qrcode "github.com/skip2/go-qrcode"
)

func WriteQrCode(article Article) (string, error) {
	path := fmt.Sprintf("qr-codes/%s.png",
		strconv.FormatUint(article.ID, 10))
	return path, qrcode.WriteFile(
		MakeLink(article), qrcode.Medium, 250, path)
}

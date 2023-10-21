package services

import (
	"fmt"
	. "web_practicum/models"

	"strconv"

	qrcode "github.com/skip2/go-qrcode"
)

func WriteQrCode(article Article) error {
	return qrcode.WriteFile(
		MakeLink(article), qrcode.Medium, 250, fmt.Sprintf("qr-codes/%s.png",
			strconv.FormatUint(article.ID, 10)))
}

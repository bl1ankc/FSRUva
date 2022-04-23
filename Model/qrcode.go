package Model

import (
	"github.com/skip2/go-qrcode"
)

func CreateQRCode(id string) string {
	err := qrcode.WriteFile(id, qrcode.Medium, 256, "./img/qrcode/"+id+".png")
	if err != nil {
		return ""
	}
	return id
}

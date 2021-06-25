package service

import (
	"wm-infoflow-api-go/service/media"
	"wm-infoflow-api-go/service/menu"
)

type Services interface {
	Menu() menu.MenusStore
	Media() media.MediasStore
}

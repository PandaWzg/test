package service

import (
	"github.com/go-redis/redis"
	gocache "github.com/patrickmn/go-cache"
	"time"

	"wm-infoflow-api-go/service/media"
	medias "wm-infoflow-api-go/service/media"
	"wm-infoflow-api-go/service/menu"
	menus "wm-infoflow-api-go/service/menu"
)

type Wrap struct {
	menu    menu.MenusStore
	media   media.MediasStore
	goCache *gocache.Cache
}

var _ Services = (*Wrap)(nil)

func NewServicesWrap(r *redis.Client) Services {
	s := new(Wrap)
	s.menu = menus.NewMenusService(s,r)
	s.media = medias.NewService(s,r)
	s.goCache = gocache.New(30*time.Minute, 10*time.Minute)
	return s
}

func (s *Wrap) Menu() menu.MenusStore {
	return s.menu
}

func (s *Wrap) Media() media.MediasStore {
	return s.media
}

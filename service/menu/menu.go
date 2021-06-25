package menu

import (
	"github.com/go-redis/redis"
	"wm-infoflow-api-go/service"
)

type MenusService struct {
	services service.Services
	redis *redis.Client
}

func NewMenusService(s service.Services,r *redis.Client) MenusStore {
	u := &MenusService{
		services: s,
		redis: r,
	}
	return u
}
func (m *MenusService) MenuList() {

}

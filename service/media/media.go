package media

import (
	"github.com/go-redis/redis"
	"wm-infoflow-api-go/service"
)

type MediasService struct {
	services service.Services
	redis    *redis.Client
}

func NewService(s service.Services, r *redis.Client) MediasStore {
	u := &MediasService{
		services: s,
		redis:    r,
	}
	return u
}
func (m *MediasService) MediaList() {

}

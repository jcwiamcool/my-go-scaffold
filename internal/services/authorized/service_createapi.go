package authorized

import (
	"github.com/jcwiamcool/my-go-scaffold/configs"
	"github.com/jcwiamcool/my-go-scaffold/internal/pkg/core"
	"github.com/jcwiamcool/my-go-scaffold/internal/repository/mysql/authorized_api"
	"github.com/jcwiamcool/my-go-scaffold/internal/repository/redis"
)

type CreateAuthorizedAPIData struct {
	BusinessKey string `json:"business_key"` // 调用方key
	Method      string `json:"method"`       // 请求方法
	API         string `json:"api"`          // 请求地址
}

func (s *service) CreateAPI(ctx core.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int32, err error) {
	model := authorized_api.NewModel()
	model.BusinessKey = authorizedAPIData.BusinessKey
	model.Method = authorizedAPIData.Method
	model.Api = authorizedAPIData.API
	model.CreatedUser = ctx.SessionUserInfo().UserName
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	s.cache.Del(configs.RedisKeyPrefixSignature+authorizedAPIData.BusinessKey, redis.WithTrace(ctx.Trace()))
	return
}

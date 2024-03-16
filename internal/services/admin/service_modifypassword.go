package admin

import (
	"github.com/jcwiamcool/my-go-scaffold/configs"
	"github.com/jcwiamcool/my-go-scaffold/internal/pkg/core"
	"github.com/jcwiamcool/my-go-scaffold/internal/pkg/password"
	"github.com/jcwiamcool/my-go-scaffold/internal/repository/mysql"
	"github.com/jcwiamcool/my-go-scaffold/internal/repository/mysql/admin"
	"github.com/jcwiamcool/my-go-scaffold/internal/repository/redis"
)

func (s *service) ModifyPassword(ctx core.Context, id int32, newPassword string) (err error) {
	data := map[string]interface{}{
		"password":     password.GeneratePassword(newPassword),
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := admin.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), redis.WithTrace(ctx.Trace()))
	return
}

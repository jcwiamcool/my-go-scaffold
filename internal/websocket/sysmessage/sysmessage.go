package sysmessage

import (
	"github.com/jcwiamcool/my-go-scaffold/internal/pkg/core"
	"github.com/jcwiamcool/my-go-scaffold/internal/repository/mysql"
	"github.com/jcwiamcool/my-go-scaffold/internal/repository/redis"
	"github.com/jcwiamcool/my-go-scaffold/internal/repository/socket"
	"github.com/jcwiamcool/my-go-scaffold/pkg/errors"

	"go.uber.org/zap"
)

var (
	err    error
	server socket.Server
)

type handler struct {
	logger *zap.Logger
	cache  redis.Repo
	db     mysql.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func GetConn() (socket.Server, error) {
	if server != nil {
		return server, nil
	}

	return nil, errors.New("conn is nil")
}

func (h *handler) Connect() core.HandlerFunc {
	return func(ctx core.Context) {
		server, err = socket.New(h.logger, h.db, h.cache, ctx.ResponseWriter(), ctx.Request(), nil)
		if err != nil {
			return
		}

		go server.OnMessage()
	}
}

package generator_handler

import (
	"github.com/jcwiamcool/my-go-scaffold/internal/pkg/core"
)

func (h *handler) HandlerView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("generator_handler", nil)
	}
}

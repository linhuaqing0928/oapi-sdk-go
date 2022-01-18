// Code generated by lark suite oapi sdk gen
package v6

import (
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
	"github.com/larksuite/oapi-sdk-go/event"
)

type ApplicationVisibilityAddedEventHandler struct {
	Fn func(*core.Context, *ApplicationVisibilityAddedEvent) error
}

func (h *ApplicationVisibilityAddedEventHandler) GetEvent() interface{} {
	return &ApplicationVisibilityAddedEvent{}
}

func (h *ApplicationVisibilityAddedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*ApplicationVisibilityAddedEvent))
}

func SetApplicationVisibilityAddedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *ApplicationVisibilityAddedEvent) error) {
	event.SetTypeHandler(conf, "application.application.visibility.added_v6", &ApplicationVisibilityAddedEventHandler{Fn: fn})
}
// Package im code generated by lark suite oapi sdk gen
package im

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v2"
)

type chatUpdatedEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *ChatUpdatedEvent) error
}

func (h *chatUpdatedEventHandler) Event() interface{} {
	return &ChatUpdatedEvent{}
}

func (h *chatUpdatedEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*ChatUpdatedEvent))
}

func (c *chats) UpdatedEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *ChatUpdatedEvent) error) {
	c.app.Webhook.EventHandler("updated", &chatUpdatedEventHandler{handler: handler})
}

type chatDisbandedEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *ChatDisbandedEvent) error
}

func (h *chatDisbandedEventHandler) Event() interface{} {
	return &ChatDisbandedEvent{}
}

func (h *chatDisbandedEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*ChatDisbandedEvent))
}

func (c *chats) DisbandedEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *ChatDisbandedEvent) error) {
	c.app.Webhook.EventHandler("disbanded", &chatDisbandedEventHandler{handler: handler})
}

type chatMemberBotAddedEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *ChatMemberBotAddedEvent) error
}

func (h *chatMemberBotAddedEventHandler) Event() interface{} {
	return &ChatMemberBotAddedEvent{}
}

func (h *chatMemberBotAddedEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*ChatMemberBotAddedEvent))
}

func (c *chatMemberBots) AddedEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *ChatMemberBotAddedEvent) error) {
	c.app.Webhook.EventHandler("added", &chatMemberBotAddedEventHandler{handler: handler})
}

type chatMemberBotDeletedEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *ChatMemberBotDeletedEvent) error
}

func (h *chatMemberBotDeletedEventHandler) Event() interface{} {
	return &ChatMemberBotDeletedEvent{}
}

func (h *chatMemberBotDeletedEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*ChatMemberBotDeletedEvent))
}

func (c *chatMemberBots) DeletedEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *ChatMemberBotDeletedEvent) error) {
	c.app.Webhook.EventHandler("deleted", &chatMemberBotDeletedEventHandler{handler: handler})
}

type chatMemberUserAddedEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *ChatMemberUserAddedEvent) error
}

func (h *chatMemberUserAddedEventHandler) Event() interface{} {
	return &ChatMemberUserAddedEvent{}
}

func (h *chatMemberUserAddedEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*ChatMemberUserAddedEvent))
}

func (c *chatMemberUsers) AddedEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *ChatMemberUserAddedEvent) error) {
	c.app.Webhook.EventHandler("added", &chatMemberUserAddedEventHandler{handler: handler})
}

type chatMemberUserWithdrawnEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *ChatMemberUserWithdrawnEvent) error
}

func (h *chatMemberUserWithdrawnEventHandler) Event() interface{} {
	return &ChatMemberUserWithdrawnEvent{}
}

func (h *chatMemberUserWithdrawnEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*ChatMemberUserWithdrawnEvent))
}

func (c *chatMemberUsers) WithdrawnEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *ChatMemberUserWithdrawnEvent) error) {
	c.app.Webhook.EventHandler("withdrawn", &chatMemberUserWithdrawnEventHandler{handler: handler})
}

type chatMemberUserDeletedEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *ChatMemberUserDeletedEvent) error
}

func (h *chatMemberUserDeletedEventHandler) Event() interface{} {
	return &ChatMemberUserDeletedEvent{}
}

func (h *chatMemberUserDeletedEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*ChatMemberUserDeletedEvent))
}

func (c *chatMemberUsers) DeletedEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *ChatMemberUserDeletedEvent) error) {
	c.app.Webhook.EventHandler("deleted", &chatMemberUserDeletedEventHandler{handler: handler})
}

type messageReceiveEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *MessageReceiveEvent) error
}

func (h *messageReceiveEventHandler) Event() interface{} {
	return &MessageReceiveEvent{}
}

func (h *messageReceiveEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*MessageReceiveEvent))
}

func (m *messages) ReceiveEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *MessageReceiveEvent) error) {
	m.app.Webhook.EventHandler("receive", &messageReceiveEventHandler{handler: handler})
}

type messageMessageReadEventHandler struct {
	handler func(context.Context, *lark.RawRequest, *MessageMessageReadEvent) error
}

func (h *messageMessageReadEventHandler) Event() interface{} {
	return &MessageMessageReadEvent{}
}

func (h *messageMessageReadEventHandler) Handle(ctx context.Context, req *lark.RawRequest, event interface{}) error {
	return h.handler(ctx, req, event.(*MessageMessageReadEvent))
}

func (m *messages) MessageReadEventHandler(handler func(ctx context.Context, req *lark.RawRequest, event *MessageMessageReadEvent) error) {
	m.app.Webhook.EventHandler("message_read", &messageMessageReadEventHandler{handler: handler})
}
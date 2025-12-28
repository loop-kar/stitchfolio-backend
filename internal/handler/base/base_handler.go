package base

import "github.com/imkarthi24/sf-backend/internal/handler"

type BaseHandler struct {
	HealthHandler       Health
	UserHandler         *handler.UserHandler
	ChannelHandler      *handler.ChannelHandler
	MasterConfigHandler *handler.MasterConfigHandler
	AdminHandler        *handler.AdminHandler
	CustomerHandler     *handler.CustomerHandler
	EnquiryHandler      *handler.EnquiryHandler
	OrderHandler        *handler.OrderHandler
	OrderItemHandler    *handler.OrderItemHandler
	MeasurementHandler  *handler.MeasurementHandler
}

func ProvideBaseHandler(health Health,
	user *handler.UserHandler,
	channelHandler *handler.ChannelHandler,
	masterConfigHandler *handler.MasterConfigHandler,
	adminHandler *handler.AdminHandler,
	customerHandler *handler.CustomerHandler,
	enquiryHandler *handler.EnquiryHandler,
	orderHandler *handler.OrderHandler,
	orderItemHandler *handler.OrderItemHandler,
	measurementHandler *handler.MeasurementHandler,
) BaseHandler {
	return BaseHandler{
		HealthHandler:       health,
		UserHandler:         user,
		ChannelHandler:      channelHandler,
		MasterConfigHandler: masterConfigHandler,
		AdminHandler:        adminHandler,
		CustomerHandler:     customerHandler,
		EnquiryHandler:      enquiryHandler,
		OrderHandler:        orderHandler,
		OrderItemHandler:    orderItemHandler,
		MeasurementHandler:  measurementHandler,
	}
}

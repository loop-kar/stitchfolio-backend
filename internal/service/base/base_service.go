package base

import "github.com/imkarthi24/sf-backend/internal/service"

type BaseService struct {
	UserService         service.UserService
	NotificationService service.NotificationService
	ChannelService      service.ChannelService
	MasterConfigService service.MasterConfigService
	CustomerService     service.CustomerService
	EnquiryService      service.EnquiryService
	OrderService        service.OrderService
	OrderItemService    service.OrderItemService
	MeasurementService  service.MeasurementService
}

func ProvideBaseService(
	user service.UserService,
	notifService service.NotificationService,
	channelService service.ChannelService,
	masterConfigService service.MasterConfigService,
	customerService service.CustomerService,
	enquiryService service.EnquiryService,
	orderService service.OrderService,
	orderItemService service.OrderItemService,
	measurementService service.MeasurementService,
) BaseService {
	return BaseService{
		UserService:         user,
		NotificationService: notifService,
		ChannelService:      channelService,
		MasterConfigService: masterConfigService,
		CustomerService:     customerService,
		EnquiryService:      enquiryService,
		OrderService:        orderService,
		OrderItemService:    orderItemService,
		MeasurementService:  measurementService,
	}
}

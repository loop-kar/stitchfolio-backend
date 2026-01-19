package service

import "github.com/imkarthi24/sf-backend/pkg/service/email"

type Service struct {
	EmailService email.EmailService
}

type ServiceOption func(*Service)

func NewService(opts ...ServiceOption) *Service {
	s := &Service{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithEmailService(emailService email.EmailService) ServiceOption {
	return func(s *Service) {
		s.EmailService = emailService
	}
}

package hostinfo

type ServiceOption func(*Service)

func WithGrpcGatewayAddr(addr string) ServiceOption {
	return func(s *Service) {
		s.gatewayAddr = addr
	}
}

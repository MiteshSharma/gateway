package model

type GatewayService struct {
}

func NewGatewayService() GatewayService {
	return GatewayService{}
}

func (gs GatewayService) GetServiceName() string {
	return "GatewayService"
}

func (gs GatewayService) GetServiceNames() []string {
	return []string{"ServiceName"}
}

func (gs GatewayService) GetServiceNameBy(route string) string {
	return "Service"
}

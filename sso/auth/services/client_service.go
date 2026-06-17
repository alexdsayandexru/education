package services

type ClientService interface {
	Verify(clientId string, clientSecret string) error
}

func NewClientService() ClientService {
	return &ClientServiceImpl{}
}

type ClientServiceImpl struct {
}

func (*ClientServiceImpl) Verify(clientId string, clientSecret string) error {
	return nil
}

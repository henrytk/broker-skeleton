package broker

import (
	"context"

	"code.cloudfoundry.org/lager"
	"github.com/henrytk/broker-skeleton/provider"
	"github.com/pivotal-cf/brokerapi"
)

type Broker struct {
	config   Config
	Provider provider.ServiceProvider
	logger   lager.Logger
}

func New(config Config, serviceProvider provider.ServiceProvider, logger lager.Logger) *Broker {
	return &Broker{
		config:   config,
		Provider: serviceProvider,
		logger:   logger,
	}
}

func (b *Broker) Services(ctx context.Context) []brokerapi.Service {
	return []brokerapi.Service{}
}

func (b *Broker) Provision(
	ctx context.Context,
	instanceID string,
	details brokerapi.ProvisionDetails,
	asyncAllowed bool,
) (brokerapi.ProvisionedServiceSpec, error) {
	b.logger.Debug("provision-start", lager.Data{
		"instance-id":   instanceID,
		"details":       details,
		"async-allowed": asyncAllowed,
	})

	if !asyncAllowed {
		return brokerapi.ProvisionedServiceSpec{}, brokerapi.ErrAsyncRequired
	}

	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (b *Broker) Deprovision(
	ctx context.Context,
	instanceID string,
	details brokerapi.DeprovisionDetails,
	asyncAllowed bool,
) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{}, nil
}

func (b *Broker) Bind(
	ctx context.Context,
	instanceID,
	bindingID string,
	details brokerapi.BindDetails,
) (brokerapi.Binding, error) {
	return brokerapi.Binding{}, nil
}

func (b *Broker) Unbind(
	ctx context.Context,
	instanceID,
	bindingID string,
	details brokerapi.UnbindDetails,
) error {
	return nil
}

func (b *Broker) Update(
	ctx context.Context,
	instanceID string,
	details brokerapi.UpdateDetails,
	asyncAllowed bool,
) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

func (b *Broker) LastOperation(
	ctx context.Context,
	instanceID,
	operationData string,
) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}

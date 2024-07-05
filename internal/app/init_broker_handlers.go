package app

import "service-one/internal/api/delivery/broker/status"

func (a *App) initBrokerHandlers() {
	a.statusBrokerHandler = status.NewHandler(a.logger, a.serviceName, a.rabbit)

	a.logger.Debug("handlers created")
}

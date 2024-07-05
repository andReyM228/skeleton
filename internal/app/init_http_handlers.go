package app

import "service-one/internal/api/delivery/http/status"

func (a *App) initHandlers() {
	a.statusHandler = status.NewHandler(a.logger, a.serviceName)

	a.logger.Debug("handlers created")
}

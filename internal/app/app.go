package app

import (
	"context"
	"embed"
	"github.com/andReyM228/lib/database"
	"github.com/andReyM228/lib/log"
	"github.com/andReyM228/lib/rabbit"
	"github.com/andReyM228/one/chain_client"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"service-one/internal/api/delivery"
	"service-one/internal/api/domain"
	"service-one/internal/config"
)

type (
	App struct {
		config config.Config
		//repo        repository.Repository
		//service     service.Service
		statusHandler       delivery.StatusHandler
		statusBrokerHandler delivery.BrokerStatusHandler
		chain               chain_client.Client
		db                  *sqlx.DB
		rabbit              rabbit.Rabbit
		logger              log.Logger
		validator           *validator.Validate
		serviceName         string
		router              *fiber.App
	}
	worker func(ctx context.Context, a *App)
)

func New(name string) App {
	return App{
		serviceName: name,
	}
}

func (a *App) Run(ctx context.Context, fs embed.FS) {
	a.initLogger()
	a.initValidator()
	a.populateConfig()
	a.initChainClient(ctx)
	a.initBroker()
	a.initRepos()
	a.initServices()
	a.initHandlers()
	a.initDatabase(fs)

	a.runWorkers(ctx)
}

func (a *App) initLogger() {
	a.logger = log.Init()
}

func (a *App) initChainClient(ctx context.Context) {
	a.chain = chain_client.NewClient(a.config.Chain)

	err := a.chain.AddAccount(ctx, domain.SignerAccount, a.config.Extra.Mnemonic)
	if err != nil {
		a.logger.Fatal(err.Error())
	}
}

func (a *App) initRepos() {
	//a.repo = repository.NewRepository(a.db, a.logger)

	a.logger.Debug("repos created")
}

func (a *App) initServices() {
	//a.service = service.NewService(a.logger)

	a.logger.Debug("services created")
}

func (a *App) initValidator() {
	a.validator = validator.New()
}

func (a *App) populateConfig() {
	cfg, err := config.ParseConfig()
	if err != nil {
		a.logger.Debugf("populateConfig: %s", err)
	}

	err = cfg.ValidateConfig(a.validator)
	if err != nil {
		a.logger.Debugf("populateConfig: %s", err)
	}

	a.config = cfg
}

func (a *App) initDatabase(fs embed.FS) {
	a.db = database.InitDatabase(a.logger, a.config.DB, fs)
}

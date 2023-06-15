//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/if1bonacci/lets-go-chat/internal/configs"
	"github.com/if1bonacci/lets-go-chat/internal/handlers"
	"github.com/if1bonacci/lets-go-chat/internal/repositories"
	"github.com/if1bonacci/lets-go-chat/internal/routing"
)

func InitializeDB() (configs.MongoDB, error) {
	wire.Build(
		configs.ProvideMongoDB,
		configs.ProvideEnv,
	)
	return configs.MongoDB{}, nil
}

func InitializeRouting() (routing.Routing, error) {
	wire.Build(
		routing.ProvideRouting,
		configs.ProvideEnv,
		configs.ProvideMongoDB,
		repositories.ProvideChatRepo,
		repositories.ProvideUserRepo,
		repositories.ProvideMessageRepo,
		handlers.ProvideAuthHandler,
		handlers.ProvideChatHandler,
		handlers.ProvideMessageHandler,
		handlers.ProvideUserHandler,
	)
	return routing.Routing{}, nil
}

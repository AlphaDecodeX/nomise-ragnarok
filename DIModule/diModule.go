package dimodule

import (
	"database/sql"
	"fmt"

	"github.com/alphadecodex/nomise-ragnarok/configuration"
	"github.com/alphadecodex/nomise-ragnarok/repository"
	"go.uber.org/dig"
)

type DIModule struct{}

func (m *DIModule) NewSQLPool(config *configuration.Configuration) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (m *DIModule) BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(func() (*configuration.Configuration, error) {
		err := configuration.ConfigureRagnarok()
		if err != nil {
			return nil, err
		}
		return configuration.RagnarokConfig, nil
	})

	container.Provide(m.NewSQLPool)

	container.Provide(func(db *sql.DB) *repository.OrderRepository {
		return &repository.OrderRepository{Db: db}
	})

	return container
}

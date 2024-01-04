package context

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rwirdemann/bikeage/application/domain"
)

type PostgresAdapter struct {
	dbpool *pgxpool.Pool
}

func (a PostgresAdapter) GetAll() ([]domain.Configuration, error) {
	rows, _ := a.dbpool.Query(context.Background(), "select * from configurations")
	defer rows.Close()
	var configurations []domain.Configuration
	for rows.Next() {
		// ...
	}
	return configurations, nil
}

func NewPostgresAdapter() *PostgresAdapter {
	return &PostgresAdapter{}
}

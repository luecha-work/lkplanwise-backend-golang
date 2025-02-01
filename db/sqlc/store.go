package db

import "github.com/jackc/pgx/v5/pgxpool"

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
}

// Store defines all functions to execute SQL db queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

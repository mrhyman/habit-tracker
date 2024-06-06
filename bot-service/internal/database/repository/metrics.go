package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

var (
	AdultAge = 18
)

func (r *UserRepositoryImpl) AdultUserMetric() (int, error) {
	var adultCount int
	query := `SELECT COUNT(*) FROM users WHERE EXTRACT(YEAR FROM AGE(CURRENT_DATE, birthday)) >= @age`
	args := pgx.NamedArgs{
		"age": AdultAge,
	}

	err := r.conn.QueryRow(context.Background(), query, args).Scan(&adultCount)

	if err != nil {
		return 0, err
	}

	return adultCount, nil
}

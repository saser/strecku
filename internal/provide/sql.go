package provide

import (
	"context"
	"database/sql"
	"time"

	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

func DBPool(
	ctx context.Context,
	logger *zap.Logger,
	connString string,
	timeout time.Duration,
) (*sql.DB, func(), error) {
	logger.Info("creating new DB pool", zap.String("connString", connString))
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, nil, xerrors.Errorf("provide DB pool: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, nil, xerrors.Errorf("provide DB pool: %w", err)
	}
	cleanup := func() {
		logger.Info("closing DB pool")
		if err := db.Close(); err != nil {
			logger.Error("closing DB pool failed", zap.Error(err))
			return
		}
		logger.Info("closed DB pool")
	}
	return db, cleanup, nil
}
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Goboolean/common/pkg/resolver"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type Client = Queries

func NewDB(c *resolver.ConfigMap) (*Client, error) {
	user, err := c.GetStringKey("USER")
	if err != nil {
		return nil, err
	}

	password, err := c.GetStringKey("PASSWORD")
	if err != nil {
		return nil, err
	}

	host, err := c.GetStringKey("HOST")
	if err != nil {
		return nil, err
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		return nil, err
	}

	database, err := c.GetStringKey("DATABASE")
	if err != nil {
		return nil, err
	}

	config, err := pgx.ParseConfig(fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database,
	))
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	db, err := pgx.ConnectConfig(context.Background(), config)

	if err != nil {
		return nil, fmt.Errorf("unable to connect to db: %s, %w", config.ConnString(), err)
	}

	q := New(db)

	return q, nil
}

func (c *Client) Ping(ctx context.Context) error {
	for {
		if err := c.db.(*pgx.Conn).Ping(ctx); err != nil {
			log.WithField("error", err).Error("Failed to ping, waiting 5 seconds")
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(5 * time.Second):
				continue
			}
		}

		return nil
	}
}

func (c *Client) Close() {
	c.db.(*pgx.Conn).Close(context.Background())
}

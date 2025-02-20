package preparer

import (
	"context"

	"github.com/Goboolean/fetch-system.IaC/internal/etcd"
	"github.com/Goboolean/fetch-system.IaC/internal/influxdb"
	"github.com/Goboolean/fetch-system.IaC/internal/kafka"
	"github.com/Goboolean/fetch-system.IaC/pkg/db"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Manager struct {
	etcd     *etcd.Client
	db       *db.Client
	conf     *kafka.Configurator
	influxdb *influxdb.Client
}

func New(
	etcd *etcd.Client,
	db *db.Client,
	conf *kafka.Configurator,
) *Manager {
	return &Manager{
		etcd: etcd,
		db:   db,
		conf: conf,
	}
}

func (m *Manager) SyncETCDToDB(ctx context.Context) ([]string, error) {
	products, err := m.db.GetAllProducts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get products from db")
	}
	log.Infof("received number of %d products", len(products))

	dtos := make([]*etcd.Product, len(products))

	for i, product := range products {
		dtos[i] = &etcd.Product{
			ID:       product.ID,
			Symbol:   product.Symbol,
			Platform: string(product.Platform),
			Locale:   string(product.Locale),
			Market:   string(product.Market),
		}
	}

	if err := m.etcd.UpsertProducts(ctx, dtos); err != nil {
		return nil, errors.Wrap(err, "Failed to upsert products to etcd")
	}

	var topics []string
	for _, product := range products {
		topics = append(topics, product.ID)
	}

	return topics, nil
}

func (m *Manager) PrepareTopics(ctx context.Context, topics []string) error {
	log.WithField("count", len(topics)).Info("Preparing topics started")

	if err := m.influxdb.CreateBuckets(ctx, topics); err != nil {
		return errors.Wrap(err, "Failed to create buckets")
	}

	log.WithField("count", len(topics)).Info("Preparing topics finished")
	return nil
}

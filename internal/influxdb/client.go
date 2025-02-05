package influxdb

import (
	"context"
	"fmt"
	"log"

	"github.com/Goboolean/common/pkg/resolver"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const (
	maxConcurrentConnections = 20
	maxErrorsLimit           = 5
)

type Client struct {
	client    influxdb2.Client
	org       string
	semaphore *semaphore.Weighted
}

func New(c *resolver.ConfigMap) (*Client, error) {
	influxURL, err := c.GetStringKey("URL")
	if err != nil {
		return nil, err
	}

	influxToken, err := c.GetStringKey("TOKEN")
	if err != nil {
		return nil, err
	}

	org, err := c.GetStringKey("ORGANIZATION")
	if err != nil {
		return nil, err
	}

	client := influxdb2.NewClient(influxURL, influxToken)
	return &Client{
		client:    client,
		org:       org,
		semaphore: semaphore.NewWeighted(maxConcurrentConnections),
	}, nil
}

func (c *Client) Ping(ctx context.Context) error {
	health, err := c.client.Health(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping InfluxDB: %w", err)
	}
	if health.Status != "pass" {
		return fmt.Errorf("InfluxDB health check failed: %s", *health.Message)
	}
	return nil
}

func (c *Client) createBucket(ctx context.Context, bucket string) error {
	bucketsAPI := c.client.BucketsAPI()
	_, err := bucketsAPI.CreateBucketWithName(ctx, &domain.Organization{Id: &c.org}, bucket)
	if err != nil {
		return fmt.Errorf("failed to create bucket %s: %w", bucket, err)
	}
	return nil
}

func (c *Client) isBucketExists(ctx context.Context, bucket string) (bool, error) {
	bucketsAPI := c.client.BucketsAPI()
	buckets, err := bucketsAPI.FindBucketsByOrgName(ctx, c.org)
	if err != nil {
		return false, fmt.Errorf("failed to list buckets: %w", err)
	}

	for _, b := range *buckets {
		if b.Name == bucket {
			return true, nil
		}
	}
	return false, nil
}

func (c *Client) CreateBucket(ctx context.Context, bucket string) error {
	exists, err := c.isBucketExists(ctx, bucket)
	if err != nil {
		return err
	}
	if exists {
		log.Printf("Bucket already exists: %s", bucket)
		return nil
	}
	return c.createBucket(ctx, bucket)
}

func (c *Client) CreateBuckets(ctx context.Context, buckets []string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	errCount := 0

	for _, bucket := range buckets {
		if errCount >= maxErrorsLimit {
			continue
		}

		bucket := bucket

		g.Go(func() error {
			if err := c.semaphore.Acquire(ctx, 1); err != nil {
				errCount++
				return fmt.Errorf("failed to acquire semaphore: %w", err)
			}
			defer c.semaphore.Release(1)

			if err := c.CreateBucket(ctx, bucket); err != nil {
				errCount++
				return fmt.Errorf("failed to create bucket %s: %w", bucket, err)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to create buckets: %w", err)
	}

	return nil
}

func (c *Client) Close() {
	c.client.Close()
}

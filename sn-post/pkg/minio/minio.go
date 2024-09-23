package minio

import (
	"context"
	"github.com/magmaheat/social-network/sn-post/configs"
	"github.com/minio/minio-go/v7"
)

type Minio struct {
}

func New(cfg configs.Config) error {
	ctx := context.Background()

	client, err := minio.New()
}

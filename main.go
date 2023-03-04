package main

import (
	"Image-loader/internal/config"
	"Image-loader/internal/repository"
	"Image-loader/internal/server"
	"Image-loader/internal/service"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	logger := logrus.New()

	cfg := &config.Config{}

	err := cfg.Process()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info(cfg.DB.Driver)

	db, err := sqlx.Connect(cfg.DB.Driver, fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s", cfg.DB.User,
		cfg.DB.Name, cfg.DB.SSLMode, cfg.DB.Password))
	if err != nil {
		logger.Fatal(err.Error())
	}

	minioClient, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.KeyID, cfg.Minio.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		logger.Fatal(err)
	}

	ok, err := minioClient.BucketExists(ctx, cfg.Minio.Bucket)
	if err != nil {
		logger.Fatal(err)
	}

	if !ok {
		err = minioClient.MakeBucket(context.Background(), cfg.Minio.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			logger.Fatal(err)
		}
	}

	userRepo := repository.NewUserRepo(db, cfg.DB)

	err = userRepo.RunMigrations()
	if err != nil {
		logger.Warning(err)
	}

	controller := service.NewController(userRepo, cfg, minioClient)

	srv := server.NewServer(":8000", logger, controller, cfg)
	srv.RegisterRoutes()

	srv.StartServer()
}

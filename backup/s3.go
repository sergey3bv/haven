package backup

import (
	"context"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type s3Backuper struct {
	client *minio.Client

	Bucket string
}

func NewS3Backuper(
	accessKeyId string,
	secretKey string,
	endpoint string,
	bucketName string,
	region string,
) (Backuper, error) {
	cred := credentials.NewStaticV4(accessKeyId, secretKey, "")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  cred,
		Region: region,
		Secure: true,
	})
	if err != nil {
		return nil, err
	}

	return &s3Backuper{
		client: client,
		Bucket: bucketName,
	}, nil
}

func (s3 s3Backuper) Backup(ctx context.Context, filePath string) error {
	size, err := fileSize(filePath)
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = s3.client.PutObject(
		ctx,
		s3.Bucket,
		filePath,
		file,
		size,
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		},
	)
	if err != nil {
		return err
	}

	return nil
}

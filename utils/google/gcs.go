package google

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var Upload *ClientUploader

const (
	ProjectID  = "octopuslab-365307"
	BucketName = "tanam-apps"
)

const MaxFileSize = 5 * 1024 * 1024 // 5 MB as an example limit

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func init() {
	ctx := context.Background()
	keyPath := "utils/google/key.json"
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(keyPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	Upload = &ClientUploader{
		cl:         client,
		bucketName: BucketName,
		projectID:  ProjectID,
	}
}

func getFileSize(file multipart.File) (int64, error) {
	size, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(ctx context.Context, file multipart.File, object string) (string, error) {
	fileSize, err := getFileSize(file)
	if err != nil {
		return "", fmt.Errorf("failed to get file size: %v", err)
	}
	if fileSize > MaxFileSize {
		return "", fmt.Errorf("file size exceeds the limit: %v", MaxFileSize)
	}

	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	uploadedFileURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s%s", c.bucketName, c.uploadPath, object)

	return uploadedFileURL, nil
}

func (c *ClientUploader) DeleteFileFromGCS(ctx context.Context, object string) error {
	err := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

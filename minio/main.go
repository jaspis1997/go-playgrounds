package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	endpoint := os.Getenv("endpoint")
	accessKeyID := os.Getenv("accesskey")
	secretAccessKey := os.Getenv("secretaccesskey")
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal(err)
	}

	bucketName := "test"

	policy, err := minioClient.GetBucketPolicy(ctx, bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(policy)

	objectName := "testdata"
	filePath := "sample.mp4"
	contentType := "video/mp4"

	// Upload the test file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

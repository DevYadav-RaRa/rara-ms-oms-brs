package s3config

import (
	"log"
	"os"
	"time"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/aws"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/aws/aws-sdk-go/service/s3"
)

//x-amz-acl
func GetPresignedUrl() string {
	AWS_BUCKET := os.Getenv("AWS_BUCKET")
	AWS_KEY := os.Getenv("AWS_KEY")
	AWS_REGION := os.Getenv("AWS_REGION")
	AWS_SECRET := os.Getenv("AWS_SECRET")
	ACL := "public-read"

	ses, err := aws.New(aws.Config{
		Region: AWS_REGION,
		ID:     AWS_KEY,
		Secret: AWS_SECRET,
	})
	if err != nil {
		framework.Logs(err.Error())
		return err.Error()
	}

	svc := s3.New(ses)

	var imageKey string = "abc.png"

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &AWS_BUCKET,
		Key:    &imageKey,
		ACL:    &ACL,
	})

	// req.HTTPRequest.Header.Set("ContentType", "text/csv")

	str, err := req.Presign(60 * time.Minute)
	if err != nil {
		framework.Logs(err.Error())
		return err.Error()
	}

	log.Println("The URL is:", str, " err:", err)
	return str
}

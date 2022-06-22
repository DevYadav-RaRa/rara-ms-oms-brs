package s3

import (
	"os"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/aws"
	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
	routing "github.com/qiangxue/fasthttp-routing"
)

func UploadCsv(c *routing.Context) error {
	ses, err := aws.New(aws.Config{
		Region: os.Getenv("AWS_REGION"),
		ID:     os.Getenv("AWS_KEY"),
		Secret: os.Getenv("AWS_SECRET"),
	})
	if err != nil {
		framework.Logs(err.Error())
	}

	// fmt.Println(sess)

	uploader := s3manager.NewUploader(ses)

	MyBucket := os.Getenv("AWS_BUCKET")
	MyRegion := os.Getenv("AWS_REGION")
	ACL := "public-read"

	// framework.Logs(string(c.Request.Body()))

	file, err := c.FormFile("file")
	if err != nil {
		return errors.New("FormFile Error: " + err.Error())
	}

	upFile, err := file.Open()
	if err != nil {
		return errors.New("FileOpen Error: " + err.Error())
	}
	filename := "upload/" + file.Filename
	framework.Logs(filename)

	//upload to the s3 bucket
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &MyBucket,
		ACL:    &ACL,
		Key:    &filename,
		Body:   upFile,
	})

	filepath := "https://" + MyBucket + "." + "s3-" + MyRegion + ".amazonaws.com/" + filename

	if err != nil {
		return errors.New("FilePath: " + filepath + " UploadFile Error: " + err.Error())
	}

	return errors.New("FilePath: " + filepath + ", Error: " + err.Error())
}

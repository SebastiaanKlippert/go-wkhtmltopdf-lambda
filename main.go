package main

import (
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(handler)
}

func handler(s3Event events.S3Event) error {
	if len(s3Event.Records) > 0 {
		return createPDF(s3Event.Records[0])
	}
	return nil
}

func createPDF(record events.S3EventRecord) error {

	// get json file from S3
	object, err := getS3Object(record.S3.Bucket.Name, record.S3.Object.Key)
	if err != nil {
		return err
	}
	defer object.Close()

	// make sure we look for the included wkhtmltopdf binary
	os.Setenv("WKHTMLTOPDF_PATH", os.Getenv("LAMBDA_TASK_ROOT"))

	// create PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGeneratorFromJSON(object)
	if err != nil {
		return err
	}

	// create PDF
	err = pdfg.Create()
	if err != nil {
		return err
	}

	// write PDF to same filename with .pdf added
	return putS3Object(record.S3.Bucket.Name, record.S3.Object.Key+".pdf", pdfg.Bytes())
}

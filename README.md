# go-wkhtmltopdf-lambda
Run wkhtmltopdf in AWS Lambda

#### Work in progress.

Uses https://github.com/SebastiaanKlippert/go-wkhtmltopdf to generate PDF files from JSON input.

# Usage

- Create a Golang AWS Lambda function in the AWS console
- As source, use the lambda.zip file from a [release](https://github.com/SebastiaanKlippert/go-wkhtmltopdf-lambda/releases), this includes the latest version of [wkhtmltopdf](https://wkhtmltopdf.org)
- The Handler name is `go_wkhtmltopdf_linux`
- If you want to build your own version, make sure you build the go executable for Linux (GOOS=linux) and make it executable (chmod +x)
- Create an S3 trigger for you Lambda function (using suffix `.json` is recommended)
- Make sure your IAM Lambda role has S3 Read and Write access
- Create a JSON file from the PDF generator, following the instructions at https://github.com/SebastiaanKlippert/go-wkhtmltopdf#saving-to-and-loading-from-json
- Upload the JSON to the S3 bucket you chose in your Lambda trigger
- The PDF is saved in the same bucket with extension `.pdf` added to the original filename


package aws

//var defaultListObjectsOutput = &s3.ListObjectsOutput{
//	Name:        aws.String(""),
//	Marker:      aws.String(""),
//	MaxKeys:     aws.Int64(1000),
//	Prefix:      aws.String(""),
//	IsTruncated: aws.Bool(false),
//}
//
//// Define a mock struct to be used in your unit tests
//type mockS3Client struct {
//	s3iface.S3API
//}
//
//// ListObjects mocks the S3API ListObjects method
//func (m *mockS3Client) ListObjects(obj *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
//	return defaultListObjectsOutput, nil
//}
//
//// GetObject mocks the S3API GetObject method
//func (m *mockS3Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
//	bytes, err := os.Open(*input.Key)
//	if err != nil {
//		return nil, err
//	}
//
//	return &s3.GetObjectOutput{
//		AcceptRanges:  aws.String("bytes"),
//		Body:          bytes,
//		ContentLength: aws.Int64(1000),
//		ContentType:   aws.String("text/plain"),
//		ETag:          aws.String("d73a503d212d9279e6b2ed8ac6bb81f3"),
//	}, nil
//}

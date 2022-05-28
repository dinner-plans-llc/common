package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

// GCP google cloud provider storage implementation
type GCP struct {
	c *storage.Client
	b string // bucket
}

// NewMgr create new google cloud provider storage instance
func NewMgr(client *storage.Client, bucket string) *GCP {
	return &GCP{c: client, b: bucket}
}

// UploadFile
func (gcp *GCP) UploadFile(ctx context.Context, w io.Writer, bucket, object string, filePath string) error {

	defer gcp.c.Close()

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	timeout, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	o := gcp.c.Bucket(bucket).Object(object)

	// set upload conditions
	// only upload if the object does not exist
	o.If(storage.Conditions{DoesNotExist: true})

	wc := o.NewWriter(timeout)
	if _, err := io.Copy(wc, f); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("wc.Close: %v", err)
	}

	return nil
}

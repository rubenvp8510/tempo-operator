package manifestutils

import "github.com/os-observability/tempo-operator/apis/tempo/v1alpha1"

// Params holds parameters used to create Tempo objects.
type Params struct {
	StorageParams  StorageParams
	ConfigChecksum string

	Tempo v1alpha1.Microservices
}

// StorageParams holds storage configuration.
type StorageParams struct {
	S3 S3
}

// S3 holds S3 configuration.
type S3 struct {
	Endpoint string
	Bucket   string
}

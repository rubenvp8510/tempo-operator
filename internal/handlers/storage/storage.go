package storage

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/grafana/tempo-operator/apis/tempo/v1alpha1"
	"github.com/grafana/tempo-operator/internal/manifests/manifestutils"
)

// GetStorageParamsForTempoStack validates and retrieves StorageParams of the TempoStack CR.
func GetStorageParamsForTempoStack(ctx context.Context, client client.Client, tempo v1alpha1.TempoStack) (manifestutils.StorageParams, field.ErrorList) {
	storagePath := field.NewPath("spec", "storage")
	secretPath := storagePath.Child("secret")
	secretNamePath := secretPath.Child("name")
	tlsPath := storagePath.Child("tls")

	storageSecret, errs := getSecret(ctx, client, tempo.Namespace, tempo.Spec.Storage.Secret.Name, secretNamePath)
	if len(errs) > 0 {
		return manifestutils.StorageParams{}, errs
	}

	storageParams := manifestutils.StorageParams{}
	switch tempo.Spec.Storage.Secret.Type {
	case v1alpha1.ObjectStorageSecretS3:
		storageParams.S3, errs = getS3Params(storageSecret, secretNamePath)
		if len(errs) > 0 {
			return manifestutils.StorageParams{}, errs
		}

		storageParams.S3.Insecure = !tempo.Spec.Storage.TLS.Enabled

		if tempo.Spec.Storage.TLS.Enabled && storageParams.S3.LongLived != nil {
			storageParams.S3.LongLived.TLS, errs = getTLSParams(ctx, client, tempo.Namespace, tempo.Spec.Storage.TLS, tlsPath.Child("caName"))
			if len(errs) > 0 {
				return manifestutils.StorageParams{}, errs
			}
		}

	case v1alpha1.ObjectStorageSecretAzure:
		storageParams.AzureStorage, errs = getAzureParams(storageSecret, secretNamePath)
		if len(errs) > 0 {
			return manifestutils.StorageParams{}, errs
		}

		if tempo.Spec.Storage.TLS.Enabled {
			return manifestutils.StorageParams{}, field.ErrorList{field.Invalid(
				tlsPath.Child("enabled"),
				tempo.Spec.Storage.TLS.Enabled,
				"custom TLS settings are not supported for Azure Storage",
			)}
		}

	case v1alpha1.ObjectStorageSecretGCS:
		storageParams.GCS, errs = getGCSParams(storageSecret, secretNamePath)
		if len(errs) > 0 {
			return manifestutils.StorageParams{}, errs
		}

		if tempo.Spec.Storage.TLS.Enabled {
			return manifestutils.StorageParams{}, field.ErrorList{field.Invalid(
				tlsPath.Child("enabled"),
				tempo.Spec.Storage.TLS.Enabled,
				"custom TLS settings are not supported for Google Cloud Storage",
			)}
		}

	case "":
		return manifestutils.StorageParams{}, field.ErrorList{field.Invalid(
			secretPath.Child("type"),
			tempo.Spec.Storage.Secret.Type,
			"storage secret type is required",
		)}

	default:
		return manifestutils.StorageParams{}, field.ErrorList{field.Invalid(
			secretPath.Child("type"),
			tempo.Spec.Storage.Secret.Type,
			fmt.Sprintf("%s is not an allowed storage secret type", tempo.Spec.Storage.Secret.Type),
		)}
	}

	return storageParams, nil
}

// GetStorageParamsForTempoMonolithic validates and retrieves StorageParams of the TempoMonolithic CR.
func GetStorageParamsForTempoMonolithic(ctx context.Context, client client.Client, tempo v1alpha1.TempoMonolithic) (manifestutils.StorageParams, field.ErrorList) {
	tracesPath := field.NewPath("spec", "storage", "traces")
	if tempo.Spec.Storage == nil {
		return manifestutils.StorageParams{}, field.ErrorList{field.Invalid(tracesPath, "", "storage not configured")}
	}

	storageParams := manifestutils.StorageParams{}
	switch tempo.Spec.Storage.Traces.Backend {
	case v1alpha1.MonolithicTracesStorageBackendMemory,
		v1alpha1.MonolithicTracesStorageBackendPV:
		// nothing to do here

	case v1alpha1.MonolithicTracesStorageBackendS3:
		secretNamePath := tracesPath.Child("s3", "secret")
		if tempo.Spec.Storage.Traces.S3 == nil {
			return manifestutils.StorageParams{}, field.ErrorList{field.Invalid(secretNamePath, "", "please specify a storage secret")}
		}

		storageSecret, errs := getSecret(ctx, client, tempo.Namespace, tempo.Spec.Storage.Traces.S3.Secret, secretNamePath)
		if len(errs) > 0 {
			return manifestutils.StorageParams{}, errs
		}

		storageParams.S3, errs = getS3Params(storageSecret, secretNamePath)
		if len(errs) > 0 {
			return manifestutils.StorageParams{}, errs
		}

		if tempo.Spec.Storage.Traces.S3.TLS != nil && tempo.Spec.Storage.Traces.S3.TLS.Enabled {
			caPath := tracesPath.Child("s3", "tls", "caName")
			storageParams.S3.LongLived.TLS, errs = getTLSParams(ctx, client, tempo.Namespace, *tempo.Spec.Storage.Traces.S3.TLS, caPath)
			if len(errs) > 0 {
				return manifestutils.StorageParams{}, errs
			}
		}

	case v1alpha1.MonolithicTracesStorageBackendAzure:
		secretNamePath := tracesPath.Child("azure").Child("secret")
		if tempo.Spec.Storage.Traces.Azure == nil {
			return manifestutils.StorageParams{}, field.ErrorList{field.Invalid(secretNamePath, "", "please specify a storage secret")}
		}

		storageSecret, errs := getSecret(ctx, client, tempo.Namespace, tempo.Spec.Storage.Traces.Azure.Secret, secretNamePath)
		if len(errs) > 0 {
			return manifestutils.StorageParams{}, errs
		}

		storageParams.AzureStorage, errs = getAzureParams(storageSecret, secretNamePath)
		if len(errs) > 0 {
			return manifestutils.StorageParams{}, errs
		}

	case v1alpha1.MonolithicTracesStorageBackendGCS:
		secretNamePath := tracesPath.Child("gcs").Child("secret")
		if tempo.Spec.Storage.Traces.GCS == nil {
			return manifestutils.StorageParams{}, field.ErrorList{field.Invalid(secretNamePath, "", "please specify a storage secret")}
		}

		storageSecret, errs := getSecret(ctx, client, tempo.Namespace, tempo.Spec.Storage.Traces.GCS.Secret, secretNamePath)
		if len(errs) > 0 {
			return manifestutils.StorageParams{}, errs
		}

		storageParams.GCS, errs = getGCSParams(storageSecret, secretNamePath)
		if len(errs) > 0 {
			return manifestutils.StorageParams{}, errs
		}

	default:
		return manifestutils.StorageParams{}, field.ErrorList{field.Invalid(
			tracesPath.Child("backend"),
			tempo.Spec.Storage.Traces.Backend,
			fmt.Sprintf("%s is not an allowed storage secret type", tempo.Spec.Storage.Traces.Backend),
		)}
	}

	return storageParams, nil
}

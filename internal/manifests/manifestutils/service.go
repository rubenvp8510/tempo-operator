package manifestutils

import (
	"fmt"
	"path"

	"github.com/ViaQ/logerr/v2/kverrors"
	"github.com/imdario/mergo"
	corev1 "k8s.io/api/core/v1"
)

func TempoServerGRPCTLSDir() string {
	return path.Join(grpcTLSDir, "server")
}

func TempoServerHTTPTLSDir() string {
	return path.Join(httpTLSDir, "server")
}

func ConfigureServiceCA(podSpec *corev1.PodSpec, caBundleName string, containers ...int) error {
	secretVolumeSpec := corev1.PodSpec{
		Volumes: []corev1.Volume{
			{
				Name: caBundleName,
				VolumeSource: corev1.VolumeSource{
					ConfigMap: &corev1.ConfigMapVolumeSource{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: caBundleName,
						},
					},
				},
			},
		},
	}

	secretContainerSpec := corev1.Container{
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      caBundleName,
				ReadOnly:  false,
				MountPath: CABundleDir,
			},
		},
	}

	if err := mergo.Merge(podSpec, secretVolumeSpec, mergo.WithAppendSlice); err != nil {
		return kverrors.Wrap(err, "failed to merge volumes")
	}

	if len(containers) > 0 {
		for _, i := range containers {
			if err := mergo.Merge(&podSpec.Containers[i], secretContainerSpec, mergo.WithAppendSlice); err != nil {
				return kverrors.Wrap(err, "failed to merge container")
			}
		}
	} else {
		if err := mergo.Merge(&podSpec.Containers[0], secretContainerSpec, mergo.WithAppendSlice); err != nil {
			return kverrors.Wrap(err, "failed to merge container")
		}
	}

	return nil
}

func ConfigureGRPCServicePKI(podSpec *corev1.PodSpec, componentName string, containers ...int) error {
	serviceName := fmt.Sprintf("%s-grpc", componentName)
	secretVolumeSpec := corev1.PodSpec{
		Volumes: []corev1.Volume{
			{
				Name: serviceName,
				VolumeSource: corev1.VolumeSource{
					Secret: &corev1.SecretVolumeSource{
						SecretName: serviceName,
					},
				},
			},
		},
	}
	secretContainerSpec := corev1.Container{
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      serviceName,
				ReadOnly:  false,
				MountPath: TempoServerGRPCTLSDir(),
			},
		},
	}

	if err := mergo.Merge(podSpec, secretVolumeSpec, mergo.WithAppendSlice); err != nil {
		return kverrors.Wrap(err, "failed to merge volumes")
	}

	if len(containers) > 0 {
		for _, i := range containers {
			if err := mergo.Merge(&podSpec.Containers[i], secretContainerSpec, mergo.WithAppendSlice); err != nil {
				return kverrors.Wrap(err, "failed to merge container")
			}
		}
	} else {
		if err := mergo.Merge(&podSpec.Containers[0], secretContainerSpec, mergo.WithAppendSlice); err != nil {
			return kverrors.Wrap(err, "failed to merge container")
		}
	}
	return nil
}

func ConfigureHTTPServicePKI(podSpec *corev1.PodSpec, componentName string, containers ...int) error {
	serviceName := fmt.Sprintf("%s-http", componentName)
	secretVolumeSpec := corev1.PodSpec{
		Volumes: []corev1.Volume{
			{
				Name: serviceName,
				VolumeSource: corev1.VolumeSource{
					Secret: &corev1.SecretVolumeSource{
						SecretName: serviceName,
					},
				},
			},
		},
	}

	secretContainerSpec := corev1.Container{
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      serviceName,
				ReadOnly:  false,
				MountPath: TempoServerHTTPTLSDir(),
			},
		},
	}

	if err := mergo.Merge(podSpec, secretVolumeSpec, mergo.WithAppendSlice); err != nil {
		return kverrors.Wrap(err, "failed to merge volumes")
	}

	if len(containers) > 0 {
		for _, i := range containers {
			if err := mergo.Merge(&podSpec.Containers[i], secretContainerSpec, mergo.WithAppendSlice); err != nil {
				return kverrors.Wrap(err, "failed to merge container")
			}
		}
	} else {
		if err := mergo.Merge(&podSpec.Containers[0], secretContainerSpec, mergo.WithAppendSlice); err != nil {
			return kverrors.Wrap(err, "failed to merge container")
		}
	}

	return nil
}

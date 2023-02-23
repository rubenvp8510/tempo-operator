package tlsprofile

import (
	"context"

	"github.com/ViaQ/logerr/v2/kverrors"
	openshiftconfigv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/library-go/pkg/crypto"
	"sigs.k8s.io/controller-runtime/pkg/client"

	configv1alpha1 "github.com/os-observability/tempo-operator/apis/config/v1alpha1"
	"github.com/os-observability/tempo-operator/internal/manifests/manifestutils"
)

// APIServerName is the apiserver resource name used to fetch it.
const APIServerName = "cluster"

// GetTLSProfileFromCluster get the TLS profile from the cluster, if it's defined.
func GetTLSProfileFromCluster(ctx context.Context, k client.Client) (*openshiftconfigv1.TLSSecurityProfile, error) {
	var apiServer openshiftconfigv1.APIServer
	if err := k.Get(ctx, client.ObjectKey{Name: APIServerName}, &apiServer); err != nil {
		return nil, kverrors.Wrap(err, "failed to lookup openshift apiServer")
	}
	return apiServer.Spec.TLSSecurityProfile, nil
}

// GetTLSSecurityProfile gets the tls profile info to apply.
func GetTLSSecurityProfile(tlsProfileType configv1alpha1.TLSProfileType) (*openshiftconfigv1.TLSSecurityProfile, error) {
	switch tlsProfileType {
	case configv1alpha1.TLSProfileOldType:
		return &openshiftconfigv1.TLSSecurityProfile{
			Type: openshiftconfigv1.TLSProfileOldType,
		}, nil
	case configv1alpha1.TLSProfileIntermediateType:
		return &openshiftconfigv1.TLSSecurityProfile{
			Type: openshiftconfigv1.TLSProfileIntermediateType,
		}, nil
	case configv1alpha1.TLSProfileModernType:
		return &openshiftconfigv1.TLSSecurityProfile{
			Type: openshiftconfigv1.TLSProfileModernType,
		}, nil
	default:
		return &openshiftconfigv1.TLSSecurityProfile{}, kverrors.New("unknow profile")
	}
}

// GetTLSSettings get the tls settings that belongs to the TLS profile specifications.
func GetTLSSettings(profile *openshiftconfigv1.TLSSecurityProfile) (manifestutils.TLSProfileOptions, error) {
	tlsSecurityProfile := &openshiftconfigv1.TLSSecurityProfile{
		Type: openshiftconfigv1.TLSProfileIntermediateType,
	}

	if profile != nil {
		tlsSecurityProfile = profile
	}

	var (
		minTLSVersion openshiftconfigv1.TLSProtocolVersion
		ciphers       []string
	)

	switch tlsSecurityProfile.Type {
	case openshiftconfigv1.TLSProfileCustomType:
		if tlsSecurityProfile.Custom == nil {
			return manifestutils.TLSProfileOptions{}, kverrors.New("missing TLS custom profile spec")
		}
		minTLSVersion = tlsSecurityProfile.Custom.MinTLSVersion
		ciphers = tlsSecurityProfile.Custom.Ciphers
	case openshiftconfigv1.TLSProfileOldType, openshiftconfigv1.TLSProfileIntermediateType, openshiftconfigv1.TLSProfileModernType:
		spec := openshiftconfigv1.TLSProfiles[tlsSecurityProfile.Type]
		minTLSVersion = spec.MinTLSVersion
		ciphers = spec.Ciphers
	default:
		return manifestutils.TLSProfileOptions{}, kverrors.New("unable to determine tls profile settings %s", tlsSecurityProfile.Type)
	}

	// need to remap all ciphers to their respective IANA names used by Go
	return manifestutils.TLSProfileOptions{
		MinTLSVersion: string(minTLSVersion),
		Ciphers:       crypto.OpenSSLToIANACipherSuites(ciphers),
	}, nil
}

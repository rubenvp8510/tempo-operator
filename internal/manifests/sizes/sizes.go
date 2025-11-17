package sizes

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	tempov1alpha1 "github.com/grafana/tempo-operator/api/tempo/v1alpha1"
)

// ComponentResources is a map of component->requests/limits.
type ComponentResources struct {
	Ingester   ResourceRequirements
	Compactor  ResourceRequirements
	WALStorage ResourceRequirements
	// these two don't need a PVCSize
	Querier       corev1.ResourceRequirements
	Distributor   corev1.ResourceRequirements
	QueryFrontend corev1.ResourceRequirements
	Gateway       corev1.ResourceRequirements
	JaegerQuery   corev1.ResourceRequirements
}

// DeepCopy creates a deep copy of the ComponentResources object, preserving all nested ResourceRequirements fields.
func (c ComponentResources) DeepCopy() ComponentResources {
	return ComponentResources{
		Ingester:      *c.Ingester.DeepCopy(),
		Compactor:     *c.Compactor.DeepCopy(),
		WALStorage:    *c.WALStorage.DeepCopy(),
		Querier:       *c.Querier.DeepCopy(),
		Distributor:   *c.Distributor.DeepCopy(),
		QueryFrontend: *c.QueryFrontend.DeepCopy(),
		Gateway:       *c.Gateway.DeepCopy(),
	}
}

// ResourceRequirements sets CPU, Memory, and PVC requirements for a component.
type ResourceRequirements struct {
	Limits          corev1.ResourceList
	Requests        corev1.ResourceList
	PVCSize         resource.Quantity
	PDBMinAvailable int
}

// DeepCopy creates a deep copy of the ResourceRequirements object, preserving all nested structures.
func (r *ResourceRequirements) DeepCopy() *ResourceRequirements {
	return &ResourceRequirements{
		Limits:          r.Limits.DeepCopy(),
		Requests:        r.Requests.DeepCopy(),
		PVCSize:         r.PVCSize.DeepCopy(),
		PDBMinAvailable: r.PDBMinAvailable,
	}
}

// resourceRequirementsTable defines the default resource requests and limits for each size.
var resourceRequirementsTable = map[tempov1alpha1.TempoStackSizeType]ComponentResources{
	tempov1alpha1.SizeOneXDemo: {
		Ingester: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
		},
		Compactor: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
		},
		WALStorage: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
		},
	},
	tempov1alpha1.SizeOneXPico: {
		Querier: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("750m"),
				corev1.ResourceMemory: resource.MustParse("1.5Gi"),
			},
		},
		Ingester: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("3Gi"),
			},
			PDBMinAvailable: 2,
		},
		Distributor: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		QueryFrontend: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		Compactor: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		Gateway: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		WALStorage: ResourceRequirements{
			PVCSize: resource.MustParse("150Gi"),
		},
	},
	tempov1alpha1.SizeOneXExtraSmall: {
		Querier: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1.5"),
				corev1.ResourceMemory: resource.MustParse("3Gi"),
			},
		},
		Ingester: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("8Gi"),
			},
			PDBMinAvailable: 1,
		},
		Distributor: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		QueryFrontend: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		Compactor: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("2Gi"),
			},
		},
		Gateway: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("500m"),
				corev1.ResourceMemory: resource.MustParse("500Mi"),
			},
		},
		WALStorage: ResourceRequirements{
			PVCSize: resource.MustParse("150Gi"),
		},
	},
	tempov1alpha1.SizeOneXSmall: {
		Querier: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("4"),
				corev1.ResourceMemory: resource.MustParse("4Gi"),
			},
		},
		Ingester: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("4"),
				corev1.ResourceMemory: resource.MustParse("20Gi"),
			},
			PDBMinAvailable: 1,
		},
		Distributor: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("2Gi"),
			},
		},
		QueryFrontend: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("4"),
				corev1.ResourceMemory: resource.MustParse("2.5Gi"),
			},
		},
		Compactor: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("4Gi"),
			},
		},
		Gateway: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		WALStorage: ResourceRequirements{
			PVCSize: resource.MustParse("150Gi"),
		},
	},
	tempov1alpha1.SizeOneXMedium: {
		Querier: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("6"),
				corev1.ResourceMemory: resource.MustParse("10Gi"),
			},
		},
		Ingester: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("6"),
				corev1.ResourceMemory: resource.MustParse("30Gi"),
			},
			PDBMinAvailable: 2,
		},
		Distributor: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("2Gi"),
			},
		},
		QueryFrontend: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("4"),
				corev1.ResourceMemory: resource.MustParse("2.5Gi"),
			},
		},
		Compactor: ResourceRequirements{
			PVCSize: resource.MustParse("10Gi"),
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("2"),
				corev1.ResourceMemory: resource.MustParse("4Gi"),
			},
		},
		Gateway: corev1.ResourceRequirements{
			Requests: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("1Gi"),
			},
		},
		WALStorage: ResourceRequirements{
			PVCSize: resource.MustParse("150Gi"),
		},
	},
}

// ResourceRequirementsForSize returns the resource configuration for a specific LokiStack size.
func ResourceRequirementsForSize(size tempov1alpha1.TempoStackSizeType, useRequestsAsLimits bool) ComponentResources {
	resources := resourceRequirementsTable[size].DeepCopy()
	if useRequestsAsLimits {
		resources.Ingester.Limits = resources.Ingester.Requests.DeepCopy()
		resources.Compactor.Limits = resources.Compactor.Requests.DeepCopy()
		resources.WALStorage.Limits = resources.WALStorage.Requests.DeepCopy()
		resources.Querier.Limits = resources.Querier.Requests.DeepCopy()
		resources.Distributor.Limits = resources.Distributor.Requests.DeepCopy()
		resources.QueryFrontend.Limits = resources.QueryFrontend.Requests.DeepCopy()
		resources.Gateway.Limits = resources.Gateway.Requests.DeepCopy()
	}
	return resources
}

// StackSizeTable defines the default configurations for each size
/* var StackSizeTable = map[tempov1alpha1.TempoStackSizeType]tempov1alpha1.TempoStackSpec{
	lokiv1.SizeOneXDemo: {
		Size: lokiv1.SizeOneXDemo,
		Replication: &lokiv1.ReplicationSpec{
			Factor: 1,
		},
		Limits: &lokiv1.LimitsSpec{
			Global: &lokiv1.LimitsTemplateSpec{
				IngestionLimits: &lokiv1.IngestionLimitSpec{
					// Defaults from Loki docs
					IngestionRate:           4,
					IngestionBurstSize:      6,
					MaxLabelNameLength:      1024,
					MaxLabelValueLength:     2048,
					MaxLabelNamesPerSeries:  30,
					MaxLineSize:             256000,
					PerStreamDesiredRate:    3,
					PerStreamRateLimit:      5,
					PerStreamRateLimitBurst: 15,
				},
				QueryLimits: &lokiv1.QueryLimitSpec{
					// Defaults from Loki docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
					CardinalityLimit:        100000,
					MaxVolumeSeries:         1000,
				},
			},
		},
		Template: &lokiv1.LokiTemplateSpec{
			Compactor: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
			Distributor: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
			Ingester: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
			Querier: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
			QueryFrontend: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
			Gateway: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			IndexGateway: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
			Ruler: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
		},
	},

	lokiv1.SizeOneXPico: {
		Size: lokiv1.SizeOneXPico,
		Replication: &lokiv1.ReplicationSpec{
			Factor: 2,
		},
		Limits: &lokiv1.LimitsSpec{
			Global: &lokiv1.LimitsTemplateSpec{
				IngestionLimits: &lokiv1.IngestionLimitSpec{
					// Defaults from Loki docs
					IngestionRate:             4,
					IngestionBurstSize:        6,
					MaxGlobalStreamsPerTenant: 10000,
					MaxLabelNameLength:        1024,
					MaxLabelValueLength:       2048,
					MaxLabelNamesPerSeries:    30,
					MaxLineSize:               256000,
					PerStreamDesiredRate:      3,
					PerStreamRateLimit:        5,
					PerStreamRateLimitBurst:   15,
				},
				QueryLimits: &lokiv1.QueryLimitSpec{
					// Defaults from Loki docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
					CardinalityLimit:        100000,
					MaxVolumeSeries:         1000,
				},
			},
		},
		Template: &lokiv1.LokiTemplateSpec{
			Compactor: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
			Distributor: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Ingester: &lokiv1.LokiComponentSpec{
				Replicas: 3,
			},
			Querier: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			QueryFrontend: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Gateway: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			IndexGateway: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Ruler: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
		},
	},

	lokiv1.SizeOneXExtraSmall: {
		Size: lokiv1.SizeOneXExtraSmall,
		Replication: &lokiv1.ReplicationSpec{
			Factor: 2,
		},
		Limits: &lokiv1.LimitsSpec{
			Global: &lokiv1.LimitsTemplateSpec{
				IngestionLimits: &lokiv1.IngestionLimitSpec{
					// Defaults from Loki docs
					IngestionRate:             4,
					IngestionBurstSize:        6,
					MaxGlobalStreamsPerTenant: 10000,
					MaxLabelNameLength:        1024,
					MaxLabelValueLength:       2048,
					MaxLabelNamesPerSeries:    30,
					MaxLineSize:               256000,
					PerStreamDesiredRate:      3,
					PerStreamRateLimit:        5,
					PerStreamRateLimitBurst:   15,
				},
				QueryLimits: &lokiv1.QueryLimitSpec{
					// Defaults from Loki docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
					CardinalityLimit:        100000,
					MaxVolumeSeries:         1000,
				},
			},
		},
		Template: &lokiv1.LokiTemplateSpec{
			Compactor: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
			Distributor: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Ingester: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Querier: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			QueryFrontend: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Gateway: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			IndexGateway: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Ruler: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
		},
	},

	lokiv1.SizeOneXSmall: {
		Size: lokiv1.SizeOneXSmall,
		Replication: &lokiv1.ReplicationSpec{
			Factor: 2,
		},
		Limits: &lokiv1.LimitsSpec{
			Global: &lokiv1.LimitsTemplateSpec{
				IngestionLimits: &lokiv1.IngestionLimitSpec{
					// Custom for 1x.small
					IngestionRate:             15,
					IngestionBurstSize:        20,
					MaxGlobalStreamsPerTenant: 10000,
					// Defaults from Loki docs
					MaxLabelNameLength:      1024,
					MaxLabelValueLength:     2048,
					MaxLabelNamesPerSeries:  30,
					MaxLineSize:             256000,
					PerStreamDesiredRate:    3,
					PerStreamRateLimit:      5,
					PerStreamRateLimitBurst: 15,
				},
				QueryLimits: &lokiv1.QueryLimitSpec{
					// Defaults from Loki docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
					CardinalityLimit:        100000,
					MaxVolumeSeries:         1000,
				},
			},
		},
		Template: &lokiv1.LokiTemplateSpec{
			Compactor: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
			Distributor: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Ingester: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Querier: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			QueryFrontend: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Gateway: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			IndexGateway: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Ruler: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
		},
	},

	lokiv1.SizeOneXMedium: {
		Size: lokiv1.SizeOneXMedium,
		Replication: &lokiv1.ReplicationSpec{
			Factor: 2,
		},
		Limits: &lokiv1.LimitsSpec{
			Global: &lokiv1.LimitsTemplateSpec{
				IngestionLimits: &lokiv1.IngestionLimitSpec{
					// Custom for 1x.medium
					IngestionRate:             50,
					IngestionBurstSize:        20,
					MaxGlobalStreamsPerTenant: 25000,
					// Defaults from Loki docs
					MaxLabelNameLength:      1024,
					MaxLabelValueLength:     2048,
					MaxLabelNamesPerSeries:  30,
					MaxLineSize:             256000,
					PerStreamDesiredRate:    3,
					PerStreamRateLimit:      5,
					PerStreamRateLimitBurst: 15,
				},
				QueryLimits: &lokiv1.QueryLimitSpec{
					// Defaults from Loki docs
					MaxEntriesLimitPerQuery: 5000,
					MaxChunksPerQuery:       2000000,
					MaxQuerySeries:          500,
					QueryTimeout:            "3m",
					CardinalityLimit:        100000,
					MaxVolumeSeries:         1000,
				},
			},
		},
		Template: &lokiv1.LokiTemplateSpec{
			Compactor: &lokiv1.LokiComponentSpec{
				Replicas: 1,
			},
			Distributor: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Ingester: &lokiv1.LokiComponentSpec{
				Replicas: 3,
			},
			Querier: &lokiv1.LokiComponentSpec{
				Replicas: 3,
			},
			QueryFrontend: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Gateway: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			IndexGateway: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
			Ruler: &lokiv1.LokiComponentSpec{
				Replicas: 2,
			},
		},
	},
} */

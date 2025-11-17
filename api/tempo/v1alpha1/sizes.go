package v1alpha1

// TempoStackSizeType declares the type for Tempo cluster scale outs.
//
// +kubebuilder:validation:Enum="legacy";"1x.demo";"1x.pico";"1x.extra-small";"1x.small";"1x.medium"
type TempoStackSizeType string

const (
	// SizeOneXDemo defines the size of a single Tempo deployment
	// with tiny resource requirements and without HA support.
	// This size is intended to run in single-node clusters on laptops,
	// it is only useful for very light testing, demonstrations, or prototypes.
	// There are no ingestion/query performance guarantees.
	// DO NOT USE THIS IN PRODUCTION!
	SizeOneXDemo TempoStackSizeType = "1x.demo"

	// SizeOneXPico defines the size of a single Tempo deployment
	// with extra small resources/limits requirements and HA support for all
	// Tempo components. This size is dedicated for setup **without** the
	// requirement for single replication factor and auto-compaction.
	//
	// FIXME: Add clear description of ingestion/query performance expectations.
	SizeOneXPico TempoStackSizeType = "1x.pico"

	// SizeOneXExtraSmall defines the size of a single Tempo deployment
	// with extra small resources/limits requirements and HA support for all
	// Tempo components. This size is dedicated for setup **without** the
	// requirement for single replication factor and auto-compaction.
	//
	// FIXME: Add clear description of ingestion/query performance expectations.
	SizeOneXExtraSmall TempoStackSizeType = "1x.extra-small"

	// SizeOneXSmall defines the size of a single Tempo deployment
	// with small resources/limits requirements and HA support for all
	// Tempo components. This size is dedicated for setup **without** the
	// requirement for single replication factor and auto-compaction.
	//
	// FIXME: Add clear description of ingestion/query performance expectations.
	SizeOneXSmall TempoStackSizeType = "1x.small"

	// SizeOneXMedium defines the size of a single Tempo deployment
	// with small resources/limits requirements and HA support for all
	// Tempo components. This size is dedicated for setup **with** the
	// requirement for single replication factor and auto-compaction.
	//
	// FIXME: Add clear description of ingestion/query performance expectations.
	SizeOneXMedium TempoStackSizeType = "1x.medium"
)

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

// Application Type
type ApplicationType string

const (
	ApplicationType_SAP_HANA = ApplicationType("SAP-HANA")
)

// Specifies whether the volume is enabled for Azure VMware Solution (AVS) datastore purpose
type AvsDataStore string

const (
	// avsDataStore is enabled
	AvsDataStoreEnabled = AvsDataStore("Enabled")
	// avsDataStore is disabled
	AvsDataStoreDisabled = AvsDataStore("Disabled")
)

// This parameter specifies who is authorized to change the ownership of a file. restricted - Only root user can change the ownership of the file. unrestricted - Non-root users can change ownership of files that they own.
type ChownMode string

const (
	ChownModeRestricted   = ChownMode("Restricted")
	ChownModeUnrestricted = ChownMode("Unrestricted")
)

// Encryption type of the capacity pool, set encryption type for data at rest for this pool and all volumes in it. This value can only be set when creating new pool.
type EncryptionType string

const (
	// EncryptionType Single, volumes will use single encryption at rest
	EncryptionTypeSingle = EncryptionType("Single")
	// EncryptionType Double, volumes will use double encryption at rest
	EncryptionTypeDouble = EncryptionType("Double")
)

// Indicates whether the local volume is the source or destination for the Volume Replication
type EndpointType string

const (
	EndpointTypeSrc = EndpointType("src")
	EndpointTypeDst = EndpointType("dst")
)

// Basic network, or Standard features available to the volume.
type NetworkFeatures string

const (
	// Basic network feature.
	NetworkFeaturesBasic = NetworkFeatures("Basic")
	// Standard network feature.
	NetworkFeaturesStandard = NetworkFeatures("Standard")
)

// The qos type of the pool
type QosType string

const (
	// qos type Auto
	QosTypeAuto = QosType("Auto")
	// qos type Manual
	QosTypeManual = QosType("Manual")
)

// Schedule
type ReplicationSchedule string

const (
	ReplicationSchedule_10minutely = ReplicationSchedule("_10minutely")
	ReplicationScheduleHourly      = ReplicationSchedule("hourly")
	ReplicationScheduleDaily       = ReplicationSchedule("daily")
)

// The security style of volume, default unix, defaults to ntfs for dual protocol or CIFS protocol
type SecurityStyle string

const (
	SecurityStyleNtfs = SecurityStyle("ntfs")
	SecurityStyleUnix = SecurityStyle("unix")
)

// The service level of the file system
type ServiceLevel string

const (
	// Standard service level
	ServiceLevelStandard = ServiceLevel("Standard")
	// Premium service level
	ServiceLevelPremium = ServiceLevel("Premium")
	// Ultra service level
	ServiceLevelUltra = ServiceLevel("Ultra")
	// Zone redundant storage service level
	ServiceLevelStandardZRS = ServiceLevel("StandardZRS")
)

func init() {
}
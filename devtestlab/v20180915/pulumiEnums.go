// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180915

// The OS type of the custom image (i.e. Windows, Linux)
type CustomImageOsType string

const (
	CustomImageOsTypeWindows = CustomImageOsType("Windows")
	CustomImageOsTypeLinux   = CustomImageOsType("Linux")
	CustomImageOsTypeNone    = CustomImageOsType("None")
)

// The status of the schedule (i.e. Enabled, Disabled)
type EnableStatus string

const (
	EnableStatusEnabled  = EnableStatus("Enabled")
	EnableStatusDisabled = EnableStatus("Disabled")
)

// The access rights to be granted to the user when provisioning an environment
type EnvironmentPermission string

const (
	EnvironmentPermissionReader      = EnvironmentPermission("Reader")
	EnvironmentPermissionContributor = EnvironmentPermission("Contributor")
)

// Caching option for a data disk (i.e. None, ReadOnly, ReadWrite).
type HostCachingOptions string

const (
	HostCachingOptionsNone      = HostCachingOptions("None")
	HostCachingOptionsReadOnly  = HostCachingOptions("ReadOnly")
	HostCachingOptionsReadWrite = HostCachingOptions("ReadWrite")
)

// The state of the Linux OS (i.e. NonDeprovisioned, DeprovisionRequested, DeprovisionApplied).
type LinuxOsState string

const (
	LinuxOsStateNonDeprovisioned     = LinuxOsState("NonDeprovisioned")
	LinuxOsStateDeprovisionRequested = LinuxOsState("DeprovisionRequested")
	LinuxOsStateDeprovisionApplied   = LinuxOsState("DeprovisionApplied")
)

// Managed identity.
type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                         = ManagedIdentityType("None")
	ManagedIdentityTypeSystemAssigned               = ManagedIdentityType("SystemAssigned")
	ManagedIdentityTypeUserAssigned                 = ManagedIdentityType("UserAssigned")
	ManagedIdentityType_SystemAssigned_UserAssigned = ManagedIdentityType("SystemAssigned,UserAssigned")
)

// The event type for which this notification is enabled (i.e. AutoShutdown, Cost)
type NotificationChannelEventType string

const (
	NotificationChannelEventTypeAutoShutdown = NotificationChannelEventType("AutoShutdown")
	NotificationChannelEventTypeCost         = NotificationChannelEventType("Cost")
)

// The evaluator type of the policy (i.e. AllowedValuesPolicy, MaxValuePolicy).
type PolicyEvaluatorType string

const (
	PolicyEvaluatorTypeAllowedValuesPolicy = PolicyEvaluatorType("AllowedValuesPolicy")
	PolicyEvaluatorTypeMaxValuePolicy      = PolicyEvaluatorType("MaxValuePolicy")
)

// The fact name of the policy (e.g. LabVmCount, LabVmSize, MaxVmsAllowedPerLab, etc.
type PolicyFactName string

const (
	PolicyFactNameUserOwnedLabVmCount         = PolicyFactName("UserOwnedLabVmCount")
	PolicyFactNameUserOwnedLabPremiumVmCount  = PolicyFactName("UserOwnedLabPremiumVmCount")
	PolicyFactNameLabVmCount                  = PolicyFactName("LabVmCount")
	PolicyFactNameLabPremiumVmCount           = PolicyFactName("LabPremiumVmCount")
	PolicyFactNameLabVmSize                   = PolicyFactName("LabVmSize")
	PolicyFactNameGalleryImage                = PolicyFactName("GalleryImage")
	PolicyFactNameUserOwnedLabVmCountInSubnet = PolicyFactName("UserOwnedLabVmCountInSubnet")
	PolicyFactNameLabTargetCost               = PolicyFactName("LabTargetCost")
	PolicyFactNameEnvironmentTemplate         = PolicyFactName("EnvironmentTemplate")
	PolicyFactNameScheduleEditPermission      = PolicyFactName("ScheduleEditPermission")
)

// The status of the policy.
type PolicyStatus string

const (
	PolicyStatusEnabled  = PolicyStatus("Enabled")
	PolicyStatusDisabled = PolicyStatus("Disabled")
)

// The setting to enable usage of premium data disks.
// When its value is 'Enabled', creation of standard or premium data disks is allowed.
// When its value is 'Disabled', only creation of standard data disks is allowed.
type PremiumDataDisk string

const (
	PremiumDataDiskDisabled = PremiumDataDisk("Disabled")
	PremiumDataDiskEnabled  = PremiumDataDisk("Enabled")
)

// The artifact source's type.
type SourceControlType string

const (
	SourceControlTypeVsoGit         = SourceControlType("VsoGit")
	SourceControlTypeGitHub         = SourceControlType("GitHub")
	SourceControlTypeStorageAccount = SourceControlType("StorageAccount")
)

// The storage type for the disk (i.e. Standard, Premium).
type StorageType string

const (
	StorageTypeStandard    = StorageType("Standard")
	StorageTypePremium     = StorageType("Premium")
	StorageTypeStandardSSD = StorageType("StandardSSD")
)

// Protocol type of the port.
type TransportProtocol string

const (
	TransportProtocolTcp = TransportProtocol("Tcp")
	TransportProtocolUdp = TransportProtocol("Udp")
)

// Indicates whether public IP addresses can be assigned to virtual machines on this subnet (i.e. Allow, Deny).
type UsagePermissionType string

const (
	UsagePermissionTypeDefault = UsagePermissionType("Default")
	UsagePermissionTypeDeny    = UsagePermissionType("Deny")
	UsagePermissionTypeAllow   = UsagePermissionType("Allow")
)

// The state of the Windows OS (i.e. NonSysprepped, SysprepRequested, SysprepApplied).
type WindowsOsState string

const (
	WindowsOsStateNonSysprepped    = WindowsOsState("NonSysprepped")
	WindowsOsStateSysprepRequested = WindowsOsState("SysprepRequested")
	WindowsOsStateSysprepApplied   = WindowsOsState("SysprepApplied")
)

func init() {
}
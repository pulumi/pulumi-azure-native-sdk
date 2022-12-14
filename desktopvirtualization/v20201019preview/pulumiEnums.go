// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201019preview

// Resource Type of ApplicationGroup.
type ApplicationGroupType string

const (
	ApplicationGroupTypeRemoteApp = ApplicationGroupType("RemoteApp")
	ApplicationGroupTypeDesktop   = ApplicationGroupType("Desktop")
)

// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
type CommandLineSetting string

const (
	CommandLineSettingDoNotAllow = CommandLineSetting("DoNotAllow")
	CommandLineSettingAllow      = CommandLineSetting("Allow")
	CommandLineSettingRequire    = CommandLineSetting("Require")
)

// HostPool type for desktop.
type HostPoolType string

const (
	HostPoolTypePersonal = HostPoolType("Personal")
	HostPoolTypePooled   = HostPoolType("Pooled")
)

// The type of the load balancer.
type LoadBalancerType string

const (
	LoadBalancerTypeBreadthFirst = LoadBalancerType("BreadthFirst")
	LoadBalancerTypeDepthFirst   = LoadBalancerType("DepthFirst")
	LoadBalancerTypePersistent   = LoadBalancerType("Persistent")
)

// PersonalDesktopAssignment type for HostPool.
type PersonalDesktopAssignmentType string

const (
	PersonalDesktopAssignmentTypeAutomatic = PersonalDesktopAssignmentType("Automatic")
	PersonalDesktopAssignmentTypeDirect    = PersonalDesktopAssignmentType("Direct")
)

// The type of preferred application group type, default to Desktop Application Group
type PreferredAppGroupType string

const (
	PreferredAppGroupTypeNone             = PreferredAppGroupType("None")
	PreferredAppGroupTypeDesktop          = PreferredAppGroupType("Desktop")
	PreferredAppGroupTypeRailApplications = PreferredAppGroupType("RailApplications")
)

// The type of resetting the token.
type RegistrationTokenOperation string

const (
	RegistrationTokenOperationDelete = RegistrationTokenOperation("Delete")
	RegistrationTokenOperationNone   = RegistrationTokenOperation("None")
	RegistrationTokenOperationUpdate = RegistrationTokenOperation("Update")
)

// Resource Type of Application.
type RemoteApplicationType string

const (
	RemoteApplicationTypeInBuilt         = RemoteApplicationType("InBuilt")
	RemoteApplicationTypeMsixApplication = RemoteApplicationType("MsixApplication")
)

// The type of single sign on Secret Type.
type SSOSecretType string

const (
	SSOSecretTypeSharedKey             = SSOSecretType("SharedKey")
	SSOSecretTypeCertificate           = SSOSecretType("Certificate")
	SSOSecretTypeSharedKeyInKeyVault   = SSOSecretType("SharedKeyInKeyVault")
	SSOSecretTypeCertificateInKeyVault = SSOSecretType("CertificateInKeyVault")
)

func init() {
}

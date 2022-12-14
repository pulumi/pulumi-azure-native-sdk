// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

type AccessKeyPermissions string

const (
	AccessKeyPermissionsRead   = AccessKeyPermissions("Read")
	AccessKeyPermissionsWrite  = AccessKeyPermissions("Write")
	AccessKeyPermissionsListen = AccessKeyPermissions("Listen")
	AccessKeyPermissionsSend   = AccessKeyPermissions("Send")
	AccessKeyPermissionsManage = AccessKeyPermissions("Manage")
)

// Optional. Indicates public network solution. If enable, enable public network access of target service with best try. Default is enable. If optOut, opt out public network access configuration.
type ActionType string

const (
	ActionTypeEnable = ActionType("enable")
	ActionTypeOptOut = ActionType("optOut")
)

// Allow caller client IP to access the target service if true. the property is used when connecting local application to target service.
type AllowType string

const (
	AllowTypeTrue  = AllowType("true")
	AllowTypeFalse = AllowType("false")
)

// The authentication type.
type AuthType string

const (
	AuthTypeSystemAssignedIdentity      = AuthType("systemAssignedIdentity")
	AuthTypeUserAssignedIdentity        = AuthType("userAssignedIdentity")
	AuthTypeServicePrincipalSecret      = AuthType("servicePrincipalSecret")
	AuthTypeServicePrincipalCertificate = AuthType("servicePrincipalCertificate")
	AuthTypeSecret                      = AuthType("secret")
	AuthTypeAccessKey                   = AuthType("accessKey")
	AuthTypeUserAccount                 = AuthType("userAccount")
)

// The azure resource type.
type AzureResourceType string

const (
	AzureResourceTypeKeyVault = AzureResourceType("KeyVault")
)

// The application client type
type ClientType string

const (
	ClientTypeNone              = ClientType("none")
	ClientTypeDotnet            = ClientType("dotnet")
	ClientTypeJava              = ClientType("java")
	ClientTypePython            = ClientType("python")
	ClientTypeGo                = ClientType("go")
	ClientTypePhp               = ClientType("php")
	ClientTypeRuby              = ClientType("ruby")
	ClientTypeDjango            = ClientType("django")
	ClientTypeNodejs            = ClientType("nodejs")
	ClientTypeSpringBoot        = ClientType("springBoot")
	ClientType_Kafka_SpringBoot = ClientType("kafka-springBoot")
)

// Indicates whether to clean up previous operation when Linker is updating or deleting
type DeleteOrUpdateBehavior string

const (
	DeleteOrUpdateBehaviorDefault       = DeleteOrUpdateBehavior("Default")
	DeleteOrUpdateBehaviorForcedCleanup = DeleteOrUpdateBehavior("ForcedCleanup")
)

// The name of action for you dryrun job.
type DryrunActionName string

const (
	DryrunActionNameCreateOrUpdate = DryrunActionName("createOrUpdate")
)

// The secret type.
type SecretType string

const (
	SecretTypeRawValue                = SecretType("rawValue")
	SecretTypeKeyVaultSecretUri       = SecretType("keyVaultSecretUri")
	SecretTypeKeyVaultSecretReference = SecretType("keyVaultSecretReference")
)

// The target service type.
type TargetServiceType string

const (
	TargetServiceTypeAzureResource            = TargetServiceType("AzureResource")
	TargetServiceTypeConfluentBootstrapServer = TargetServiceType("ConfluentBootstrapServer")
	TargetServiceTypeConfluentSchemaRegistry  = TargetServiceType("ConfluentSchemaRegistry")
	TargetServiceTypeSelfHostedServer         = TargetServiceType("SelfHostedServer")
)

// Type of VNet solution.
type VNetSolutionType string

const (
	VNetSolutionTypeServiceEndpoint = VNetSolutionType("serviceEndpoint")
	VNetSolutionTypePrivateLink     = VNetSolutionType("privateLink")
)

func init() {
}

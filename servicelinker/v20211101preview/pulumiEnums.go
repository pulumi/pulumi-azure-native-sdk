// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101preview

// The authentication type.
type AuthType string

const (
	AuthTypeSystemAssignedIdentity      = AuthType("systemAssignedIdentity")
	AuthTypeUserAssignedIdentity        = AuthType("userAssignedIdentity")
	AuthTypeServicePrincipalSecret      = AuthType("servicePrincipalSecret")
	AuthTypeServicePrincipalCertificate = AuthType("servicePrincipalCertificate")
	AuthTypeSecret                      = AuthType("secret")
)

// The application client type
type ClientType string

const (
	ClientTypeNone       = ClientType("none")
	ClientTypeDotnet     = ClientType("dotnet")
	ClientTypeJava       = ClientType("java")
	ClientTypePython     = ClientType("python")
	ClientTypeGo         = ClientType("go")
	ClientTypePhp        = ClientType("php")
	ClientTypeRuby       = ClientType("ruby")
	ClientTypeDjango     = ClientType("django")
	ClientTypeNodejs     = ClientType("nodejs")
	ClientTypeSpringBoot = ClientType("springBoot")
)

// Type of VNet solution.
type VNetSolutionType string

const (
	VNetSolutionTypeServiceEndpoint = VNetSolutionType("serviceEndpoint")
	VNetSolutionTypePrivateLink     = VNetSolutionType("privateLink")
)

func init() {
}
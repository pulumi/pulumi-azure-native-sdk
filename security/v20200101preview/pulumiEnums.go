// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

// Connect to your cloud account, for AWS use either account credentials or role-based authentication. For GCP use account organization credentials.
type AuthenticationType string

const (
	// AWS cloud account connector user credentials authentication
	AuthenticationTypeAwsCreds = AuthenticationType("awsCreds")
	// AWS account connector assume role authentication
	AuthenticationTypeAwsAssumeRole = AuthenticationType("awsAssumeRole")
	// GCP account connector service to service authentication
	AuthenticationTypeGcpCredentials = AuthenticationType("gcpCredentials")
)

// Whether or not to automatically install Azure Arc (hybrid compute) agents on machines
type AutoProvision string

const (
	// Install missing Azure Arc agents on machines automatically
	AutoProvisionOn = AutoProvision("On")
	// Do not install Azure Arc agent on the machines automatically
	AutoProvisionOff = AutoProvision("Off")
)

// Defines the minimal alert severity which will be sent as email notifications
type MinimalSeverity string

const (
	// Get notifications on new alerts with High severity
	MinimalSeverityHigh = MinimalSeverity("High")
	// Get notifications on new alerts with medium or high severity
	MinimalSeverityMedium = MinimalSeverity("Medium")
	// Don't get notifications on new alerts with low, medium or high severity
	MinimalSeverityLow = MinimalSeverity("Low")
)

// A possible role to configure sending security notification alerts to
type Roles string

const (
	// If enabled, send notification on new alerts to the account admins
	RolesAccountAdmin = Roles("AccountAdmin")
	// If enabled, send notification on new alerts to the service admins
	RolesServiceAdmin = Roles("ServiceAdmin")
	// If enabled, send notification on new alerts to the subscription owners
	RolesOwner = Roles("Owner")
	// If enabled, send notification on new alerts to the subscription contributors
	RolesContributor = Roles("Contributor")
)

// Defines whether to send email notifications from AMicrosoft Defender for Cloud to persons with specific RBAC roles on the subscription.
type State string

const (
	// Send notification on new alerts to the subscription's admins
	StateOn = State("On")
	// Don't send notification on new alerts to the subscription's admins
	StateOff = State("Off")
)

func init() {
}
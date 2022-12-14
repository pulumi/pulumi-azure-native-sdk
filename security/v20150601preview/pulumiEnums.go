// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150601preview

type Protocol string

const (
	ProtocolTCP = Protocol("TCP")
	ProtocolUDP = Protocol("UDP")
	ProtocolAll = Protocol("*")
)

// The status of the port
type Status string

const (
	StatusRevoked   = Status("Revoked")
	StatusInitiated = Status("Initiated")
)

// A description of why the `status` has its value
type StatusReason string

const (
	StatusReasonExpired               = StatusReason("Expired")
	StatusReasonUserRequested         = StatusReason("UserRequested")
	StatusReasonNewerRequestInitiated = StatusReason("NewerRequestInitiated")
)

func init() {
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201preview

// Specifies if the IP is exposed to the public internet.
type ContainerGroupIpAddressType string

const (
	ContainerGroupIpAddressTypePublic = ContainerGroupIpAddressType("Public")
)

// The protocol associated with the port.
type ContainerGroupNetworkProtocol string

const (
	ContainerGroupNetworkProtocolTCP = ContainerGroupNetworkProtocol("TCP")
	ContainerGroupNetworkProtocolUDP = ContainerGroupNetworkProtocol("UDP")
)

// Restart policy for all containers within the container group.
// - `Always` Always restart
// - `OnFailure` Restart on failure
// - `Never` Never restart
type ContainerGroupRestartPolicy string

const (
	ContainerGroupRestartPolicyAlways    = ContainerGroupRestartPolicy("Always")
	ContainerGroupRestartPolicyOnFailure = ContainerGroupRestartPolicy("OnFailure")
	ContainerGroupRestartPolicyNever     = ContainerGroupRestartPolicy("Never")
)

// The protocol associated with the port.
type ContainerNetworkProtocol string

const (
	ContainerNetworkProtocolTCP = ContainerNetworkProtocol("TCP")
	ContainerNetworkProtocolUDP = ContainerNetworkProtocol("UDP")
)

// The operating system type required by the containers in the container group.
type OperatingSystemTypes string

const (
	OperatingSystemTypesWindows = OperatingSystemTypes("Windows")
	OperatingSystemTypesLinux   = OperatingSystemTypes("Linux")
)

func init() {
}
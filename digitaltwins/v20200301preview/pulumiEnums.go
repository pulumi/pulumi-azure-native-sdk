// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

// The name of the SKU.
type DigitalTwinsSku string

const (
	DigitalTwinsSkuF1 = DigitalTwinsSku("F1")
)

// The type of Digital Twins endpoint
type EndpointType string

const (
	EndpointTypeEventHub   = EndpointType("EventHub")
	EndpointTypeEventGrid  = EndpointType("EventGrid")
	EndpointTypeServiceBus = EndpointType("ServiceBus")
)

func init() {
}

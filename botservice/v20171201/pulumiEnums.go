// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201

// Required. Gets or sets the Kind of the resource.
type Kind string

const (
	KindSdk      = Kind("sdk")
	KindDesigner = Kind("designer")
	KindBot      = Kind("bot")
	KindFunction = Kind("function")
)

// The sku name
type SkuName string

const (
	SkuNameF0 = SkuName("F0")
	SkuNameS1 = SkuName("S1")
)

func init() {
}
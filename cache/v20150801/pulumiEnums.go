// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

// Which family to use. Valid values: (C, P).
type SkuFamily string

const (
	SkuFamilyC = SkuFamily("C")
	SkuFamilyP = SkuFamily("P")
)

// What type of Redis cache to deploy. Valid values: (Basic, Standard, Premium).
type SkuName string

const (
	SkuNameBasic    = SkuName("Basic")
	SkuNameStandard = SkuName("Standard")
	SkuNamePremium  = SkuName("Premium")
)

func init() {
}
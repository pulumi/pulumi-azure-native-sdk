// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned,UserAssigned")
)

// Managed identity type to use for accessing encryption key Url
type Type string

const (
	TypeSystemAssigned = Type("SystemAssigned")
	TypeUserAssigned   = Type("UserAssigned")
)

func init() {
}

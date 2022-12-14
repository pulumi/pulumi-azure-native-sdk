// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

// Type of the managed identity
type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                         = ManagedIdentityType("None")
	ManagedIdentityTypeSystemAssigned               = ManagedIdentityType("SystemAssigned")
	ManagedIdentityTypeUserAssigned                 = ManagedIdentityType("UserAssigned")
	ManagedIdentityType_SystemAssigned_UserAssigned = ManagedIdentityType("SystemAssigned,UserAssigned")
)

// Runtime version
type RuntimeVersion string

const (
	RuntimeVersion_Java_8     = RuntimeVersion("Java_8")
	RuntimeVersion_Java_11    = RuntimeVersion("Java_11")
	RuntimeVersion_NetCore_31 = RuntimeVersion("NetCore_31")
)

// The type of the storage.
type StorageType string

const (
	StorageTypeStorageAccount = StorageType("StorageAccount")
)

// The type of the underlying resource to mount as a persistent disk.
type Type string

const (
	TypeAzureFileVolume = Type("AzureFileVolume")
)

// Type of the source uploaded
type UserSourceType string

const (
	UserSourceTypeJar        = UserSourceType("Jar")
	UserSourceTypeNetCoreZip = UserSourceType("NetCoreZip")
	UserSourceTypeSource     = UserSourceType("Source")
	UserSourceTypeContainer  = UserSourceType("Container")
)

func init() {
}

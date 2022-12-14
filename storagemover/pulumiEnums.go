// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagemover

// Strategy to use for copy.
type CopyMode string

const (
	CopyModeAdditive = CopyMode("Additive")
	CopyModeMirror   = CopyMode("Mirror")
)

// The Endpoint resource type.
type EndpointType string

const (
	EndpointTypeAzureStorageBlobContainer = EndpointType("AzureStorageBlobContainer")
	EndpointTypeNfsMount                  = EndpointType("NfsMount")
)

// The NFS protocol version.
type NfsVersion string

const (
	NfsVersionNFSauto = NfsVersion("NFSauto")
	NfsVersionNFSv3   = NfsVersion("NFSv3")
	NfsVersionNFSv4   = NfsVersion("NFSv4")
)

func init() {
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a server.
type Server struct {
	pulumi.CustomResourceState

	// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
	AdministratorLogin pulumi.StringPtrOutput `pulumi:"administratorLogin"`
	// Status showing whether the server data encryption is enabled with customer-managed keys.
	ByokEnforcement pulumi.StringOutput `pulumi:"byokEnforcement"`
	// Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate pulumi.StringPtrOutput `pulumi:"earliestRestoreDate"`
	// The fully qualified domain name of a server.
	FullyQualifiedDomainName pulumi.StringPtrOutput `pulumi:"fullyQualifiedDomainName"`
	// The Azure Active Directory identity of the server.
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// Status showing whether the server enabled infrastructure encryption.
	InfrastructureEncryption pulumi.StringPtrOutput `pulumi:"infrastructureEncryption"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The master server id of a replica server.
	MasterServerId pulumi.StringPtrOutput `pulumi:"masterServerId"`
	// Enforce a minimal Tls version for the server.
	MinimalTlsVersion pulumi.StringPtrOutput `pulumi:"minimalTlsVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections on a server
	PrivateEndpointConnections ServerPrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The maximum number of replicas that a master server can have.
	ReplicaCapacity pulumi.IntPtrOutput `pulumi:"replicaCapacity"`
	// The replication role of the server.
	ReplicationRole pulumi.StringPtrOutput `pulumi:"replicationRole"`
	// The SKU (pricing tier) of the server.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Enable ssl enforcement or not when connect to server.
	SslEnforcement pulumi.StringPtrOutput `pulumi:"sslEnforcement"`
	// Storage profile of a server.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// A state of a server that is visible to user.
	UserVisibleState pulumi.StringPtrOutput `pulumi:"userVisibleState"`
	// Server version.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbforpostgresql:Server"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20171201preview:Server"),
		},
	})
	opts = append(opts, aliases)
	var resource Server
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20171201:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20171201:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
}

type ServerState struct {
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// The Azure Active Directory identity of the server.
	Identity *ResourceIdentity `pulumi:"identity"`
	// The location the resource resides in.
	Location *string `pulumi:"location"`
	// Properties of the server.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName *string `pulumi:"serverName"`
	// The SKU (pricing tier) of the server.
	Sku *Sku `pulumi:"sku"`
	// Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The Azure Active Directory identity of the server.
	Identity ResourceIdentityPtrInput
	// The location the resource resides in.
	Location pulumi.StringPtrInput
	// Properties of the server.
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringPtrInput
	// The SKU (pricing tier) of the server.
	Sku SkuPtrInput
	// Application-specific metadata in the form of key-value pairs.
	Tags pulumi.StringMapInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
func (o ServerOutput) AdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.AdministratorLogin }).(pulumi.StringPtrOutput)
}

// Status showing whether the server data encryption is enabled with customer-managed keys.
func (o ServerOutput) ByokEnforcement() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ByokEnforcement }).(pulumi.StringOutput)
}

// Earliest restore point creation time (ISO8601 format)
func (o ServerOutput) EarliestRestoreDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.EarliestRestoreDate }).(pulumi.StringPtrOutput)
}

// The fully qualified domain name of a server.
func (o ServerOutput) FullyQualifiedDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.FullyQualifiedDomainName }).(pulumi.StringPtrOutput)
}

// The Azure Active Directory identity of the server.
func (o ServerOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Server) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Status showing whether the server enabled infrastructure encryption.
func (o ServerOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o ServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The master server id of a replica server.
func (o ServerOutput) MasterServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.MasterServerId }).(pulumi.StringPtrOutput)
}

// Enforce a minimal Tls version for the server.
func (o ServerOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections on a server
func (o ServerOutput) PrivateEndpointConnections() ServerPrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Server) ServerPrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(ServerPrivateEndpointConnectionResponseArrayOutput)
}

// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
func (o ServerOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The maximum number of replicas that a master server can have.
func (o ServerOutput) ReplicaCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.IntPtrOutput { return v.ReplicaCapacity }).(pulumi.IntPtrOutput)
}

// The replication role of the server.
func (o ServerOutput) ReplicationRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.ReplicationRole }).(pulumi.StringPtrOutput)
}

// The SKU (pricing tier) of the server.
func (o ServerOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Server) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Enable ssl enforcement or not when connect to server.
func (o ServerOutput) SslEnforcement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.SslEnforcement }).(pulumi.StringPtrOutput)
}

// Storage profile of a server.
func (o ServerOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *Server) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// Resource tags.
func (o ServerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Server) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A state of a server that is visible to user.
func (o ServerOutput) UserVisibleState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.UserVisibleState }).(pulumi.StringPtrOutput)
}

// Server version.
func (o ServerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerOutput{})
}
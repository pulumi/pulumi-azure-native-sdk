// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnsresolver

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a DNS resolver domain list.
//
// Uses Azure REST API version 2023-07-01-preview.
type DnsResolverDomainList struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The domains in the domain list.
	Domains pulumi.StringArrayOutput `pulumi:"domains"`
	// ETag of the DNS resolver domain list.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The current provisioning state of the DNS resolver domain list. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resourceGuid property of the DNS resolver domain list resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDnsResolverDomainList registers a new resource with the given unique name, arguments, and options.
func NewDnsResolverDomainList(ctx *pulumi.Context,
	name string, args *DnsResolverDomainListArgs, opts ...pulumi.ResourceOption) (*DnsResolverDomainList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domains == nil {
		return nil, errors.New("invalid value for required argument 'Domains'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dnsresolver/v20230701preview:DnsResolverDomainList"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230701preview:DnsResolverDomainList"),
		},
		{
			Type: pulumi.String("azure-native:network:DnsResolverDomainList"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DnsResolverDomainList
	err := ctx.RegisterResource("azure-native:dnsresolver:DnsResolverDomainList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsResolverDomainList gets an existing DnsResolverDomainList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsResolverDomainList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsResolverDomainListState, opts ...pulumi.ResourceOption) (*DnsResolverDomainList, error) {
	var resource DnsResolverDomainList
	err := ctx.ReadResource("azure-native:dnsresolver:DnsResolverDomainList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsResolverDomainList resources.
type dnsResolverDomainListState struct {
}

type DnsResolverDomainListState struct {
}

func (DnsResolverDomainListState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsResolverDomainListState)(nil)).Elem()
}

type dnsResolverDomainListArgs struct {
	// The name of the DNS resolver domain list.
	DnsResolverDomainListName *string `pulumi:"dnsResolverDomainListName"`
	// The domains in the domain list.
	Domains []string `pulumi:"domains"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DnsResolverDomainList resource.
type DnsResolverDomainListArgs struct {
	// The name of the DNS resolver domain list.
	DnsResolverDomainListName pulumi.StringPtrInput
	// The domains in the domain list.
	Domains pulumi.StringArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DnsResolverDomainListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsResolverDomainListArgs)(nil)).Elem()
}

type DnsResolverDomainListInput interface {
	pulumi.Input

	ToDnsResolverDomainListOutput() DnsResolverDomainListOutput
	ToDnsResolverDomainListOutputWithContext(ctx context.Context) DnsResolverDomainListOutput
}

func (*DnsResolverDomainList) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsResolverDomainList)(nil)).Elem()
}

func (i *DnsResolverDomainList) ToDnsResolverDomainListOutput() DnsResolverDomainListOutput {
	return i.ToDnsResolverDomainListOutputWithContext(context.Background())
}

func (i *DnsResolverDomainList) ToDnsResolverDomainListOutputWithContext(ctx context.Context) DnsResolverDomainListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsResolverDomainListOutput)
}

type DnsResolverDomainListOutput struct{ *pulumi.OutputState }

func (DnsResolverDomainListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsResolverDomainList)(nil)).Elem()
}

func (o DnsResolverDomainListOutput) ToDnsResolverDomainListOutput() DnsResolverDomainListOutput {
	return o
}

func (o DnsResolverDomainListOutput) ToDnsResolverDomainListOutputWithContext(ctx context.Context) DnsResolverDomainListOutput {
	return o
}

// The Azure API version of the resource.
func (o DnsResolverDomainListOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolverDomainList) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The domains in the domain list.
func (o DnsResolverDomainListOutput) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DnsResolverDomainList) pulumi.StringArrayOutput { return v.Domains }).(pulumi.StringArrayOutput)
}

// ETag of the DNS resolver domain list.
func (o DnsResolverDomainListOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolverDomainList) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o DnsResolverDomainListOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolverDomainList) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o DnsResolverDomainListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolverDomainList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the DNS resolver domain list. This is a read-only property and any attempt to set this value will be ignored.
func (o DnsResolverDomainListOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolverDomainList) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resourceGuid property of the DNS resolver domain list resource.
func (o DnsResolverDomainListOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolverDomainList) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DnsResolverDomainListOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DnsResolverDomainList) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DnsResolverDomainListOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DnsResolverDomainList) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DnsResolverDomainListOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsResolverDomainList) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DnsResolverDomainListOutput{})
}

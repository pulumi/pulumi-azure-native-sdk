// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Rule Collection Group resource.
//
// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
//
// Other available API versions: 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type FirewallPolicyRuleCollectionGroup struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Priority of the Firewall Policy Rule Collection Group resource.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The provisioning state of the firewall policy rule collection group resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Group of Firewall Policy rule collections.
	RuleCollections pulumi.ArrayOutput `pulumi:"ruleCollections"`
	// A read-only string that represents the size of the FirewallPolicyRuleCollectionGroupProperties in MB. (ex 1.2MB)
	Size pulumi.StringOutput `pulumi:"size"`
	// Rule Group type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallPolicyRuleCollectionGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context,
	name string, args *FirewallPolicyRuleCollectionGroupArgs, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleCollectionGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicyName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20200501:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230901:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20231101:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240101:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240301:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240501:FirewallPolicyRuleCollectionGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240701:FirewallPolicyRuleCollectionGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FirewallPolicyRuleCollectionGroup
	err := ctx.RegisterResource("azure-native:network:FirewallPolicyRuleCollectionGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicyRuleCollectionGroup gets an existing FirewallPolicyRuleCollectionGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyRuleCollectionGroupState, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleCollectionGroup, error) {
	var resource FirewallPolicyRuleCollectionGroup
	err := ctx.ReadResource("azure-native:network:FirewallPolicyRuleCollectionGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicyRuleCollectionGroup resources.
type firewallPolicyRuleCollectionGroupState struct {
}

type FirewallPolicyRuleCollectionGroupState struct {
}

func (FirewallPolicyRuleCollectionGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleCollectionGroupState)(nil)).Elem()
}

type firewallPolicyRuleCollectionGroupArgs struct {
	// The name of the Firewall Policy.
	FirewallPolicyName string `pulumi:"firewallPolicyName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Priority of the Firewall Policy Rule Collection Group resource.
	Priority *int `pulumi:"priority"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the FirewallPolicyRuleCollectionGroup.
	RuleCollectionGroupName *string `pulumi:"ruleCollectionGroupName"`
	// Group of Firewall Policy rule collections.
	RuleCollections []interface{} `pulumi:"ruleCollections"`
}

// The set of arguments for constructing a FirewallPolicyRuleCollectionGroup resource.
type FirewallPolicyRuleCollectionGroupArgs struct {
	// The name of the Firewall Policy.
	FirewallPolicyName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Priority of the Firewall Policy Rule Collection Group resource.
	Priority pulumi.IntPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the FirewallPolicyRuleCollectionGroup.
	RuleCollectionGroupName pulumi.StringPtrInput
	// Group of Firewall Policy rule collections.
	RuleCollections pulumi.ArrayInput
}

func (FirewallPolicyRuleCollectionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleCollectionGroupArgs)(nil)).Elem()
}

type FirewallPolicyRuleCollectionGroupInput interface {
	pulumi.Input

	ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput
	ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput
}

func (*FirewallPolicyRuleCollectionGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleCollectionGroup)(nil)).Elem()
}

func (i *FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput {
	return i.ToFirewallPolicyRuleCollectionGroupOutputWithContext(context.Background())
}

func (i *FirewallPolicyRuleCollectionGroup) ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleCollectionGroupOutput)
}

type FirewallPolicyRuleCollectionGroupOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleCollectionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleCollectionGroup)(nil)).Elem()
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupOutput() FirewallPolicyRuleCollectionGroupOutput {
	return o
}

func (o FirewallPolicyRuleCollectionGroupOutput) ToFirewallPolicyRuleCollectionGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionGroupOutput {
	return o
}

// The Azure API version of the resource.
func (o FirewallPolicyRuleCollectionGroupOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o FirewallPolicyRuleCollectionGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o FirewallPolicyRuleCollectionGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Priority of the Firewall Policy Rule Collection Group resource.
func (o FirewallPolicyRuleCollectionGroupOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The provisioning state of the firewall policy rule collection group resource.
func (o FirewallPolicyRuleCollectionGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Group of Firewall Policy rule collections.
func (o FirewallPolicyRuleCollectionGroupOutput) RuleCollections() pulumi.ArrayOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.ArrayOutput { return v.RuleCollections }).(pulumi.ArrayOutput)
}

// A read-only string that represents the size of the FirewallPolicyRuleCollectionGroupProperties in MB. (ex 1.2MB)
func (o FirewallPolicyRuleCollectionGroupOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.StringOutput { return v.Size }).(pulumi.StringOutput)
}

// Rule Group type.
func (o FirewallPolicyRuleCollectionGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyRuleCollectionGroupOutput{})
}

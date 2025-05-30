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

// Network security user rule.
//
// Uses Azure REST API version 2022-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-04-01-preview.
//
// Other available API versions: 2021-02-01-preview, 2022-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type UserRule struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// A description for this rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayOutput `pulumi:"destinationPortRanges"`
	// The destination address prefixes. CIDR or destination IP ranges.
	Destinations AddressPrefixItemResponseArrayOutput `pulumi:"destinations"`
	// Indicates if the traffic matched against the rule in inbound or outbound.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Whether the rule is custom or default.
	// Expected value is 'Custom'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network protocol this rule applies to.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The provisioning state of the security configuration user rule resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayOutput `pulumi:"sourcePortRanges"`
	// The CIDR or source IP ranges.
	Sources AddressPrefixItemResponseArrayOutput `pulumi:"sources"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUserRule registers a new resource with the given unique name, arguments, and options.
func NewUserRule(ctx *pulumi.Context,
	name string, args *UserRuleArgs, opts ...pulumi.ResourceOption) (*UserRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationName'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'RuleCollectionName'")
	}
	args.Kind = pulumi.String("Custom")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210201preview:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240301:SecurityUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240301:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240501:SecurityUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240501:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240701:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network:SecurityUserRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource UserRule
	err := ctx.RegisterResource("azure-native:network:UserRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserRule gets an existing UserRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserRuleState, opts ...pulumi.ResourceOption) (*UserRule, error) {
	var resource UserRule
	err := ctx.ReadResource("azure-native:network:UserRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserRule resources.
type userRuleState struct {
}

type UserRuleState struct {
}

func (UserRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*userRuleState)(nil)).Elem()
}

type userRuleArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// A description for this rule.
	Description *string `pulumi:"description"`
	// The destination port ranges.
	DestinationPortRanges []string `pulumi:"destinationPortRanges"`
	// The destination address prefixes. CIDR or destination IP ranges.
	Destinations []AddressPrefixItem `pulumi:"destinations"`
	// Indicates if the traffic matched against the rule in inbound or outbound.
	Direction string `pulumi:"direction"`
	// Whether the rule is custom or default.
	// Expected value is 'Custom'.
	Kind string `pulumi:"kind"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// Network protocol this rule applies to.
	Protocol string `pulumi:"protocol"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	// The name of the rule.
	RuleName *string `pulumi:"ruleName"`
	// The source port ranges.
	SourcePortRanges []string `pulumi:"sourcePortRanges"`
	// The CIDR or source IP ranges.
	Sources []AddressPrefixItem `pulumi:"sources"`
}

// The set of arguments for constructing a UserRule resource.
type UserRuleArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringInput
	// A description for this rule.
	Description pulumi.StringPtrInput
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayInput
	// The destination address prefixes. CIDR or destination IP ranges.
	Destinations AddressPrefixItemArrayInput
	// Indicates if the traffic matched against the rule in inbound or outbound.
	Direction pulumi.StringInput
	// Whether the rule is custom or default.
	// Expected value is 'Custom'.
	Kind pulumi.StringInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// Network protocol this rule applies to.
	Protocol pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName pulumi.StringInput
	// The name of the rule.
	RuleName pulumi.StringPtrInput
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayInput
	// The CIDR or source IP ranges.
	Sources AddressPrefixItemArrayInput
}

func (UserRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userRuleArgs)(nil)).Elem()
}

type UserRuleInput interface {
	pulumi.Input

	ToUserRuleOutput() UserRuleOutput
	ToUserRuleOutputWithContext(ctx context.Context) UserRuleOutput
}

func (*UserRule) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRule)(nil)).Elem()
}

func (i *UserRule) ToUserRuleOutput() UserRuleOutput {
	return i.ToUserRuleOutputWithContext(context.Background())
}

func (i *UserRule) ToUserRuleOutputWithContext(ctx context.Context) UserRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRuleOutput)
}

type UserRuleOutput struct{ *pulumi.OutputState }

func (UserRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserRule)(nil)).Elem()
}

func (o UserRuleOutput) ToUserRuleOutput() UserRuleOutput {
	return o
}

func (o UserRuleOutput) ToUserRuleOutputWithContext(ctx context.Context) UserRuleOutput {
	return o
}

// The Azure API version of the resource.
func (o UserRuleOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// A description for this rule.
func (o UserRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination port ranges.
func (o UserRuleOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringArrayOutput { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

// The destination address prefixes. CIDR or destination IP ranges.
func (o UserRuleOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v *UserRule) AddressPrefixItemResponseArrayOutput { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

// Indicates if the traffic matched against the rule in inbound or outbound.
func (o UserRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o UserRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Whether the rule is custom or default.
// Expected value is 'Custom'.
func (o UserRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource name.
func (o UserRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network protocol this rule applies to.
func (o UserRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The provisioning state of the security configuration user rule resource.
func (o UserRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The source port ranges.
func (o UserRuleOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringArrayOutput { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

// The CIDR or source IP ranges.
func (o UserRuleOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v *UserRule) AddressPrefixItemResponseArrayOutput { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
}

// The system metadata related to this resource.
func (o UserRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *UserRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o UserRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UserRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UserRuleOutput{})
}

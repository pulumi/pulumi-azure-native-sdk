// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network security rule.
type SecurityRule struct {
	pulumi.CustomResourceState

	// The network traffic is allowed or denied.
	Access pulumi.StringOutput `pulumi:"access"`
	// A description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix pulumi.StringPtrOutput `pulumi:"destinationAddressPrefix"`
	// The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes pulumi.StringArrayOutput `pulumi:"destinationAddressPrefixes"`
	// The application security group specified as destination.
	DestinationApplicationSecurityGroups ApplicationSecurityGroupResponseArrayOutput `pulumi:"destinationApplicationSecurityGroups"`
	// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange pulumi.StringPtrOutput `pulumi:"destinationPortRange"`
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayOutput `pulumi:"destinationPortRanges"`
	// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', 'Icmp', 'Esp', and '*'.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix pulumi.StringPtrOutput `pulumi:"sourceAddressPrefix"`
	// The CIDR or source IP ranges.
	SourceAddressPrefixes pulumi.StringArrayOutput `pulumi:"sourceAddressPrefixes"`
	// The application security group specified as source.
	SourceApplicationSecurityGroups ApplicationSecurityGroupResponseArrayOutput `pulumi:"sourceApplicationSecurityGroups"`
	// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange pulumi.StringPtrOutput `pulumi:"sourcePortRange"`
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayOutput `pulumi:"sourcePortRanges"`
}

// NewSecurityRule registers a new resource with the given unique name, arguments, and options.
func NewSecurityRule(ctx *pulumi.Context,
	name string, args *SecurityRuleArgs, opts ...pulumi.ResourceOption) (*SecurityRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Access == nil {
		return nil, errors.New("invalid value for required argument 'Access'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.NetworkSecurityGroupName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityGroupName'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:SecurityRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:SecurityRule"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityRule
	err := ctx.RegisterResource("azure-native:network/v20190201:SecurityRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityRule gets an existing SecurityRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityRuleState, opts ...pulumi.ResourceOption) (*SecurityRule, error) {
	var resource SecurityRule
	err := ctx.ReadResource("azure-native:network/v20190201:SecurityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityRule resources.
type securityRuleState struct {
}

type SecurityRuleState struct {
}

func (SecurityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityRuleState)(nil)).Elem()
}

type securityRuleArgs struct {
	// The network traffic is allowed or denied.
	Access string `pulumi:"access"`
	// A description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	// The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes []string `pulumi:"destinationAddressPrefixes"`
	// The application security group specified as destination.
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupType `pulumi:"destinationApplicationSecurityGroups"`
	// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// The destination port ranges.
	DestinationPortRanges []string `pulumi:"destinationPortRanges"`
	// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction string `pulumi:"direction"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The name of the network security group.
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *int `pulumi:"priority"`
	// Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', 'Icmp', 'Esp', and '*'.
	Protocol string `pulumi:"protocol"`
	// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the security rule.
	SecurityRuleName *string `pulumi:"securityRuleName"`
	// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix *string `pulumi:"sourceAddressPrefix"`
	// The CIDR or source IP ranges.
	SourceAddressPrefixes []string `pulumi:"sourceAddressPrefixes"`
	// The application security group specified as source.
	SourceApplicationSecurityGroups []ApplicationSecurityGroupType `pulumi:"sourceApplicationSecurityGroups"`
	// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// The source port ranges.
	SourcePortRanges []string `pulumi:"sourcePortRanges"`
}

// The set of arguments for constructing a SecurityRule resource.
type SecurityRuleArgs struct {
	// The network traffic is allowed or denied.
	Access pulumi.StringInput
	// A description for this rule. Restricted to 140 chars.
	Description pulumi.StringPtrInput
	// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix pulumi.StringPtrInput
	// The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes pulumi.StringArrayInput
	// The application security group specified as destination.
	DestinationApplicationSecurityGroups ApplicationSecurityGroupTypeArrayInput
	// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange pulumi.StringPtrInput
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayInput
	// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The name of the network security group.
	NetworkSecurityGroupName pulumi.StringInput
	// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntPtrInput
	// Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', 'Icmp', 'Esp', and '*'.
	Protocol pulumi.StringInput
	// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the security rule.
	SecurityRuleName pulumi.StringPtrInput
	// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix pulumi.StringPtrInput
	// The CIDR or source IP ranges.
	SourceAddressPrefixes pulumi.StringArrayInput
	// The application security group specified as source.
	SourceApplicationSecurityGroups ApplicationSecurityGroupTypeArrayInput
	// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange pulumi.StringPtrInput
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayInput
}

func (SecurityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityRuleArgs)(nil)).Elem()
}

type SecurityRuleInput interface {
	pulumi.Input

	ToSecurityRuleOutput() SecurityRuleOutput
	ToSecurityRuleOutputWithContext(ctx context.Context) SecurityRuleOutput
}

func (*SecurityRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityRule)(nil)).Elem()
}

func (i *SecurityRule) ToSecurityRuleOutput() SecurityRuleOutput {
	return i.ToSecurityRuleOutputWithContext(context.Background())
}

func (i *SecurityRule) ToSecurityRuleOutputWithContext(ctx context.Context) SecurityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityRuleOutput)
}

type SecurityRuleOutput struct{ *pulumi.OutputState }

func (SecurityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityRule)(nil)).Elem()
}

func (o SecurityRuleOutput) ToSecurityRuleOutput() SecurityRuleOutput {
	return o
}

func (o SecurityRuleOutput) ToSecurityRuleOutputWithContext(ctx context.Context) SecurityRuleOutput {
	return o
}

// The network traffic is allowed or denied.
func (o SecurityRuleOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringOutput { return v.Access }).(pulumi.StringOutput)
}

// A description for this rule. Restricted to 140 chars.
func (o SecurityRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
func (o SecurityRuleOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

// The destination address prefixes. CIDR or destination IP ranges.
func (o SecurityRuleOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringArrayOutput { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

// The application security group specified as destination.
func (o SecurityRuleOutput) DestinationApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v *SecurityRule) ApplicationSecurityGroupResponseArrayOutput {
		return v.DestinationApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
func (o SecurityRuleOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

// The destination port ranges.
func (o SecurityRuleOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringArrayOutput { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
func (o SecurityRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o SecurityRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o SecurityRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
func (o SecurityRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', 'Icmp', 'Esp', and '*'.
func (o SecurityRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
func (o SecurityRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
func (o SecurityRuleOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

// The CIDR or source IP ranges.
func (o SecurityRuleOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringArrayOutput { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

// The application security group specified as source.
func (o SecurityRuleOutput) SourceApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v *SecurityRule) ApplicationSecurityGroupResponseArrayOutput {
		return v.SourceApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
func (o SecurityRuleOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringPtrOutput { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

// The source port ranges.
func (o SecurityRuleOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityRule) pulumi.StringArrayOutput { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityRuleOutput{})
}
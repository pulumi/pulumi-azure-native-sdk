// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210827

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a database script.
type Script struct {
	pulumi.CustomResourceState

	// Flag that indicates whether to continue if one of the command fails.
	ContinueOnErrors pulumi.BoolPtrOutput `pulumi:"continueOnErrors"`
	// A unique string. If changed the script will be applied again.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The url to the KQL script blob file.
	ScriptUrl pulumi.StringOutput `pulumi:"scriptUrl"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewScript registers a new resource with the given unique name, arguments, and options.
func NewScript(ctx *pulumi.Context,
	name string, args *ScriptArgs, opts ...pulumi.ResourceOption) (*Script, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScriptUrl == nil {
		return nil, errors.New("invalid value for required argument 'ScriptUrl'")
	}
	if args.ScriptUrlSasToken == nil {
		return nil, errors.New("invalid value for required argument 'ScriptUrlSasToken'")
	}
	if isZero(args.ContinueOnErrors) {
		args.ContinueOnErrors = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:Script"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:Script"),
		},
	})
	opts = append(opts, aliases)
	var resource Script
	err := ctx.RegisterResource("azure-native:kusto/v20210827:Script", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScript gets an existing Script resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScriptState, opts ...pulumi.ResourceOption) (*Script, error) {
	var resource Script
	err := ctx.ReadResource("azure-native:kusto/v20210827:Script", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Script resources.
type scriptState struct {
}

type ScriptState struct {
}

func (ScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptState)(nil)).Elem()
}

type scriptArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// Flag that indicates whether to continue if one of the command fails.
	ContinueOnErrors *bool `pulumi:"continueOnErrors"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// A unique string. If changed the script will be applied again.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Kusto database script.
	ScriptName *string `pulumi:"scriptName"`
	// The url to the KQL script blob file.
	ScriptUrl string `pulumi:"scriptUrl"`
	// The SaS token.
	ScriptUrlSasToken string `pulumi:"scriptUrlSasToken"`
}

// The set of arguments for constructing a Script resource.
type ScriptArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// Flag that indicates whether to continue if one of the command fails.
	ContinueOnErrors pulumi.BoolPtrInput
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput
	// A unique string. If changed the script will be applied again.
	ForceUpdateTag pulumi.StringPtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
	// The name of the Kusto database script.
	ScriptName pulumi.StringPtrInput
	// The url to the KQL script blob file.
	ScriptUrl pulumi.StringInput
	// The SaS token.
	ScriptUrlSasToken pulumi.StringInput
}

func (ScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptArgs)(nil)).Elem()
}

type ScriptInput interface {
	pulumi.Input

	ToScriptOutput() ScriptOutput
	ToScriptOutputWithContext(ctx context.Context) ScriptOutput
}

func (*Script) ElementType() reflect.Type {
	return reflect.TypeOf((**Script)(nil)).Elem()
}

func (i *Script) ToScriptOutput() ScriptOutput {
	return i.ToScriptOutputWithContext(context.Background())
}

func (i *Script) ToScriptOutputWithContext(ctx context.Context) ScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptOutput)
}

type ScriptOutput struct{ *pulumi.OutputState }

func (ScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Script)(nil)).Elem()
}

func (o ScriptOutput) ToScriptOutput() ScriptOutput {
	return o
}

func (o ScriptOutput) ToScriptOutputWithContext(ctx context.Context) ScriptOutput {
	return o
}

// Flag that indicates whether to continue if one of the command fails.
func (o ScriptOutput) ContinueOnErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Script) pulumi.BoolPtrOutput { return v.ContinueOnErrors }).(pulumi.BoolPtrOutput)
}

// A unique string. If changed the script will be applied again.
func (o ScriptOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Script) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o ScriptOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The url to the KQL script blob file.
func (o ScriptOutput) ScriptUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.ScriptUrl }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ScriptOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Script) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScriptOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Script) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScriptOutput{})
}
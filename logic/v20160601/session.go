// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The integration account session.
type Session struct {
	pulumi.CustomResourceState

	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The session content.
	Content pulumi.AnyOutput `pulumi:"content"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSession registers a new resource with the given unique name, arguments, and options.
func NewSession(ctx *pulumi.Context,
	name string, args *SessionArgs, opts ...pulumi.ResourceOption) (*Session, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Session"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Session"),
		},
		{
			Type: pulumi.String("azure-native:logic:Session"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Session
	err := ctx.RegisterResource("azure-native:logic/v20160601:Session", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSession gets an existing Session resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSession(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SessionState, opts ...pulumi.ResourceOption) (*Session, error) {
	var resource Session
	err := ctx.ReadResource("azure-native:logic/v20160601:Session", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Session resources.
type sessionState struct {
}

type SessionState struct {
}

func (SessionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sessionState)(nil)).Elem()
}

type sessionArgs struct {
	// The session content.
	Content interface{} `pulumi:"content"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The integration account session name.
	SessionName *string `pulumi:"sessionName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Session resource.
type SessionArgs struct {
	// The session content.
	Content pulumi.Input
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The integration account session name.
	SessionName pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (SessionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sessionArgs)(nil)).Elem()
}

type SessionInput interface {
	pulumi.Input

	ToSessionOutput() SessionOutput
	ToSessionOutputWithContext(ctx context.Context) SessionOutput
}

func (*Session) ElementType() reflect.Type {
	return reflect.TypeOf((**Session)(nil)).Elem()
}

func (i *Session) ToSessionOutput() SessionOutput {
	return i.ToSessionOutputWithContext(context.Background())
}

func (i *Session) ToSessionOutputWithContext(ctx context.Context) SessionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SessionOutput)
}

type SessionOutput struct{ *pulumi.OutputState }

func (SessionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Session)(nil)).Elem()
}

func (o SessionOutput) ToSessionOutput() SessionOutput {
	return o
}

func (o SessionOutput) ToSessionOutputWithContext(ctx context.Context) SessionOutput {
	return o
}

// The changed time.
func (o SessionOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Session) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

// The session content.
func (o SessionOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v *Session) pulumi.AnyOutput { return v.Content }).(pulumi.AnyOutput)
}

// The created time.
func (o SessionOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Session) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The resource location.
func (o SessionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Session) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o SessionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Session) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource tags.
func (o SessionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Session) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o SessionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Session) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SessionOutput{})
}

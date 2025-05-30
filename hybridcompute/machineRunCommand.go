// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridcompute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Run Command
//
// Uses Azure REST API version 2024-07-31-preview. In version 2.x of the Azure Native provider, it used API version 2023-10-03-preview.
//
// Other available API versions: 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13, 2025-02-19-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcompute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type MachineRunCommand struct {
	pulumi.CustomResourceState

	// Optional. If set to true, provisioning will complete as soon as script starts and will not wait for script to complete.
	AsyncExecution pulumi.BoolPtrOutput `pulumi:"asyncExecution"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	ErrorBlobManagedIdentity RunCommandManagedIdentityResponsePtrOutput `pulumi:"errorBlobManagedIdentity"`
	// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
	ErrorBlobUri pulumi.StringPtrOutput `pulumi:"errorBlobUri"`
	// The machine run command instance view.
	InstanceView MachineRunCommandInstanceViewResponseOutput `pulumi:"instanceView"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	OutputBlobManagedIdentity RunCommandManagedIdentityResponsePtrOutput `pulumi:"outputBlobManagedIdentity"`
	// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
	OutputBlobUri pulumi.StringPtrOutput `pulumi:"outputBlobUri"`
	// The parameters used by the script.
	Parameters RunCommandInputParameterResponseArrayOutput `pulumi:"parameters"`
	// The parameters used by the script.
	ProtectedParameters RunCommandInputParameterResponseArrayOutput `pulumi:"protectedParameters"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Specifies the user account password on the machine when executing the run command.
	RunAsPassword pulumi.StringPtrOutput `pulumi:"runAsPassword"`
	// Specifies the user account on the machine when executing the run command.
	RunAsUser pulumi.StringPtrOutput `pulumi:"runAsUser"`
	// The source of the run command script.
	Source MachineRunCommandScriptSourceResponsePtrOutput `pulumi:"source"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds pulumi.IntPtrOutput `pulumi:"timeoutInSeconds"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMachineRunCommand registers a new resource with the given unique name, arguments, and options.
func NewMachineRunCommand(ctx *pulumi.Context,
	name string, args *MachineRunCommandArgs, opts ...pulumi.ResourceOption) (*MachineRunCommand, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineName == nil {
		return nil, errors.New("invalid value for required argument 'MachineName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.AsyncExecution == nil {
		args.AsyncExecution = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute/v20231003preview:MachineRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240331preview:MachineRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240520preview:MachineRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240731preview:MachineRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240910preview:MachineRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20241110preview:MachineRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20250113:MachineRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20250219preview:MachineRunCommand"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MachineRunCommand
	err := ctx.RegisterResource("azure-native:hybridcompute:MachineRunCommand", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineRunCommand gets an existing MachineRunCommand resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineRunCommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineRunCommandState, opts ...pulumi.ResourceOption) (*MachineRunCommand, error) {
	var resource MachineRunCommand
	err := ctx.ReadResource("azure-native:hybridcompute:MachineRunCommand", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineRunCommand resources.
type machineRunCommandState struct {
}

type MachineRunCommandState struct {
}

func (MachineRunCommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineRunCommandState)(nil)).Elem()
}

type machineRunCommandArgs struct {
	// Optional. If set to true, provisioning will complete as soon as script starts and will not wait for script to complete.
	AsyncExecution *bool `pulumi:"asyncExecution"`
	// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	ErrorBlobManagedIdentity *RunCommandManagedIdentity `pulumi:"errorBlobManagedIdentity"`
	// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
	ErrorBlobUri *string `pulumi:"errorBlobUri"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the hybrid machine.
	MachineName string `pulumi:"machineName"`
	// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	OutputBlobManagedIdentity *RunCommandManagedIdentity `pulumi:"outputBlobManagedIdentity"`
	// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
	OutputBlobUri *string `pulumi:"outputBlobUri"`
	// The parameters used by the script.
	Parameters []RunCommandInputParameter `pulumi:"parameters"`
	// The parameters used by the script.
	ProtectedParameters []RunCommandInputParameter `pulumi:"protectedParameters"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the user account password on the machine when executing the run command.
	RunAsPassword *string `pulumi:"runAsPassword"`
	// Specifies the user account on the machine when executing the run command.
	RunAsUser *string `pulumi:"runAsUser"`
	// The name of the run command.
	RunCommandName *string `pulumi:"runCommandName"`
	// The source of the run command script.
	Source *MachineRunCommandScriptSource `pulumi:"source"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
}

// The set of arguments for constructing a MachineRunCommand resource.
type MachineRunCommandArgs struct {
	// Optional. If set to true, provisioning will complete as soon as script starts and will not wait for script to complete.
	AsyncExecution pulumi.BoolPtrInput
	// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	ErrorBlobManagedIdentity RunCommandManagedIdentityPtrInput
	// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
	ErrorBlobUri pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the hybrid machine.
	MachineName pulumi.StringInput
	// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	OutputBlobManagedIdentity RunCommandManagedIdentityPtrInput
	// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
	OutputBlobUri pulumi.StringPtrInput
	// The parameters used by the script.
	Parameters RunCommandInputParameterArrayInput
	// The parameters used by the script.
	ProtectedParameters RunCommandInputParameterArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Specifies the user account password on the machine when executing the run command.
	RunAsPassword pulumi.StringPtrInput
	// Specifies the user account on the machine when executing the run command.
	RunAsUser pulumi.StringPtrInput
	// The name of the run command.
	RunCommandName pulumi.StringPtrInput
	// The source of the run command script.
	Source MachineRunCommandScriptSourcePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds pulumi.IntPtrInput
}

func (MachineRunCommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineRunCommandArgs)(nil)).Elem()
}

type MachineRunCommandInput interface {
	pulumi.Input

	ToMachineRunCommandOutput() MachineRunCommandOutput
	ToMachineRunCommandOutputWithContext(ctx context.Context) MachineRunCommandOutput
}

func (*MachineRunCommand) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineRunCommand)(nil)).Elem()
}

func (i *MachineRunCommand) ToMachineRunCommandOutput() MachineRunCommandOutput {
	return i.ToMachineRunCommandOutputWithContext(context.Background())
}

func (i *MachineRunCommand) ToMachineRunCommandOutputWithContext(ctx context.Context) MachineRunCommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineRunCommandOutput)
}

type MachineRunCommandOutput struct{ *pulumi.OutputState }

func (MachineRunCommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineRunCommand)(nil)).Elem()
}

func (o MachineRunCommandOutput) ToMachineRunCommandOutput() MachineRunCommandOutput {
	return o
}

func (o MachineRunCommandOutput) ToMachineRunCommandOutputWithContext(ctx context.Context) MachineRunCommandOutput {
	return o
}

// Optional. If set to true, provisioning will complete as soon as script starts and will not wait for script to complete.
func (o MachineRunCommandOutput) AsyncExecution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.BoolPtrOutput { return v.AsyncExecution }).(pulumi.BoolPtrOutput)
}

// The Azure API version of the resource.
func (o MachineRunCommandOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
func (o MachineRunCommandOutput) ErrorBlobManagedIdentity() RunCommandManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *MachineRunCommand) RunCommandManagedIdentityResponsePtrOutput {
		return v.ErrorBlobManagedIdentity
	}).(RunCommandManagedIdentityResponsePtrOutput)
}

// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
func (o MachineRunCommandOutput) ErrorBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.StringPtrOutput { return v.ErrorBlobUri }).(pulumi.StringPtrOutput)
}

// The machine run command instance view.
func (o MachineRunCommandOutput) InstanceView() MachineRunCommandInstanceViewResponseOutput {
	return o.ApplyT(func(v *MachineRunCommand) MachineRunCommandInstanceViewResponseOutput { return v.InstanceView }).(MachineRunCommandInstanceViewResponseOutput)
}

// The geo-location where the resource lives
func (o MachineRunCommandOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o MachineRunCommandOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
func (o MachineRunCommandOutput) OutputBlobManagedIdentity() RunCommandManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *MachineRunCommand) RunCommandManagedIdentityResponsePtrOutput {
		return v.OutputBlobManagedIdentity
	}).(RunCommandManagedIdentityResponsePtrOutput)
}

// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
func (o MachineRunCommandOutput) OutputBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.StringPtrOutput { return v.OutputBlobUri }).(pulumi.StringPtrOutput)
}

// The parameters used by the script.
func (o MachineRunCommandOutput) Parameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v *MachineRunCommand) RunCommandInputParameterResponseArrayOutput { return v.Parameters }).(RunCommandInputParameterResponseArrayOutput)
}

// The parameters used by the script.
func (o MachineRunCommandOutput) ProtectedParameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v *MachineRunCommand) RunCommandInputParameterResponseArrayOutput { return v.ProtectedParameters }).(RunCommandInputParameterResponseArrayOutput)
}

// The provisioning state, which only appears in the response.
func (o MachineRunCommandOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies the user account password on the machine when executing the run command.
func (o MachineRunCommandOutput) RunAsPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.StringPtrOutput { return v.RunAsPassword }).(pulumi.StringPtrOutput)
}

// Specifies the user account on the machine when executing the run command.
func (o MachineRunCommandOutput) RunAsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.StringPtrOutput { return v.RunAsUser }).(pulumi.StringPtrOutput)
}

// The source of the run command script.
func (o MachineRunCommandOutput) Source() MachineRunCommandScriptSourceResponsePtrOutput {
	return o.ApplyT(func(v *MachineRunCommand) MachineRunCommandScriptSourceResponsePtrOutput { return v.Source }).(MachineRunCommandScriptSourceResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MachineRunCommandOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MachineRunCommand) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MachineRunCommandOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The timeout in seconds to execute the run command.
func (o MachineRunCommandOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.IntPtrOutput { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MachineRunCommandOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineRunCommand) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineRunCommandOutput{})
}

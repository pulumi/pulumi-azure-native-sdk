// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The operation to get the VMSS VM run command.
//
// Uses Azure REST API version 2024-11-01.
//
// Other available API versions: 2022-08-01, 2022-11-01, 2023-03-01, 2023-07-01, 2023-09-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native compute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupVirtualMachineScaleSetVMRunCommand(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetVMRunCommandArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetVMRunCommandResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualMachineScaleSetVMRunCommandResult
	err := ctx.Invoke("azure-native:compute:getVirtualMachineScaleSetVMRunCommand", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualMachineScaleSetVMRunCommandArgs struct {
	// The expand expression to apply on the operation.
	Expand *string `pulumi:"expand"`
	// The instance ID of the virtual machine.
	InstanceId string `pulumi:"instanceId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine run command.
	RunCommandName string `pulumi:"runCommandName"`
	// The name of the VM scale set.
	VmScaleSetName string `pulumi:"vmScaleSetName"`
}

// Describes a Virtual Machine run command.
type LookupVirtualMachineScaleSetVMRunCommandResult struct {
	// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
	AsyncExecution *bool `pulumi:"asyncExecution"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	ErrorBlobManagedIdentity *RunCommandManagedIdentityResponse `pulumi:"errorBlobManagedIdentity"`
	// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
	ErrorBlobUri *string `pulumi:"errorBlobUri"`
	// Resource Id
	Id string `pulumi:"id"`
	// The virtual machine run command instance view.
	InstanceView VirtualMachineRunCommandInstanceViewResponse `pulumi:"instanceView"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	OutputBlobManagedIdentity *RunCommandManagedIdentityResponse `pulumi:"outputBlobManagedIdentity"`
	// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
	OutputBlobUri *string `pulumi:"outputBlobUri"`
	// The parameters used by the script.
	Parameters []RunCommandInputParameterResponse `pulumi:"parameters"`
	// The parameters used by the script.
	ProtectedParameters []RunCommandInputParameterResponse `pulumi:"protectedParameters"`
	// The provisioning state, which only appears in the response. If treatFailureAsDeploymentFailure set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If treatFailureAsDeploymentFailure set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
	ProvisioningState string `pulumi:"provisioningState"`
	// Specifies the user account password on the VM when executing the run command.
	RunAsPassword *string `pulumi:"runAsPassword"`
	// Specifies the user account on the VM when executing the run command.
	RunAsUser *string `pulumi:"runAsUser"`
	// The source of the run command script.
	Source *VirtualMachineRunCommandScriptSourceResponse `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
	// Optional. If set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
	TreatFailureAsDeploymentFailure *bool `pulumi:"treatFailureAsDeploymentFailure"`
	// Resource type
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupVirtualMachineScaleSetVMRunCommandResult
func (val *LookupVirtualMachineScaleSetVMRunCommandResult) Defaults() *LookupVirtualMachineScaleSetVMRunCommandResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AsyncExecution == nil {
		asyncExecution_ := false
		tmp.AsyncExecution = &asyncExecution_
	}
	if tmp.TreatFailureAsDeploymentFailure == nil {
		treatFailureAsDeploymentFailure_ := false
		tmp.TreatFailureAsDeploymentFailure = &treatFailureAsDeploymentFailure_
	}
	return &tmp
}
func LookupVirtualMachineScaleSetVMRunCommandOutput(ctx *pulumi.Context, args LookupVirtualMachineScaleSetVMRunCommandOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineScaleSetVMRunCommandResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineScaleSetVMRunCommandResultOutput, error) {
			args := v.(LookupVirtualMachineScaleSetVMRunCommandArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:compute:getVirtualMachineScaleSetVMRunCommand", args, LookupVirtualMachineScaleSetVMRunCommandResultOutput{}, options).(LookupVirtualMachineScaleSetVMRunCommandResultOutput), nil
		}).(LookupVirtualMachineScaleSetVMRunCommandResultOutput)
}

type LookupVirtualMachineScaleSetVMRunCommandOutputArgs struct {
	// The expand expression to apply on the operation.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The instance ID of the virtual machine.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the virtual machine run command.
	RunCommandName pulumi.StringInput `pulumi:"runCommandName"`
	// The name of the VM scale set.
	VmScaleSetName pulumi.StringInput `pulumi:"vmScaleSetName"`
}

func (LookupVirtualMachineScaleSetVMRunCommandOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetVMRunCommandArgs)(nil)).Elem()
}

// Describes a Virtual Machine run command.
type LookupVirtualMachineScaleSetVMRunCommandResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineScaleSetVMRunCommandResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetVMRunCommandResult)(nil)).Elem()
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ToLookupVirtualMachineScaleSetVMRunCommandResultOutput() LookupVirtualMachineScaleSetVMRunCommandResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ToLookupVirtualMachineScaleSetVMRunCommandResultOutputWithContext(ctx context.Context) LookupVirtualMachineScaleSetVMRunCommandResultOutput {
	return o
}

// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) AsyncExecution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *bool { return v.AsyncExecution }).(pulumi.BoolPtrOutput)
}

// The Azure API version of the resource.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ErrorBlobManagedIdentity() RunCommandManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *RunCommandManagedIdentityResponse {
		return v.ErrorBlobManagedIdentity
	}).(RunCommandManagedIdentityResponsePtrOutput)
}

// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ErrorBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *string { return v.ErrorBlobUri }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.Id }).(pulumi.StringOutput)
}

// The virtual machine run command instance view.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) InstanceView() VirtualMachineRunCommandInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) VirtualMachineRunCommandInstanceViewResponse {
		return v.InstanceView
	}).(VirtualMachineRunCommandInstanceViewResponseOutput)
}

// Resource location
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.Name }).(pulumi.StringOutput)
}

// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) OutputBlobManagedIdentity() RunCommandManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *RunCommandManagedIdentityResponse {
		return v.OutputBlobManagedIdentity
	}).(RunCommandManagedIdentityResponsePtrOutput)
}

// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) OutputBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *string { return v.OutputBlobUri }).(pulumi.StringPtrOutput)
}

// The parameters used by the script.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Parameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) []RunCommandInputParameterResponse {
		return v.Parameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

// The parameters used by the script.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ProtectedParameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) []RunCommandInputParameterResponse {
		return v.ProtectedParameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

// The provisioning state, which only appears in the response. If treatFailureAsDeploymentFailure set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If treatFailureAsDeploymentFailure set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies the user account password on the VM when executing the run command.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) RunAsPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *string { return v.RunAsPassword }).(pulumi.StringPtrOutput)
}

// Specifies the user account on the VM when executing the run command.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) RunAsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *string { return v.RunAsUser }).(pulumi.StringPtrOutput)
}

// The source of the run command script.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Source() VirtualMachineRunCommandScriptSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *VirtualMachineRunCommandScriptSourceResponse {
		return v.Source
	}).(VirtualMachineRunCommandScriptSourceResponsePtrOutput)
}

// Resource tags
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The timeout in seconds to execute the run command.
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *int { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

// Optional. If set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) TreatFailureAsDeploymentFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) *bool { return v.TreatFailureAsDeploymentFailure }).(pulumi.BoolPtrOutput)
}

// Resource type
func (o LookupVirtualMachineScaleSetVMRunCommandResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMRunCommandResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineScaleSetVMRunCommandResultOutput{})
}

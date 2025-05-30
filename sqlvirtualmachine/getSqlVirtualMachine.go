// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlvirtualmachine

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a SQL virtual machine.
//
// Uses Azure REST API version 2023-10-01.
//
// Other available API versions: 2022-02-01, 2022-07-01-preview, 2022-08-01-preview, 2023-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sqlvirtualmachine [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupSqlVirtualMachine(ctx *pulumi.Context, args *LookupSqlVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupSqlVirtualMachineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlVirtualMachineResult
	err := ctx.Invoke("azure-native:sqlvirtualmachine:getSqlVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSqlVirtualMachineArgs struct {
	// The child resources to include in the response.
	Expand *string `pulumi:"expand"`
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SQL virtual machine.
	SqlVirtualMachineName string `pulumi:"sqlVirtualMachineName"`
}

// A SQL virtual machine.
type LookupSqlVirtualMachineResult struct {
	// Additional VM Patching solution enabled on the Virtual Machine
	AdditionalVmPatch string `pulumi:"additionalVmPatch"`
	// SQL best practices Assessment Settings.
	AssessmentSettings *AssessmentSettingsResponse `pulumi:"assessmentSettings"`
	// Auto backup settings for SQL Server.
	AutoBackupSettings *AutoBackupSettingsResponse `pulumi:"autoBackupSettings"`
	// Auto patching settings for applying critical security updates to SQL virtual machine.
	AutoPatchingSettings *AutoPatchingSettingsResponse `pulumi:"autoPatchingSettings"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Enable automatic upgrade of Sql IaaS extension Agent.
	EnableAutomaticUpgrade *bool `pulumi:"enableAutomaticUpgrade"`
	// Resource ID.
	Id string `pulumi:"id"`
	// DO NOT USE. This value will be deprecated. Azure Active Directory identity of the server.
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// Key vault credential settings.
	KeyVaultCredentialSettings *KeyVaultCredentialSettingsResponse `pulumi:"keyVaultCredentialSettings"`
	// SQL IaaS Agent least privilege mode.
	LeastPrivilegeMode *string `pulumi:"leastPrivilegeMode"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Operating System of the current SQL Virtual Machine.
	OsType string `pulumi:"osType"`
	// Provisioning state to track the async operation status.
	ProvisioningState string `pulumi:"provisioningState"`
	// SQL Server configuration management settings.
	ServerConfigurationsManagementSettings *ServerConfigurationsManagementSettingsResponse `pulumi:"serverConfigurationsManagementSettings"`
	// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
	SqlImageOffer *string `pulumi:"sqlImageOffer"`
	// SQL Server edition type.
	SqlImageSku *string `pulumi:"sqlImageSku"`
	// SQL Server Management type. NOTE: This parameter is not used anymore. API will automatically detect the Sql Management, refrain from using it.
	SqlManagement *string `pulumi:"sqlManagement"`
	// SQL Server license type.
	SqlServerLicenseType *string `pulumi:"sqlServerLicenseType"`
	// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
	SqlVirtualMachineGroupResourceId *string `pulumi:"sqlVirtualMachineGroupResourceId"`
	// Storage Configuration Settings.
	StorageConfigurationSettings *StorageConfigurationSettingsResponse `pulumi:"storageConfigurationSettings"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Troubleshooting status
	TroubleshootingStatus TroubleshootingStatusResponse `pulumi:"troubleshootingStatus"`
	// Resource type.
	Type string `pulumi:"type"`
	// Virtual Machine Identity details used for Sql IaaS extension configurations.
	VirtualMachineIdentitySettings *VirtualMachineIdentityResponse `pulumi:"virtualMachineIdentitySettings"`
	// ARM Resource id of underlying virtual machine created from SQL marketplace image.
	VirtualMachineResourceId *string `pulumi:"virtualMachineResourceId"`
	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcDomainCredentials *WsfcDomainCredentialsResponse `pulumi:"wsfcDomainCredentials"`
	// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
	WsfcStaticIp *string `pulumi:"wsfcStaticIp"`
}

// Defaults sets the appropriate defaults for LookupSqlVirtualMachineResult
func (val *LookupSqlVirtualMachineResult) Defaults() *LookupSqlVirtualMachineResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.AutoPatchingSettings = tmp.AutoPatchingSettings.Defaults()

	if tmp.EnableAutomaticUpgrade == nil {
		enableAutomaticUpgrade_ := false
		tmp.EnableAutomaticUpgrade = &enableAutomaticUpgrade_
	}
	if tmp.LeastPrivilegeMode == nil {
		leastPrivilegeMode_ := "NotSet"
		tmp.LeastPrivilegeMode = &leastPrivilegeMode_
	}
	tmp.TroubleshootingStatus = *tmp.TroubleshootingStatus.Defaults()

	return &tmp
}
func LookupSqlVirtualMachineOutput(ctx *pulumi.Context, args LookupSqlVirtualMachineOutputArgs, opts ...pulumi.InvokeOption) LookupSqlVirtualMachineResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSqlVirtualMachineResultOutput, error) {
			args := v.(LookupSqlVirtualMachineArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:sqlvirtualmachine:getSqlVirtualMachine", args, LookupSqlVirtualMachineResultOutput{}, options).(LookupSqlVirtualMachineResultOutput), nil
		}).(LookupSqlVirtualMachineResultOutput)
}

type LookupSqlVirtualMachineOutputArgs struct {
	// The child resources to include in the response.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the SQL virtual machine.
	SqlVirtualMachineName pulumi.StringInput `pulumi:"sqlVirtualMachineName"`
}

func (LookupSqlVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlVirtualMachineArgs)(nil)).Elem()
}

// A SQL virtual machine.
type LookupSqlVirtualMachineResultOutput struct{ *pulumi.OutputState }

func (LookupSqlVirtualMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlVirtualMachineResult)(nil)).Elem()
}

func (o LookupSqlVirtualMachineResultOutput) ToLookupSqlVirtualMachineResultOutput() LookupSqlVirtualMachineResultOutput {
	return o
}

func (o LookupSqlVirtualMachineResultOutput) ToLookupSqlVirtualMachineResultOutputWithContext(ctx context.Context) LookupSqlVirtualMachineResultOutput {
	return o
}

// Additional VM Patching solution enabled on the Virtual Machine
func (o LookupSqlVirtualMachineResultOutput) AdditionalVmPatch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.AdditionalVmPatch }).(pulumi.StringOutput)
}

// SQL best practices Assessment Settings.
func (o LookupSqlVirtualMachineResultOutput) AssessmentSettings() AssessmentSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *AssessmentSettingsResponse { return v.AssessmentSettings }).(AssessmentSettingsResponsePtrOutput)
}

// Auto backup settings for SQL Server.
func (o LookupSqlVirtualMachineResultOutput) AutoBackupSettings() AutoBackupSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *AutoBackupSettingsResponse { return v.AutoBackupSettings }).(AutoBackupSettingsResponsePtrOutput)
}

// Auto patching settings for applying critical security updates to SQL virtual machine.
func (o LookupSqlVirtualMachineResultOutput) AutoPatchingSettings() AutoPatchingSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *AutoPatchingSettingsResponse { return v.AutoPatchingSettings }).(AutoPatchingSettingsResponsePtrOutput)
}

// The Azure API version of the resource.
func (o LookupSqlVirtualMachineResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Enable automatic upgrade of Sql IaaS extension Agent.
func (o LookupSqlVirtualMachineResultOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

// Resource ID.
func (o LookupSqlVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

// DO NOT USE. This value will be deprecated. Azure Active Directory identity of the server.
func (o LookupSqlVirtualMachineResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Key vault credential settings.
func (o LookupSqlVirtualMachineResultOutput) KeyVaultCredentialSettings() KeyVaultCredentialSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *KeyVaultCredentialSettingsResponse {
		return v.KeyVaultCredentialSettings
	}).(KeyVaultCredentialSettingsResponsePtrOutput)
}

// SQL IaaS Agent least privilege mode.
func (o LookupSqlVirtualMachineResultOutput) LeastPrivilegeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.LeastPrivilegeMode }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupSqlVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupSqlVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

// Operating System of the current SQL Virtual Machine.
func (o LookupSqlVirtualMachineResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.OsType }).(pulumi.StringOutput)
}

// Provisioning state to track the async operation status.
func (o LookupSqlVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SQL Server configuration management settings.
func (o LookupSqlVirtualMachineResultOutput) ServerConfigurationsManagementSettings() ServerConfigurationsManagementSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *ServerConfigurationsManagementSettingsResponse {
		return v.ServerConfigurationsManagementSettings
	}).(ServerConfigurationsManagementSettingsResponsePtrOutput)
}

// SQL image offer. Examples include SQL2016-WS2016, SQL2017-WS2016.
func (o LookupSqlVirtualMachineResultOutput) SqlImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.SqlImageOffer }).(pulumi.StringPtrOutput)
}

// SQL Server edition type.
func (o LookupSqlVirtualMachineResultOutput) SqlImageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.SqlImageSku }).(pulumi.StringPtrOutput)
}

// SQL Server Management type. NOTE: This parameter is not used anymore. API will automatically detect the Sql Management, refrain from using it.
func (o LookupSqlVirtualMachineResultOutput) SqlManagement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.SqlManagement }).(pulumi.StringPtrOutput)
}

// SQL Server license type.
func (o LookupSqlVirtualMachineResultOutput) SqlServerLicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.SqlServerLicenseType }).(pulumi.StringPtrOutput)
}

// ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.
func (o LookupSqlVirtualMachineResultOutput) SqlVirtualMachineGroupResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.SqlVirtualMachineGroupResourceId }).(pulumi.StringPtrOutput)
}

// Storage Configuration Settings.
func (o LookupSqlVirtualMachineResultOutput) StorageConfigurationSettings() StorageConfigurationSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *StorageConfigurationSettingsResponse {
		return v.StorageConfigurationSettings
	}).(StorageConfigurationSettingsResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupSqlVirtualMachineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSqlVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Troubleshooting status
func (o LookupSqlVirtualMachineResultOutput) TroubleshootingStatus() TroubleshootingStatusResponseOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) TroubleshootingStatusResponse { return v.TroubleshootingStatus }).(TroubleshootingStatusResponseOutput)
}

// Resource type.
func (o LookupSqlVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

// Virtual Machine Identity details used for Sql IaaS extension configurations.
func (o LookupSqlVirtualMachineResultOutput) VirtualMachineIdentitySettings() VirtualMachineIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *VirtualMachineIdentityResponse {
		return v.VirtualMachineIdentitySettings
	}).(VirtualMachineIdentityResponsePtrOutput)
}

// ARM Resource id of underlying virtual machine created from SQL marketplace image.
func (o LookupSqlVirtualMachineResultOutput) VirtualMachineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.VirtualMachineResourceId }).(pulumi.StringPtrOutput)
}

// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
func (o LookupSqlVirtualMachineResultOutput) WsfcDomainCredentials() WsfcDomainCredentialsResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *WsfcDomainCredentialsResponse { return v.WsfcDomainCredentials }).(WsfcDomainCredentialsResponsePtrOutput)
}

// Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.
func (o LookupSqlVirtualMachineResultOutput) WsfcStaticIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlVirtualMachineResult) *string { return v.WsfcStaticIp }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlVirtualMachineResultOutput{})
}

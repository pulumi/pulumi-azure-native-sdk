// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

// Rest API method for target endpoint.
type AzureFunctionActivityMethod string

const (
	AzureFunctionActivityMethodGET     = AzureFunctionActivityMethod("GET")
	AzureFunctionActivityMethodPOST    = AzureFunctionActivityMethod("POST")
	AzureFunctionActivityMethodPUT     = AzureFunctionActivityMethod("PUT")
	AzureFunctionActivityMethodDELETE  = AzureFunctionActivityMethod("DELETE")
	AzureFunctionActivityMethodOPTIONS = AzureFunctionActivityMethod("OPTIONS")
	AzureFunctionActivityMethodHEAD    = AzureFunctionActivityMethod("HEAD")
	AzureFunctionActivityMethodTRACE   = AzureFunctionActivityMethod("TRACE")
)

// Specify the write behavior when upserting documents into Azure Search Index.
type AzureSearchIndexWriteBehaviorType string

const (
	AzureSearchIndexWriteBehaviorTypeMerge  = AzureSearchIndexWriteBehaviorType("Merge")
	AzureSearchIndexWriteBehaviorTypeUpload = AzureSearchIndexWriteBehaviorType("Upload")
)

// Big data pool reference type.
type BigDataPoolReferenceType string

const (
	BigDataPoolReferenceTypeBigDataPoolReference = BigDataPoolReferenceType("BigDataPoolReference")
)

type BlobEventTypes string

const (
	BlobEventTypes_Microsoft_Storage_BlobCreated = BlobEventTypes("Microsoft.Storage.BlobCreated")
	BlobEventTypes_Microsoft_Storage_BlobDeleted = BlobEventTypes("Microsoft.Storage.BlobDeleted")
)

// The consistency level specifies how many Cassandra servers must respond to a read request before returning data to the client application. Cassandra checks the specified number of Cassandra servers for data to satisfy the read request. Must be one of cassandraSourceReadConsistencyLevels. The default value is 'ONE'. It is case-insensitive.
type CassandraSourceReadConsistencyLevels string

const (
	CassandraSourceReadConsistencyLevelsALL           = CassandraSourceReadConsistencyLevels("ALL")
	CassandraSourceReadConsistencyLevels_EACH_QUORUM  = CassandraSourceReadConsistencyLevels("EACH_QUORUM")
	CassandraSourceReadConsistencyLevelsQUORUM        = CassandraSourceReadConsistencyLevels("QUORUM")
	CassandraSourceReadConsistencyLevels_LOCAL_QUORUM = CassandraSourceReadConsistencyLevels("LOCAL_QUORUM")
	CassandraSourceReadConsistencyLevelsONE           = CassandraSourceReadConsistencyLevels("ONE")
	CassandraSourceReadConsistencyLevelsTWO           = CassandraSourceReadConsistencyLevels("TWO")
	CassandraSourceReadConsistencyLevelsTHREE         = CassandraSourceReadConsistencyLevels("THREE")
	CassandraSourceReadConsistencyLevels_LOCAL_ONE    = CassandraSourceReadConsistencyLevels("LOCAL_ONE")
	CassandraSourceReadConsistencyLevelsSERIAL        = CassandraSourceReadConsistencyLevels("SERIAL")
	CassandraSourceReadConsistencyLevels_LOCAL_SERIAL = CassandraSourceReadConsistencyLevels("LOCAL_SERIAL")
)

// The connection mode used to access CosmosDB account. Type: string (or Expression with resultType string).
type CosmosDbConnectionMode string

const (
	CosmosDbConnectionModeGateway = CosmosDbConnectionMode("Gateway")
	CosmosDbConnectionModeDirect  = CosmosDbConnectionMode("Direct")
)

// The service principal credential type to use in Server-To-Server authentication. 'ServicePrincipalKey' for key/secret, 'ServicePrincipalCert' for certificate. Type: string (or Expression with resultType string).
type CosmosDbServicePrincipalCredentialType string

const (
	CosmosDbServicePrincipalCredentialTypeServicePrincipalKey  = CosmosDbServicePrincipalCredentialType("ServicePrincipalKey")
	CosmosDbServicePrincipalCredentialTypeServicePrincipalCert = CosmosDbServicePrincipalCredentialType("ServicePrincipalCert")
)

// Credential reference type.
type CredentialReferenceType string

const (
	CredentialReferenceTypeCredentialReference = CredentialReferenceType("CredentialReference")
)

// Compute type of the cluster which will execute data flow job.
type DataFlowComputeType string

const (
	DataFlowComputeTypeGeneral          = DataFlowComputeType("General")
	DataFlowComputeTypeMemoryOptimized  = DataFlowComputeType("MemoryOptimized")
	DataFlowComputeTypeComputeOptimized = DataFlowComputeType("ComputeOptimized")
)

// Data flow reference type.
type DataFlowReferenceType string

const (
	DataFlowReferenceTypeDataFlowReference = DataFlowReferenceType("DataFlowReference")
)

// The day of the week.
type DayOfWeek string

const (
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
)

type DaysOfWeek string

const (
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
)

// AuthenticationType to be used for connection. It is mutually exclusive with connectionString property.
type Db2AuthenticationType string

const (
	Db2AuthenticationTypeBasic = Db2AuthenticationType("Basic")
)

type DependencyCondition string

const (
	DependencyConditionSucceeded = DependencyCondition("Succeeded")
	DependencyConditionFailed    = DependencyCondition("Failed")
	DependencyConditionSkipped   = DependencyCondition("Skipped")
	DependencyConditionCompleted = DependencyCondition("Completed")
)

// The write behavior for the operation.
type DynamicsSinkWriteBehavior string

const (
	DynamicsSinkWriteBehaviorUpsert = DynamicsSinkWriteBehavior("Upsert")
)

// The identity type.
type FactoryIdentityType string

const (
	FactoryIdentityTypeSystemAssigned               = FactoryIdentityType("SystemAssigned")
	FactoryIdentityTypeUserAssigned                 = FactoryIdentityType("UserAssigned")
	FactoryIdentityType_SystemAssigned_UserAssigned = FactoryIdentityType("SystemAssigned,UserAssigned")
)

// The authentication type to be used to connect to the FTP server.
type FtpAuthenticationType string

const (
	FtpAuthenticationTypeBasic     = FtpAuthenticationType("Basic")
	FtpAuthenticationTypeAnonymous = FtpAuthenticationType("Anonymous")
)

// Global Parameter type.
type GlobalParameterType string

const (
	GlobalParameterTypeObject = GlobalParameterType("Object")
	GlobalParameterTypeString = GlobalParameterType("String")
	GlobalParameterTypeInt    = GlobalParameterType("Int")
	GlobalParameterTypeFloat  = GlobalParameterType("Float")
	GlobalParameterTypeBool   = GlobalParameterType("Bool")
	GlobalParameterTypeArray  = GlobalParameterType("Array")
)

// The OAuth 2.0 authentication mechanism used for authentication. ServiceAuthentication can only be used on self-hosted IR.
type GoogleAdWordsAuthenticationType string

const (
	GoogleAdWordsAuthenticationTypeServiceAuthentication = GoogleAdWordsAuthenticationType("ServiceAuthentication")
	GoogleAdWordsAuthenticationTypeUserAuthentication    = GoogleAdWordsAuthenticationType("UserAuthentication")
)

// The OAuth 2.0 authentication mechanism used for authentication. ServiceAuthentication can only be used on self-hosted IR.
type GoogleBigQueryAuthenticationType string

const (
	GoogleBigQueryAuthenticationTypeServiceAuthentication = GoogleBigQueryAuthenticationType("ServiceAuthentication")
	GoogleBigQueryAuthenticationTypeUserAuthentication    = GoogleBigQueryAuthenticationType("UserAuthentication")
)

// The authentication mechanism to use to connect to the HBase server.
type HBaseAuthenticationType string

const (
	HBaseAuthenticationTypeAnonymous = HBaseAuthenticationType("Anonymous")
	HBaseAuthenticationTypeBasic     = HBaseAuthenticationType("Basic")
)

// Debug info option.
type HDInsightActivityDebugInfoOption string

const (
	HDInsightActivityDebugInfoOptionNone    = HDInsightActivityDebugInfoOption("None")
	HDInsightActivityDebugInfoOptionAlways  = HDInsightActivityDebugInfoOption("Always")
	HDInsightActivityDebugInfoOptionFailure = HDInsightActivityDebugInfoOption("Failure")
)

// The authentication method used to access the Hive server.
type HiveAuthenticationType string

const (
	HiveAuthenticationTypeAnonymous                    = HiveAuthenticationType("Anonymous")
	HiveAuthenticationTypeUsername                     = HiveAuthenticationType("Username")
	HiveAuthenticationTypeUsernameAndPassword          = HiveAuthenticationType("UsernameAndPassword")
	HiveAuthenticationTypeWindowsAzureHDInsightService = HiveAuthenticationType("WindowsAzureHDInsightService")
)

// The type of Hive server.
type HiveServerType string

const (
	HiveServerTypeHiveServer1      = HiveServerType("HiveServer1")
	HiveServerTypeHiveServer2      = HiveServerType("HiveServer2")
	HiveServerTypeHiveThriftServer = HiveServerType("HiveThriftServer")
)

// The transport protocol to use in the Thrift layer.
type HiveThriftTransportProtocol string

const (
	HiveThriftTransportProtocolBinary = HiveThriftTransportProtocol("Binary")
	HiveThriftTransportProtocolSASL   = HiveThriftTransportProtocol("SASL")
	HiveThriftTransportProtocol_HTTP_ = HiveThriftTransportProtocol("HTTP ")
)

// The authentication type to be used to connect to the HTTP server.
type HttpAuthenticationType string

const (
	HttpAuthenticationTypeBasic             = HttpAuthenticationType("Basic")
	HttpAuthenticationTypeAnonymous         = HttpAuthenticationType("Anonymous")
	HttpAuthenticationTypeDigest            = HttpAuthenticationType("Digest")
	HttpAuthenticationTypeWindows           = HttpAuthenticationType("Windows")
	HttpAuthenticationTypeClientCertificate = HttpAuthenticationType("ClientCertificate")
)

// The authentication type to use.
type ImpalaAuthenticationType string

const (
	ImpalaAuthenticationTypeAnonymous           = ImpalaAuthenticationType("Anonymous")
	ImpalaAuthenticationTypeSASLUsername        = ImpalaAuthenticationType("SASLUsername")
	ImpalaAuthenticationTypeUsernameAndPassword = ImpalaAuthenticationType("UsernameAndPassword")
)

// The edition for the SSIS Integration Runtime
type IntegrationRuntimeEdition string

const (
	IntegrationRuntimeEditionStandard   = IntegrationRuntimeEdition("Standard")
	IntegrationRuntimeEditionEnterprise = IntegrationRuntimeEdition("Enterprise")
)

// The type of this referenced entity.
type IntegrationRuntimeEntityReferenceType string

const (
	IntegrationRuntimeEntityReferenceTypeIntegrationRuntimeReference = IntegrationRuntimeEntityReferenceType("IntegrationRuntimeReference")
	IntegrationRuntimeEntityReferenceTypeLinkedServiceReference      = IntegrationRuntimeEntityReferenceType("LinkedServiceReference")
)

// License type for bringing your own license scenario.
type IntegrationRuntimeLicenseType string

const (
	IntegrationRuntimeLicenseTypeBasePrice       = IntegrationRuntimeLicenseType("BasePrice")
	IntegrationRuntimeLicenseTypeLicenseIncluded = IntegrationRuntimeLicenseType("LicenseIncluded")
)

// The pricing tier for the catalog database. The valid values could be found in https://azure.microsoft.com/en-us/pricing/details/sql-database/
type IntegrationRuntimeSsisCatalogPricingTier string

const (
	IntegrationRuntimeSsisCatalogPricingTierBasic     = IntegrationRuntimeSsisCatalogPricingTier("Basic")
	IntegrationRuntimeSsisCatalogPricingTierStandard  = IntegrationRuntimeSsisCatalogPricingTier("Standard")
	IntegrationRuntimeSsisCatalogPricingTierPremium   = IntegrationRuntimeSsisCatalogPricingTier("Premium")
	IntegrationRuntimeSsisCatalogPricingTierPremiumRS = IntegrationRuntimeSsisCatalogPricingTier("PremiumRS")
)

// Type of integration runtime.
type IntegrationRuntimeType string

const (
	IntegrationRuntimeTypeManaged    = IntegrationRuntimeType("Managed")
	IntegrationRuntimeTypeSelfHosted = IntegrationRuntimeType("SelfHosted")
)

// Managed Virtual Network reference type.
type ManagedVirtualNetworkReferenceType string

const (
	ManagedVirtualNetworkReferenceTypeManagedVirtualNetworkReference = ManagedVirtualNetworkReferenceType("ManagedVirtualNetworkReference")
)

// The authentication type to be used to connect to the MongoDB database.
type MongoDbAuthenticationType string

const (
	MongoDbAuthenticationTypeBasic     = MongoDbAuthenticationType("Basic")
	MongoDbAuthenticationTypeAnonymous = MongoDbAuthenticationType("Anonymous")
)

// Notebook parameter type.
type NotebookParameterType string

const (
	NotebookParameterTypeString = NotebookParameterType("string")
	NotebookParameterTypeInt    = NotebookParameterType("int")
	NotebookParameterTypeFloat  = NotebookParameterType("float")
	NotebookParameterTypeBool   = NotebookParameterType("bool")
)

// Synapse notebook reference type.
type NotebookReferenceType string

const (
	NotebookReferenceTypeNotebookReference = NotebookReferenceType("NotebookReference")
)

// Specify the credential type (key or cert) is used for service principal.
type ODataAadServicePrincipalCredentialType string

const (
	ODataAadServicePrincipalCredentialTypeServicePrincipalKey  = ODataAadServicePrincipalCredentialType("ServicePrincipalKey")
	ODataAadServicePrincipalCredentialTypeServicePrincipalCert = ODataAadServicePrincipalCredentialType("ServicePrincipalCert")
)

// Type of authentication used to connect to the OData service.
type ODataAuthenticationType string

const (
	ODataAuthenticationTypeBasic                  = ODataAuthenticationType("Basic")
	ODataAuthenticationTypeAnonymous              = ODataAuthenticationType("Anonymous")
	ODataAuthenticationTypeWindows                = ODataAuthenticationType("Windows")
	ODataAuthenticationTypeAadServicePrincipal    = ODataAuthenticationType("AadServicePrincipal")
	ODataAuthenticationTypeManagedServiceIdentity = ODataAuthenticationType("ManagedServiceIdentity")
)

// Parameter type.
type ParameterType string

const (
	ParameterTypeObject       = ParameterType("Object")
	ParameterTypeString       = ParameterType("String")
	ParameterTypeInt          = ParameterType("Int")
	ParameterTypeFloat        = ParameterType("Float")
	ParameterTypeBool         = ParameterType("Bool")
	ParameterTypeArray        = ParameterType("Array")
	ParameterTypeSecureString = ParameterType("SecureString")
)

// The authentication mechanism used to connect to the Phoenix server.
type PhoenixAuthenticationType string

const (
	PhoenixAuthenticationTypeAnonymous                    = PhoenixAuthenticationType("Anonymous")
	PhoenixAuthenticationTypeUsernameAndPassword          = PhoenixAuthenticationType("UsernameAndPassword")
	PhoenixAuthenticationTypeWindowsAzureHDInsightService = PhoenixAuthenticationType("WindowsAzureHDInsightService")
)

// Reject type.
type PolybaseSettingsRejectType string

const (
	PolybaseSettingsRejectTypeValue      = PolybaseSettingsRejectType("value")
	PolybaseSettingsRejectTypePercentage = PolybaseSettingsRejectType("percentage")
)

// The authentication mechanism used to connect to the Presto server.
type PrestoAuthenticationType string

const (
	PrestoAuthenticationTypeAnonymous = PrestoAuthenticationType("Anonymous")
	PrestoAuthenticationTypeLDAP      = PrestoAuthenticationType("LDAP")
)

// Whether or not public network access is allowed for the data factory.
type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

// The frequency.
type RecurrenceFrequency string

const (
	RecurrenceFrequencyNotSpecified = RecurrenceFrequency("NotSpecified")
	RecurrenceFrequencyMinute       = RecurrenceFrequency("Minute")
	RecurrenceFrequencyHour         = RecurrenceFrequency("Hour")
	RecurrenceFrequencyDay          = RecurrenceFrequency("Day")
	RecurrenceFrequencyWeek         = RecurrenceFrequency("Week")
	RecurrenceFrequencyMonth        = RecurrenceFrequency("Month")
	RecurrenceFrequencyYear         = RecurrenceFrequency("Year")
)

// Type of authentication used to connect to the REST service.
type RestServiceAuthenticationType string

const (
	RestServiceAuthenticationTypeAnonymous              = RestServiceAuthenticationType("Anonymous")
	RestServiceAuthenticationTypeBasic                  = RestServiceAuthenticationType("Basic")
	RestServiceAuthenticationTypeAadServicePrincipal    = RestServiceAuthenticationType("AadServicePrincipal")
	RestServiceAuthenticationTypeManagedServiceIdentity = RestServiceAuthenticationType("ManagedServiceIdentity")
	RestServiceAuthenticationTypeOAuth2ClientCredential = RestServiceAuthenticationType("OAuth2ClientCredential")
)

// The write behavior for the operation. Default is Insert.
type SalesforceSinkWriteBehavior string

const (
	SalesforceSinkWriteBehaviorInsert = SalesforceSinkWriteBehavior("Insert")
	SalesforceSinkWriteBehaviorUpsert = SalesforceSinkWriteBehavior("Upsert")
)

// The read behavior for the operation. Default is Query.
type SalesforceSourceReadBehavior string

const (
	SalesforceSourceReadBehaviorQuery    = SalesforceSourceReadBehavior("Query")
	SalesforceSourceReadBehaviorQueryAll = SalesforceSourceReadBehavior("QueryAll")
)

// The write behavior for the operation. Default is 'Insert'.
type SapCloudForCustomerSinkWriteBehavior string

const (
	SapCloudForCustomerSinkWriteBehaviorInsert = SapCloudForCustomerSinkWriteBehavior("Insert")
	SapCloudForCustomerSinkWriteBehaviorUpdate = SapCloudForCustomerSinkWriteBehavior("Update")
)

// The authentication type to be used to connect to the SAP HANA server.
type SapHanaAuthenticationType string

const (
	SapHanaAuthenticationTypeBasic   = SapHanaAuthenticationType("Basic")
	SapHanaAuthenticationTypeWindows = SapHanaAuthenticationType("Windows")
)

// The destination of logs. Type: string.
type ScriptActivityLogDestination string

const (
	ScriptActivityLogDestinationActivityOutput = ScriptActivityLogDestination("ActivityOutput")
	ScriptActivityLogDestinationExternalStore  = ScriptActivityLogDestination("ExternalStore")
)

// The direction of the parameter.
type ScriptActivityParameterDirection string

const (
	ScriptActivityParameterDirectionValueInput       = ScriptActivityParameterDirection("Input")
	ScriptActivityParameterDirectionValueOutput      = ScriptActivityParameterDirection("Output")
	ScriptActivityParameterDirectionValueInputOutput = ScriptActivityParameterDirection("InputOutput")
)

// The type of the parameter.
type ScriptActivityParameterType string

const (
	ScriptActivityParameterTypeBoolean        = ScriptActivityParameterType("Boolean")
	ScriptActivityParameterTypeDateTime       = ScriptActivityParameterType("DateTime")
	ScriptActivityParameterTypeDateTimeOffset = ScriptActivityParameterType("DateTimeOffset")
	ScriptActivityParameterTypeDecimal        = ScriptActivityParameterType("Decimal")
	ScriptActivityParameterTypeDouble         = ScriptActivityParameterType("Double")
	ScriptActivityParameterTypeGuid           = ScriptActivityParameterType("Guid")
	ScriptActivityParameterTypeInt16          = ScriptActivityParameterType("Int16")
	ScriptActivityParameterTypeInt32          = ScriptActivityParameterType("Int32")
	ScriptActivityParameterTypeInt64          = ScriptActivityParameterType("Int64")
	ScriptActivityParameterTypeSingle         = ScriptActivityParameterType("Single")
	ScriptActivityParameterTypeString         = ScriptActivityParameterType("String")
	ScriptActivityParameterTypeTimespan       = ScriptActivityParameterType("Timespan")
)

// The type of the query. Type: string.
type ScriptType string

const (
	ScriptTypeQuery    = ScriptType("Query")
	ScriptTypeNonQuery = ScriptType("NonQuery")
)

// The authentication type to use.
type ServiceNowAuthenticationType string

const (
	ServiceNowAuthenticationTypeBasic  = ServiceNowAuthenticationType("Basic")
	ServiceNowAuthenticationTypeOAuth2 = ServiceNowAuthenticationType("OAuth2")
)

// The authentication type to be used to connect to the FTP server.
type SftpAuthenticationType string

const (
	SftpAuthenticationTypeBasic        = SftpAuthenticationType("Basic")
	SftpAuthenticationTypeSshPublicKey = SftpAuthenticationType("SshPublicKey")
	SftpAuthenticationTypeMultiFactor  = SftpAuthenticationType("MultiFactor")
)

// The authentication method used to access the Spark server.
type SparkAuthenticationType string

const (
	SparkAuthenticationTypeAnonymous                    = SparkAuthenticationType("Anonymous")
	SparkAuthenticationTypeUsername                     = SparkAuthenticationType("Username")
	SparkAuthenticationTypeUsernameAndPassword          = SparkAuthenticationType("UsernameAndPassword")
	SparkAuthenticationTypeWindowsAzureHDInsightService = SparkAuthenticationType("WindowsAzureHDInsightService")
)

// Synapse spark job reference type.
type SparkJobReferenceType string

const (
	SparkJobReferenceTypeSparkJobDefinitionReference = SparkJobReferenceType("SparkJobDefinitionReference")
)

// The type of Spark server.
type SparkServerType string

const (
	SparkServerTypeSharkServer       = SparkServerType("SharkServer")
	SparkServerTypeSharkServer2      = SparkServerType("SharkServer2")
	SparkServerTypeSparkThriftServer = SparkServerType("SparkThriftServer")
)

// The transport protocol to use in the Thrift layer.
type SparkThriftTransportProtocol string

const (
	SparkThriftTransportProtocolBinary = SparkThriftTransportProtocol("Binary")
	SparkThriftTransportProtocolSASL   = SparkThriftTransportProtocol("SASL")
	SparkThriftTransportProtocol_HTTP_ = SparkThriftTransportProtocol("HTTP ")
)

// Sql always encrypted AKV authentication type. Type: string (or Expression with resultType string).
type SqlAlwaysEncryptedAkvAuthType string

const (
	SqlAlwaysEncryptedAkvAuthTypeServicePrincipal            = SqlAlwaysEncryptedAkvAuthType("ServicePrincipal")
	SqlAlwaysEncryptedAkvAuthTypeManagedIdentity             = SqlAlwaysEncryptedAkvAuthType("ManagedIdentity")
	SqlAlwaysEncryptedAkvAuthTypeUserAssignedManagedIdentity = SqlAlwaysEncryptedAkvAuthType("UserAssignedManagedIdentity")
)

// The type of SSIS log location.
type SsisLogLocationType string

const (
	SsisLogLocationTypeFile = SsisLogLocationType("File")
)

// The type of SSIS package location.
type SsisPackageLocationType string

const (
	SsisPackageLocationTypeSSISDB        = SsisPackageLocationType("SSISDB")
	SsisPackageLocationTypeFile          = SsisPackageLocationType("File")
	SsisPackageLocationTypeInlinePackage = SsisPackageLocationType("InlinePackage")
	SsisPackageLocationTypePackageStore  = SsisPackageLocationType("PackageStore")
)

// Stored procedure parameter type.
type StoredProcedureParameterType string

const (
	StoredProcedureParameterTypeString  = StoredProcedureParameterType("String")
	StoredProcedureParameterTypeInt     = StoredProcedureParameterType("Int")
	StoredProcedureParameterTypeInt64   = StoredProcedureParameterType("Int64")
	StoredProcedureParameterTypeDecimal = StoredProcedureParameterType("Decimal")
	StoredProcedureParameterTypeGuid    = StoredProcedureParameterType("Guid")
	StoredProcedureParameterTypeBoolean = StoredProcedureParameterType("Boolean")
	StoredProcedureParameterTypeDate    = StoredProcedureParameterType("Date")
)

// AuthenticationType to be used for connection.
type SybaseAuthenticationType string

const (
	SybaseAuthenticationTypeBasic   = SybaseAuthenticationType("Basic")
	SybaseAuthenticationTypeWindows = SybaseAuthenticationType("Windows")
)

// The authentication type to use.
type TeamDeskAuthenticationType string

const (
	TeamDeskAuthenticationTypeBasic = TeamDeskAuthenticationType("Basic")
	TeamDeskAuthenticationTypeToken = TeamDeskAuthenticationType("Token")
)

// AuthenticationType to be used for connection.
type TeradataAuthenticationType string

const (
	TeradataAuthenticationTypeBasic   = TeradataAuthenticationType("Basic")
	TeradataAuthenticationTypeWindows = TeradataAuthenticationType("Windows")
)

// Trigger reference type.
type TriggerReferenceType string

const (
	TriggerReferenceTypeTriggerReference = TriggerReferenceType("TriggerReference")
)

// The frequency of the time windows.
type TumblingWindowFrequency string

const (
	TumblingWindowFrequencyMinute = TumblingWindowFrequency("Minute")
	TumblingWindowFrequencyHour   = TumblingWindowFrequency("Hour")
	TumblingWindowFrequencyMonth  = TumblingWindowFrequency("Month")
)

// Linked service reference type.
type Type string

const (
	TypeLinkedServiceReference = Type("LinkedServiceReference")
)

// Variable type.
type VariableType string

const (
	VariableTypeString = VariableType("String")
	VariableTypeBool   = VariableType("Bool")
	VariableTypeArray  = VariableType("Array")
)

// Rest API method for target endpoint.
type WebActivityMethod string

const (
	WebActivityMethodGET    = WebActivityMethod("GET")
	WebActivityMethodPOST   = WebActivityMethod("POST")
	WebActivityMethodPUT    = WebActivityMethod("PUT")
	WebActivityMethodDELETE = WebActivityMethod("DELETE")
)

// Type of authentication used to connect to the web table source.
type WebAuthenticationType string

const (
	WebAuthenticationTypeBasic             = WebAuthenticationType("Basic")
	WebAuthenticationTypeAnonymous         = WebAuthenticationType("Anonymous")
	WebAuthenticationTypeClientCertificate = WebAuthenticationType("ClientCertificate")
)

// Rest API method for target endpoint.
type WebHookActivityMethod string

const (
	WebHookActivityMethodPOST = WebHookActivityMethod("POST")
)

// The authentication type to use.
type ZendeskAuthenticationType string

const (
	ZendeskAuthenticationTypeBasic = ZendeskAuthenticationType("Basic")
	ZendeskAuthenticationTypeToken = ZendeskAuthenticationType("Token")
)

func init() {
}
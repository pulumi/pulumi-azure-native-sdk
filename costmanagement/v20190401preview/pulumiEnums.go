// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401preview

// Show costs accumulated over time.
type AccumulatedType string

const (
	AccumulatedTypeTrue  = AccumulatedType("true")
	AccumulatedTypeFalse = AccumulatedType("false")
)

// The category of the budget, whether the budget tracks cost or usage.
type CategoryType string

const (
	CategoryTypeCost  = CategoryType("Cost")
	CategoryTypeUsage = CategoryType("Usage")
)

// Chart type of the main view in Cost Analysis. Required.
type ChartType string

const (
	ChartTypeArea          = ChartType("Area")
	ChartTypeLine          = ChartType("Line")
	ChartTypeStackedColumn = ChartType("StackedColumn")
	ChartTypeGroupedColumn = ChartType("GroupedColumn")
	ChartTypeTable         = ChartType("Table")
)

// The name of the aggregation function to use.
type FunctionType string

const (
	FunctionTypeSum = FunctionType("Sum")
)

// The granularity of rows in the report.
type GranularityType string

const (
	GranularityTypeDaily   = GranularityType("Daily")
	GranularityTypeMonthly = GranularityType("Monthly")
)

// KPI type (Forecast, Budget).
type KpiTypeType string

const (
	KpiTypeTypeForecast = KpiTypeType("Forecast")
	KpiTypeTypeBudget   = KpiTypeType("Budget")
)

// Metric to use when displaying costs.
type MetricType string

const (
	MetricTypeActualCost    = MetricType("ActualCost")
	MetricTypeAmortizedCost = MetricType("AmortizedCost")
	MetricTypeAHUB          = MetricType("AHUB")
)

// The comparison operator.
type NotificationOperatorType string

const (
	NotificationOperatorTypeEqualTo              = NotificationOperatorType("EqualTo")
	NotificationOperatorTypeGreaterThan          = NotificationOperatorType("GreaterThan")
	NotificationOperatorTypeGreaterThanOrEqualTo = NotificationOperatorType("GreaterThanOrEqualTo")
)

// The operator to use for comparison.
type OperatorType string

const (
	OperatorTypeIn       = OperatorType("In")
	OperatorTypeContains = OperatorType("Contains")
)

// Data type to show in view.
type PivotTypeType string

const (
	PivotTypeTypeDimension = PivotTypeType("Dimension")
	PivotTypeTypeTagKey    = PivotTypeType("TagKey")
)

// Has type of the column to group.
type ReportConfigColumnType string

const (
	ReportConfigColumnTypeTag       = ReportConfigColumnType("Tag")
	ReportConfigColumnTypeDimension = ReportConfigColumnType("Dimension")
)

// The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.
type ReportType string

const (
	ReportTypeUsage = ReportType("Usage")
)

// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
type TimeGrainType string

const (
	TimeGrainTypeMonthly   = TimeGrainType("Monthly")
	TimeGrainTypeQuarterly = TimeGrainType("Quarterly")
	TimeGrainTypeAnnually  = TimeGrainType("Annually")
)

// The time frame for pulling data for the report. If custom, then a specific time period must be provided.
type TimeframeType string

const (
	TimeframeTypeWeekToDate  = TimeframeType("WeekToDate")
	TimeframeTypeMonthToDate = TimeframeType("MonthToDate")
	TimeframeTypeYearToDate  = TimeframeType("YearToDate")
	TimeframeTypeCustom      = TimeframeType("Custom")
)

func init() {
}
// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101

// The type of the query.
type ExportType string

const (
	ExportTypeUsage = ExportType("Usage")
)

// The format of the export being delivered.
type FormatType string

const (
	FormatTypeCsv = FormatType("Csv")
)

// The name of the aggregation function to use.
type FunctionType string

const (
	FunctionTypeSum = FunctionType("Sum")
)

// The granularity of rows in the query.
type GranularityType string

const (
	GranularityTypeDaily  = GranularityType("Daily")
	GranularityTypeHourly = GranularityType("Hourly")
)

// The operator to use for comparison.
type OperatorType string

const (
	OperatorTypeIn = OperatorType("In")
)

// Has type of the column to group.
type QueryColumnType string

const (
	QueryColumnTypeTag       = QueryColumnType("Tag")
	QueryColumnTypeDimension = QueryColumnType("Dimension")
)

// The schedule recurrence.
type RecurrenceType string

const (
	RecurrenceTypeDaily    = RecurrenceType("Daily")
	RecurrenceTypeWeekly   = RecurrenceType("Weekly")
	RecurrenceTypeMonthly  = RecurrenceType("Monthly")
	RecurrenceTypeAnnually = RecurrenceType("Annually")
)

// The sorting direction
type SortDirection string

const (
	SortDirectionAscending  = SortDirection("Ascending")
	SortDirectionDescending = SortDirection("Descending")
)

// The status of the schedule. Whether active or not. If inactive, the export's scheduled execution is paused.
type StatusType string

const (
	StatusTypeActive   = StatusType("Active")
	StatusTypeInactive = StatusType("Inactive")
)

// The time frame for pulling data for the query. If custom, then a specific time period must be provided.
type TimeframeType string

const (
	TimeframeTypeWeekToDate          = TimeframeType("WeekToDate")
	TimeframeTypeMonthToDate         = TimeframeType("MonthToDate")
	TimeframeTypeYearToDate          = TimeframeType("YearToDate")
	TimeframeTypeTheLastWeek         = TimeframeType("TheLastWeek")
	TimeframeTypeTheLastMonth        = TimeframeType("TheLastMonth")
	TimeframeTypeTheLastYear         = TimeframeType("TheLastYear")
	TimeframeTypeCustom              = TimeframeType("Custom")
	TimeframeTypeBillingMonthToDate  = TimeframeType("BillingMonthToDate")
	TimeframeTypeTheLastBillingMonth = TimeframeType("TheLastBillingMonth")
)

func init() {
}
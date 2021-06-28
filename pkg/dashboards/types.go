// Code generated by tutone: DO NOT EDIT
package dashboards

import (
	"github.com/newrelic/newrelic-client-go/pkg/entities"
	"github.com/newrelic/newrelic-client-go/pkg/nrdb"
	"github.com/newrelic/newrelic-client-go/pkg/nrtime"
)

// DashboardCreateErrorType - Expected error types that can be returned by create operation
type DashboardCreateErrorType string

var DashboardCreateErrorTypeTypes = struct {
	// Invalid input error
	INVALID_INPUT DashboardCreateErrorType
}{
	// Invalid input error
	INVALID_INPUT: "INVALID_INPUT",
}

// DashboardDeleteErrorType - Expected error types that can be returned by delete operation
type DashboardDeleteErrorType string

var DashboardDeleteErrorTypeTypes = struct {
	// Dashboard not found in the system
	DASHBOARD_NOT_FOUND DashboardDeleteErrorType
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION DashboardDeleteErrorType
}{
	// Dashboard not found in the system
	DASHBOARD_NOT_FOUND: "DASHBOARD_NOT_FOUND",
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION: "FORBIDDEN_OPERATION",
}

// DashboardDeleteResultStatus - Result status of delete operation.
type DashboardDeleteResultStatus string

var DashboardDeleteResultStatusTypes = struct {
	// FAILURE.
	FAILURE DashboardDeleteResultStatus
	// SUCCESS.
	SUCCESS DashboardDeleteResultStatus
}{
	// FAILURE.
	FAILURE: "FAILURE",
	// SUCCESS.
	SUCCESS: "SUCCESS",
}

// DashboardUpdateErrorType - Expected error types that can be returned by update operation
type DashboardUpdateErrorType string

var DashboardUpdateErrorTypeTypes = struct {
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION DashboardUpdateErrorType
	// Invalid input error
	INVALID_INPUT DashboardUpdateErrorType
}{
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION: "FORBIDDEN_OPERATION",
	// Invalid input error
	INVALID_INPUT: "INVALID_INPUT",
}

// DashboardUpdatePageErrorType - Expected error types that can be returned by updatePage operation
type DashboardUpdatePageErrorType string

var DashboardUpdatePageErrorTypeTypes = struct {
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION DashboardUpdatePageErrorType
	// Invalid input error
	INVALID_INPUT DashboardUpdatePageErrorType
	// Page not found in the system
	PAGE_NOT_FOUND DashboardUpdatePageErrorType
}{
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION: "FORBIDDEN_OPERATION",
	// Invalid input error
	INVALID_INPUT: "INVALID_INPUT",
	// Page not found in the system
	PAGE_NOT_FOUND: "PAGE_NOT_FOUND",
}

// DashboardUpdateWidgetsInPageErrorType - Expected error types that can be returned by updateWidgetsInPage operation
type DashboardUpdateWidgetsInPageErrorType string

var DashboardUpdateWidgetsInPageErrorTypeTypes = struct {
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION DashboardUpdateWidgetsInPageErrorType
	// Invalid input error
	INVALID_INPUT DashboardUpdateWidgetsInPageErrorType
	// Page not found in the system
	PAGE_NOT_FOUND DashboardUpdateWidgetsInPageErrorType
	// Widget not found in the system
	WIDGET_NOT_FOUND DashboardUpdateWidgetsInPageErrorType
}{
	// User is not allowed to execute the operation
	FORBIDDEN_OPERATION: "FORBIDDEN_OPERATION",
	// Invalid input error
	INVALID_INPUT: "INVALID_INPUT",
	// Page not found in the system
	PAGE_NOT_FOUND: "PAGE_NOT_FOUND",
	// Widget not found in the system
	WIDGET_NOT_FOUND: "WIDGET_NOT_FOUND",
}

// DashboardAreaWidgetConfigurationInput - Configuration for visualization type 'viz.area'
type DashboardAreaWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
}

// DashboardBarWidgetConfigurationInput - Configuration for visualization type 'viz.bar'
type DashboardBarWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
}

// DashboardBillboardWidgetConfigurationInput - Configuration for visualization type 'viz.billboard'
type DashboardBillboardWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
	// thresholds
	Thresholds []DashboardBillboardWidgetThresholdInput `json:"thresholds,omitempty"`
}

// DashboardCreateError - Expected errors that can be returned by create operation
type DashboardCreateError struct {
	// Error description
	Description string `json:"description,omitempty"`
	// Error type
	Type DashboardCreateErrorType `json:"type"`
}

// DashboardCreateResult - Result of create operation.
type DashboardCreateResult struct {
	// Dashboard creation result
	EntityResult DashboardEntityResult `json:"entityResult,omitempty"`
	// Expected errors while processing request
	Errors []DashboardCreateError `json:"errors,omitempty"`
}

// DashboardDeleteError - Expected error types that can be returned by delete operation
type DashboardDeleteError struct {
	// Error description
	Description string `json:"description,omitempty"`
	// Error type
	Type DashboardDeleteErrorType `json:"type"`
}

// DashboardDeleteResult - Result of delete operation.
type DashboardDeleteResult struct {
	// Expected errors while processing request
	Errors []DashboardDeleteError `json:"errors,omitempty"`
	// The status of the attempted delete.
	Status DashboardDeleteResultStatus `json:"status,omitempty"`
}

// DashboardEntityResult - Public schema - `DashboardEntity` result representation for mutations. It's a subset of the `DashboardEntity` that inherits from the Entity type, but a complete different type.
type DashboardEntityResult struct {
	// Account ID.
	AccountID int `json:"accountId,omitempty"`
	// Dashboard creation timestamp.
	CreatedAt nrtime.DateTime `json:"createdAt,omitempty"`
	// Dashboard description.
	Description string `json:"description,omitempty"`
	// Unique entity identifier.
	GUID entities.EntityGUID `json:"guid,omitempty"`
	// Dashboard name.
	Name string `json:"name,omitempty"`
	// Dashboard owner
	Owner entities.DashboardOwnerInfo `json:"owner,omitempty"`
	// Dashboard pages.
	Pages []entities.DashboardPage `json:"pages,omitempty"`
	// Dashboard permissions configuration.
	Permissions entities.DashboardPermissions `json:"permissions,omitempty"`
	// Dashboard update timestamp.
	UpdatedAt nrtime.DateTime `json:"updatedAt,omitempty"`
}

// DashboardInput - Dashboard input
type DashboardInput struct {
	// Dashboard description.
	Description string `json:"description,omitempty"`
	// Dashboard name.
	Name string `json:"name"`
	// Dashboard page input.
	Pages []DashboardPageInput `json:"pages,omitempty"`
	// Dashboard permissions configuration.
	Permissions entities.DashboardPermissions `json:"permissions"`
}

// DashboardLineWidgetConfigurationInput - Configuration for visualization type 'viz.line'
type DashboardLineWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
}

// DashboardMarkdownWidgetConfigurationInput - Configuration for visualization type 'viz.markdown'
type DashboardMarkdownWidgetConfigurationInput struct {
	// Markdown content of the widget
	Text string `json:"text"`
}

// DashboardPageInput - Page input
type DashboardPageInput struct {
	// Page description.
	Description string `json:"description,omitempty"`
	// Unique entity identifier of the Page to be updated. When null, it means a new Page will be created.
	GUID entities.EntityGUID `json:"guid,omitempty"`
	// Page name.
	Name string `json:"name"`
	// Page widgets.
	Widgets []DashboardWidgetInput `json:"widgets,omitempty"`
}

// DashboardPieWidgetConfigurationInput - Configuration for visualization type 'viz.pie'
type DashboardPieWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
}

// DashboardSnapshotURLInput - Parameters that affect the data and the rendering of the dashboards returned by the snapshot url mutation.
type DashboardSnapshotURLInput struct {
	// Period of time from which the data to be displayed on the dashboard will be obtained.
	TimeWindow DashboardSnapshotURLTimeWindowInput `json:"timeWindow,omitempty"`
}

// DashboardSnapshotURLTimeWindowInput - Period of time from which the data to be displayed on the dashboard will be obtained.
type DashboardSnapshotURLTimeWindowInput struct {
	// The starting time of the time window. If specified, an endTime or a duration must also be specified.
	BeginTime nrtime.EpochMilliseconds `json:"beginTime,omitempty"`
	// The duration of the time window.
	Duration nrtime.Milliseconds `json:"duration,omitempty"`
	// The end time of the time window. If specified, a beginTime or a duration must also be specified.
	EndTime nrtime.EpochMilliseconds `json:"endTime,omitempty"`
}

// DashboardTableWidgetConfigurationInput - Configuration for visualization type 'viz.table'
type DashboardTableWidgetConfigurationInput struct {
	// nrql queries
	NRQLQueries []DashboardWidgetNRQLQueryInput `json:"nrqlQueries,omitempty"`
}

// DashboardUpdateError - Expected errors that can be returned by update operation
type DashboardUpdateError struct {
	// Error description
	Description string `json:"description,omitempty"`
	// Error type
	Type DashboardUpdateErrorType `json:"type"`
}

// DashboardUpdatePageError - Expected errors that can be returned by updatePage operation
type DashboardUpdatePageError struct {
	// Error description
	Description string `json:"description,omitempty"`
	// Error type
	Type DashboardUpdatePageErrorType `json:"type"`
}

// DashboardUpdatePageInput - Page input used when updating an individual page
type DashboardUpdatePageInput struct {
	// Page description.
	Description string `json:"description,omitempty"`
	// Page name.
	Name string `json:"name"`
	// Page widgets.
	Widgets []DashboardWidgetInput `json:"widgets,omitempty"`
}

// DashboardUpdatePageResult - Result of updatePage operation.
type DashboardUpdatePageResult struct {
	// Expected errors while processing request. No errors means successful request.
	Errors []DashboardUpdatePageError `json:"errors,omitempty"`
}

// DashboardUpdateResult - Result of update operation.
type DashboardUpdateResult struct {
	// Dashboard update result
	EntityResult DashboardEntityResult `json:"entityResult,omitempty"`
	// Expected errors while processing request
	Errors []DashboardUpdateError `json:"errors,omitempty"`
}

// DashboardUpdateWidgetInput - Input type used when updating widgets
type DashboardUpdateWidgetInput struct {
	// Typed configuration for the widget
	Configuration DashboardWidgetConfigurationInput `json:"configuration,omitempty"`
	// Id of the widget to be updated.
	ID string `json:"id"`
	// layout
	Layout DashboardWidgetLayoutInput `json:"layout,omitempty"`
	// Related entities. Currently only supports Dashboard entities, but may allow other cases in the future.
	LinkedEntityGUIDs []entities.EntityGUID `json:"linkedEntityGuids"`
	// Untyped scalar of configuration for the widget
	RawConfiguration entities.DashboardWidgetRawConfiguration `json:"rawConfiguration,omitempty"`
	// title
	Title string `json:"title,omitempty"`
	// Specifies how this widget will be visualized. If null, the WidgetConfigurationInput will be used to determine the visualization.
	Visualization DashboardWidgetVisualizationInput `json:"visualization,omitempty"`
}

// DashboardUpdateWidgetsInPageError - Expected errors that can be returned by updateWidgetsInPage operation
type DashboardUpdateWidgetsInPageError struct {
	// Error description
	Description string `json:"description,omitempty"`
	// Error type
	Type DashboardUpdateWidgetsInPageErrorType `json:"type"`
}

// DashboardUpdateWidgetsInPageResult - Result of updateWidgetsInPage operation.
type DashboardUpdateWidgetsInPageResult struct {
	// Expected errors while processing request. No errors means successful request.
	Errors []DashboardUpdateWidgetsInPageError `json:"errors,omitempty"`
}

// DashboardWidgetConfigurationInput - Typed configuration for known visualizations. At most one may be populated.
type DashboardWidgetConfigurationInput struct {
	// Configuration for visualization type 'viz.area'
	Area *DashboardAreaWidgetConfigurationInput `json:"area,omitempty"`
	// Configuration for visualization type 'viz.bar'
	Bar *DashboardBarWidgetConfigurationInput `json:"bar,omitempty"`
	// Configuration for visualization type 'viz.billboard'
	Billboard *DashboardBillboardWidgetConfigurationInput `json:"billboard,omitempty"`
	// Configuration for visualization type 'viz.line'
	Line *DashboardLineWidgetConfigurationInput `json:"line,omitempty"`
	// Configuration for visualization type 'viz.markdown'
	Markdown *DashboardMarkdownWidgetConfigurationInput `json:"markdown,omitempty"`
	// Configuration for visualization type 'viz.pie'
	Pie *DashboardPieWidgetConfigurationInput `json:"pie,omitempty"`
	// Configuration for visualization type 'viz.table'
	Table *DashboardTableWidgetConfigurationInput `json:"table,omitempty"`
}

// DashboardWidgetInput - Widget input
type DashboardWidgetInput struct {
	// Typed configuration for the widget
	Configuration DashboardWidgetConfigurationInput `json:"configuration,omitempty"`
	// Id of the widget. If null, a new widget will be created and added to a dashboard.
	ID string `json:"id,omitempty"`
	// layout
	Layout DashboardWidgetLayoutInput `json:"layout,omitempty"`
	// Related entities. Currently only supports Dashboard entities, but may allow other cases in the future.
	LinkedEntityGUIDs []entities.EntityGUID `json:"linkedEntityGuids"`
	// Untyped scalar of configuration for the widget
	RawConfiguration entities.DashboardWidgetRawConfiguration `json:"rawConfiguration,omitempty"`
	// title
	Title string `json:"title,omitempty"`
	// Specifies how this widget will be visualized. If null, the WidgetConfigurationInput will be used to determine the visualization.
	Visualization DashboardWidgetVisualizationInput `json:"visualization,omitempty"`
}

// DashboardWidgetLayoutInput - Widget layout input
type DashboardWidgetLayoutInput struct {
	// column.
	Column int `json:"column,omitempty"`
	// height.
	Height int `json:"height,omitempty"`
	// row.
	Row int `json:"row,omitempty"`
	// width.
	Width int `json:"width,omitempty"`
}

// DashboardWidgetNRQLQueryInput - NRQL query used by a widget
type DashboardWidgetNRQLQueryInput struct {
	// accountId
	AccountID int `json:"accountId"`
	// NRQL formatted query
	Query nrdb.NRQL `json:"query"`
}

// DashboardWidgetVisualizationInput - Visualization configuration
type DashboardWidgetVisualizationInput struct {
	// Nerdpack artifact ID
	ID string `json:"id,omitempty"`
}

// Float - The `Float` scalar type represents signed double-precision fractional
// values as specified by
// [IEEE 754](http://en.wikipedia.org/wiki/IEEE_floating_point).
type Float string

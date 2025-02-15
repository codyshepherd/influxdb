// Code generated by the feature package; DO NOT EDIT.

package feature

var appMetrics = MakeBoolFlag(
	"App Metrics",
	"appMetrics",
	"Bucky, Monitoring Team",
	false,
	Permanent,
	true,
)

// AppMetrics - Send UI Telementry to Tools cluster - should always be false in OSS
func AppMetrics() BoolFlag {
	return appMetrics
}

var backendExample = MakeBoolFlag(
	"Backend Example",
	"backendExample",
	"Gavin Cabbage",
	false,
	Permanent,
	false,
)

// BackendExample - A permanent backend example boolean flag
func BackendExample() BoolFlag {
	return backendExample
}

var communityTemplates = MakeBoolFlag(
	"Community Templates",
	"communityTemplates",
	"Bucky",
	true,
	Permanent,
	true,
)

// CommunityTemplates - Replace current template uploading functionality with community driven templates
func CommunityTemplates() BoolFlag {
	return communityTemplates
}

var frontendExample = MakeIntFlag(
	"Frontend Example",
	"frontendExample",
	"Gavin Cabbage",
	42,
	Temporary,
	true,
)

// FrontendExample - A temporary frontend example integer flag
func FrontendExample() IntFlag {
	return frontendExample
}

var groupWindowAggregateTranspose = MakeBoolFlag(
	"Group Window Aggregate Transpose",
	"groupWindowAggregateTranspose",
	"Query Team",
	false,
	Temporary,
	false,
)

// GroupWindowAggregateTranspose - Enables the GroupWindowAggregateTransposeRule for all enabled window aggregates
func GroupWindowAggregateTranspose() BoolFlag {
	return groupWindowAggregateTranspose
}

var newLabels = MakeBoolFlag(
	"New Label Package",
	"newLabels",
	"Alirie Gray",
	false,
	Temporary,
	false,
)

// NewLabelPackage - Enables the refactored labels api
func NewLabelPackage() BoolFlag {
	return newLabels
}

var memoryOptimizedFill = MakeBoolFlag(
	"Memory Optimized Fill",
	"memoryOptimizedFill",
	"Query Team",
	false,
	Temporary,
	false,
)

// MemoryOptimizedFill - Enable the memory optimized fill()
func MemoryOptimizedFill() BoolFlag {
	return memoryOptimizedFill
}

var memoryOptimizedSchemaMutation = MakeBoolFlag(
	"Memory Optimized Schema Mutation",
	"memoryOptimizedSchemaMutation",
	"Query Team",
	false,
	Temporary,
	false,
)

// MemoryOptimizedSchemaMutation - Enable the memory optimized schema mutation functions
func MemoryOptimizedSchemaMutation() BoolFlag {
	return memoryOptimizedSchemaMutation
}

var queryTracing = MakeBoolFlag(
	"Query Tracing",
	"queryTracing",
	"Query Team",
	false,
	Permanent,
	false,
)

// QueryTracing - Turn on query tracing for queries that are sampled
func QueryTracing() BoolFlag {
	return queryTracing
}

var bandPlotType = MakeBoolFlag(
	"Band Plot Type",
	"bandPlotType",
	"Monitoring Team",
	false,
	Temporary,
	true,
)

// BandPlotType - Enables the creation of a band plot in Dashboards
func BandPlotType() BoolFlag {
	return bandPlotType
}

var mosaicGraphType = MakeBoolFlag(
	"Mosaic Graph Type",
	"mosaicGraphType",
	"Monitoring Team",
	false,
	Temporary,
	true,
)

// MosaicGraphType - Enables the creation of a mosaic graph in Dashboards
func MosaicGraphType() BoolFlag {
	return mosaicGraphType
}

var notebooks = MakeBoolFlag(
	"Notebooks",
	"notebooks",
	"Monitoring Team",
	false,
	Temporary,
	true,
)

// Notebooks - Determine if the notebook feature's route and navbar icon are visible to the user
func Notebooks() BoolFlag {
	return notebooks
}

var injectLatestSuccessTime = MakeBoolFlag(
	"Inject Latest Success Time",
	"injectLatestSuccessTime",
	"Compute Team",
	false,
	Temporary,
	false,
)

// InjectLatestSuccessTime - Inject the latest successful task run timestamp into a Task query extern when executing.
func InjectLatestSuccessTime() BoolFlag {
	return injectLatestSuccessTime
}

var enforceOrgDashboardLimits = MakeBoolFlag(
	"Enforce Organization Dashboard Limits",
	"enforceOrgDashboardLimits",
	"Compute Team",
	false,
	Temporary,
	false,
)

// EnforceOrganizationDashboardLimits - Enforces the default limit params for the dashboards api when orgs are set
func EnforceOrganizationDashboardLimits() BoolFlag {
	return enforceOrgDashboardLimits
}

var timeFilterFlags = MakeBoolFlag(
	"Time Filter Flags",
	"timeFilterFlags",
	"Compute Team",
	false,
	Temporary,
	true,
)

// TimeFilterFlags - Filter task run list based on before and after flags
func TimeFilterFlags() BoolFlag {
	return timeFilterFlags
}

var all = []Flag{
	appMetrics,
	backendExample,
	communityTemplates,
	frontendExample,
	groupWindowAggregateTranspose,
	newLabels,
	memoryOptimizedFill,
	memoryOptimizedSchemaMutation,
	queryTracing,
	bandPlotType,
	mosaicGraphType,
	notebooks,
	injectLatestSuccessTime,
	enforceOrgDashboardLimits,
	timeFilterFlags,
}

var byKey = map[string]Flag{
	"appMetrics":                    appMetrics,
	"backendExample":                backendExample,
	"communityTemplates":            communityTemplates,
	"frontendExample":               frontendExample,
	"groupWindowAggregateTranspose": groupWindowAggregateTranspose,
	"newLabels":                     newLabels,
	"memoryOptimizedFill":           memoryOptimizedFill,
	"memoryOptimizedSchemaMutation": memoryOptimizedSchemaMutation,
	"queryTracing":                  queryTracing,
	"bandPlotType":                  bandPlotType,
	"mosaicGraphType":               mosaicGraphType,
	"notebooks":                     notebooks,
	"injectLatestSuccessTime":       injectLatestSuccessTime,
	"enforceOrgDashboardLimits":     enforceOrgDashboardLimits,
	"timeFilterFlags":               timeFilterFlags,
}

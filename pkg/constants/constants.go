package constants

const (
	DATE_SEPERATOR = "-"
)

// Context Keys
const (
	TRANSACTION_KEY          = "TXN"
	PAGINATION_KEY           = "PGN"
	FILTER_KEY               = "FILTER"
	SESSION                  = "SESSION"
	LOGGER                   = "LOGGER"
	INTERNAL_SKIP_PAGINATION = "INTERNAL_SKIP_PAGINATION"

	CORRELATION_ID = "correlation.id" //in this case to match newrelic standards

	FILTER_QUERY_PARAM = "filters"
)

// GORM Transaction Keys
const (
	USER_ID    = "USER_ID"
	CHANNEL_ID = "CHANNEL_ID"
)

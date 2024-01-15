package summary

import "github.com/theo-bot/service4.1-video/business/data/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByUserID, order.ASC)

// Set of fields that the results can be ordered by. These are the names
// that should be used by the application layer.
const (
	OrderByUserID   = "userid"
	OrderByUserName = "userName"
)

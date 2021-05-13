package forms

import (
	"shin-psmapi/utils"
	"time"
)

type GetAllServisanForm struct {
	SearchQuery string                 `form:"q"`
	SearchBy    utils.ServisanSearchBy `form:"by" binding:"servisan_search_by_col"`
	MinDate     time.Time              `form:"min_date"`
	MaxDate     time.Time              `form:"max_date"`
}

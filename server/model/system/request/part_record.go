package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PartRecordSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	PartModel      *string     `json:"partModel" form:"partModel"`
	PartNumber     *string     `json:"partNumber" form:"partNumber"`
	request.PageInfo
}


package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EquipmentSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      EquipmentNumber  *string `json:"equipmentNumber" form:"equipmentNumber"` 
      EquipmentName  *string `json:"equipmentName" form:"equipmentName"` 
      OperationalStatus  *string `json:"operationalStatus" form:"operationalStatus"` 
      NunningState  *string `json:"nunningState" form:"nunningState"` 
      SctiveState  *string `json:"sctiveState" form:"sctiveState"` 
      SoftwareVersion  *string `json:"softwareVersion" form:"softwareVersion"` 
      AfterSalesPersonnel  *string `json:"afterSalesPersonnel" form:"afterSalesPersonnel"` 
    request.PageInfo
}

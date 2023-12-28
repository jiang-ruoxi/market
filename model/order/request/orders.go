package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type OrdersSearch struct{
    order.Orders
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}

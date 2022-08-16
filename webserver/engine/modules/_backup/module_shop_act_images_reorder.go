package modules

import (
	"context"
	"encoding/json"

	"server/engine/utils"
	"server/engine/wrapper"
)

type OrderItem struct {
	Id    int
	Order int
}

type Orders struct {
	Items []OrderItem
}

func (this *Modules) RegisterAction_ShopImagesReorder() *Action {
	return this.newAction(AInfo{
		Mount:     "shop-images-reorder",
		WantAdmin: true,
	}, func(wrap *wrapper.Wrapper) {
		pf_data := utils.Trim(wrap.R.FormValue("data"))

		var orders Orders
		if err := json.Unmarshal([]byte(pf_data), &orders); err != nil {
			wrap.MsgError(err.Error())
			return
		}

		if len(orders.Items) > 0 {
			if err := wrap.DB.Transaction(wrap.R.Context(), func(ctx context.Context, tx *wrapper.Tx) error {
				for _, value := range orders.Items {
					if _, err := tx.Exec(ctx, "UPDATE shop_product_images SET ord = ? WHERE id = ?;", value.Order, value.Id); err != nil {
						return err
					}
				}
				return nil
			}); err != nil {
				wrap.MsgError(err.Error())
				return
			}
		}

		wrap.RecreateProductXmlFile()

		wrap.ResetCacheBlocks()

		// Remove loading effect
		wrap.Write(`$('#list-images').removeClass('loading');`)
	})
}

package modules

import (
	"context"

	"server/engine/utils"
	"server/engine/wrapper"
)

func (this *Modules) RegisterAction_IndexDelete() *Action {
	return this.newAction(AInfo{
		Mount:     "index-delete",
		WantAdmin: true,
	}, func(wrap *wrapper.Wrapper) {
		pf_id := utils.Trim(wrap.R.FormValue("id"))

		if !utils.IsNumeric(pf_id) {
			wrap.MsgError(`Inner system error`)
			return
		}

		err := wrap.DB.Transaction(wrap.R.Context(), func(ctx context.Context, tx *wrapper.Tx) error {
			// Process
			if _, err := tx.Exec(ctx, "DELETE FROM pages WHERE id = ?;", utils.StrToInt(pf_id)); err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			wrap.MsgError(err.Error())
			return
		}

		wrap.ResetCacheBlocks()

		// Reload current page
		wrap.Write(`window.location.reload(false);`)
	})
}

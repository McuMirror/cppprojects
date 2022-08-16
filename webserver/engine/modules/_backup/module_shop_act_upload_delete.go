package modules

import (
	"context"
	"os"

	"server/engine/utils"
	"server/engine/wrapper"
)

func (this *Modules) RegisterAction_ShopUploadDelete() *Action {
	return this.newAction(AInfo{
		Mount:     "shop-upload-delete",
		WantAdmin: true,
	}, func(wrap *wrapper.Wrapper) {
		pf_id := utils.Trim(wrap.R.FormValue("id"))
		pf_file := utils.Trim(wrap.R.FormValue("file"))

		if !utils.IsNumeric(pf_id) {
			wrap.MsgError(`Inner system error`)
			return
		}

		if pf_file == "" {
			wrap.MsgError(`Inner system error`)
			return
		}

		if err := wrap.DB.Transaction(wrap.R.Context(), func(ctx context.Context, tx *wrapper.Tx) error {
			// Block rows
			if _, err := tx.Exec(ctx, "SELECT id FROM shop_products WHERE id = ? FOR UPDATE;", utils.StrToInt(pf_id)); err != nil {
				return err
			}
			if _, err := tx.Exec(ctx, "SELECT product_id FROM shop_product_images WHERE product_id = ? FOR UPDATE;", utils.StrToInt(pf_id)); err != nil {
				return err
			}

			// Delete row
			if _, err := tx.Exec(ctx, "DELETE FROM shop_product_images WHERE product_id = ? AND filename = ?;", utils.StrToInt(pf_id), pf_file); err != nil {
				return err
			}

			// Delete file
			target_file_full := wrap.DHtdocs + string(os.PathSeparator) + "products" + string(os.PathSeparator) + "images" + string(os.PathSeparator) + pf_id + string(os.PathSeparator) + pf_file
			_ = os.Remove(target_file_full)

			// Delete thumbnails
			if err := wrap.RemoveProductImageThumbnails(pf_id, "thumb-*-"+pf_file); err != nil {
				return err
			}

			return nil
		}); err != nil {
			wrap.MsgError(err.Error())
			return
		}

		wrap.RecreateProductXmlFile()

		wrap.ResetCacheBlocks()

		wrap.Write(`$('#list-images a').each(function(i, e) { if($(e).attr('title') == '` + pf_file + `') { $(e).parent().remove(); return; } });`)
	})
}

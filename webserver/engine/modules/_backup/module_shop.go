package modules

import (
	"html"
	"net/http"
	"strings"
	"time"

	"server/engine/assets"
	"server/engine/basket"
	"server/engine/builder"
	"server/engine/consts"
	"server/engine/fetdata"
	"server/engine/sqlw"
	"server/engine/utils"
	"server/engine/wrapper"
)

func (this *Modules) shop_GetCurrencySelectOptions(wrap *wrapper.Wrapper, id int) string {
	result := ``
	rows, err := wrap.DB.Query(
		wrap.R.Context(),
		`SELECT
			id,
			code
		FROM
			shop_currencies
		ORDER BY
			id ASC
		;`,
	)
	if err == nil {
		defer rows.Close()
		values := make([]string, 2)
		scan := make([]interface{}, len(values))
		for i := range values {
			scan[i] = &values[i]
		}
		idStr := utils.IntToStr(id)
		for rows.Next() {
			err = rows.Scan(scan...)
			if *wrap.LogCpError(&err) == nil {
				selected := ""
				if string(values[0]) == idStr {
					selected = " selected"
				}
				result += `<option title="` + html.EscapeString(string(values[1])) + `" value="` + html.EscapeString(string(values[0])) + `"` + selected + `>` + html.EscapeString(string(values[1])) + `</option>`
			}
		}
	}
	return result
}

func (this *Modules) shop_GetProductValuesInputs(wrap *wrapper.Wrapper, product_id int) string {
	result := ``
	rows, err := wrap.DB.Query(
		wrap.R.Context(),
		`SELECT
			shop_filters.id,
			shop_filters.name,
			shop_filters_values.id,
			shop_filters_values.name,
			IF(shop_filter_product_values.filter_value_id > 0, 1, 0) as selected
		FROM
			shop_filters_values
			LEFT JOIN shop_filters ON shop_filters.id = shop_filters_values.filter_id
			LEFT JOIN shop_filter_product_values ON
				shop_filter_product_values.filter_value_id = shop_filters_values.id AND
				shop_filter_product_values.product_id = `+utils.IntToStr(product_id)+`
			LEFT JOIN (
				SELECT
					shop_filters_values.filter_id,
					shop_filter_product_values.product_id
				FROM
					shop_filter_product_values
					LEFT JOIN shop_filters_values ON shop_filters_values.id = shop_filter_product_values.filter_value_id 
				WHERE
					shop_filter_product_values.product_id = `+utils.IntToStr(product_id)+`
				GROUP BY
					shop_filters_values.filter_id
			) as filter_used ON filter_used.filter_id = shop_filters.id
		WHERE
			filter_used.filter_id IS NOT NULL
		ORDER BY
			shop_filters.name ASC,
			shop_filters_values.name ASC
		;`,
	)

	filter_ids := []int{}
	filter_names := map[int]string{}
	filter_values := map[int][]string{}

	if err == nil {
		defer rows.Close()
		values := make([]string, 5)
		scan := make([]interface{}, len(values))
		for i := range values {
			scan[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scan...)
			if *wrap.LogCpError(&err) == nil {
				filter_id := utils.StrToInt(string(values[0]))
				if !utils.InArrayInt(filter_ids, filter_id) {
					filter_ids = append(filter_ids, filter_id)
				}
				filter_names[filter_id] = html.EscapeString(string(values[1]))
				selected := ``
				if utils.StrToInt(string(values[4])) == 1 {
					selected = ` selected`
				}
				filter_values[filter_id] = append(filter_values[filter_id], `<option value="`+html.EscapeString(string(values[2]))+`"`+selected+`>`+html.EscapeString(string(values[3]))+`</option>`)
			}
		}
	}
	for _, filter_id := range filter_ids {
		result += `<div class="form-group" id="prod_attr_` + utils.IntToStr(filter_id) + `">` +
			`<div><b>` + filter_names[filter_id] + `</b></div>` +
			`<div class="position-relative">` +
			`<select class="selectpicker form-control" name="value.` + utils.IntToStr(filter_id) + `" autocomplete="off" required multiple>` +
			strings.Join(filter_values[filter_id], "") +
			`</select>` +
			`<button type="button" class="btn btn-danger btn-dynamic-remove" onclick="fave.ShopProductsRemove(this);">&times;</button>` +
			`</div>` +
			`</div>`
	}
	return result
}

func (this *Modules) shop_GetFilterValuesInputs(wrap *wrapper.Wrapper, filter_id int) string {
	result := ``
	rows, err := wrap.DB.Query(
		wrap.R.Context(),
		`SELECT
			id,
			name
		FROM
			shop_filters_values
		WHERE
			filter_id = ?
		ORDER BY
			name ASC
		;`,
		filter_id,
	)
	if err == nil {
		defer rows.Close()
		values := make([]string, 2)
		scan := make([]interface{}, len(values))
		for i := range values {
			scan[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scan...)
			if *wrap.LogCpError(&err) == nil {
				result += `<div class="form-group position-relative"><input class="form-control" type="text" name="value.` + html.EscapeString(string(values[0])) + `" value="` + html.EscapeString(string(values[1])) + `" placeholder="" autocomplete="off" required><button type="button" class="btn btn-danger btn-dynamic-remove" onclick="fave.ShopAttributesRemove(this);">&times;</button></div>`
			}
		}
	}
	return result
}

func (this *Modules) shop_GetAllAttributesSelectOptions(wrap *wrapper.Wrapper) string {
	result := ``
	rows, err := wrap.DB.Query(
		wrap.R.Context(),
		`SELECT
			id,
			name,
			filter
		FROM
			shop_filters
		ORDER BY
			name ASC
		;`,
	)
	result += `<option title="&mdash;" value="0">&mdash;</option>`
	if err == nil {
		defer rows.Close()
		values := make([]string, 3)
		scan := make([]interface{}, len(values))
		for i := range values {
			scan[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scan...)
			if *wrap.LogCpError(&err) == nil {
				result += `<option title="` + html.EscapeString(string(values[1])) + `" value="` + html.EscapeString(string(values[0])) + `">` + html.EscapeString(string(values[1])) + `</option>`
			}
		}
	}
	return result
}

func (this *Modules) shop_GetAllCurrencies(wrap *wrapper.Wrapper) map[int]string {
	result := map[int]string{}
	rows, err := wrap.DB.Query(
		wrap.R.Context(),
		`SELECT
			id,
			code
		FROM
			shop_currencies
		ORDER BY
			id ASC
		;`,
	)
	if err == nil {
		defer rows.Close()
		values := make([]string, 2)
		scan := make([]interface{}, len(values))
		for i := range values {
			scan[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scan...)
			if *wrap.LogCpError(&err) == nil {
				result[utils.StrToInt(string(values[0]))] = html.EscapeString(string(values[1]))
			}
		}
	}
	return result
}

func (this *Modules) shop_GetAllProductImages(wrap *wrapper.Wrapper, product_id int) string {
	result := ``
	rows, err := wrap.DB.Query(
		wrap.R.Context(),
		`SELECT
			id,
			product_id,
			filename
		FROM
			shop_product_images
		WHERE
			product_id = ?
		ORDER BY
			ord ASC
		;`,
		product_id,
	)
	if err == nil {
		defer rows.Close()
		values := make([]string, 3)
		scan := make([]interface{}, len(values))
		for i := range values {
			scan[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scan...)
			if *wrap.LogCpError(&err) == nil {
				result += `<div class="attached-img" data-id="` + html.EscapeString(string(values[0])) + `"><a href="/products/images/` + html.EscapeString(string(values[1])) + `/` + html.EscapeString(string(values[2])) + `" title="` + html.EscapeString(string(values[2])) + `" target="_blank"><img id="pimg_` + string(values[1]) + `_` + strings.Replace(string(values[2]), ".", "_", -1) + `" src="/products/images/` + string(values[1]) + `/thumb-0-` + string(values[2]) + `" onerror="WaitForFave(function(){fave.ShopProductsRetryImage(this, 'pimg_` + string(values[1]) + `_` + strings.Replace(string(values[2]), ".", "_", -1) + `');});" /></a><a class="remove" onclick="fave.ShopProductsDeleteImage(this, ` + html.EscapeString(string(values[1])) + `, '` + html.EscapeString(string(values[2])) + `');"><svg viewBox="1 1 11 14" width="10" height="12" class="sicon" version="1.1"><path fill-rule="evenodd" d="M11 2H9c0-.55-.45-1-1-1H5c-.55 0-1 .45-1 1H2c-.55 0-1 .45-1 1v1c0 .55.45 1 1 1v9c0 .55.45 1 1 1h7c.55 0 1-.45 1-1V5c.55 0 1-.45 1-1V3c0-.55-.45-1-1-1zm-1 12H3V5h1v8h1V5h1v8h1V5h1v8h1V5h1v9zm1-10H2V3h9v1z"></path></svg></a></div>`
			}
		}
	}
	return result
}

func (this *Modules) shop_GetSubProducts(wrap *wrapper.Wrapper, id int) string {
	result := ``
	rows, err := wrap.DB.Query(
		wrap.R.Context(),
		`SELECT
			id,
			name
		FROM
			shop_products
		WHERE
			parent_id = ?
		ORDER BY
			id DESC
		;`,
		id,
	)
	if err == nil {
		defer rows.Close()
		values := make([]string, 2)
		scan := make([]interface{}, len(values))
		for i := range values {
			scan[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scan...)
			if *wrap.LogCpError(&err) == nil {
				result += `<div><a href="/cp/` + wrap.CurrModule + `/modify/` + html.EscapeString(string(values[0])) + `/">` + html.EscapeString(string(values[1])) + ` ` + html.EscapeString(string(values[0])) + `</a> <a class="ico delete" title="Delete" href="javascript:fave.ActionDataTableDelete(this,'shop-detach','` + html.EscapeString(string(values[0])) + `','Are you sure want to detach product?');"><svg viewBox="0 0 16 16" width="16" height="16" class="sicon" version="1.1"><path fill-rule="evenodd" d="M11 2H9c0-.55-.45-1-1-1H5c-.55 0-1 .45-1 1H2c-.55 0-1 .45-1 1v1c0 .55.45 1 1 1v9c0 .55.45 1 1 1h7c.55 0 1-.45 1-1V5c.55 0 1-.45 1-1V3c0-.55-.45-1-1-1zm-1 12H3V5h1v8h1V5h1v8h1V5h1v8h1V5h1v9zm1-10H2V3h9v1z"></path></svg></a></div>`
			}
		}
	}
	return result
}

func (this *Modules) shop_GetParentProduct(wrap *wrapper.Wrapper, id int) string {
	result := ``
	rows, err := wrap.DB.Query(
		wrap.R.Context(),
		`SELECT
			id,
			name
		FROM
			shop_products
		WHERE
			id = ?
		;`,
		id,
	)
	if err == nil {
		defer rows.Close()
		values := make([]string, 2)
		scan := make([]interface{}, len(values))
		for i := range values {
			scan[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scan...)
			if *wrap.LogCpError(&err) == nil {
				result += `<div><a href="/cp/` + wrap.CurrModule + `/modify/` + html.EscapeString(string(values[0])) + `/">` + html.EscapeString(string(values[1])) + ` ` + html.EscapeString(string(values[0])) + `</a></div>`
			}
		}
	}
	return result
}

func (this *Modules) shop_GetOrderStatus(status int) string {
	if status == 0 {
		return `<span style="color:#f0ad4e;">New</span>`
	} else if status == 1 {
		return `<span style="color:#28a745;">Confirmed</span>`
	} else if status == 2 {
		return `<span style="color:#f0ad4e;">In progress</span>`
	} else if status == 3 {
		return `<span style="color:#d9534f;">Canceled</span>`
	} else if status == 4 {
		return `<span style="color:#6c757d;">Completed</span>`
	}
	return "Unknown"
}

func (this *Modules) RegisterModule_Shop() *Module {
	return this.newModule(MInfo{
		Mount:  "shop",
		Name:   "Магазин",
		Order:  2,
		System: false,
		Icon:   "<i class=\"material-icons notranslate\">shop</i>", //assets.SysSvgIconShop,
		Sub: &[]MISub{
			{Mount: "default", Name: "Список продуктов", Show: true, Icon: "<i class=\"material-icons notranslate\">list</i>" /*assets.SysSvgIconList*/},
			{Mount: "add", Name: "Добавить продукт", Show: true, Icon: "<i class=\"material-icons notranslate\">add</i>" /*assets.SysSvgIconPlus*/},
			{Mount: "modify", Name: "Редактировать", Show: false},
			{Sep: true, Show: true},
			{Mount: "categories", Name: "Список категорий", Show: true, Icon: "<i class=\"material-icons notranslate\">list</i>" /*assets.SysSvgIconList*/},
			{Mount: "categories-add", Name: "Добавить категорию", Show: true, Icon: "<i class=\"material-icons notranslate\">add</i>" /*assets.SysSvgIconPlus*/},
			{Mount: "categories-modify", Name: "Редактировать", Show: false},
			{Sep: true, Show: true},
			{Mount: "attributes", Name: "Список атрибутов", Show: true, Icon: "<i class=\"material-icons notranslate\">list</i>" /*assets.SysSvgIconList*/},
			{Mount: "attributes-add", Name: "Добавить атрибут", Show: true, Icon: "<i class=\"material-icons notranslate\">add</i>" /*assets.SysSvgIconPlus*/},
			{Mount: "attributes-modify", Name: "Редактировать", Show: false},
			{Sep: true, Show: true},
			{Mount: "currencies", Name: "Список валют", Show: true, Icon: "<i class=\"material-icons notranslate\">list</i>" /*assets.SysSvgIconList*/},
			{Mount: "currencies-add", Name: "Добавить валюту", Show: true, Icon: "<i class=\"material-icons notranslate\">add</i>" /*assets.SysSvgIconPlus*/},
			{Mount: "currencies-modify", Name: "Редактировать", Show: false},
			{Sep: true, Show: true},
			{Mount: "orders", Name: "Список заказов", Show: true, Icon: "<i class=\"material-icons notranslate\">list</i>" /*assets.SysSvgIconList*/},
			{Mount: "orders-modify", Name: "Просмотр заказа", Show: false},
		},
	}, func(wrap *wrapper.Wrapper) {
		if len(wrap.UrlArgs) == 3 && wrap.UrlArgs[0] == "shop" && wrap.UrlArgs[1] == "category" && wrap.UrlArgs[2] != "" {
			// Shop category
			row := &utils.Sql_shop_category{}
			rou := &utils.Sql_user{}
			err := wrap.DB.QueryRow(
				wrap.R.Context(),
				`SELECT
					main.id,
					main.user,
					main.name,
					main.alias,
					main.lft,
					main.rgt,
					main.depth,
					parent.id AS parent_id,
					users.id,
					users.first_name,
					users.last_name,
					users.email,
					users.admin,
					users.active
				FROM
					(
						SELECT
							node.id,
							node.user,
							node.name,
							node.alias,
							node.lft,
							node.rgt,
							(COUNT(parent.id) - 1) AS depth
						FROM
							shop_cats AS node,
							shop_cats AS parent
						WHERE
							node.lft BETWEEN parent.lft AND parent.rgt
						GROUP BY
							node.id
						ORDER BY
							node.lft ASC
					) AS main
					LEFT JOIN (
						SELECT
							node.id,
							node.user,
							node.name,
							node.alias,
							node.lft,
							node.rgt,
							(COUNT(parent.id) - 0) AS depth
						FROM
							shop_cats AS node,
							shop_cats AS parent
						WHERE
							node.lft BETWEEN parent.lft AND parent.rgt
						GROUP BY
							node.id
						ORDER BY
							node.lft ASC
					) AS parent ON
					parent.depth = main.depth AND
					main.lft > parent.lft AND
					main.rgt < parent.rgt
					LEFT JOIN users ON users.id = main.user
				WHERE
					main.id > 1 AND
					main.alias = ?
				ORDER BY
					main.lft ASC
				;`,
				wrap.UrlArgs[2],
			).Scan(
				&row.A_id,
				&row.A_user,
				&row.A_name,
				&row.A_alias,
				&row.A_lft,
				&row.A_rgt,
				&row.A_depth,
				&row.A_parent,
				&rou.A_id,
				&rou.A_first_name,
				&rou.A_last_name,
				&rou.A_email,
				&rou.A_admin,
				&rou.A_active,
			)

			if err != nil && err != wrapper.ErrNoRows {
				// System error 500
				wrap.LogCpError(&err)
				utils.SystemErrorPageEngine(wrap.W, err)
				return
			} else if err == wrapper.ErrNoRows {
				// User error 404 page
				wrap.RenderFrontEnd("404", fetdata.New(wrap, true, nil, nil), http.StatusNotFound)
				return
			}

			// Fix url
			if wrap.R.URL.Path[len(wrap.R.URL.Path)-1] != '/' {
				http.Redirect(wrap.W, wrap.R, wrap.R.URL.Path+"/"+utils.ExtractGetParams(wrap.R.RequestURI), 301)
				return
			}

			// Render template
			wrap.RenderFrontEnd("shop-category", fetdata.New(wrap, false, row, rou), http.StatusOK)
			return
		} else if len(wrap.UrlArgs) >= 3 && wrap.UrlArgs[0] == "shop" && wrap.UrlArgs[1] == "basket" && (wrap.UrlArgs[2] == "info" || wrap.UrlArgs[2] == "plus" || wrap.UrlArgs[2] == "minus" || wrap.UrlArgs[2] == "remove" || wrap.UrlArgs[2] == "currency") {
			SBParam := basket.SBParam{
				R:         wrap.R,
				DB:        wrap.DB,
				Host:      wrap.CurrHost,
				Config:    wrap.Config,
				SessionId: wrap.GetSessionId(),
			}
			if wrap.UrlArgs[2] == "info" {
				wrap.W.WriteHeader(http.StatusOK)
				wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				wrap.W.Header().Set("Content-Type", "application/json; charset=utf-8")
				wrap.W.Write([]byte(wrap.ShopBasket.Info(&SBParam)))
				wrap.S.SetString("LastBasketAction", wrap.UrlArgs[2])
				return
			} else if wrap.UrlArgs[2] == "plus" && len(wrap.UrlArgs) == 4 && utils.IsNumeric(wrap.UrlArgs[3]) {
				wrap.W.WriteHeader(http.StatusOK)
				wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				wrap.W.Header().Set("Content-Type", "application/json; charset=utf-8")
				wrap.W.Write([]byte(wrap.ShopBasket.Plus(&SBParam, utils.StrToInt(wrap.UrlArgs[3]))))
				wrap.S.SetString("LastBasketAction", wrap.UrlArgs[2])
				return
			} else if wrap.UrlArgs[2] == "minus" && len(wrap.UrlArgs) == 4 && utils.IsNumeric(wrap.UrlArgs[3]) {
				wrap.W.WriteHeader(http.StatusOK)
				wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				wrap.W.Header().Set("Content-Type", "application/json; charset=utf-8")
				wrap.W.Write([]byte(wrap.ShopBasket.Minus(&SBParam, utils.StrToInt(wrap.UrlArgs[3]))))
				wrap.S.SetString("LastBasketAction", wrap.UrlArgs[2])
				return
			} else if wrap.UrlArgs[2] == "remove" && len(wrap.UrlArgs) == 4 && utils.IsNumeric(wrap.UrlArgs[3]) {
				wrap.W.WriteHeader(http.StatusOK)
				wrap.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				wrap.W.Header().Set("Content-Type", "application/json; charset=utf-8")
				wrap.W.Write([]byte(wrap.ShopBasket.Remove(&SBParam, utils.StrToInt(wrap.UrlArgs[3]))))
				wrap.S.SetString("LastBasketAction", wrap.UrlArgs[2])
				return
			} else if wrap.UrlArgs[2] == "currency" && len(wrap.UrlArgs) == 4 && utils.IsNumeric(wrap.UrlArgs[3]) {
				http.SetCookie(wrap.W, &http.Cookie{
					Name:     "currency",
					Value:    wrap.UrlArgs[3],
					Path:     "/",
					Expires:  time.Now().Add(7 * 24 * time.Hour),
					HttpOnly: true,
				})
				redirectUrl := wrap.R.Referer()
				if redirectUrl == "" {
					redirectUrl = "/"
				}
				http.Redirect(wrap.W, wrap.R, redirectUrl, 302)
				wrap.S.SetString("LastBasketAction", wrap.UrlArgs[2])
				return
			}
		} else if len(wrap.UrlArgs) == 2 && wrap.UrlArgs[0] == "shop" && wrap.UrlArgs[1] != "" {
			// Shop product
			row := &utils.Sql_shop_product{}
			rou := &utils.Sql_user{}
			err := wrap.DB.QueryRow(
				wrap.R.Context(),
				// `SELECT
				// 	shop_products.id,
				// 	shop_products.parent_id,
				// 	shop_products.user,
				// 	shop_products.currency,
				// 	shop_products.price,
				// 	shop_products.price_old,
				// 	shop_products.name,
				// 	shop_products.alias,
				// 	shop_products.vendor,
				// 	shop_products.quantity,
				// 	shop_products.category,
				// 	shop_products.briefly,
				// 	shop_products.content,
				// 	UNIX_TIMESTAMP(shop_products.datetime) as datetime,
				// 	shop_products.active,
				// 	users.id,
				// 	users.first_name,
				// 	users.last_name,
				// 	users.email,
				// 	users.admin,
				// 	users.active
				// FROM
				// 	shop_products
				// 	LEFT JOIN users ON users.id = shop_products.user
				// WHERE
				// 	shop_products.active = 1 and
				// 	shop_products.alias = ?
				// LIMIT 1;`,
				`SELECT
					shop_products.id,
					shop_products.parent_id,
					shop_products.user,
					shop_products.currency,
					shop_products.price,
					shop_products.price_old,
					shop_products.name,
					shop_products.alias,
					shop_products.vendor,
					shop_products.quantity,
					shop_products.category,
					shop_products.briefly,
					shop_products.content,
					strftime('%s', shop_products.datetime) as datetime,
					shop_products.active,
					users.id,
					users.first_name,
					users.last_name,
					users.email,
					users.admin,
					users.active
				FROM
					shop_products
					LEFT JOIN users ON users.id = shop_products.user
				WHERE
					shop_products.active = 1 and
					shop_products.alias = ?
				LIMIT 1;`,
				wrap.UrlArgs[1],
			).Scan(
				&row.A_id,
				&row.A_parent,
				&row.A_user,
				&row.A_currency,
				&row.A_price,
				&row.A_price_old,
				&row.A_name,
				&row.A_alias,
				&row.A_vendor,
				&row.A_quantity,
				&row.A_category,
				&row.A_briefly,
				&row.A_content,
				&row.A_datetime,
				&row.A_active,
				&rou.A_id,
				&rou.A_first_name,
				&rou.A_last_name,
				&rou.A_email,
				&rou.A_admin,
				&rou.A_active,
			)

			if err != nil && err != wrapper.ErrNoRows {
				// System error 500
				wrap.LogCpError(&err)
				utils.SystemErrorPageEngine(wrap.W, err)
				return
			} else if err == wrapper.ErrNoRows {
				// User error 404 page
				wrap.RenderFrontEnd("404", fetdata.New(wrap, true, nil, nil), http.StatusNotFound)
				return
			}

			// Fix url
			if wrap.R.URL.Path[len(wrap.R.URL.Path)-1] != '/' {
				http.Redirect(wrap.W, wrap.R, wrap.R.URL.Path+"/"+utils.ExtractGetParams(wrap.R.RequestURI), 301)
				return
			}

			// Render template
			wrap.RenderFrontEnd("shop-product", fetdata.New(wrap, false, row, rou), http.StatusOK)
			return
		} else if len(wrap.UrlArgs) == 1 && wrap.UrlArgs[0] == "shop" {
			// Shop

			// Fix url
			if wrap.R.URL.Path[len(wrap.R.URL.Path)-1] != '/' {
				http.Redirect(wrap.W, wrap.R, wrap.R.URL.Path+"/"+utils.ExtractGetParams(wrap.R.RequestURI), 301)
				return
			}

			// Render template
			wrap.RenderFrontEnd("shop", fetdata.New(wrap, false, nil, nil), http.StatusOK)
			return
		} else if (*wrap.Config).Engine.MainModule == 2 {
			// Render template
			wrap.RenderFrontEnd("shop", fetdata.New(wrap, false, nil, nil), http.StatusOK)
			return
		}

		// User error 404 page
		wrap.RenderFrontEnd("404", fetdata.New(wrap, true, nil, nil), http.StatusNotFound)
	}, func(wrap *wrapper.Wrapper) (string, string, string) {
		content := ""
		sidebar := ""
		if wrap.CurrSubModule == "" || wrap.CurrSubModule == "default" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "List of products"},
			})

			// Load currencies
			currencies := this.shop_GetAllCurrencies(wrap)

			content += builder.DataTable(
				wrap,
				"shop_products",
				"id",
				"DESC",
				&[]builder.DataTableRow{
					{
						DBField: "id",
					},
					{
						DBField:     "name",
						NameInTable: "Product / URL",
						CallBack: func(values *[]string) string {
							name := `<a href="/cp/` + wrap.CurrModule + `/modify/` + (*values)[0] + `/">` + html.EscapeString((*values)[1]) + ` ` + html.EscapeString((*values)[0]) + `</a>`
							alias := html.EscapeString((*values)[2])
							parent := ``
							outofstock := ``
							if (*values)[7] != "" {
								parent = `<div class="parent">&uarr;<small><a href="/cp/` + wrap.CurrModule + `/modify/` + (*values)[7] + `/">` + html.EscapeString((*values)[8]) + ` ` + (*values)[7] + `</a></small></div>`
							}
							if utils.StrToInt((*values)[10]) <= 0 {
								outofstock = `<div><span class="badge badge-primary">Out of stock</span></div>`
							}
							return `<div>` + name + `</div><div><small>/shop/` + alias + `/</small></div>` + parent + outofstock
						},
					},
					{
						DBField: "alias",
					},
					{
						DBField: "currency",
					},
					{
						DBField:     "price",
						NameInTable: "Price",
						Classes:     "d-none d-md-table-cell",
						CallBack: func(values *[]string) string {
							price_old := ""
							price_promo := ""
							price_styles := ""
							if utils.StrToFloat64((*values)[9]) > 0 {
								price_old = `<div title="Old price"><strike>` + utils.Float64ToStr(utils.StrToFloat64((*values)[9])) + `</strike></div>`
								price_styles = ` style="color:#fb3f4c;"`
							}
							if utils.StrToFloat64((*values)[11]) > 0 {
								price_promo = `<div style="color:#ffc107;" title="Promo price">` + utils.Float64ToStr(utils.StrToFloat64((*values)[11])) + `</div>`
							}
							return price_promo + price_old + `<div` + price_styles + `title="Current price">` + utils.Float64ToStr(utils.StrToFloat64((*values)[4])) + `</div>` +
								`<div><small>` + currencies[utils.StrToInt((*values)[3])] + `</small></div>`
						},
					},
					{
						DBField: "datetime",
						//DBExp:       "UNIX_TIMESTAMP(`datetime`)",
						DBExp:       "strftime('%s', datetime)",
						NameInTable: "Date / Time",
						Classes:     "d-none d-lg-table-cell",
						CallBack: func(values *[]string) string {
							t := int64(utils.StrToInt((*values)[5]))
							return `<div>` + utils.UnixTimestampToFormat(t, "02.01.2006") + `</div>` +
								`<div><small>` + utils.UnixTimestampToFormat(t, "15:04:05") + `</small></div>`
						},
					},
					{
						DBField:     "active",
						NameInTable: "Active",
						Classes:     "d-none d-sm-table-cell",
						CallBack: func(values *[]string) string {
							return builder.CheckBox(utils.StrToInt((*values)[6]))
						},
					},
					{
						DBField: "parent_id",
					},
					{
						DBField: "pname",
						DBExp:   "spp.name",
					},
					{
						DBField: "price_old",
					},
					{
						DBField: "quantity",
					},
					{
						DBField: "price_promo",
					},
				},
				func(values *[]string) string {
					return builder.DataTableAction(&[]builder.DataTableActionRow{
						{
							Icon:   assets.SysSvgIconView,
							Href:   `/shop/` + (*values)[2] + `/`,
							Hint:   "View",
							Target: "_blank",
						},
						{
							Icon: assets.SysSvgIconEdit,
							Href: "/cp/" + wrap.CurrModule + "/modify/" + (*values)[0] + "/",
							Hint: "Edit",
						},
						{
							Icon: assets.SysSvgIconRemove,
							Href: "javascript:fave.ActionDataTableDelete(this,'shop-delete','" +
								(*values)[0] + "','Are you sure want to delete product?');",
							Hint:    "Delete",
							Classes: "delete",
						},
					})
				},
				"/cp/"+wrap.CurrModule+"/",
				func() (int, error) {
					var count int
					return count, wrap.DB.QueryRow(
						wrap.R.Context(),
						"SELECT COUNT(*) FROM `shop_products`;",
					).Scan(&count)
				},
				func(limit_offset int, pear_page int) (*sqlw.Rows, error) {
					return wrap.DB.Query(
						wrap.R.Context(),
						// `SELECT
						// 	shop_products.id,
						// 	shop_products.name,
						// 	shop_products.alias,
						// 	shop_products.currency,
						// 	shop_products.price,
						// 	UNIX_TIMESTAMP(`+"`shop_products`.`datetime`"+`) AS datetime,
						// 	shop_products.active,
						// 	shop_products.parent_id,
						// 	spp.name AS pname,
						// 	shop_products.price_old,
						// 	shop_products.quantity,
						// 	shop_products.price_promo
						// FROM
						// 	shop_products
						// 	LEFT JOIN shop_products AS spp ON spp.id = shop_products.parent_id
						// ORDER BY
						// 	shop_products.id DESC
						// LIMIT ?, ?;`,
						`SELECT
							shop_products.id,
							shop_products.name,
							shop_products.alias,
							shop_products.currency,
							shop_products.price,
							strftime('%s', shop_products.datetime) as datetime,
							shop_products.active,
							shop_products.parent_id,
							spp.name AS pname,
							shop_products.price_old,
							shop_products.quantity,
							shop_products.price_promo
						FROM
							shop_products
							LEFT JOIN shop_products AS spp ON spp.id = shop_products.parent_id
						ORDER BY
							shop_products.id DESC
						LIMIT ?, ?;`,
						limit_offset,
						pear_page,
					)
				},
				true,
			)
		} else if wrap.CurrSubModule == "categories" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Categories", Link: "/cp/" + wrap.CurrModule + "/" + wrap.CurrSubModule + "/"},
				{Name: "List of categories"},
			})
			content += builder.DataTable(
				wrap,
				"shop_cats",
				"id",
				"ASC",
				&[]builder.DataTableRow{
					{
						DBField: "id",
					},
					{
						DBField: "user",
					},
					{
						DBField:     "name",
						NameInTable: "Category",
						CallBack: func(values *[]string) string {
							depth := utils.StrToInt((*values)[4]) - 1
							if depth < 0 {
								depth = 0
							}
							sub := strings.Repeat("&mdash; ", depth)
							name := `<a href="/cp/` + wrap.CurrModule + `/categories-modify/` + (*values)[0] + `/">` + sub + html.EscapeString((*values)[2]) + `</a>`
							return `<div>` + name + `</div>`
						},
					},
					{
						DBField: "alias",
					},
					{
						DBField: "depth",
					},
				},
				func(values *[]string) string {
					return builder.DataTableAction(&[]builder.DataTableActionRow{
						{
							Icon:   assets.SysSvgIconView,
							Href:   `/shop/category/` + (*values)[3] + `/`,
							Hint:   "View",
							Target: "_blank",
						},
						{
							Icon: assets.SysSvgIconEdit,
							Href: "/cp/" + wrap.CurrModule + "/categories-modify/" + (*values)[0] + "/",
							Hint: "Edit",
						},
						{
							Icon: assets.SysSvgIconRemove,
							Href: "javascript:fave.ActionDataTableDelete(this,'shop-categories-delete','" +
								(*values)[0] + "','Are you sure want to delete category?');",
							Hint:    "Delete",
							Classes: "delete",
						},
					})
				},
				"/cp/"+wrap.CurrModule+"/"+wrap.CurrSubModule+"/",
				nil,
				func(limit_offset int, pear_page int) (*sqlw.Rows, error) {
					return wrap.DB.Query(
						wrap.R.Context(),
						`SELECT
							node.id,
							node.user,
							node.name,
							node.alias,
							(COUNT(parent.id) - 1) AS depth
						FROM
							shop_cats AS node,
							shop_cats AS parent
						WHERE
							node.lft BETWEEN parent.lft AND parent.rgt AND
							node.id > 1
						GROUP BY
							node.id
						ORDER BY
							node.lft ASC
						;`,
					)
				},
				false,
			)
		} else if wrap.CurrSubModule == "attributes" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Attributes", Link: "/cp/" + wrap.CurrModule + "/" + wrap.CurrSubModule + "/"},
				{Name: "List of attributes"},
			})
			content += builder.DataTable(
				wrap,
				"shop_filters",
				"id",
				"DESC",
				&[]builder.DataTableRow{
					{
						DBField: "id",
					},
					{
						DBField:     "name",
						NameInTable: "Name",
						CallBack: func(values *[]string) string {
							name := `<a href="/cp/` + wrap.CurrModule + `/attributes-modify/` + (*values)[0] + `/">` + html.EscapeString((*values)[1]) + `</a>`
							return `<div>` + name + `</div><div><small>` + html.EscapeString((*values)[2]) + `</small></div>`
						},
					},
					{
						DBField: "filter",
					},
				},
				func(values *[]string) string {
					return builder.DataTableAction(&[]builder.DataTableActionRow{
						{
							Icon: assets.SysSvgIconEdit,
							Href: "/cp/" + wrap.CurrModule + "/attributes-modify/" + (*values)[0] + "/",
							Hint: "Edit",
						},
						{
							Icon: assets.SysSvgIconRemove,
							Href: "javascript:fave.ActionDataTableDelete(this,'shop-attributes-delete','" +
								(*values)[0] + "','Are you sure want to delete attribute?');",
							Hint:    "Delete",
							Classes: "delete",
						},
					})
				},
				"/cp/"+wrap.CurrModule+"/"+wrap.CurrSubModule+"/",
				nil,
				nil,
				true,
			)
		} else if wrap.CurrSubModule == "currencies" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Currencies", Link: "/cp/" + wrap.CurrModule + "/" + wrap.CurrSubModule + "/"},
				{Name: "List of currencies"},
			})
			content += builder.DataTable(
				wrap,
				"shop_currencies",
				"id",
				"DESC",
				&[]builder.DataTableRow{
					{
						DBField: "id",
					},
					{
						DBField:     "name",
						NameInTable: "Name",
						CallBack: func(values *[]string) string {
							name := `<a href="/cp/` + wrap.CurrModule + `/currencies-modify/` + (*values)[0] + `/">` + html.EscapeString((*values)[1]) + ` (` + (*values)[3] + `, ` + (*values)[4] + `)</a>`
							return `<div>` + name + `</div>`
						},
					},
					{
						DBField:     "coefficient",
						NameInTable: "Coefficient",
						Classes:     "d-none d-md-table-cell",
						CallBack: func(values *[]string) string {
							return utils.Float64ToStrF(utils.StrToFloat64((*values)[2]), "%.4f")
						},
					},
					{
						DBField: "code",
					},
					{
						DBField: "symbol",
					},
				},
				func(values *[]string) string {
					return builder.DataTableAction(&[]builder.DataTableActionRow{
						{
							Icon: assets.SysSvgIconEdit,
							Href: "/cp/" + wrap.CurrModule + "/currencies-modify/" + (*values)[0] + "/",
							Hint: "Edit",
						},
						{
							Icon: assets.SysSvgIconRemove,
							Href: "javascript:fave.ActionDataTableDelete(this,'shop-currencies-delete','" +
								(*values)[0] + "','Are you sure want to delete currency?');",
							Hint:    "Delete",
							Classes: "delete",
						},
					})
				},
				"/cp/"+wrap.CurrModule+"/"+wrap.CurrSubModule+"/",
				nil,
				nil,
				true,
			)
		} else if wrap.CurrSubModule == "orders" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Orders", Link: "/cp/" + wrap.CurrModule + "/" + wrap.CurrSubModule + "/"},
				{Name: "List of orders"},
			})
			content += builder.DataTable(
				wrap,
				"shop_orders",
				"id",
				"DESC",
				&[]builder.DataTableRow{
					{
						DBField:     "id",
						NameInTable: "Order #",
						Classes:     "d-none d-lg-table-cell",
					},
					{
						DBField: "client_phone",
					},
					{
						DBField: "update_datetime",
					},
					{
						DBField: "currency_id",
					},
					{
						DBField: "currency_name",
					},
					{
						DBField: "currency_coefficient",
					},
					{
						DBField: "currency_code",
					},
					{
						DBField: "currency_symbol",
					},
					{
						DBField:     "client_last_name",
						NameInTable: "Client / Contact",
						CallBack: func(values *[]string) string {
							link := "/cp/" + wrap.CurrModule + "/orders-modify/" + (*values)[0] + "/"

							last_name := html.EscapeString((*values)[8])
							first_name := html.EscapeString((*values)[9])
							middle_name := html.EscapeString((*values)[10])

							phone := html.EscapeString((*values)[1])
							email := html.EscapeString((*values)[12])

							order_id := `<span class="d-inline d-sm-none">#` + (*values)[0] + `. </span>`

							name := ""
							if last_name != "" {
								name += " " + last_name
							}
							if first_name != "" {
								name += " " + first_name
							}
							if middle_name != "" {
								name += " " + middle_name
							}
							name = `<a href="` + link + `">` + order_id + utils.Trim(name) + `</a>`

							contact := ""
							if email != "" {
								contact += " " + email
							}
							if phone != "" {
								contact += " (" + phone + ")"
							}
							contact = `<a href="` + link + `">` + utils.Trim(contact) + `</a>`

							return `<div>` + name + `</div><div><small>` + contact + `</small></div>`
						},
					},
					{
						DBField: "client_first_name",
					},
					{
						DBField: "client_middle_name",
					},
					{
						DBField: "create_datetime",
						// DBExp:       "UNIX_TIMESTAMP(`create_datetime`)",
						DBExp:       "strftime('%s', datetime)",
						NameInTable: "Date / Time",
						Classes:     "d-none d-lg-table-cell",
						CallBack: func(values *[]string) string {
							t := int64(utils.StrToInt((*values)[11]))
							return `<div>` + utils.UnixTimestampToFormat(t, "02.01.2006") + `</div>` +
								`<div><small>` + utils.UnixTimestampToFormat(t, "15:04:05") + `</small></div>`
						},
					},
					{
						DBField:     "client_email",
						NameInTable: "Status / Total",
						CallBack: func(values *[]string) string {
							status := this.shop_GetOrderStatus(utils.StrToInt((*values)[15]))
							total := utils.Float64ToStr(utils.StrToFloat64((*values)[16])) + " " + html.EscapeString((*values)[6])
							return `<div>` + status + `</div><div><small>` + total + `</small></div>`
						},
					},
					{
						DBField: "client_delivery_comment",
					},
					{
						DBField: "client_order_comment",
					},
					{
						DBField: "status",
					},
					{
						DBField: "total",
					},
				},
				nil,
				"/cp/"+wrap.CurrModule+"/"+wrap.CurrSubModule+"/",
				nil,
				func(limit_offset int, pear_page int) (*sqlw.Rows, error) {
					return wrap.DB.Query(
						wrap.R.Context(),
						`SELECT
							shop_orders.id,
							shop_orders.client_phone,
							shop_orders.update_datetime,
							shop_orders.currency_id,
							shop_orders.currency_name,
							shop_orders.currency_coefficient,
							shop_orders.currency_code,
							shop_orders.currency_symbol,
							shop_orders.client_last_name,
							shop_orders.client_first_name,
							shop_orders.client_middle_name,
							strftime('%s', shop_orders.create_datetime) as create_datetime, 
							shop_orders.client_email,
							shop_orders.client_delivery_comment,
							shop_orders.client_order_comment,
							shop_orders.status,
							shop_order_total.total
						FROM
							shop_orders
							LEFT JOIN (
								SELECT
									order_id,
									SUM(price * quantity) as total
								FROM
									shop_order_products
								GROUP BY
									order_id
							) as shop_order_total ON shop_order_total.order_id = shop_orders.id
						ORDER BY
							shop_orders.status ASC,
							shop_orders.id DESC
						LIMIT ?, ?;`,
						limit_offset,
						pear_page,
					)
				},
				true,
			)
		} else if wrap.CurrSubModule == "add" || wrap.CurrSubModule == "modify" {
			if wrap.CurrSubModule == "add" {
				content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
					{Name: "Add new product"},
				})
			} else {
				if len(wrap.UrlArgs) >= 3 && utils.IsNumeric(wrap.UrlArgs[2]) {
					content += `<div class="product-copy"><a title="Duplicate product" href="javascript:fave.ShopProductsDuplicate(this, ` + wrap.UrlArgs[2] + `);">` + assets.SysSvgIconCopy + `</a></div>`
					content += `<div class="product-another"><a title="Duplicate product and attach" href="javascript:fave.ShopProductsDuplicateWithAttach(this, ` + wrap.UrlArgs[2] + `);">` + assets.SysSvgIconPlus + `</a></div>`
				}
				content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
					{Name: "Modify product"},
				})
			}

			data := utils.Sql_shop_product{
				A_id:        0,
				A_user:      0,
				A_currency:  0,
				A_price:     0,
				A_price_old: 0,
				A_gname:     "",
				A_name:      "",
				A_alias:     "",
				A_vendor:    "",
				A_quantity:  0,
				A_category:  0,
				A_briefly:   "",
				A_content:   "",
				A_datetime:  0,
				A_active:    0,
				A_custom1:   "",
				A_custom2:   "",
			}

			if wrap.CurrSubModule == "modify" {
				if len(wrap.UrlArgs) != 3 {
					return "", "", ""
				}
				if !utils.IsNumeric(wrap.UrlArgs[2]) {
					return "", "", ""
				}
				err := wrap.DB.QueryRow(
					wrap.R.Context(),
					`SELECT
						id,
						parent_id,
						user,
						currency,
						price,
						price_old,
						price_promo,
						gname,
						name,
						alias,
						vendor,
						quantity,
						category,
						briefly,
						content,
						active,
						custom1,
						custom2
					FROM
						shop_products
					WHERE
						id = ?
					LIMIT 1;`,
					utils.StrToInt(wrap.UrlArgs[2]),
				).Scan(
					&data.A_id,
					&data.A_parent,
					&data.A_user,
					&data.A_currency,
					&data.A_price,
					&data.A_price_old,
					&data.A_price_promo,
					&data.A_gname,
					&data.A_name,
					&data.A_alias,
					&data.A_vendor,
					&data.A_quantity,
					&data.A_category,
					&data.A_briefly,
					&data.A_content,
					&data.A_active,
					&data.A_custom1,
					&data.A_custom2,
				)
				if *wrap.LogCpError(&err) != nil {
					return "", "", ""
				}
			}

			if data.A_parent_id() > 0 {
				content += `<style>.product-another{display:none}</style>`
			}

			// All product current categories
			var selids []int
			if data.A_id > 0 {
				rows, err := wrap.DB.Query(wrap.R.Context(), "SELECT category_id FROM shop_cat_product_rel WHERE product_id = ?;", data.A_id)
				if err == nil {
					defer rows.Close()
					values := make([]int, 1)
					scan := make([]interface{}, len(values))
					for i := range values {
						scan[i] = &values[i]
					}
					for rows.Next() {
						err = rows.Scan(scan...)
						if *wrap.LogCpError(&err) == nil {
							selids = append(selids, int(values[0]))
						}
					}
				}
			}

			// Sub products
			sub_products := ``
			if data.A_id >= 1 && data.A_parent_id() <= 0 {
				sub_products = this.shop_GetSubProducts(wrap, data.A_id)
			}

			btn_caption := "Add"
			if wrap.CurrSubModule == "modify" {
				btn_caption = "Save"
			}

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "shop-modify",
				},
				{
					Kind:  builder.DFKHidden,
					Name:  "id",
					Value: utils.IntToStr(data.A_id),
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						if data.A_id >= 1 && data.A_parent_id() <= 0 {
							return `<div class="form-group nf">` +
								`<div class="row">` +
								`<div class="col-md-3">` +
								`<label>Sub products</label>` +
								`</div>` +
								`<div class="col-md-9">` +
								`<div class="list-wrapper">` +
								sub_products +
								`<div><a href="javascript:fave.ShopAttachProduct(` + utils.IntToStr(data.A_id) + `);"><b>Attach product</b></a></div>` +
								`</div>` +
								`</div>` +
								`</div>` +
								`</div>`
						}
						return ""
					},
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						if data.A_id >= 1 && data.A_parent_id() != 0 {
							return `<div class="form-group nf">` +
								`<div class="row">` +
								`<div class="col-md-3">` +
								`<label>Parent</label>` +
								`</div>` +
								`<div class="col-md-9">` +
								`<div class="list-wrapper">` +
								this.shop_GetParentProduct(wrap, data.A_parent_id()) +
								`</div>` +
								`</div>` +
								`</div>` +
								`</div>`
						}
						return ""
					},
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						if data.A_id >= 1 && data.A_parent_id() <= 0 && sub_products != "" {
							return `<div class="form-group nf">` +
								`<div class="row">` +
								`<div class="col-md-3">` +
								`<label for="lbl_gname">Group name</label>` +
								`</div>` +
								`<div class="col-md-9">` +
								`<div><input class="form-control" type="text" id="lbl_gname" name="gname" value="` + html.EscapeString(data.A_gname) + `" maxlength="255" autocomplete="off"></div>` +
								`</div>` +
								`</div>` +
								`</div>`
						}
						return ""
					},
				},
				{
					Kind:     builder.DFKText,
					Caption:  "Product name",
					Name:     "name",
					Value:    data.A_name,
					Required: true,
					Min:      "1",
					Max:      "255",
				},
				{
					Kind:    builder.DFKText,
					Caption: "Product price",
					Name:    "price",
					Value:   "0",
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group n6">` +
							`<div class="row">` +
							`<div class="col-md-3">` +
							`<label for="lbl_price">Product price</label>` +
							`</div>` +
							`<div class="col-md-9">` +
							`<div>` +
							`<div class="row">` +
							`<div class="col-md-8">` +
							`<div class="row">` +
							`<div class="col-md-12">` +
							`<div><input class="form-control" type="number" step="0.01" id="lbl_price" name="price" value="` + utils.Float64ToStr(data.A_price) + `" placeholder="" autocomplete="off" required></div>` +
							`</div>` +
							`</div>` +
							`<div class="d-md-none mb-3"></div>` +
							`</div>` +
							`<div class="col-md-4">` +
							`<select class="selectpicker form-control" id="lbl_currency" name="currency" data-live-search="true">` +
							this.shop_GetCurrencySelectOptions(wrap, data.A_currency) +
							`</select>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>`
					},
				},
				{
					Kind:    builder.DFKText,
					Caption: "Old/Promo price",
					Name:    "price_old",
					Value:   "0",
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group n6">` +
							`<div class="row">` +
							`<div class="col-md-3">` +
							`<label for="lbl_price_old">Old/Promo price</label>` +
							`</div>` +
							`<div class="col-md-9">` +
							`<div>` +
							`<div class="row">` +
							`<div class="col-md-12">` +
							`<div class="row">` +
							`<div class="col-md-6">` +
							`<div><input class="form-control" type="number" id="lbl_price_old" name="price_old" value="` + utils.Float64ToStr(data.A_price_old) + `" placeholder="" autocomplete="off"></div>` +
							`<div class="d-md-none mb-3"></div>` +
							`</div>` +
							`<div class="col-md-6">` +
							`<div><input class="form-control" type="number" id="lbl_price_promo" name="price_promo" value="` + utils.Float64ToStr(data.A_price_promo) + `" placeholder="" autocomplete="off"></div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>`
					},
				},
				{
					Kind:    builder.DFKText,
					Caption: "Product alias",
					Name:    "alias",
					Value:   data.A_alias,
					Hint:    "Example: mobile-phone",
					Max:     "255",
				},
				{
					Kind:    builder.DFKText,
					Caption: "Vendor/Count",
					Name:    "vendor",
					Value:   "0",
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group n8">` +
							`<div class="row">` +
							`<div class="col-md-3">` +
							`<label for="lbl_vendor">Vendor/Count</label>` +
							`</div>` +
							`<div class="col-md-9">` +
							`<div>` +
							`<div class="row">` +
							`<div class="col-md-8">` +
							`<div><input class="form-control" type="text" id="lbl_vendor" name="vendor" value="` + html.EscapeString(data.A_vendor) + `" placeholder="" autocomplete="off"></div>` +
							`<div class="d-md-none mb-3"></div>` +
							`</div>` +
							`<div class="col-md-4">` +
							`<input class="form-control" type="number" step="1" id="lbl_quantity" name="quantity" value="` + utils.IntToStr(data.A_quantity) + `" placeholder="" autocomplete="off">` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>`
					},
				},
				{
					Kind:    builder.DFKText,
					Caption: "Category",
					Name:    "category",
					Value:   "0",
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group n9">` +
							`<div class="row">` +
							`<div class="col-md-3">` +
							`<label for="lbl_category">Category</label>` +
							`</div>` +
							`<div class="col-md-9">` +
							`<div>` +
							`<select class="selectpicker form-control" id="lbl_category" name="category" data-live-search="true">` +
							`<option title="Nothing selected" value="0">&mdash;</option>` +
							this.shop_GetCategorySelectOptions(wrap, 0, data.A_category, []int{}) +
							`</select>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>`
					},
				},
				{
					Kind:    builder.DFKText,
					Caption: "Categories",
					Name:    "cats",
					Value:   "0",
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group n10">` +
							`<div class="row">` +
							`<div class="col-md-3">` +
							`<label for="lbl_cats">Categories</label>` +
							`</div>` +
							`<div class="col-md-9">` +
							`<div>` +
							`<select class="selectpicker form-control" id="lbl_cats" name="cats[]" data-live-search="true" multiple>` +
							this.shop_GetCategorySelectOptions(wrap, 0, 0, selids) +
							`</select>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>`
					},
				},
				{
					Kind:    builder.DFKText,
					Caption: "Attributes",
					Name:    "",
					Value:   "",
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group n11">` +
							`<div class="row">` +
							`<div class="col-md-3">` +
							`<label>Attributes</label>` +
							`</div>` +
							`<div class="col-md-9">` +
							`<div class="list-wrapper">` +
							`<div id="list">` +
							this.shop_GetProductValuesInputs(wrap, data.A_id) +
							`</div>` +
							`<div class="list-button position-relative">` +
							`<select class="selectpicker form-control" id="lbl_attributes" data-live-search="true" onchange="fave.ShopProductsAdd();">` +
							this.shop_GetAllAttributesSelectOptions(wrap) +
							`</select>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>`
					},
				},
				{
					Kind:    builder.DFKTextArea,
					Caption: "Briefly",
					Name:    "briefly",
					Value:   data.A_briefly,
					Classes: "briefly wysiwyg",
				},
				{
					Kind:    builder.DFKTextArea,
					Caption: "Product content",
					Name:    "content",
					Value:   data.A_content,
					Classes: "wysiwyg",
				},
				{
					Kind:    builder.DFKText,
					Caption: "Product images",
					Name:    "",
					Value:   "",
					CallBack: func(field *builder.DataFormField) string {
						if data.A_id == 0 {
							return ``
						}
						return `<div class="form-group n14">` +
							`<div class="row">` +
							`<div class="col-md-3">` +
							`<label>Product images</label>` +
							`</div>` +
							`<div class="col-md-9">` +
							`<div class="list-wrapper">` +
							`<div id="list-images">` +
							this.shop_GetAllProductImages(wrap, data.A_id) +
							`</div>` +
							`<div id="img-upload-block" class="list-button position-relative">` +
							`<div id="upload-msg">Uploading...</div>` +
							`<input class="form-control ignore-lost-data" type="file" id="file" name="file" onchange="fave.ShopProductsUploadImage('shop-upload-image', ` + utils.IntToStr(data.A_id) + `, 'file');" style="font-size:12px;" multiple />` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`<script>WaitForFave(function(){Sortable.create(document.getElementById('list-images'),{animation:0,onEnd:function(evt){var orderData=[];$('#list-images div.attached-img').each(function(i,v){orderData.push({Id:$(v).data('id'),Order:i+1});});$('#list-images').addClass('loading');fave.ShopProductsImageReorder('shop-images-reorder',{Items:orderData});},});});</script>`
					},
				},
				{
					Kind:    builder.DFKText,
					Caption: "Custom field 1",
					Name:    "",
					Value:   "",
					CallBack: func(field *builder.DataFormField) string {
						if (*wrap.Config).Shop.CustomFields.Field1.Enabled <= 0 {
							return ``
						}
						return `<div class="form-group nf">
							<div class="row">
								<div class="col-md-3">
									<label for="lbl_custom1">` + (*wrap.Config).Shop.CustomFields.Field1.Caption + `</label>
								</div>
								<div class="col-md-9">
									<div><input class="form-control" type="text" id="lbl_custom1" name="custom1" value="` + html.EscapeString(data.A_custom1) + `" maxlength="2048" autocomplete="off"></div>
								</div>
							</div>
						</div>`
					},
				},
				{
					Kind:    builder.DFKText,
					Caption: "Custom field 2",
					Name:    "",
					Value:   "",
					CallBack: func(field *builder.DataFormField) string {
						if (*wrap.Config).Shop.CustomFields.Field2.Enabled <= 0 {
							return ``
						}
						return `<div class="form-group nf">
							<div class="row">
								<div class="col-md-3">
									<label for="lbl_custom2">` + (*wrap.Config).Shop.CustomFields.Field2.Caption + `</label>
								</div>
								<div class="col-md-9">
									<div><input class="form-control" type="text" id="lbl_custom2" name="custom2" value="` + html.EscapeString(data.A_custom2) + `" maxlength="2048" autocomplete="off"></div>
								</div>
							</div>
						</div>`
					},
				},
				{
					Kind:    builder.DFKCheckBox,
					Caption: "Active",
					Name:    "active",
					Value:   utils.IntToStr(data.A_active),
				},
				{
					Kind:   builder.DFKSubmit,
					Value:  btn_caption,
					Target: "add-edit-button",
				},
			})

			if wrap.CurrSubModule == "add" {
				sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Add</button>`
			} else {
				sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
			}
		} else if wrap.CurrSubModule == "categories-add" || wrap.CurrSubModule == "categories-modify" {
			if wrap.CurrSubModule == "categories-add" {
				content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
					{Name: "Categories", Link: "/cp/" + wrap.CurrModule + "/categories/"},
					{Name: "Add new category"},
				})
			} else {
				content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
					{Name: "Categories", Link: "/cp/" + wrap.CurrModule + "/categories/"},
					{Name: "Modify category"},
				})
			}

			data := utils.Sql_shop_category{
				A_id:    0,
				A_user:  0,
				A_name:  "",
				A_alias: "",
				A_lft:   0,
				A_rgt:   0,
			}

			if wrap.CurrSubModule == "categories-modify" {
				if len(wrap.UrlArgs) != 3 {
					return "", "", ""
				}
				if !utils.IsNumeric(wrap.UrlArgs[2]) {
					return "", "", ""
				}
				err := wrap.DB.QueryRow(
					wrap.R.Context(),
					`SELECT
						id,
						user,
						name,
						alias,
						lft,
						rgt
					FROM
						shop_cats
					WHERE
						id = ?
					LIMIT 1;`,
					utils.StrToInt(wrap.UrlArgs[2]),
				).Scan(
					&data.A_id,
					&data.A_user,
					&data.A_name,
					&data.A_alias,
					&data.A_lft,
					&data.A_rgt,
				)
				if *wrap.LogCpError(&err) != nil {
					return "", "", ""
				}
			}

			btn_caption := "Add"
			if wrap.CurrSubModule == "categories-modify" {
				btn_caption = "Save"
			}

			parentId := 0
			if wrap.CurrSubModule == "categories-modify" {
				parentId = this.shop_GetCategoryParentId(wrap, data.A_id)
			}

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "shop-categories-modify",
				},
				{
					Kind:  builder.DFKHidden,
					Name:  "id",
					Value: utils.IntToStr(data.A_id),
				},
				{
					Kind:    builder.DFKText,
					Caption: "Parent",
					Name:    "parent",
					Value:   "0",
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group n2">` +
							`<div class="row">` +
							`<div class="col-md-3">` +
							`<label for="lbl_parent">Parent</label>` +
							`</div>` +
							`<div class="col-md-9">` +
							`<div>` +
							`<select class="selectpicker form-control" id="lbl_parent" name="parent" data-live-search="true">` +
							`<option title="Nothing selected" value="0">&mdash;</option>` +
							this.shop_GetCategorySelectOptions(wrap, data.A_id, parentId, []int{}) +
							`</select>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>`
					},
				},
				{
					Kind:     builder.DFKText,
					Caption:  "Name",
					Name:     "name",
					Value:    data.A_name,
					Required: true,
					Min:      "1",
					Max:      "255",
				},
				{
					Kind:    builder.DFKText,
					Caption: "Alias",
					Name:    "alias",
					Value:   data.A_alias,
					Hint:    "Example: popular-products",
					Max:     "255",
				},
				{
					Kind:   builder.DFKSubmit,
					Value:  btn_caption,
					Target: "add-edit-button",
				},
			})

			if wrap.CurrSubModule == "categories-add" {
				sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Add</button>`
			} else {
				sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
			}
		} else if wrap.CurrSubModule == "attributes-add" || wrap.CurrSubModule == "attributes-modify" {
			if wrap.CurrSubModule == "attributes-add" {
				content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
					{Name: "Attributes", Link: "/cp/" + wrap.CurrModule + "/attributes/"},
					{Name: "Add new attribute"},
				})
			} else {
				content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
					{Name: "Attributes", Link: "/cp/" + wrap.CurrModule + "/attributes/"},
					{Name: "Modify attribute"},
				})
			}

			data := utils.Sql_shop_filter{
				A_id:     0,
				A_name:   "",
				A_filter: "",
			}

			if wrap.CurrSubModule == "attributes-modify" {
				if len(wrap.UrlArgs) != 3 {
					return "", "", ""
				}
				if !utils.IsNumeric(wrap.UrlArgs[2]) {
					return "", "", ""
				}
				err := wrap.DB.QueryRow(
					wrap.R.Context(),
					`SELECT
						id,
						name,
						filter
					FROM
						shop_filters
					WHERE
						id = ?
					LIMIT 1;`,
					utils.StrToInt(wrap.UrlArgs[2]),
				).Scan(
					&data.A_id,
					&data.A_name,
					&data.A_filter,
				)
				if *wrap.LogCpError(&err) != nil {
					return "", "", ""
				}
			}

			btn_caption := "Add"
			if wrap.CurrSubModule == "attributes-modify" {
				btn_caption = "Save"
			}

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "shop-attributes-modify",
				},
				{
					Kind:  builder.DFKHidden,
					Name:  "id",
					Value: utils.IntToStr(data.A_id),
				},
				{
					Kind:     builder.DFKText,
					Caption:  "Attribute name",
					Name:     "name",
					Value:    data.A_name,
					Required: true,
					Min:      "1",
					Max:      "255",
				},
				{
					Kind:     builder.DFKText,
					Caption:  "Attribute in filter",
					Name:     "filter",
					Value:    data.A_filter,
					Required: true,
					Min:      "1",
					Max:      "255",
				},
				{
					Kind:    builder.DFKText,
					Caption: "Attribute values",
					Name:    "",
					Value:   "",
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group n4">` +
							`<div class="row">` +
							`<div class="col-md-3">` +
							`<label class="col-md-3">Attribute values</label>` +
							`</div>` +
							`<div class="col-md-9">` +
							`<div class="list-wrapper">` +
							`<div id="list">` +
							this.shop_GetFilterValuesInputs(wrap, data.A_id) +
							`</div>` +
							`<div class="list-button"><button type="button" class="btn btn-success" onclick="fave.ShopAttributesAdd();">Add attribute value</button></div>` +
							`</div>` +
							`</div>` +
							`</div>` +
							`</div>`
					},
				},
				{
					Kind:   builder.DFKSubmit,
					Value:  btn_caption,
					Target: "add-edit-button",
				},
			})

			if wrap.CurrSubModule == "attributes-add" {
				sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Add</button>`
			} else {
				sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
			}
		} else if wrap.CurrSubModule == "currencies-add" || wrap.CurrSubModule == "currencies-modify" {
			if wrap.CurrSubModule == "currencies-add" {
				content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
					{Name: "Currencies", Link: "/cp/" + wrap.CurrModule + "/currencies/"},
					{Name: "Add new currency"},
				})
			} else {
				content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
					{Name: "Currencies", Link: "/cp/" + wrap.CurrModule + "/currencies/"},
					{Name: "Modify currency"},
				})
			}

			data := utils.Sql_shop_currency{
				A_id:          0,
				A_name:        "",
				A_coefficient: 0,
				A_code:        "",
				A_symbol:      "",
			}

			if wrap.CurrSubModule == "currencies-modify" {
				if len(wrap.UrlArgs) != 3 {
					return "", "", ""
				}
				if !utils.IsNumeric(wrap.UrlArgs[2]) {
					return "", "", ""
				}
				err := wrap.DB.QueryRow(
					wrap.R.Context(),
					`SELECT
						id,
						name,
						coefficient,
						code,
						symbol
					FROM
						shop_currencies
					WHERE
						id = ?
					LIMIT 1;`,
					utils.StrToInt(wrap.UrlArgs[2]),
				).Scan(
					&data.A_id,
					&data.A_name,
					&data.A_coefficient,
					&data.A_code,
					&data.A_symbol,
				)
				if *wrap.LogCpError(&err) != nil {
					return "", "", ""
				}
			}

			btn_caption := "Add"
			if wrap.CurrSubModule == "currencies-modify" {
				btn_caption = "Save"
			}

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "shop-currencies-modify",
				},
				{
					Kind:  builder.DFKHidden,
					Name:  "id",
					Value: utils.IntToStr(data.A_id),
				},
				{
					Kind:     builder.DFKText,
					Caption:  "Currency name",
					Name:     "name",
					Value:    data.A_name,
					Required: true,
					Min:      "1",
					Max:      "255",
				},
				{
					Kind:    builder.DFKText,
					Caption: "Currency coefficient",
					Name:    "coefficient",
					Value:   "0",
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group n3">` +
							`<div class="row">` +
							`<div class="col-md-3">` +
							`<label for="lbl_coefficient">Currency coefficient</label>` +
							`</div>` +
							`<div class="col-md-9">` +
							`<div><input class="form-control" type="number" step="0.0001" id="lbl_coefficient" name="coefficient" value="` + utils.Float64ToStrF(data.A_coefficient, "%.4f") + `" placeholder="" autocomplete="off" required></div>` +
							`</div>` +
							`</div>` +
							`</div>`
					},
				},
				{
					Kind:     builder.DFKText,
					Caption:  "Currency code",
					Name:     "code",
					Value:    data.A_code,
					Required: true,
					Min:      "1",
					Max:      "10",
				},
				{
					Kind:     builder.DFKText,
					Caption:  "Currency symbol",
					Name:     "symbol",
					Value:    data.A_symbol,
					Required: true,
					Min:      "1",
					Max:      "5",
				},
				{
					Kind:   builder.DFKSubmit,
					Value:  btn_caption,
					Target: "add-edit-button",
				},
			})

			if wrap.CurrSubModule == "currencies-add" {
				sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Add</button>`
			} else {
				sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
			}
		} else if wrap.CurrSubModule == "orders-modify" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Orders", Link: "/cp/" + wrap.CurrModule + "/orders/"},
				{Name: "Modify order"},
			})

			if len(wrap.UrlArgs) != 3 {
				return "", "", ""
			}
			if !utils.IsNumeric(wrap.UrlArgs[2]) {
				return "", "", ""
			}

			var curr_order_id int
			var curr_order_client_phone string
			var curr_order_update_datetime string
			var curr_order_currency_id int
			var curr_order_currency_name string
			var curr_order_currency_coefficient float64
			var curr_order_currency_code string
			var curr_order_currency_symbol string
			var curr_order_client_last_name string
			var curr_order_client_first_name string
			var curr_order_client_middle_name string
			var curr_order_create_datetime int
			var curr_order_client_email string
			var curr_order_client_delivery_comment string
			var curr_order_client_order_comment string
			var curr_order_status int
			var curr_order_total float64

			err := wrap.DB.QueryRow(
				wrap.R.Context(),
				`SELECT
					shop_orders.id,
					shop_orders.client_phone,
					shop_orders.update_datetime,
					shop_orders.currency_id,
					shop_orders.currency_name,
					shop_orders.currency_coefficient,
					shop_orders.currency_code,
					shop_orders.currency_symbol,
					shop_orders.client_last_name,
					shop_orders.client_first_name,
					shop_orders.client_middle_name,
					strftime('%s', shop_orders.create_datetime) as create_datetime, 
					shop_orders.client_email,
					shop_orders.client_delivery_comment,
					shop_orders.client_order_comment,
					shop_orders.status,
					shop_order_total.total
				FROM
					shop_orders
					LEFT JOIN (
						SELECT
							order_id,
							SUM(price * quantity) as total
						FROM
							shop_order_products
						GROUP BY
							order_id
					) as shop_order_total ON shop_order_total.order_id = shop_orders.id
				WHERE
					shop_orders.id = ?
				LIMIT 1;`,
				utils.StrToInt(wrap.UrlArgs[2]),
			).Scan(
				&curr_order_id,
				&curr_order_client_phone,
				&curr_order_update_datetime,
				&curr_order_currency_id,
				&curr_order_currency_name,
				&curr_order_currency_coefficient,
				&curr_order_currency_code,
				&curr_order_currency_symbol,
				&curr_order_client_last_name,
				&curr_order_client_first_name,
				&curr_order_client_middle_name,
				&curr_order_create_datetime,
				&curr_order_client_email,
				&curr_order_client_delivery_comment,
				&curr_order_client_order_comment,
				&curr_order_status,
				&curr_order_total,
			)
			if *wrap.LogCpError(&err) != nil {
				return "", "", ""
			}

			last_name := html.EscapeString(curr_order_client_last_name)
			first_name := html.EscapeString(curr_order_client_first_name)
			middle_name := html.EscapeString(curr_order_client_middle_name)

			phone := html.EscapeString(curr_order_client_phone)
			email := html.EscapeString(curr_order_client_email)

			order_id := `<span class="d-inline d-sm-none">#` + utils.IntToStr(curr_order_id) + `. </span>`

			name := ""
			if last_name != "" {
				name += " " + last_name
			}
			if first_name != "" {
				name += " " + first_name
			}
			if middle_name != "" {
				name += " " + middle_name
			}
			name = order_id + utils.Trim(name)

			contact := ""
			if email != "" {
				contact += " " + email
			}
			if phone != "" {
				contact += " (" + phone + ")"
			}
			contact = utils.Trim(contact)

			content += `<table id="cp-table-shop_orders" class="table data-table table-striped table-bordered table-hover table_shop_orders">
				<thead>
					<tr>
						<th scope="col" class="col_id d-none d-lg-table-cell">Order #</th>
						<th scope="col" class="col_client_last_name">Client / Contact</th>
						<th scope="col" class="col_create_datetime d-none d-lg-table-cell">Date / Time</th>
						<th scope="col" class="col_client_email">Status / Total</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td class="col_id d-none d-lg-table-cell">` + utils.IntToStr(curr_order_id) + `</td>
						<td class="col_client_last_name">
							<div>` + name + `</div>
							<div><small>` + contact + `</small></div>
							<div>
								<small>
									<a href="javascript:fave.ShopSetOrderStatus(this,'` + utils.IntToStr(curr_order_id) + `','0','Are you sure want to change order status?');">` + this.shop_GetOrderStatus(0) + `</a> | <a href="javascript:fave.ShopSetOrderStatus(this,'` + utils.IntToStr(curr_order_id) + `','1','Are you sure want to change order status?');">` + this.shop_GetOrderStatus(1) + `</a> | <a href="javascript:fave.ShopSetOrderStatus(this,'` + utils.IntToStr(curr_order_id) + `','2','Are you sure want to change order status?');">` + this.shop_GetOrderStatus(2) + `</a> | <a href="javascript:fave.ShopSetOrderStatus(this,'` + utils.IntToStr(curr_order_id) + `','3','Are you sure want to change order status?');">` + this.shop_GetOrderStatus(3) + `</a> | <a href="javascript:fave.ShopSetOrderStatus(this,'` + utils.IntToStr(curr_order_id) + `','4','Are you sure want to change order status?');">` + this.shop_GetOrderStatus(4) + `</a>
								</small>
							</div>
						</td>
						<td class="col_create_datetime d-none d-lg-table-cell">
							<div>` + utils.UnixTimestampToFormat(int64(curr_order_create_datetime), "02.01.2006") + `</div>
							<div><small>` + utils.UnixTimestampToFormat(int64(curr_order_create_datetime), "15:04:05") + `</small></div>
						</td>
						<td class="col_client_email">
							<div>` + this.shop_GetOrderStatus(curr_order_status) + `</div>
							<div><small>` + utils.Float64ToStr(curr_order_total) + " " + html.EscapeString(curr_order_currency_code) + `</small></div>
						</td>
					</tr>
				</tbody>
			</table>`

			content += `<table id="cp-table-shop_products" class="table data-table table-striped table-bordered table-hover mt-3 table_shop_products">
				<thead>
					<tr>
						<th scope="col" class="col_name">Product</th>
						<th scope="col" class="col_price d-none d-md-table-cell">Price</th>
						<th scope="col" class="col_quantity">Quantity</th>
						<th scope="col" class="col_total d-none d-md-table-cell">Total</th>
					</tr>
				</thead>
				<tbody>
			`
			rows, err := wrap.DB.Query(
				wrap.R.Context(),
				`SELECT
					shop_order_products.product_id,
					shop_products.name,
					shop_products.alias,
					shop_order_products.price,
					shop_order_products.quantity,
					(shop_order_products.price * shop_order_products.quantity) as total
				FROM
					shop_order_products
					LEFT JOIN shop_products ON shop_products.id = shop_order_products.product_id
				WHERE
					shop_order_products.order_id = ?
				;`,
				curr_order_id,
			)
			if err == nil {
				defer rows.Close()
				var curr_product_id int
				var curr_product_name string
				var curr_product_alias string
				var curr_product_price float64
				var curr_product_quantity int
				var curr_product_total float64

				for rows.Next() {
					err = rows.Scan(
						&curr_product_id,
						&curr_product_name,
						&curr_product_alias,
						&curr_product_price,
						&curr_product_quantity,
						&curr_product_total,
					)
					if *wrap.LogCpError(&err) == nil {
						content += `<tr>
							<td class="col_name">
								<div><a href="/cp/shop/modify/` + utils.IntToStr(curr_product_id) + `/">` + html.EscapeString(curr_product_name) + ` ` + utils.IntToStr(curr_product_id) + `</a></div>
								<div><small><a href="/shop/` + html.EscapeString(curr_product_alias) + `/" target="_blank">/shop/` + html.EscapeString(curr_product_alias) + `/</a></small></div>
							</td>
							<td class="col_price d-none d-md-table-cell">
								<div>` + utils.Float64ToStr(curr_product_price) + `</div>
								<div><small>` + html.EscapeString(curr_order_currency_code) + `</small></div>
							</td>
							<td class="col_quantity">
								` + utils.IntToStr(curr_product_quantity) + `
							</td>
							<td class="col_total d-none d-md-table-cell">
								<div>` + utils.Float64ToStr(curr_product_total) + `</div>
								<div><small>` + html.EscapeString(curr_order_currency_code) + `</small></div>
							</td>
						</tr>`
					}
				}
			}
			content += `</tbody>
			</table>`

			// Delivery
			content += `<div class="card mt-3">
				<div class="card-header"><b>Delivery</b></div>
				<ul class="list-group list-group-flush">
					<li class="list-group-item">`
			if utils.Trim(curr_order_client_delivery_comment) != "" {
				content += html.EscapeString(curr_order_client_delivery_comment)
			} else {
				content += `NO SET`
			}
			content += `</li>
				</ul>
			</div>`

			// Comment
			content += `<div class="card mt-3">
				<div class="card-header"><b>Comment</b></div>
				<ul class="list-group list-group-flush">
					<li class="list-group-item">`
			if utils.Trim(curr_order_client_order_comment) != "" {
				content += html.EscapeString(curr_order_client_order_comment)
			} else {
				content += `NO SET`
			}
			content += `</li>
				</ul>
			</div>`
		}
		return this.getSidebarModules(wrap), content, sidebar
	})
}

package modules

import (
	"html"
	"io/ioutil"
	"os"

	"server/engine/assets"
	"server/engine/builder"
	"server/engine/consts"
	"server/engine/utils"
	"server/engine/wrapper"
)

func (this *Modules) RegisterModule_Settings() *Module {
	return this.newModule(MInfo{
		Mount:  "settings",
		Name:   "Установки",
		Order:  801,
		System: true,
		Icon:   "<i class=\"material-icons notranslate\">settings</i>", //assets.SysSvgIconGear,
		Sub: &[]MISub{
			{Mount: "default", Name: "Основные", Show: true, Icon: "<i class=\"material-icons notranslate\">settings</i>" /*assets.SysSvgIconGear*/},
			{Mount: "robots-txt", Name: "Robots.txt", Show: true, Icon: "<i class=\"material-icons notranslate\">bug_report</i>" /*assets.SysSvgIconGear*/},
			{Mount: "pagination", Name: "Страницы", Show: true, Icon: assets.SysSvgIconGear},
			{Mount: "thumbnails", Name: "Thumbnails", Show: true, Icon: assets.SysSvgIconThumbnails},
			{Mount: "domains", Name: "Domains", Show: true, Icon: assets.SysSvgIconApi},
			{Mount: "smtp", Name: "SMTP", Show: true, Icon: assets.SysSvgIconEmail},
			{Mount: "shop", Name: "Shop", Show: true, Icon: assets.SysSvgIconShop},
			{Mount: "api", Name: "API", Show: true, Icon: assets.SysSvgIconApi},
		},
	}, nil, func(wrap *wrapper.Wrapper) (string, string, string) {
		content := ""
		sidebar := ""

		if wrap.CurrSubModule == "" || wrap.CurrSubModule == "default" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "General"},
			})

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "settings-general",
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						modules_list := ``
						modules_list += `<select class="selectpicker" data-size="7" data-style="select-with-transition" id="lbl_module-at-home" name="module-at-home">`
						modules_list += `<option value="0"`
						if (*wrap.Config).Engine.MainModule == 0 {
							modules_list += ` selected`
						}
						modules_list += `>Страницы</option>`
						modules_list += `<option value="1"`
						if (*wrap.Config).Engine.MainModule == 1 {
							modules_list += ` selected`
						}
						modules_list += `>Панели</option>`
						modules_list += `</select>`
						//modules_list += `<button type="button" class="btn dropdown-toggle bs-placeholder select-with-transition" data-toggle="dropdown" role="button" title="Choose City"><div class="filter-option"><div class="filter-option-inner"><div class="filter-option-inner-inner">Choose City</div></div> </div></button>`
						//modules_list += `<div class="dropdown-menu " role="combobox"><div class="inner show" role="listbox" aria-expanded="false" tabindex="-1"><ul class="dropdown-menu inner show"></ul></div></div>`

						return `<div class="form-group n3">` +
							`<div class="row">` +
							`<label for="lbl_module-at-home" class="col-sm-3 col-form-label">Модуль домашней страницы</label>` +
							`<div class="col-sm-9">` +
							modules_list +
							`</div>` +
							`</div>` +
							`</div>`
					},
				},
				{
					Kind:    builder.DFKCheckBox,
					Caption: "Maintenance",
					Name:    "maintenance",
					Value:   utils.IntToStr((*wrap.Config).Engine.Maintenance),
					Hint:    "Close web site for maintenance",
				},
				{
					Kind:    builder.DFKCheckBox,
					Caption: "Blog is enabled",
					Name:    "mod-enabled-blog",
					Value:   utils.IntToStr((*wrap.Config).Modules.Enabled.Blog),
					Hint:    "Module can be enabled or fully disabled",
				},
				// {
				// 	Kind:    builder.DFKCheckBox,
				// 	Caption: "Shop is enabled",
				// 	Name:    "mod-enabled-shop",
				// 	Value:   utils.IntToStr((*wrap.Config).Modules.Enabled.Shop),
				// 	Hint:    "Module can be enabled or fully disabled",
				// },
				{
					Kind:    builder.DFKCheckBox,
					Caption: "Разрешить пустые пароли",
					Name:    "mod-enabled-security", // TODO
					Value:   utils.IntToStr((*wrap.Config).Modules.Enabled.EmptyPassword),
					Hint:    "Запретить или разрешить пустые пароли",
				},
				{
					Kind:   builder.DFKSubmit,
					Value:  "Save",
					Target: "add-edit-button",
				},
			})

			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
		} else if wrap.CurrSubModule == "robots-txt" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Robots.txt"},
			})

			fcont := []byte(``)
			fcont, _ = ioutil.ReadFile(wrap.DTemplate + string(os.PathSeparator) + "robots.txt")

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "settings-robots-txt",
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group last"><div class="row"><div class="col-12"><textarea class="form-control autosize" id="lbl_content" name="content" placeholder="" autocomplete="off">` + html.EscapeString(string(fcont)) + `</textarea></div></div></div>`
					},
				},
				{
					Kind: builder.DFKSubmit,
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="row d-lg-none"><div class="col-12"><div class="pt-3"><button type="submit" class="btn btn-primary" data-target="add-edit-button">Save</button></div></div></div>`
					},
				},
			})

			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
		} else if wrap.CurrSubModule == "pagination" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Pagination"},
			})

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "settings-pagination",
				},
				{
					Kind:     builder.DFKNumber,
					Caption:  "Blog main page",
					Name:     "blog-index",
					Min:      "1",
					Max:      "100",
					Required: true,
					Value:    utils.IntToStr((*wrap.Config).Blog.Pagination.Index),
				},
				{
					Kind:     builder.DFKNumber,
					Caption:  "Blog category page",
					Name:     "blog-category",
					Min:      "1",
					Max:      "100",
					Required: true,
					Value:    utils.IntToStr((*wrap.Config).Blog.Pagination.Category),
				},
				{
					Kind:    builder.DFKText,
					Caption: "",
					Name:    "",
					Value:   "",
					CallBack: func(field *builder.DataFormField) string {
						return `<hr>`
					},
				},
				// {
				// 	Kind:     builder.DFKNumber,
				// 	Caption:  "Shop main page",
				// 	Name:     "shop-index",
				// 	Min:      "1",
				// 	Max:      "100",
				// 	Required: true,
				// 	Value:    utils.IntToStr((*wrap.Config).Shop.Pagination.Index),
				// },
				// {
				// 	Kind:     builder.DFKNumber,
				// 	Caption:  "Shop category page",
				// 	Name:     "shop-category",
				// 	Min:      "1",
				// 	Max:      "100",
				// 	Required: true,
				// 	Value:    utils.IntToStr((*wrap.Config).Shop.Pagination.Category),
				// },
				{
					Kind:   builder.DFKSubmit,
					Value:  "Save",
					Target: "add-edit-button",
				},
			})

			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
		} else if wrap.CurrSubModule == "thumbnails" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Thumbnails"},
			})

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "settings-thumbnails",
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						// resize_list := ``
						// resize_list += `<select class="form-control" name="shop-thumbnail-r-1">`
						// resize_list += `<option value="0"`
						// if (*wrap.Config).Shop.Thumbnails.Thumbnail1[2] == 0 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Crop</option>`
						// resize_list += `<option value="1"`
						// if (*wrap.Config).Shop.Thumbnails.Thumbnail1[2] == 1 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Resize</option>`
						// resize_list += `<option value="2"`
						// if (*wrap.Config).Shop.Thumbnails.Thumbnail1[2] == 2 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Fit into size</option>`
						// resize_list += `</select>`
						// return `<div class="form-group n3">` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<label>Shop thumbnail 1</label>` +
						// 	`</div>` +
						// 	`<div class="col-md-9">` +
						// 	`<div>` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<div><input class="form-control" type="number" name="shop-thumbnail-w-1" value="` + utils.IntToStr((*wrap.Config).Shop.Thumbnails.Thumbnail1[0]) + `" min="100" max="1000" placeholder="" autocomplete="off" required></div>` +
						// 	`<div class="d-md-none mb-3"></div>` +
						// 	`</div>` +
						// 	`<div class="col-md-3">` +
						// 	`<div><input class="form-control" type="number" name="shop-thumbnail-h-1" value="` + utils.IntToStr((*wrap.Config).Shop.Thumbnails.Thumbnail1[1]) + `" min="100" max="1000" placeholder="" autocomplete="off" required></div>` +
						// 	`<div class="d-md-none mb-3"></div>` +
						// 	`</div>` +
						// 	`<div class="col-md-6">` +
						// 	resize_list +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>`
						return ""
					},
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						// resize_list := ``
						// resize_list += `<select class="form-control" name="shop-thumbnail-r-2">`
						// resize_list += `<option value="0"`
						// if (*wrap.Config).Shop.Thumbnails.Thumbnail2[2] == 0 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Crop</option>`
						// resize_list += `<option value="1"`
						// if (*wrap.Config).Shop.Thumbnails.Thumbnail2[2] == 1 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Resize</option>`
						// resize_list += `<option value="2"`
						// if (*wrap.Config).Shop.Thumbnails.Thumbnail2[2] == 2 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Fit into size</option>`
						// resize_list += `</select>`
						// return `<div class="form-group n3">` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<label>Shop thumbnail 2</label>` +
						// 	`</div>` +
						// 	`<div class="col-md-9">` +
						// 	`<div>` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<div><input class="form-control" type="number" name="shop-thumbnail-w-2" value="` + utils.IntToStr((*wrap.Config).Shop.Thumbnails.Thumbnail2[0]) + `" min="100" max="1000" placeholder="" autocomplete="off" required></div>` +
						// 	`<div class="d-md-none mb-3"></div>` +
						// 	`</div>` +
						// 	`<div class="col-md-3">` +
						// 	`<div><input class="form-control" type="number" name="shop-thumbnail-h-2" value="` + utils.IntToStr((*wrap.Config).Shop.Thumbnails.Thumbnail2[1]) + `" min="100" max="1000" placeholder="" autocomplete="off" required></div>` +
						// 	`<div class="d-md-none mb-3"></div>` +
						// 	`</div>` +
						// 	`<div class="col-md-6">` +
						// 	resize_list +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>`
						return ""
					},
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						// resize_list := ``
						// resize_list += `<select class="form-control" name="shop-thumbnail-r-3">`
						// resize_list += `<option value="0"`
						// if (*wrap.Config).Shop.Thumbnails.Thumbnail3[2] == 0 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Crop</option>`
						// resize_list += `<option value="1"`
						// if (*wrap.Config).Shop.Thumbnails.Thumbnail3[2] == 1 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Resize</option>`
						// resize_list += `<option value="2"`
						// if (*wrap.Config).Shop.Thumbnails.Thumbnail3[2] == 2 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Fit into size</option>`
						// resize_list += `</select>`
						// return `<div class="form-group n3">` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<label>Shop thumbnail 3</label>` +
						// 	`</div>` +
						// 	`<div class="col-md-9">` +
						// 	`<div>` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<div><input class="form-control" type="number" name="shop-thumbnail-w-3" value="` + utils.IntToStr((*wrap.Config).Shop.Thumbnails.Thumbnail3[0]) + `" min="100" max="1000" placeholder="" autocomplete="off" required></div>` +
						// 	`<div class="d-md-none mb-3"></div>` +
						// 	`</div>` +
						// 	`<div class="col-md-3">` +
						// 	`<div><input class="form-control" type="number" name="shop-thumbnail-h-3" value="` + utils.IntToStr((*wrap.Config).Shop.Thumbnails.Thumbnail3[1]) + `" min="100" max="1000" placeholder="" autocomplete="off" required></div>` +
						// 	`<div class="d-md-none mb-3"></div>` +
						// 	`</div>` +
						// 	`<div class="col-md-6">` +
						// 	resize_list +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>`
						return ""
					},
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						// resize_list := ``
						// resize_list += `<select class="form-control" name="shop-thumbnail-r-full">`
						// resize_list += `<option value="0"`
						// if (*wrap.Config).Shop.Thumbnails.ThumbnailFull[2] == 0 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Crop</option>`
						// resize_list += `<option value="1"`
						// if (*wrap.Config).Shop.Thumbnails.ThumbnailFull[2] == 1 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Resize</option>`
						// resize_list += `<option value="2"`
						// if (*wrap.Config).Shop.Thumbnails.ThumbnailFull[2] == 2 {
						// 	resize_list += ` selected`
						// }
						// resize_list += `>Fit into size</option>`
						// resize_list += `</select>`
						// return `<div class="form-group n3">` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<label>Shop thumbnail full</label>` +
						// 	`</div>` +
						// 	`<div class="col-md-9">` +
						// 	`<div>` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<div><input class="form-control" type="number" name="shop-thumbnail-w-full" value="` + utils.IntToStr((*wrap.Config).Shop.Thumbnails.ThumbnailFull[0]) + `" min="100" max="1000" placeholder="" autocomplete="off" required></div>` +
						// 	`<div class="d-md-none mb-3"></div>` +
						// 	`</div>` +
						// 	`<div class="col-md-3">` +
						// 	`<div><input class="form-control" type="number" name="shop-thumbnail-h-full" value="` + utils.IntToStr((*wrap.Config).Shop.Thumbnails.ThumbnailFull[1]) + `" min="100" max="1000" placeholder="" autocomplete="off" required></div>` +
						// 	`<div class="d-md-none mb-3"></div>` +
						// 	`</div>` +
						// 	`<div class="col-md-6">` +
						// 	resize_list +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>`
						return ""
					},
				},
				{
					Kind:   builder.DFKSubmit,
					Value:  "Save",
					Target: "add-edit-button",
				},
			})

			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
		} else if wrap.CurrSubModule == "domains" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Domains"},
			})

			fcont := []byte(``)
			fcont, _ = ioutil.ReadFile(wrap.DConfig + string(os.PathSeparator) + ".domains")

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "settings-domains",
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="form-group last"><div class="row"><div class="col-12"><textarea class="form-control autosize" id="lbl_content" name="content" placeholder="" autocomplete="off">` + html.EscapeString(string(fcont)) + `</textarea></div></div></div>`
					},
				},
				{
					Kind: builder.DFKSubmit,
					CallBack: func(field *builder.DataFormField) string {
						return `<div class="row d-lg-none"><div class="col-12"><div class="pt-3"><button type="submit" class="btn btn-primary" data-target="add-edit-button">Save</button></div></div></div>`
					},
				},
			})

			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
		} else if wrap.CurrSubModule == "smtp" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "SMTP"},
			})

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "settings-smtp",
				},
				{
					Kind:    builder.DFKText,
					Caption: "SMTP server host",
					Name:    "smtp-host",
					Value:   (*wrap.Config).Smtp.Host,
					Hint:    "Example: smtp.gmail.com",
				},
				{
					Kind:     builder.DFKNumber,
					Caption:  "SMTP server port",
					Name:     "smtp-port",
					Min:      "0",
					Max:      "9999",
					Required: true,
					Value:    utils.IntToStr((*wrap.Config).Smtp.Port),
					Hint:     "Example: 587",
				},
				{
					Kind:    builder.DFKText,
					Caption: "SMTP user login",
					Name:    "smtp-login",
					Value:   (*wrap.Config).Smtp.Login,
					Hint:    "Example: example@gmail.com",
				},
				{
					Kind:    builder.DFKPassword,
					Caption: "SMTP user password",
					Name:    "smtp-password",
					Value:   "",
					Hint:    "Leave this field empty if you don't want change password",
				},
				{
					Kind:    builder.DFKText,
					Caption: "Email address for testing",
					Name:    "smtp-test-email",
					Value:   "",
					Hint:    "To this email address will be send test email message if settings are correct",
				},
				{
					Kind:   builder.DFKSubmit,
					Value:  "Save",
					Target: "add-edit-button",
				},
			})

			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
		} else if wrap.CurrSubModule == "shop" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "Shop"},
			})

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "settings-shop",
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						// price_format_list := ``
						// price_format_list += `<select class="form-control" id="lbl_price-fomat" name="price-fomat">`
						// price_format_list += `<option value="0"`
						// if (*wrap.Config).Shop.Price.Format == 0 {
						// 	price_format_list += ` selected`
						// }
						// price_format_list += `>100</option>`
						// price_format_list += `<option value="1"`
						// if (*wrap.Config).Shop.Price.Format == 1 {
						// 	price_format_list += ` selected`
						// }
						// price_format_list += `>100.0</option>`
						// price_format_list += `<option value="2"`
						// if (*wrap.Config).Shop.Price.Format == 2 {
						// 	price_format_list += ` selected`
						// }
						// price_format_list += `>100.00</option>`
						// price_format_list += `<option value="3"`
						// if (*wrap.Config).Shop.Price.Format == 3 {
						// 	price_format_list += ` selected`
						// }
						// price_format_list += `>100.000</option>`
						// price_format_list += `<option value="4"`
						// if (*wrap.Config).Shop.Price.Format == 4 {
						// 	price_format_list += ` selected`
						// }
						// price_format_list += `>100.0000</option>`
						// price_format_list += `</select>`

						// price_round_list := ``
						// price_round_list += `<select class="form-control" id="lbl_price-round" name="price-round">`
						// price_round_list += `<option value="0"`
						// if (*wrap.Config).Shop.Price.Round == 0 {
						// 	price_round_list += ` selected`
						// }
						// price_round_list += `>Don't round</option>`
						// price_round_list += `<option value="1"`
						// if (*wrap.Config).Shop.Price.Round == 1 {
						// 	price_round_list += ` selected`
						// }
						// price_round_list += `>Round to ceil</option>`
						// price_round_list += `<option value="2"`
						// if (*wrap.Config).Shop.Price.Round == 2 {
						// 	price_round_list += ` selected`
						// }
						// price_round_list += `>Round to floor</option>`
						// price_round_list += `</select>`

						// return `<div class="form-group n2">` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<label for="lbl_price-fomat">Price format</label>` +
						// 	`</div>` +
						// 	`<div class="col-md-9">` +
						// 	`<div>` +
						// 	price_format_list +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`<div class="form-group n3">` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<label for="lbl_price-round">Price round</label>` +
						// 	`</div>` +
						// 	`<div class="col-md-9">` +
						// 	`<div>` +
						// 	price_round_list +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>`
						return ""
					},
				},
				{
					Kind: builder.DFKText,
					CallBack: func(field *builder.DataFormField) string {
						// make_checkbox := func(name, caption string, value int) string {
						// 	checked := ""
						// 	if value > 0 {
						// 		checked = " checked"
						// 	}
						// 	return `<div class="checkbox-clickable"><input class="form-control" type="checkbox" id="lbl_` +
						// 		name + `" name="` + name + `" value="1" "="" autocomplete="off"` + checked +
						// 		`><label for="lbl_` + name + `">` + caption + `</label></div>`
						// }

						// checkboxes := ""
						// checkboxes += make_checkbox("require-last-name", "Last Name", (*wrap.Config).Shop.Orders.RequiredFields.LastName)
						// checkboxes += make_checkbox("require-first-name", "First Name", (*wrap.Config).Shop.Orders.RequiredFields.FirstName)
						// checkboxes += make_checkbox("require-middle-name", "Middle Name", (*wrap.Config).Shop.Orders.RequiredFields.MiddleName)
						// checkboxes += make_checkbox("require-mobile-phone", "Mobile Phone", (*wrap.Config).Shop.Orders.RequiredFields.MobilePhone)
						// checkboxes += make_checkbox("require-email-address", "Email Address", (*wrap.Config).Shop.Orders.RequiredFields.EmailAddress)
						// checkboxes += make_checkbox("require-delivery", "Delivery", (*wrap.Config).Shop.Orders.RequiredFields.Delivery)
						// checkboxes += make_checkbox("require-comment", "Comment", (*wrap.Config).Shop.Orders.RequiredFields.Comment)

						// return `<div class="form-group n4">` +
						// 	`<div class="row">` +
						// 	`<div class="col-md-3">` +
						// 	`<label for="lbl_price-fomat">Order require fields</label>` +
						// 	`</div>` +
						// 	`<div class="col-md-9">` +
						// 	`<div>` +
						// 	checkboxes +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>` +
						// 	`</div>`
						return ""
					},
				},
				// {
				// 	Kind:    builder.DFKCheckBox,
				// 	Caption: "Accept orders",
				// 	Name:    "accept-orders",
				// 	Value:   utils.IntToStr((*wrap.Config).Shop.Orders.Enabled),
				// },
				// {
				// 	Kind:    builder.DFKText,
				// 	Caption: "New order notify email",
				// 	Name:    "new-order-notify-email",
				// 	Value:   (*wrap.Config).Shop.Orders.NotifyEmail,
				// 	Hint:    "Example: example@gmail.com",
				// },
				// {
				// 	Kind:    builder.DFKText,
				// 	Caption: "New order email theme (CP)",
				// 	Name:    "new-order-email-theme-cp",
				// 	Value:   (*wrap.Config).Shop.Orders.NewOrderEmailThemeCp,
				// },
				// {
				// 	Kind:    builder.DFKText,
				// 	Caption: "New order email theme (User)",
				// 	Name:    "new-order-email-theme-user",
				// 	Value:   (*wrap.Config).Shop.Orders.NewOrderEmailThemeUser,
				// },
				// {
				// 	Kind:    builder.DFKCheckBox,
				// 	Caption: "Custom field 1 enabled",
				// 	Name:    "custom-field-1-enabled",
				// 	Value:   utils.IntToStr((*wrap.Config).Shop.CustomFields.Field1.Enabled),
				// },
				// {
				// 	Kind:    builder.DFKText,
				// 	Caption: "Custom field 1 caption",
				// 	Name:    "custom-field-1-caption",
				// 	Value:   (*wrap.Config).Shop.CustomFields.Field1.Caption,
				// 	Hint:    "Caption for product custom field",
				// },
				// {
				// 	Kind:    builder.DFKCheckBox,
				// 	Caption: "Custom field 2 enabled",
				// 	Name:    "custom-field-2-enabled",
				// 	Value:   utils.IntToStr((*wrap.Config).Shop.CustomFields.Field2.Enabled),
				// },
				// {
				// 	Kind:    builder.DFKText,
				// 	Caption: "Custom field 2 caption",
				// 	Name:    "custom-field-2-caption",
				// 	Value:   (*wrap.Config).Shop.CustomFields.Field2.Caption,
				// 	Hint:    "Caption for product custom field",
				// },
				{
					Kind:   builder.DFKSubmit,
					Value:  "Save",
					Target: "add-edit-button",
				},
			})

			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
		} else if wrap.CurrSubModule == "api" {
			content += this.getBreadCrumbs(wrap, &[]consts.BreadCrumb{
				{Name: "API"},
			})

			content += builder.DataForm(wrap, []builder.DataFormField{
				{
					Kind:  builder.DFKHidden,
					Name:  "action",
					Value: "settings-api",
				},
				{
					Kind:    builder.DFKCheckBox,
					Caption: "XML enabled",
					Name:    "xml-enabled",
					Value:   utils.IntToStr((*wrap.Config).Api.Xml.Enabled),
					Hint:    "XML: <a href=\"/api/products/\" target=\"_blank\">/api/products/</a>",
				},
				{
					Kind:    builder.DFKText,
					Caption: "XML name",
					Name:    "xml-name",
					Value:   (*wrap.Config).Api.Xml.Name,
				},
				{
					Kind:    builder.DFKText,
					Caption: "XML company",
					Name:    "xml-company",
					Value:   (*wrap.Config).Api.Xml.Company,
				},
				{
					Kind:    builder.DFKText,
					Caption: "XML url",
					Name:    "xml-url",
					Value:   (*wrap.Config).Api.Xml.Url,
				},
				{
					Kind:   builder.DFKSubmit,
					Value:  "Save",
					Target: "add-edit-button",
				},
			})

			sidebar += `<button class="btn btn-primary btn-sidebar" id="add-edit-button">Save</button>`
		}
		return this.getSidebarModules(wrap), content, sidebar
	})
}

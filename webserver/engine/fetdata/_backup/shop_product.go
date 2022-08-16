package fetdata

import (
	"html/template"
	"strings"
	"time"

	"server/engine/utils"
	"server/engine/wrapper"
)

type ShopProductVarItem struct {
	Link     string
	Name     string
	Selected bool
}

type ShopProduct struct {
	wrap   *wrapper.Wrapper
	object *utils.Sql_shop_product

	user     *User
	currency *ShopCurrency
	category *ShopCategory

	images []*ShopProductImage
	specs  []*ShopProductSpec
	vars   []*ShopProductVarItem
}

func (this *ShopProduct) load() *ShopProduct {
	if this == nil {
		return this
	}
	if rows, err := this.wrap.DB.Query(
		this.wrap.R.Context(),
		`SELECT
			shop_product_images.product_id,
			shop_product_images.filename
		FROM
			shop_product_images
		WHERE
			shop_product_images.product_id = ?
		ORDER BY
			shop_product_images.ord ASC
		;`,
		this.object.A_id,
	); err == nil {
		defer rows.Close()
		for rows.Next() {
			img := utils.Sql_shop_product_image{}
			if err := rows.Scan(
				&img.A_product_id,
				&img.A_filename,
			); *this.wrap.LogCpError(&err) == nil {
				this.images = append(this.images, &ShopProductImage{wrap: this.wrap, object: &img})
			}
		}
	}

	// Get images from parent
	if len(this.images) <= 0 && this.object.A_parent_id() > 0 {
		if rows, err := this.wrap.DB.Query(
			this.wrap.R.Context(),
			`SELECT
				shop_product_images.product_id,
				shop_product_images.filename
			FROM
				shop_product_images
			WHERE
				shop_product_images.product_id = ?
			ORDER BY
				shop_product_images.ord ASC
			;`,
			this.object.A_parent_id(),
		); err == nil {
			defer rows.Close()
			for rows.Next() {
				img := utils.Sql_shop_product_image{}
				if err := rows.Scan(
					&img.A_product_id,
					&img.A_filename,
				); *this.wrap.LogCpError(&err) == nil {
					this.images = append(this.images, &ShopProductImage{wrap: this.wrap, object: &img})
				}
			}
		}
	}

	filter_ids := []int{}
	filter_names := map[int]string{}
	filter_values := map[int][]string{}
	if rows, err := this.wrap.DB.Query(
		this.wrap.R.Context(),
		`SELECT
			shop_filters.id,
			shop_filters.filter,
			shop_filters_values.name
		FROM
			shop_filter_product_values
			LEFT JOIN shop_filters_values ON shop_filters_values.id = shop_filter_product_values.filter_value_id
			LEFT JOIN shop_filters ON shop_filters.id = shop_filters_values.filter_id
		WHERE
			shop_filter_product_values.product_id = ?
		ORDER BY
			shop_filters.filter ASC,
			shop_filters_values.name ASC
		;`,
		this.object.A_id,
	); err == nil {
		defer rows.Close()
		values := make([]string, 3)
		scan := make([]interface{}, len(values))
		for i := range values {
			scan[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scan...)
			if *this.wrap.LogCpError(&err) == nil {
				if !utils.InArrayInt(filter_ids, utils.StrToInt(string(values[0]))) {
					filter_ids = append(filter_ids, utils.StrToInt(string(values[0])))
				}
				filter_names[utils.StrToInt(string(values[0]))] = string(values[1])
				filter_values[utils.StrToInt(string(values[0]))] = append(filter_values[utils.StrToInt(string(values[0]))], string(values[2]))
			}
		}
	}
	for _, filter_id := range filter_ids {
		this.specs = append(this.specs, &ShopProductSpec{wrap: this.wrap, object: &utils.Sql_shop_product_spec{
			A_product_id:   this.object.A_id,
			A_filter_id:    filter_id,
			A_filter_name:  filter_names[filter_id],
			A_filter_value: strings.Join(filter_values[filter_id], ", "),
		}})
	}

	// Variations
	if rows, err := this.wrap.DB.Query(
		this.wrap.R.Context(),
		`SELECT
			shop_products.id,
			shop_products.name,
			shop_products.alias
		FROM
			shop_products
		WHERE
			shop_products.active = 1 AND
			(
				(shop_products.id = ? OR shop_products.parent_id = ?) OR
				(
					(shop_products.id = ?) OR
					(shop_products.parent_id IS NOT NULL AND shop_products.parent_id = ?)
				)
			)
		ORDER BY
			shop_products.name ASC
		;`,
		this.object.A_id,
		this.object.A_id,
		this.object.A_parent_id(),
		this.object.A_parent_id(),
	); err == nil {
		defer rows.Close()
		for rows.Next() {
			var tmp_id int
			var tmp_name string
			var tmp_alias string
			if err := rows.Scan(
				&tmp_id,
				&tmp_name,
				&tmp_alias,
			); *this.wrap.LogCpError(&err) == nil {
				selected := false
				if tmp_id == this.object.A_id {
					selected = true
				}
				this.vars = append(this.vars, &ShopProductVarItem{
					Link:     "/shop/" + tmp_alias + "/",
					Name:     tmp_name + " " + utils.IntToStr(tmp_id),
					Selected: selected,
				})
			}
		}
	}

	return this
}

func (this *ShopProduct) Id() int {
	if this == nil {
		return 0
	}
	return this.object.A_id
}

func (this *ShopProduct) User() *User {
	if this == nil {
		return nil
	}
	if this.user != nil {
		return this.user
	}
	this.user = (&User{wrap: this.wrap}).load()
	this.user.loadById(this.object.A_user)
	return this.user
}

func (this *ShopProduct) Currency() *ShopCurrency {
	if this == nil {
		return nil
	}
	if this.currency != nil {
		return this.currency
	}
	this.currency = (&ShopCurrency{wrap: this.wrap}).load()
	this.currency.loadById(this.object.A_currency)
	return this.currency
}

func (this *ShopProduct) Price() float64 {
	if this == nil {
		return 0
	}
	if this.Currency() == nil {
		return this.object.A_price
	}
	if this.wrap.ShopGetCurrentCurrency() == nil {
		return this.object.A_price
	}
	if this.wrap.ShopGetCurrentCurrency().A_id == this.Currency().Id() {
		return this.object.A_price
	}
	if this.Currency().Id() == 1 {
		return this.object.A_price * this.wrap.ShopGetCurrentCurrency().A_coefficient
	} else {
		if c, ok := (*this.wrap.ShopGetAllCurrencies())[this.Currency().Id()]; ok == true {
			return this.object.A_price / c.A_coefficient
		} else {
			return this.object.A_price
		}
	}
}

func (this *ShopProduct) PriceOld() float64 {
	if this == nil {
		return 0
	}
	if this.Currency() == nil {
		return this.object.A_price_old
	}
	if this.wrap.ShopGetCurrentCurrency() == nil {
		return this.object.A_price_old
	}
	if this.wrap.ShopGetCurrentCurrency().A_id == this.Currency().Id() {
		return this.object.A_price_old
	}
	if this.Currency().Id() == 1 {
		return this.object.A_price_old * this.wrap.ShopGetCurrentCurrency().A_coefficient
	} else {
		if c, ok := (*this.wrap.ShopGetAllCurrencies())[this.Currency().Id()]; ok == true {
			return this.object.A_price_old / c.A_coefficient
		} else {
			return this.object.A_price_old
		}
	}
}

func (this *ShopProduct) PriceNice() string {
	return utils.FormatProductPrice(
		this.Price(),
		(*this.wrap.Config).Shop.Price.Format,
		(*this.wrap.Config).Shop.Price.Round,
	)
}

func (this *ShopProduct) PriceOldNice() string {
	return utils.FormatProductPrice(
		this.PriceOld(),
		(*this.wrap.Config).Shop.Price.Format,
		(*this.wrap.Config).Shop.Price.Round,
	)
}

func (this *ShopProduct) PriceFormat(format string) string {
	return utils.Float64ToStrF(this.Price(), format)
}

func (this *ShopProduct) Group() string {
	if this == nil {
		return ""
	}
	return this.object.A_gname
}

func (this *ShopProduct) Name() string {
	if this == nil {
		return ""
	}
	return this.object.A_name
}

func (this *ShopProduct) Alias() string {
	if this == nil {
		return ""
	}
	return this.object.A_alias
}

func (this *ShopProduct) Vendor() string {
	if this == nil {
		return ""
	}
	return this.object.A_vendor
}

func (this *ShopProduct) Quantity() int {
	if this == nil {
		return 0
	}
	return this.object.A_quantity
}

func (this *ShopProduct) Category() *ShopCategory {
	if this == nil {
		return nil
	}
	if this.category != nil {
		return this.category
	}
	this.category = (&ShopCategory{wrap: this.wrap}).load(nil)
	this.category.loadById(this.object.A_category)
	return this.category
}

func (this *ShopProduct) Briefly() template.HTML {
	if this == nil {
		return template.HTML("")
	}
	return template.HTML(this.object.A_briefly)
}

func (this *ShopProduct) Content() template.HTML {
	if this == nil {
		return template.HTML("")
	}
	return template.HTML(this.object.A_content)
}

func (this *ShopProduct) DateTimeUnix() int {
	if this == nil {
		return 0
	}
	return this.object.A_datetime
}

func (this *ShopProduct) DateTimeFormat(format string) string {
	if this == nil {
		return ""
	}
	return time.Unix(int64(this.object.A_datetime), 0).Format(format)
}

func (this *ShopProduct) Active() bool {
	if this == nil {
		return false
	}
	return this.object.A_active > 0
}

func (this *ShopProduct) Permalink() string {
	if this == nil {
		return ""
	}
	return "/shop/" + this.object.A_alias + "/"
}

func (this *ShopProduct) Image() *ShopProductImage {
	if this == nil {
		return nil
	}
	if len(this.images) <= 0 {
		return nil
	}
	return this.images[0]
}

func (this *ShopProduct) HaveImages() bool {
	if this == nil {
		return false
	}
	if len(this.images) <= 0 {
		return false
	}
	return true
}

func (this *ShopProduct) Images() []*ShopProductImage {
	if this == nil {
		return []*ShopProductImage{}
	}
	return this.images
}

func (this *ShopProduct) ImagesCount() int {
	if this == nil {
		return 0
	}
	return len(this.images)
}

func (this *ShopProduct) HaveSpecs() bool {
	if this == nil {
		return false
	}
	if len(this.specs) <= 0 {
		return false
	}
	return true
}

func (this *ShopProduct) Specs() []*ShopProductSpec {
	if this == nil {
		return []*ShopProductSpec{}
	}
	return this.specs
}

func (this *ShopProduct) SpecsCount() int {
	if this == nil {
		return 0
	}
	return len(this.specs)
}

func (this *ShopProduct) HaveVariations() bool {
	if this == nil {
		return false
	}
	if len(this.vars) <= 1 {
		return false
	}
	return true
}

func (this *ShopProduct) Variations() []*ShopProductVarItem {
	if this == nil {
		return []*ShopProductVarItem{}
	}
	return this.vars
}

func (this *ShopProduct) VariationsCount() int {
	if this == nil {
		return 0
	}
	return len(this.vars)
}

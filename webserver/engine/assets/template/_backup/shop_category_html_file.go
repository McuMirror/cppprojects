package template

var VarShopCategoryHtmlFile = []byte(`{{template "header.html" .}}
<div class="mb-4">
	{{if $.Data.Shop.HaveProducts}}
		<div class="grid-products">
			{{range $.Data.Shop.Products}}
				<div class="card card-product">
					<div class="card-img-link">
						<a href="{{.Permalink}}">
							{{if .HaveImages }}
								<img class="card-img-top" src="{{.Image.Thumbnail1}}" alt="{{$.Data.EscapeString .Name}}">
							{{else}}
								<img class="card-img-top" src="{{$.Data.ImagePlaceholderHref}}" alt="{{$.Data.EscapeString .Name}}">
							{{end}}
						</a>
					</div>
					<div class="card-body">
						<h5 class="card-title">
							<a href="{{.Permalink}}">
								{{if ne .Group ""}}
									{{.Group}}
								{{else}}
									{{.Name}}
								{{end}}
							</a>
						</h5>
						<div class="card-text">{{.Briefly}}</div>
					</div>
					<div class="card-footer">
						{{if le .Quantity 0}}<span class="badge badge-primary">Out of stock</span>{{end}}
						<a href="{{.Permalink}}" class="btn btn-success">View</a>
						<span class="price{{if gt .PriceOld 0.00}} price_red{{end}}">{{.PriceNice}} {{$.Data.Shop.CurrentCurrency.Code}}</span>
						{{if gt .PriceOld 0.00}}<span class="price price_old"><strike>{{.PriceOldNice}} {{$.Data.Shop.CurrentCurrency.Code}}</strike></span>{{else}}<span class="price price_old">&nbsp;</span>{{end}}
					</div>
				</div>
			{{end}}
		</div>
	{{else}}
		<div class="card">
			<div class="card-body">
				Sorry, no products matched your criteria
			</div>
		</div>
	{{end}}
</div>
{{if $.Data.Shop.HaveProducts}}
	{{if gt $.Data.Shop.ProductsMaxPage 1 }}
		<nav>
			<ul class="pagination mb-4">
				{{if $.Data.Shop.PaginationPrev}}
					<li class="page-item{{if $.Data.Shop.PaginationPrev.Current}} disabled{{end}}">
						<a class="page-link" href="{{$.Data.Shop.PaginationPrev.Link}}">Previous</a>
					</li>
				{{end}}
				{{range $.Data.Shop.Pagination}}
					<li class="page-item{{if .Current}} active{{end}}">
						<a class="page-link" href="{{.Link}}">{{.Num}}</a>
					</li>
				{{end}}
				{{if $.Data.Shop.PaginationNext}}
					<li class="page-item{{if $.Data.Shop.PaginationNext.Current}} disabled{{end}}">
						<a class="page-link" href="{{$.Data.Shop.PaginationNext.Link}}">Next</a>
					</li>
				{{end}}
			</ul>
		</nav>
	{{end}}
{{end}}
{{template "footer.html" .}}`)

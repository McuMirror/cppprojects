package template

var VarBlogCategoryHtmlFile = []byte(`{{template "header.html" .}}
<div class="card mb-4">
	<div class="post">
		<div class="card-body">
			<b>Category author:</b> {{$.Data.Blog.Category.User.FirstName}} {{$.Data.Blog.Category.User.LastName}}
		</div>
	</div>
</div>
<div class="card mb-4">
	{{if $.Data.Blog.HavePosts}}
		{{range $.Data.Blog.Posts}}
			<div class="post">
				<div class="card-body">
					<h2 class="card-title">
						<a href="{{.Permalink}}">
							{{.Name}}
						</a>
					</h2>
					<div class="post-content">
						{{.Briefly}}
					</div>
					<div class="post-date">
						<div><small>Published on {{.DateTimeFormat "02/01/2006, 15:04:05"}}</small></div>
						<div>Author: {{.User.FirstName}} {{.User.LastName}}</div>
					</div>
				</div>
			</div>
		{{end}}
	{{else}}
		<div class="card-body">
			Sorry, no posts matched your criteria
		</div>
	{{end}}
</div>
{{if $.Data.Blog.HavePosts}}
	{{if gt $.Data.Blog.PostsMaxPage 1 }}
		<nav>
			<ul class="pagination mb-4">
				{{if $.Data.Blog.PaginationPrev}}
					<li class="page-item{{if $.Data.Blog.PaginationPrev.Current}} disabled{{end}}">
						<a class="page-link" href="{{$.Data.Blog.PaginationPrev.Link}}">Previous</a>
					</li>
				{{end}}
				{{range $.Data.Blog.Pagination}}
					<li class="page-item{{if .Current}} active{{end}}">
						<a class="page-link" href="{{.Link}}">{{.Num}}</a>
					</li>
				{{end}}
				{{if $.Data.Blog.PaginationNext}}
					<li class="page-item{{if $.Data.Blog.PaginationNext.Current}} disabled{{end}}">
						<a class="page-link" href="{{$.Data.Blog.PaginationNext.Link}}">Next</a>
					</li>
				{{end}}
			</ul>
		</nav>
	{{end}}
{{end}}
{{template "footer.html" .}}`)

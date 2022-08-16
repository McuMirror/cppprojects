package template

var VarSidebarRightHtmlFile = []byte(`{{if $.Data.ModuleBlogEnabled}}
	<div class="card mb-4">
		<h5 class="card-header">Blog categories</h5>
		<div class="card-body">
			<ul class="m-0 p-0 pl-4">
				{{$.Data.CachedBlock5}}
			</ul>
		</div>
	</div>
{{end}}
<div class="card mb-4">
	<h5 class="card-header">Useful links</h5>
	<div class="card-body">
		<ul class="m-0 p-0 pl-4">
			<li><a href="https://github.com/vladimirok5959/server" target="_blank">Project on GitHub</a></li>
			<li><a href="https://github.com/vladimirok5959/server/wiki" target="_blank">Wiki on GitHub</a></li>
		</ul>
	</div>
</div>`)

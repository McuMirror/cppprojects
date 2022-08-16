package template

var VarFooterHtmlFile = []byte(`						</div>
						{{if or (eq $.Data.Module "index") (eq $.Data.Module "404") (eq $.Data.Module "blog") (eq $.Data.Module "blog-post") (eq $.Data.Module "blog-category")}}
							<div class="col-sm-5 col-md-4 col-lg-3">
								{{template "sidebar-right.html" .}}
							</div>
						{{end}}
					</div>
				</div>
			</div>
		</div>
		<footer class="bg-light py-4">
			<div class="container">
				<p class="m-0 text-center text-black">
					Your Website © {{if eq ($.Data.DateTimeFormat "2006") "2019"}}
						{{$.Data.DateTimeFormat "2006"}}
					{{else}}
						2019-{{$.Data.DateTimeFormat "2006"}}
					{{end}}
				</p>
			</div>
		</footer>
		<!-- Optional JavaScript -->
		<!-- jQuery first, then Popper.js, then Bootstrap JS -->
		<script src="{{$.System.PathJsJquery}}"></script>
		<script src="{{$.System.PathJsPopper}}"></script>
		<script src="{{$.System.PathJsBootstrap}}"></script>
		<script src="{{$.System.PathJsLightGallery}}"></script>
		<script src="{{$.System.PathJsCpScripts}}"></script>

		<!-- Template JavaScript file from template folder -->
		<script src="{{$.System.PathThemeScripts}}"></script>
	</body>
</html>`)

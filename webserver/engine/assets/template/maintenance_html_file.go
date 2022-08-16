package template

var VarMaintenanceHtmlFile = []byte(`<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8" />
		<meta name="theme-color" content="#205081" />
		<title>Maintenance</title>
		<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
		<meta name="viewport" content="width=device-width, initial-scale=0.8, maximum-scale=0.8" />
		<link rel="shortcut icon" href="{{$.System.PathIcoFav}}" type="image/x-icon" />
		<link rel="stylesheet" type="text/css" media="all" href="{{$.System.PathCssStyles}}" />
	</head>
	<body>
		<div class="wrapper">
			<div class="logo">
				<div class="svg">
					<img src="{{$.System.PathSvgLogo}}" width="150" height="150" />
				</div>
			</div>
			<h1>We are currently down for maintenance</h1>
			<h2>
				<script>document.write(document.location.host);</script>
				<noscript>fave.pro</noscript>
			</h2>
		</div>
	</body>
</html>`)

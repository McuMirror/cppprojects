package template

var VarEmailNewOrderUserHtmlFile = []byte(`<html>
	<head>
		<title>{{$.Else.Subject}}</title>
	</head>
	<body>
		<h2>Your contacts</h2>
		<table border="1">
			<tbody>
				{{if ne $.Client.LastName "" }}
					<tr>
						<td><b>Last&nbsp;name</b>&nbsp;&nbsp;&nbsp;</td>
						<td>{{$.Client.LastName}}</td>
					</tr>
				{{end}}
				{{if ne $.Client.FirstName "" }}
					<tr>
						<td><b>First&nbsp;name</b>&nbsp;&nbsp;&nbsp;</td>
						<td>{{$.Client.FirstName}}</td>
					</tr>
				{{end}}
				{{if ne $.Client.MiddleName "" }}
					<tr>
						<td><b>Middle&nbsp;name</b>&nbsp;&nbsp;&nbsp;</td>
						<td>{{$.Client.MiddleName}}</td>
					</tr>
				{{end}}
				{{if ne $.Client.Phone "" }}
					<tr>
						<td><b>Phone</b>&nbsp;&nbsp;&nbsp;</td>
						<td>{{$.Client.Phone}}</td>
					</tr>
				{{end}}
				{{if ne $.Client.Email "" }}
					<tr>
						<td><b>Email</b>&nbsp;&nbsp;&nbsp;</td>
						<td>{{$.Client.Email}}</td>
					</tr>
				{{end}}
			</tbody>
		</table>
		{{if ne $.Client.DeliveryComment "" }}
			<div>&nbsp;</div>
			<h2>Delivery</h2>
			<div>{{$.Client.DeliveryComment}}</div>
		{{end}}
		{{if ne $.Client.OrderComment "" }}
			<div>&nbsp;</div>
			<h2>Order comment</h2>
			<div>{{$.Client.OrderComment}}</div>
		{{end}}
		<div>&nbsp;</div>
		<h2>Order products</h2>
		<div>
			<table border="1" width="100%">
				<tbody>
					{{range $.Basket.Products}}
						<tr>
							<td>
								{{.RenderName}}
							</td>
							<td>
								{{.RenderPrice}}&nbsp;{{$.Basket.Currency.Code}}&nbsp;x&nbsp;{{.RenderQuantity}}
							</td>
							<td>
								{{.RenderSum}} {{$.Basket.Currency.Code}}
							</td>
						</tr>
					{{end}}
				</tbody>
			</table>
		</div>
		<h2>Total: {{$.Basket.RenderTotalSum}} {{$.Basket.Currency.Code}}</h2>
		<div>&nbsp;</div>
		<div><b>Thank you for choosing our shop!</b></div>
	</body>
</html>`)

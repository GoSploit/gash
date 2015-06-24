<html>
<head>

<link href="/static/css/bootstrap.css" rel="stylesheet" type="text/css" >

</head>
<body>
<div class="container">
<div class="row">
<div class="col-md-3">
</div>
<div class="col-md-9">
<table class="table table-bordered table-hover" id="projects">
	<tr>
		<th>Attack Type</th>
		<th>Number of Hashes</th>
	</tr>
	<tbody data-bind="foreach: projects">
		<tr>
			<td data-bind="text: attackType"></td>
			<td data-bind="text: hashes"></td>
		</tr>
	</tbody>
</table>
</div>
</div>
</div>
<script src="/static/js/knockout-3.3.0.js"></script>
<script src="/static/js/zepto.min.js"></script>
<script src="/static/js/app.js"></script>
</body>
</html>
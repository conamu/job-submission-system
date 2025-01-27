package view

var PageView = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Job Status</title>
    <script src="https://unpkg.com/htmx.org@1.6.1"></script>
    <style>
        .scrollable-table {
            max-height: 400px;
            overflow-y: auto;
            display: block;
        }
        table {
            width: 100%;
            border-collapse: collapse;
        }
        th, td {
            border: 1px solid #ddd;
            padding: 8px;
        }
        th {
            background-color: #f2f2f2;
        }
    </style>
</head>
	<body>
		{{ template "jobView" . }}
	</body>
</html>
`

var JobView = `
<h1>Job Status</h1>
    <div class="scrollable-table">
        <table>
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Status</th>
                </tr>
            </thead>
            <tbody hx-get="/status" hx-trigger="every 2s" hx-target="body">
                {{ range . }}
					<tr>
						<td>{{ .Id }}</td>
						<td>{{ .Status }}</td>
					</tr>
                {{ end }}
            </tbody>
        </table>
    </div>
`

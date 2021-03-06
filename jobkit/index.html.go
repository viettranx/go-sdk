package jobkit

var indexTemplate = `
{{ define "index" }}
{{ template "header" . }}
<div class="container">
	<table class="u-full-width">
		<thead>
			<tr>
				<th>Job Name</th>
				<th>Status</th>
				<th>Current</th>
				<th>Next Run</th>
				<th>Last Ran</th>
				<th>Last Result</th>
				<th>Last Elapsed</th>
			</tr>
		</thead>
		<tbody>
		{{ range $index, $job := .ViewModel.Jobs }}
			<tr>
				<td> <!-- job name -->
					{{ $job.Name }}
				</td>
				<td> <!-- job status -->
				{{ if $job.Disabled }}
					<span class="danger">Disabled</span>
				{{else}}
					<span class="primary">Enabled</span>
				{{end}}
				</td>
				<td> <!-- job status -->
				{{ if $job.Current }}
					{{ since_utc $job.Current.Started }}
				{{else}}
					<span>-</span>
				{{end}}
				</td>
				<td> <!-- next run-->
					{{ $job.NextRuntime | rfc3339 }}
				</td>
				<td> <!-- last run -->
				{{ if $job.Last }}
					{{ $job.Last.Started | rfc3339 }}
				{{ else }}
					<span class="none">-</span>
				{{ end }}
				</td>
				<td> <!-- last status -->
				{{ if $job.Last }}
					{{ if $job.Last.Err }}
						{{ $job.Last.Err }}
					{{ else }}
					<span class="none">Success</span>
					{{ end }}
				{{ else }}
					<span class="none">-</span>
				{{ end }}
				</td>
				<td><!-- last elapsed -->
				{{ if $job.Last }}
					{{ $job.Last.Elapsed }}
				{{ else }}
					<span class="none">-</span>
				{{ end }}
				</td>
			</tr>
			<tr>
				<td colspan=7>
					<h4>History</h4>
					<table class="u-full-width small-text">
						<thead>
							<tr>
								<th>Invocation</th>
								<th>Started</th>
								<th>Finished</th>
								<th>Timeout</th>
								<th>Cancelled</th>
								<th>Elapsed</th>
								<th>Error</th>
							</tr>
						</thead>
						<tbody>
						{{ range $index, $ji := $job.History | reverse }}
						<tr class="{{ if $ji.Status | eq "failed" }}failed{{ else if $ji.Status | eq "cancelled"}}cancelled{{else}}ok{{end}}">
							<td>{{ $ji.ID }}</td>
							<td>{{ $ji.Started | rfc3339 }}</td>
							<td>{{ if $ji.Finished.IsZero }}-{{ else }}{{ $ji.Finished | rfc3339 }}{{ end }}</td>
							<td>{{ if $ji.Timeout.IsZero }}-{{ else }}{{ $ji.Timeout | rfc3339 }}{{ end }}</td>
							<td>{{ if $ji.Cancelled.IsZero }}-{{ else }}{{ $ji.Cancelled | rfc3339 }}{{ end }}</td>
							<td>{{ $ji.Elapsed }}</td>
							<td>{{ if $ji.Err }}<code>{{ $ji.Err }}</code>{{ else }}-{{end}}</td>
						</tr>
						{{ else }}
						<tr>
							<td colspan=7>No History</td>
						</tr>
						{{ end }}
						</tbody>
					</table>
				</td>
			</tr>
		{{ else }}
			<tr><td colspan=7>No Jobs Loaded</td></tr>
		{{ end }}
		</tbody>
	</table>
</div>
{{ template "footer" . }}
{{ end }}
`

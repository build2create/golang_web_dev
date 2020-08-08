package main

func simpleTemplate() {
	const paymentConfirmation = `
Dear {{.Name}},
{{if .Success}}
Your transaction is successful.
{{- else}}
Transaction failure. Will recieve a message in some time.
{{- end}}
{{debited .Amount -}}
Thank you for Shopping {{.}}.
{{end}}
Best wishes,
Josie
`
}

package jwksformatter

import (
	"testing"
)

const csvTemplate string = `
Key ID,Serial,Use,Expires,Subject,Issuer
{{range .Keys}}
{{- .Kid}},{{.Serial}},{{.Use}},{{(.Expires "2006/02/01")}},{{.Subject}},{{.Issuer}}
{{end}}
`

const markdownTemplate string = `
| Key ID | Serial | Use | Expires | Subject | Issuer |
| ------ | ------ | --- | ------- | ------- | ------ |
{{range .Keys -}}
| {{- .Kid}} | {{.Serial}} | {{.Use}} | **{{(.Expires "2006/02/01")}}** | {{.Subject}} | {{.Issuer}} |
{{end}}
`

const icsTemplate string = `
BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//jwksformatter//NONSGML jwksformatter//EN
CALSCALE:GREGORIAN
METHOD:PUBLISH
{{range .Keys -}}
BEGIN:VEVENT
UUID:{{.UUID}}
SUMMARY:Certificate {{.Use}}/{{.Kid}} expires
STATUS:CONFIRMED
TRANSP:TRANSPARENT
DTSTART:{{.Expires "20060201T150405Z"}}
DTEND:{{.Expires "20060201T150405Z"}}
DTSTAMP:{{.Expires "20060201T150405Z"}}
CATEGORIES:Certificate Rotation
LOCATION:London,Europe
DESCRIPTION:{{- .Kid}},{{.Serial}},{{.Use}},{{.Issuer}}
END:VEVENT
{{end -}}
END:VCALENDAR
`

func TestLoad(t *testing.T) {
	var keyset JWKS
	var key JWK
	var err error
	var res string

	if err = keyset.Load("http://prd-keystore-obd.s3-website-eu-west-1.amazonaws.com/00158000016i44jAAA/00158000016i44jAAA.jwks"); err != nil {
		t.Errorf("TestLoad: expected no error, got: %v", err)
	}

	if key, err = keyset.Get("jM5B5CmZj0j0mpo900mS9zzV7Ck"); err != nil {
		t.Errorf("TestLoad(jM5B5CmZj0j0mpo900mS9zzV7Ck): expected key, got: %v", err)
	}

	if key.Use != "sig" {
		t.Errorf("Expected key to be of use sig, got: %v", key.Use)
	}

	if res, err = keyset.Format(csvTemplate); err != nil {
		t.Errorf("Format: expected no error, got: %v", err)
	}
	t.Log("Resulting CSV:\n", res)

	if res, err = keyset.Format(markdownTemplate); err != nil {
		t.Errorf("Format: expected no error, got: %v", err)
	}
	t.Log("Resulting Markdown:\n", res)

	if res, err = keyset.Format(icsTemplate); err != nil {
		t.Errorf("Format: expected no error, got: %v", err)
	}
	t.Log("Resulting iCalendar:\n", res)
}

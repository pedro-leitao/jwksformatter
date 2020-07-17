# A simple JWKS (JSON Web Key Store) templated formatter

Ever forgot when your certificates are due to expire ?

`./jwksformatter -uri <where your JWKS lives> -templ templates/iCalendar.tmpl > certs.ics`

And just throw `certs.ics` at Outlook or whatever calendaring application you use. Otherwise just use `templates/csv.tmpl` and import the output CSV into Excel or Google Sheets, or use `template/markdown.tmpl` and have a nice Markdown table like this:

| Key ID | Serial | Use | Expires | Subject | Issuer |
| ------ | ------ | --- | ------- | ------- | ------ |
| `ilETvqhDFWyTZJ924ML0B83x9B8` | `1509897935` | `tls` | **2020/14/08** | `CN=1qt06EuzHQOOu2CKno304T,OU=00158000016i44jAAA,O=OpenBanking,C=GB` | `CN=OpenBanking Issuing CA,O=OpenBanking,C=GB` |
| `PqkeW8paWBdtH9Kz4_4UaSjTHxg` | `1509907324` | `tls` | **2021/20/03** | `CN=ienhptqgd6uBQJWrwKwQsI,OU=00158000016i44jAAA,O=OpenBanking,C=GB` | `CN=OpenBanking Issuing CA,O=OpenBanking,C=GB` |
| `ZuZDsKg2Vs6VntPIIRD03sV42iw` | `1509911102` | `tls` | **2021/11/06** | `CN=6cchGkvTUg0JY7Vkj857Os,OU=00158000016i44jAAA,O=OpenBanking,C=GB` | `CN=OpenBanking Issuing CA,O=OpenBanking,C=GB` |

Note this tool only works with JWKS' which include the public certificate (`x5c`) as it needs to extract the `Valid To` attribute of the certificate.
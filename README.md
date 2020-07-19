# A simple JWKS (JSON Web Key Store) templated formatter

Ever forgot when your certificates are due to expire, and you have a published [OpenID Connect](https://openid.net/connect/) [JWKS](https://tools.ietf.org/html/rfc7517) ?

`./jwksformatter -template templates/iCalendar.tmpl -uri https://api.byu.edu/.well-known/byucerts`

And just throw `certs.ics` at Outlook or whatever calendaring application you use. Otherwise just use `templates/csv.tmpl` and import the output CSV into Excel or Google Sheets, or use `template/markdown.tmpl` and have a nice Markdown table like this:

| Key ID | Serial | Use | Expires | Subject | Issuer |
| ------ | ------ | --- | ------- | ------- | ------ |
| `f86d695a3866466fbf57a5fea5f56cc9` | `3776378436100916921094765356845154712` | `sig` | **2021/06/01** | `CN=wso2-is.byu.edu,O=Brigham Young University,L=Provo,ST=Utah,C=US` | `CN=DigiCert SHA2 High Assurance Server CA,OU=www.digicert.com,O=DigiCert Inc,C=US` |

Note that this tool only works with JWKS' which include the public certificate (`x5c`) as it needs to extract the `Valid To` attribute of the certificate. You can also define your own [Go template](https://golang.org/pkg/text/template/) to provide your own output format.
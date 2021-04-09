package templates

import "text/template"

var (
	Templates = template.Must(template.Must(template.Must(template.
		New("autoconfig").Parse(thunderbird)).
		New("autodiscover").Parse(outlook)).
		New("mobileconfig").Parse(applemail),
	)
)

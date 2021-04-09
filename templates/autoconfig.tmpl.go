package templates

const thunderbird = `
{{define "thunderbird"}}<?xml version="1.0" encoding="UTF-8"?>
<clientConfig version="1.1">
	<emailProvider id="{{ .Domain }}">
	    <domain>{{ .Domain }}</domain>

	    <displayName>%EMAILADDRESS%</displayName>
	    <displayShortName>%EMAILLOCALPART%</displayShortName>

	    <incomingServer type="imap">
			<hostname>{{ .IMAP.Host }}</hostname>
			<port>{{ .IMAP.Port }}</port>
			<socketType>{{ if .IMAP.STARTTLS }}STARTTLS{{ else }}SSL{{ end }}</socketType>
			<authentication>password-cleartext</authentication>
			<username>%EMAILADDRESS%</username>
		</incomingServer>

	    <outgoingServer type="smtp">
			<hostname>{{ .SMTP.Host }}</hostname>
			<port>{{ .SMTP.Port }}</port>
			<socketType>{{ if .SMTP.STARTTLS }}STARTTLS{{ else }}SSL{{ end }}</socketType>
			<authentication>password-cleartext</authentication>
			<username>%EMAILADDRESS%</username>
	    </outgoingServer>
	</emailProvider>
</clientConfig>{{end}}
`

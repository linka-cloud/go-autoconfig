package templates

const outlook = `
{{ define "outlook" }}<?xml version="1.0" encoding="utf-8"?>
<Autodiscover xmlns="http://schemas.microsoft.com/exchange/autodiscover/responseschema/2006">
	<Response xmlns="{{ .Schema }}">
		<User>
			<DisplayName>{{ .Email }}</DisplayName>
		</User>

		<Account>
			<AccountType>email</AccountType>
			<Action>settings</Action>

			<Protocol>
				<Type>IMAP</Type>
				<TTL>1</TTL>

				<Server>{{ .IMAP.Host }}</Server>
				<Port>{{ .IMAP.Port }}</Port>

				<LoginName>{{ .Email }}</LoginName>

				<DomainRequired>on</DomainRequired>
				<DomainName>{{ .Domain }}</DomainName>

				<SPA>off</SPA>
				{{ if .IMAP.STARTTLS }}<TLS>on</TLS>{{ else }}<SSL>on</SSL>{{ end }}
				<AuthRequired>on</AuthRequired>
			</Protocol>
		</Account>

		<Account>
			<AccountType>email</AccountType>
			<Action>settings</Action>

			<Protocol>
				<Type>SMTP</Type>
				<TTL>1</TTL>

				<Server>{{ .SMTP.Host }}</Server>
				<Port>{{ .SMTP.Port }}</Port>

				<LoginName>{{ .Email }}</LoginName>

				<DomainRequired>on</DomainRequired>
				<DomainName>{{ .Domain }}</DomainName>

				<SPA>off</SPA>
				{{ if .SMTP.STARTTLS }}<TLS>on</TLS>{{ else }}<SSL>on</SSL>{{ end }}
				<AuthRequired>on</AuthRequired>
			</Protocol>
		</Account>
	</Response>
</Autodiscover>{{ end }}
`

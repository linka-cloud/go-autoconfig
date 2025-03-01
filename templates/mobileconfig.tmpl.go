package templates

const applemail = `
{{ define "applemail" }}<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>HasRemovalPasscode</key>
	<false/>
	<key>PayloadContent</key>
	<array>
		<dict>
			<key>EmailAccountDescription</key>
			<string>{{ .Email }}</string>
			<key>EmailAccountName</key>
			<string>{{ .Email }}</string>
			<key>EmailAccountType</key>
			<string>EmailTypeIMAP</string>
			<key>EmailAddress</key>
			<string>{{ .Email }}</string>
			<key>IncomingMailServerAuthentication</key>
			<string>EmailAuthPassword</string>
			<key>IncomingMailServerHostName</key>
			<string>{{ .IMAP.Host }}</string>
			<key>IncomingMailServerPortNumber</key>
			<integer>{{ .IMAP.Port }}</integer>
			<key>IncomingMailServerUseSSL</key>
			<true/>
			<key>IncomingMailServerUsername</key>
			<string>{{ .Email }}</string>
			<key>OutgoingMailServerAuthentication</key>
			<string>EmailAuthPassword</string>
			<key>OutgoingMailServerHostName</key>
			<string>{{ .SMTP.Host }}</string>
			<key>OutgoingMailServerPortNumber</key>
			<integer>{{ .SMTP.Port }}</integer>
			<key>OutgoingMailServerUseSSL</key>
			<true/>
			<key>OutgoingMailServerUsername</key>
			<string>{{ .Email }}</string>
			<key>OutgoingPasswordSameAsIncomingPassword</key>
			<true/>
			<key>PayloadDescription</key>
			<string>Configure Email Settings</string>
			<key>PayloadDisplayName</key>
			<string>{{ .Email }}</string>
			<key>PayloadIdentifier</key>
			<string>com.l11r.goautoconfig.com.apple.mail.managed.7A981A9E-D5D0-4EF8-87FE-39FD6A506FAC</string>
			<key>PayloadType</key>
			<string>com.apple.mail.managed</string>
			<key>PayloadUUID</key>
			<string>7A981A9E-D5D0-4EF8-87FE-39FD6A506FAC</string>
			<key>PayloadVersion</key>
			<real>1</real>
			<key>SMIMEEnablePerMessageSwitch</key>
			<false/>
			<key>SMIMEEnabled</key>
			<false/>
			<key>disableMailRecentsSyncing</key>
			<false/>
		</dict>
	</array>
	<key>PayloadDescription</key>
	<string>Configure Email Settings</string>
	<key>PayloadDisplayName</key>
	<string>{{ .Email }}</string>
	<key>PayloadIdentifier</key>
	<string>com.l11r.goautoconfig</string>
	<key>PayloadOrganization</key>
	<string>{{ .Domain }}</string>
	<key>PayloadRemovalDisallowed</key>
	<false/>
	<key>PayloadType</key>
	<string>Configuration</string>
	<key>PayloadUUID</key>
	<string>48C88203-4DB9-49E8-B593-4831903605A0</string>
	<key>PayloadVersion</key>
	<integer>1</integer>
</dict>
</plist>{{ end }}
`

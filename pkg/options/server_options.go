package options

import "github.com/J005/interactsh/pkg/server"

type CLIServerOptions struct {
	Debug              bool
	Domain             string
	DnsPort            int
	IPAddress          string
	ListenIP           string
	HttpPort           int
	HttpsPort          int
	Hostmaster         string
	LdapWithFullLogger bool
	Eviction           int
	Responder          bool
	Smb                bool
	SmbPort            int
	SmtpPort           int
	SmtpsPort          int
	SmtpAutoTLSPort    int
	FtpPort            int
	LdapPort           int
	Ftp                bool
	Auth               bool
	Token              string
	OriginURL          string
	RootTLD            bool
	FTPDirectory       string
	SkipAcme           bool
}

func (cliServerOptions *CLIServerOptions) AsServerOptions() *server.Options {
	return &server.Options{
		Domain:          cliServerOptions.Domain,
		DnsPort:         cliServerOptions.DnsPort,
		IPAddress:       cliServerOptions.IPAddress,
		ListenIP:        cliServerOptions.ListenIP,
		HttpPort:        cliServerOptions.HttpPort,
		HttpsPort:       cliServerOptions.HttpsPort,
		Hostmaster:      cliServerOptions.Hostmaster,
		SmbPort:         cliServerOptions.SmbPort,
		SmtpPort:        cliServerOptions.SmtpPort,
		SmtpsPort:       cliServerOptions.SmtpsPort,
		SmtpAutoTLSPort: cliServerOptions.SmtpAutoTLSPort,
		FtpPort:         cliServerOptions.FtpPort,
		LdapPort:        cliServerOptions.LdapPort,
		Auth:            cliServerOptions.Auth,
		Token:           cliServerOptions.Token,
		OriginURL:       cliServerOptions.OriginURL,
		RootTLD:         cliServerOptions.RootTLD,
		FTPDirectory:    cliServerOptions.FTPDirectory,
	}
}

package configs

import (
	. "files_manager/utilities"
	"fmt"
	"strings"
)

type application struct {
	// Port where the application will run default is "7777"
	Port string

	// SecurePort where the application will run when TLS is on
	// default is "7443"
	SecurePort string

	// Host where the application will run, Leave empty to bind to 0.0.0.0
	Host string

	// UploadDir where uploaded files will be saved
	UploadDir string

	// AutoTLS to activate TLS
	AutoTLS bool

	// Domains name for AutoTLS feature
	Domains []string

	// Emails for AutoTLS feature
	Emails []string
}

var Application = application{
	Port:       GetEnvOrDefault("PORT", "7777"),
	Host:       GetEnvOrDefault("HOST", ""),
	UploadDir:  GetEnvOrDefault("UPLOAD_DIR", "./uploads"),
	Domains:    strings.Split(GetEnvOrDefault("DOMAIN", "example1.com example2.com"), " "),
	Emails:     strings.Split(GetEnvOrDefault("EMAIL", "mail@example1.com support@example2.com"), " "),
	SecurePort: GetEnvOrDefault("SECURE_PORT", "7443"),
	AutoTLS:    false,
}

/// DefaultAddress is the Host : Port
func (a application) DefaultAddress() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

/// SecureAddress is the Host : SecurePort when working with AutoTLS
func (a application) SecureAddress() string {
	return fmt.Sprintf("%s:%s", a.Host, a.SecurePort)
}

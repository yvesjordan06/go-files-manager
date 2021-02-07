package configs

import (
	. "files_manager/utilities"
	"fmt"
)

type application struct {
	Port      string
	Host      string
	UploadDir string
}

var Application = application{
	Port:      GetEnvOrDefault("PORT", "7777"),
	Host:      GetEnvOrDefault("HOST", ""),
	UploadDir: GetEnvOrDefault("UPLOAD_DIR", "./uploads"),
}

func (a application) HostAndPort() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

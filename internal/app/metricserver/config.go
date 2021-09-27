package metricserver

// Config ...
type Config struct {
	BindAddr            string `toml:"bind_addr"`
	CheckFolderInterval int    `toml:"check_folder_interval"`
	UploadFolder        string `toml:"upload_folder"`
	MetricsPath         string `toml:"metrics_path"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:            ":3000",
		CheckFolderInterval: 2,
		UploadFolder:        "upload",
		MetricsPath:         "/metrics",
	}
}

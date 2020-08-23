package main

//DBTable :
type DBTable struct {
	Name       string
	ScriptPath string
}

//DBSource :
type DBSource struct {
	Name             string
	ConnectionString string
	DownloadDuration int64
	DownloadLimit    int64

	MinimumNeedToCrawlCount int64
	PageNeedToCrawlCount    int64
}

//DBDestination :
type DBDestination struct {
	Name             string
	DatabaseEngine   string
	ConnectionString string
	Tables           []DBTable
}

//Configuration :
type Configuration struct {
	DBSources        []DBSource
	DBDestination    DBDestination
	BackupscriptPath string
	LogPath          string
}

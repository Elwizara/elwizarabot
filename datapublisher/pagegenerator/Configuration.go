package main

// Template : html template
type Template struct {
	Version  string
	Path     string
	DumpPath string
}

//Configuration :
type Configuration struct {
	DBConnectionString   string
	DBName               string
	GeneratedDataPath    string
	DumpDirectory        string
	JSONDirectory        string
	CollectionFileName   string
	LanguageRatePath     string
	DurationBetweenLangs int
	LogPath              string
	ScrapySplash         string
	DumpPagesURL         string
	CommitScriptPath     string
	CommitDuration       int64
	ElwizaraPostURL      string
	ElwizaraImageURL     string
	AllowShortURL        bool
	ElwizaraShortURL     string
	Template             Template
}

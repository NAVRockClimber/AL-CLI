package newproject

type appManifest struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Publisher        string       `json:"publisher"`
	Version          string       `json:"version"`
	Description      string       `json:"description"`
	Brief            string       `json:"brief"`
	PrivacyStatement string       `json:"privacyStatement"`
	EULA             string       `json:"EULA"`
	Help             string       `json:"help"`
	URL              string       `json:"url"`
	Logo             string       `json:"logo"`
	Dependencies     []dependency `json:"dependencies"`
	Platform         string       `json:"platform"`
	IDRanges         []idRange    `json:"idRanges"`
	SupportedLocales []string     `json:"supportedLocales"`
	Features         []string     `json:"features"`
	ShowMyCode       bool         `json:"showMyCode"`
	Runtime          string       `json:"runtime"`
}

type dependency struct {
	ID        string `json:"id"`
	Publisher string `json:"publisher"`
	Name      string `json:"name"`
	Version   string `json:"version"`
}

type idRange struct {
	From int `json:"from"`
	To   int `json:"to"`
}

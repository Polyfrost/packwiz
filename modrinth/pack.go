package modrinth

type Pack struct {
	FormatVersion uint32 `json:"formatVersion"`
	Game          string `json:"game"`
	VersionID     string `json:"versionId"`
	Name          string `json:"name"`
	Summary       string `json:"summary,omitempty"`
	// POLYFROST ADDED FIELDS START NOW
	Enabled             bool              `json:"enabled"`
	Id                  string            `json:"id"`
	Category            string            `json:"category"`
	PolyFormat          string            `json:"polyFormat"`
	UpdateUrl           string            `json:"updateUrl"`
	JavaVersionOverride int               `json:"javaVersionOverride,omitempty"`
	Files               []PackFile        `json:"files"`
	Dependencies        map[string]string `json:"dependencies"`
}

type PackFile struct {
	Path   string            `json:"path"`
	Hashes map[string]string `json:"hashes"`
	Env    *struct {
		Client string `json:"client"`
		Server string `json:"server"`
	} `json:"env"`
	Id        string   `json:"id"`
	Enabled   bool     `json:"enabled"`
	Hidden    bool     `json:"hidden"`
	Downloads []string `json:"downloads"`
	FileSize  uint32   `json:"fileSize"`
	Overrides *struct {
		Icon        string   `json:"icon,omitempty"`
		Name        string   `json:"name,omitempty"`
		Authors     []string `json:"authors,omitempty"`
		Description string   `json:"description,omitempty"`
	} `json:"overrides,omitempty"`
}

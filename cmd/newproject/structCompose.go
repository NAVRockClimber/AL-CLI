package newproject

type simpleCompose struct {
	Version  string `yaml:"version"`
	Services struct {
		BC struct {
			Image         string   `yaml:"image"`
			ContainerName string   `yaml:"container_name"`
			Volumes       []string `yaml:"volumes"`
			MemLimit      string   `yaml:"mem_limit"`
			Environment   struct {
				AcceptEula     string `yaml:"accept_eula"`
				AcceptOutdated string `yaml:"accept_outdated"`
				UseSSL         string `yaml:"usessl"`
				Licensefile    string `yaml:"licensefile"`
				Username       string `yaml:"username"`
				Password       string `yaml:"password"`
			} `yaml:"environment"`
		} `yaml:"amp"`
	} `yaml:"services"`
}

type traefikCompose struct {
	Version  string `yaml:"version"`
	Services struct {
		Amp struct {
			Image         string   `yaml:"image"`
			ContainerName string   `yaml:"container_name"`
			Volumes       []string `yaml:"volumes"`
			MemLimit      string   `yaml:"mem_limit"`
			Environment   []string `yaml:"environment"`
			Labels        []string `yaml:"labels"`
			Networks      []string `yaml:"networks"`
		} `yaml:"amp"`
	} `yaml:"services"`
	Networks struct {
		Proxy struct {
			External bool `yaml:"external"`
		} `yaml:"proxy"`
	} `yaml:"networks"`
}

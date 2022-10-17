package models

type Config struct {
	Title string `yaml:"title"`
	Quote struct {
		Text   string `yaml:"text"`
		Author string `yaml:"author"`
	} `yaml:"quote"`
	Logo      string `yaml:"logo"`
	Name      string `yaml:"name"`
	Alias     string `yaml:"alias"`
	Desc      string `yaml:"desc"`
	TechStack []struct {
		Name string `yaml:"name"`
		Icon string `yaml:"icon"`
		Url  string `yaml:"url"`
	} `yaml:"techStack"`
	Certs []struct {
		Id       string `yaml:"id"`
		Provider string `yaml:"provider"`
		Name     string `yaml:"name"`
		Image    string `yaml:"image"`
		Issuer   string `yaml:"issuer"`
		Url      string `yaml:"url"`
	} `yaml:"certs"`
}

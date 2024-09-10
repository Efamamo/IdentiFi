package domain

type Location struct {
	name      string
	link      string
	image_url string
}

type LocationConfig struct {
	Name     string
	Link     string
	ImageURL string
}

func NewLocation(config LocationConfig) *Location {
	return &Location{
		name:      config.Name,
		link:      config.Link,
		image_url: config.ImageURL,
	}
}

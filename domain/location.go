package domain

type Location struct {
	Id         string `json:"id"`
	Name       string `binding:"required" json:"name"`
	GoogleLink string `binding:"required" json:"google_link"`
	ImageURL   string `binding:"required" json:"image_url"`
}

type LocationUpdate struct {
	Name       string `json:"name"`
	GoogleLink string `json:"google_link"`
	ImageURL   string `json:"image_url"`
}

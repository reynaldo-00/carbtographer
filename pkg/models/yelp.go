package models

// YelpResponse from Yelps GraphQL API
type ParentYelpResponse struct {
	Data struct {
		Search struct {
			Total    int `json:"total"`
			Business []struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				Coordinates struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"coordinates"`
				Photos   []string `json:"photos"`
				Distance float64  `json:"distance"`
			} `json:"business"`
		} `json:"search"`
	} `json:"data"`
}

// Business db struct
type Business struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Photos   []string `json:"photos"`
	Distance float64  `json:"distance"`
}

// Search db response
type Search struct {
	Total    int        `json:"total"`
	Business []Business `json:"business"`
}

// YelpResponse db response
type YelpResponse struct {
	Search Search `json:"search"`
}
package models

type Artist struct {
	ID             int       `json:"id"`
	Image          string    `json:"image"`
	Name           string    `json:"name"`
	Members        []string  `json:"members"`
	CreationDate   int       `json:"creationDate"`
	FirstAlbum     string    `json:"firstAlbum"`
	LocationURL    string    `json:"locations"`
	Location       *Location `json:"-"`
	ConcertDateURL string    `json:"concertDates"`
	ConcertDate    *Date     `json:"-"`
	RelationURL    string    `json:"relations"`
	Relation       *Relation `json:"-"`
}

type RelationIndex struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type DateIndex struct {
	Index []Date `json:"index"`
}
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type LocationIndex struct {
	Index []Location `json:"index"`
}
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
	Dates     *Date    `json:"-"`
}

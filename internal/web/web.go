package web

import (
	"encoding/json"
	"errors"
	"fmt"
	mods "groupie-tracker/internal/models"
	"net/http"
)

var (
	Client *http.Client
)

var ErrNotFound = errors.New("artist not found")

func GetJson(client *http.Client, url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetArtists(client *http.Client, url string) ([]mods.Artist, error) {
	var artists []mods.Artist
	err := GetJson(client, url, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil

}

func GetArtist(client *http.Client, url string, artistID int) (*mods.Artist, error) {
	url = fmt.Sprintf("%s/%d", url, artistID)

	var artist mods.Artist
	err := GetJson(client, url, &artist)
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func GetDates(client *http.Client, url string) ([]mods.Date, error) {
	var datesIndex mods.DateIndex
	err := GetJson(client, url, &datesIndex)
	if err != nil {
		return nil, err
	}
	return datesIndex.Index, nil
}

func GetDate(client *http.Client, url string) (*mods.Date, error) {
	var date mods.Date
	err := GetJson(client, url, &date)
	if err != nil {
		return nil, err
	}
	return &date, nil
}

func GetLocations(client *http.Client, url string) ([]mods.Location, error) {
	var locationsIndex mods.LocationIndex
	err := GetJson(client, url, &locationsIndex)
	if err != nil {
		return nil, err
	}
	return locationsIndex.Index, nil
}

func GetLocation(client *http.Client, url string) (*mods.Location, error) {
	var location mods.Location
	err := GetJson(client, url, &location)
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func GetRelations(client *http.Client, url string) ([]mods.Relation, error) {
	var relationsIndex mods.RelationIndex

	err := GetJson(client, url, &relationsIndex)
	if err != nil {
		return nil, err
	}
	return relationsIndex.Index, nil
}

func GetRelation(client *http.Client, url string) (*mods.Relation, error) {
	var relation mods.Relation

	err := GetJson(client, url, &relation)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}

func FetchLocation(client *http.Client, url string, ch chan<- *mods.Location) {
	location, err := GetLocation(client, url)
	if err != nil {
		ch <- nil
		return
	}
	ch <- location
}

func FetchDate(client *http.Client, url string, ch chan<- *mods.Date) {
	date, err := GetDate(client, url)
	if err != nil {
		ch <- nil
		return
	}
	ch <- date
}

func FetchRelation(client *http.Client, url string, ch chan<- *mods.Relation) {
	relation, err := GetRelation(client, url)
	if err != nil {
		ch <- nil
		return
	}
	ch <- relation
}

func FetchLocationDate(client *http.Client, url string, ch chan<- *mods.Date) {
	date, err := GetDate(client, url)
	if err != nil {
		ch <- nil
		return
	}
	ch <- date
}

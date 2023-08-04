package main

import (
	"errors"
	config "groupie-tracker/internal"
	"groupie-tracker/internal/helper"
	mods "groupie-tracker/internal/models"
	"groupie-tracker/internal/responses"
	"groupie-tracker/internal/web"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	err := tpl.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		responses.ErrorResponses(w, http.StatusInternalServerError)
		return
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responses.ErrorResponses(w, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		responses.ErrorResponses(w, http.StatusBadRequest)
		return
	}

	artists, err := web.GetArtists(web.Client, config.ArtistsUrl)
	if err != nil {
		responses.ErrorResponses(w, http.StatusInternalServerError)
		return
	}

	locationCh := make(chan *mods.Location)
	dateCh := make(chan *mods.Date)
	relationCh := make(chan *mods.Relation)

	for i := 0; i < len(artists); i++ {
		go web.FetchLocation(web.Client, artists[i].LocationURL, locationCh)
		go web.FetchDate(web.Client, artists[i].ConcertDateURL, dateCh)
		go web.FetchRelation(web.Client, artists[i].RelationURL, relationCh)
	}

	for i := 0; i < len(artists); i++ {
		artists[i].Location = <-locationCh
		artists[i].ConcertDate = <-dateCh
		artists[i].Relation = <-relationCh
	}

	for i := 0; i < len(artists); i++ {
		go web.FetchDate(web.Client, artists[i].Location.DatesURL, dateCh)
	}

	for i := 0; i < len(artists); i++ {
		artists[i].Location.Dates = <-dateCh
	}

	data := map[string]interface{}{
		"Artists": artists,
	}
	renderTemplate(w, "index.html", data)
}

func detailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responses.ErrorResponses(w, http.StatusMethodNotAllowed)
		return
	}

	artistID, err := helper.ExtractArtistIDFromURL(r.URL.Path)
	if err != nil {
		responses.ErrorResponses(w, http.StatusBadRequest)
		return
	}

	artist, err := web.GetArtist(web.Client, config.ArtistsUrl, artistID)
	if err != nil {
		if errors.Is(err, web.ErrNotFound) {
			responses.ErrorResponses(w, http.StatusNotFound)
			return
		} else {
			log.Println("Internal Server Error:", err)
			responses.ErrorResponses(w, http.StatusInternalServerError)
			return
		}
	}
	if artist.ID == 0 {
		responses.ErrorResponses(w, http.StatusNotFound)
		return
	}
	artist.Location, _ = web.GetLocation(web.Client, artist.LocationURL)
	artist.ConcertDate, _ = web.GetDate(web.Client, artist.ConcertDateURL)
	artist.Relation, _ = web.GetRelation(web.Client, artist.RelationURL)
	artist.Location.Dates, _ = web.GetDate(web.Client, artist.Location.DatesURL)

	data := map[string]interface{}{
		"Artist": artist,
	}
	renderTemplate(w, "detail.html", data)
}

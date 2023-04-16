package animerecommender

import (
	"fmt"
	"io"
	"net/http"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/log"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/rwxrob/json"
	"github.com/rwxrob/to"
	"github.com/spf13/viper"

	"github.com/tr00datp00nar/fn"
)

func getConfig() (apiKey string) {
	fn.GetViperConfig(
		"config",
		"yaml",
		"$XDG_CONFIG_HOME/c",
	)
	apiKey = viper.GetString("animerecommender.api_key")
	return apiKey
}

type Recommendation struct {
	Data []struct {
		AverageScore int    `json:"averageScore"`
		BannerImage  string `json:"bannerImage"`
		CoverImage   struct {
			Large string `json:"large"`
		} `json:"coverImage"`
		Description string  `json:"description"`
		Format      string  `json:"format"`
		GenVec      []int   `json:"gen_vec"`
		ID          int     `json:"id"`
		IDMal       int     `json:"idMal"`
		Key         float64 `json:"key"`
		SeasonYear  int     `json:"seasonYear"`
		Simi        float64 `json:"simi"`
		Synonyms    []any   `json:"synonyms"`
		TagVec      []int   `json:"tag_vec"`
		Title       struct {
			English       string `json:"english"`
			Native        string `json:"native"`
			Romaji        string `json:"romaji"`
			UserPreferred string `json:"userPreferred"`
		} `json:"title"`
		Trailer struct {
			ID   string `json:"id"`
			Site string `json:"site"`
		} `json:"trailer"`
		Type string `json:"type"`
	} `json:"data"`
}

func getRecommendation(query string) {

	url := "https://anime-recommender.p.rapidapi.com/?anime_title=" + query + "&number_of_anime=50"
	apiKey := getConfig()

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "anime-recommender.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	response := body
	var rec Recommendation
	err = json.Unmarshal(response, &rec)
	if err != nil {
		log.Fatal(err)
	}

	var title string

	idx, err := fuzzyfinder.FindMulti(
		rec.Data,
		func(i int) string {
			bTitle := rec.Data[i].Title
			if len(bTitle.English) == 0 {
				if len(bTitle.Romaji) == 0 {
					title = bTitle.Native
				} else {
					title = bTitle.Romaji
				}
			} else {
				title = bTitle.English
			}
			return title
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			bTitle := rec.Data[i].Title
			if i == -1 {
				return ""
			}

			if len(bTitle.English) == 0 {
				if len(bTitle.Romaji) == 0 {
					title = bTitle.Native
				} else {
					title = bTitle.Romaji
				}
			} else {
				title = bTitle.English
			}
			desc, _ := to.Wrapped(rec.Data[i].Description, 50)
			return fmt.Sprintf("%s\nMal ID: %d\nAverage Score: %d\n\nDescription: %s",
				title,
				rec.Data[i].IDMal,
				rec.Data[i].AverageScore,
				desc)
		}))
	if err != nil {
		log.Fatal(err)
	}

	for _, indx := range idx {

		sel := rec.Data[indx].Title.English

		clipboard.WriteAll(to.String(sel))

		id := to.String(rec.Data[indx].IDMal)
		link := "https://myanimelist.net/anime/" + id

		fn.OpenURL(link)
	}
}

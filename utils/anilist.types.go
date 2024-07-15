package utils

type Anime struct {
	Count           int   `json:"count"`
	MinutesWatched  int32 `json:"minutesWatched"`
	EpisodesWatched int   `json:"episodesWatched"`
}

type Statistics struct {
	Anime Anime `json:"anime"`
}

type User struct {
	Name        string     `json:"name"`
	BannerImage string     `json:"bannerImage"`
	Statistics  Statistics `json:"statistics"`
}

type Media struct {
	Title struct {
		Romaji        string `json:"romanji"`
		English       string `json:"english"`
		Native        string `json:"native"`
		UserPreferred string `json:"userPreferred"`
	} `json:"title"`
	CoverImage struct {
		Large string `json:"large"`
	}
}

type Entry struct {
	Status   string `json:"status"`
	Progress int    `json:"progress"`
	Media    Media  `json:"media"`
}

type Entries struct {
	Entries []Entry `json:"entries"`
}

type MediaListCollection struct {
	Lists []Entries `json:"lists"`
}

type Response struct {
	User                User                `json:"user"`
	MediaListCollection MediaListCollection `json:"mediaListCollection"`
}

func (response *Response) TruncateResponse() {
	response.MediaListCollection.Lists = response.MediaListCollection.Lists[:1]
	response.MediaListCollection.Lists[0].Entries = response.MediaListCollection.Lists[0].Entries[:1]
}

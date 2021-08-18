package diaries

type Diary struct {
	ID    int    `json:"id"`
	Title string `json:"body"`
	Body  string `json:"body"`
}

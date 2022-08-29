package main

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Entry        string        `json:"entry"`
		Explanations []string      `json:"explanations"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
		Prons        struct {
		} `json:"prons"`
		Type string `json:"type"`
	} `json:"dictionary"`
}

func main() {

}

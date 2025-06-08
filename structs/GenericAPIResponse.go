package structs

type GenericAPIResponse struct {
	Count    int `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (g GenericAPIResponse) GetNames() ([]string) {
	arr := make([]string, 0)

	for _, i := range g.Results {
		arr = append(arr, i.Name)
	}

	return arr
}
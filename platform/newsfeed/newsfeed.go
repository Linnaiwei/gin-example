package newsfeed

type Item struct {
	Title string `json:"tittle"`
	Post string `json:"post"`
}

type Repo struct {
	Items []item
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
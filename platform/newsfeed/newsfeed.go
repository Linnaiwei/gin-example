package newsfeed

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

//Item News feed item
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

//Repo Slice of News feed item
type Repo struct {
	Items []Item
}

//New Return a pointer of Repo
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

//Add Append a item to Repo
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

//GetAll Get all items from Repo
func (r *Repo) GetAll() []Item {
	return r.Items
}

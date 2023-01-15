package forum

type ForumMap struct {
	URL      string
	UserName string
	Secret   string

	TableLineSelector string
	TitleSelector     string
	SizeSelector      string
	SeedSelector      string
	LeachSelector     string
	NextPageSelector  string
}

func NewForumMap() *ForumMap {
	return &ForumMap{}
}

func (f *ForumMap) Login() {

}

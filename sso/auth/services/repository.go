package services

var (
	Repo = Repository{
		dic: make(map[string]string),
	}
)

type Repository struct {
	dic map[string]string
}

func (t *Repository) Set(id, secret string) {
	t.dic[id] = secret
}

func (t *Repository) Get(id string) string {
	return t.dic[id]
}

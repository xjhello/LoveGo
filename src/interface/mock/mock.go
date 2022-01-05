package mock

type Retriever struct {
	Contexts string
	Timeout  int32
}

func (r Retriever) Get(url string) string {
	return r.Contexts
}

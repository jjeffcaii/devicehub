package core

type Indexer interface {
	Search(device *string) string
}

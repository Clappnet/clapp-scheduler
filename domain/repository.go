package domain

type Repository interface {
	Ping() error
	Close() error

	ListCollections() ([]*CollectionItem, error)
	CreateCollection(c *Collection) error

	ListJobs(collectionID string) ([]*JobItem, error)
}

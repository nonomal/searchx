package bleve

import (
	"github.com/iyear/searchx/pkg/storage/search"
)

func (b *Bleve) Index(items []*search.Item) error {
	batch := b.index.NewBatch()
	for _, item := range items {
		if err := batch.Index(item.ID, item.Data); err != nil {
			return err
		}
	}
	return b.index.Batch(batch)
}

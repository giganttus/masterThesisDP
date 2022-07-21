package middleware

import (
	"context"
	"github.com/go-pg/pg"
	"net/http"
	"time"
	"trackingApp/graph/models"
)

const itemloaderKey = "itemloader"

func DataloaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemloader := ItemLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*models.Item, []error) {

				var items []*models.Item

				err := db.Model(&items).Where("id in (?)", pg.In(ids)).Select()

				if err != nil {
					return nil, []error{err}
				}

				i := make(map[string]*models.Item, len(items))

				for _, item := range items {
					i[item.ID] = item
				}

				result := make([]*models.Item, len(ids))

				for j, id := range ids {
					result[j] = i[id]
				}

				return result, nil
			},
		}

		ctx := context.WithValue(r.Context(), itemloaderKey, &itemloader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetItemLoader(ctx context.Context) *ItemLoader {
	return ctx.Value(itemloaderKey).(*ItemLoader)
}

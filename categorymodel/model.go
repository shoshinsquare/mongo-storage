package categorymodel

import "github.com/ponlv/go-kit/mongodb"

type Category struct {
	mongodb.DefaultModel `json:",inline" bson:",inline,omitnested"`
	CreatedAt            float64 `json:"created_at" bson:"created_at,omitempty" `
	UpdatedAt            float64 `json:"updated_at" bson:"updated_at,omitempty" `
	DeletedAt            float64 `json:"deleted_at" bson:"deleted_at,omitempty"`

	Title    string `json:"title" bson:"title,omitempty"`
	Slug     string `json:"slug" bson:"slug,omitempty"`
	Priority int    `json:"priority" bson:"priority,omitempty"`

	IsDelete bool `json:"is_delete" bson:"is_delete"`
}

func (Category) CollectionName() string {
	return "category"
}

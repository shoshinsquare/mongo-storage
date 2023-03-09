package nftmodel

import (
	"time"

	"github.com/ponlv/go-kit/mongodb"
)

type NFT struct {
	mongodb.DefaultModel `json:",inline" bson:",inline,omitnested"`
	CreatedAt            time.Time `json:"created_at" bson:"created_at,omitempty" `
	UpdatedAt            time.Time `json:"updated_at" bson:"updated_at,omitempty" `
	DeletedAt            time.Time `json:"deleted_at" bson:"deleted_at,omitempty"`

	// usermodel information
	Title       string `json:"title,omitempty" bson:"title,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	ObjectID    string `json:"object_id,omitempty" bson:"object_id,omitempty"`
	Owner       string `json:"owner,omitempty" bson:"owner,omitempty"`
	RoundID     string `json:"round_id,omitempty" bson:"round_id,omitempty"`
	Type        string `json:"type,omitempty" bson:"type,omitempty"`
	URL         string `json:"url,omitempty" bson:"url,omitempty"`
	TxDigest    string `json:"tx_digest,omitempty" bson:"tx_digest,omitempty"`
}

func (NFT) CollectionName() string {
	return "nft"
}

type (
	Status int
	Role   int
)

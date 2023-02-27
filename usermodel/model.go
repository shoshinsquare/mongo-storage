package usermodel

import "github.com/ponlv/go-kit/mongodb"

type User struct {
	mongodb.DefaultModel `json:",inline" bson:",inline,omitnested"`
	CreatedAt            int64   `json:"created_at" bson:"created_at,omitempty" `
	UpdatedAt            int64   `json:"updated_at" bson:"updated_at,omitempty" `
	DeletedAt            float64 `json:"deleted_at" bson:"deleted_at,omitempty"`

	// usermodel information
	Phone         string `json:"phone,omitempty" bson:"phone,omitempty"`
	FullName      string `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Username      string `json:"username,omitempty" bson:"username,omitempty"`
	Email         string `json:"email,omitempty" bson:"email,omitempty"`
	HomeAddress   string `json:"home_address,omitempty" bson:"home_address,omitempty"`
	Avatar        string `json:"avatar,omitempty" bson:"avatar,omitempty"`
	CoverPhotoURL string `json:"cover_photo_url,omitempty" bson:"cover_photo_url,omitempty"`
	CountryCode   string `json:"country_code,omitempty" bson:"country_code,omitempty"`
	Status        Status `json:"status,omitempty" bson:"status,omitempty"`
	GoogleId      string `json:"google_id,omitempty" bson:"google_id,omitempty"`
	Role          Role   `json:"role,omitempty" bson:"role,omitempty"`
	// usermodel state
	IsDeleted  bool `json:"is_deleted,omitempty" bson:"is_deleted,omitempty"`
	IsVerified bool `json:"is_verified" bson:"is_verified"`
}

func (User) CollectionName() string {
	return "user"
}

type (
	Status int
	Role   int
)

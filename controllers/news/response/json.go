package response

import (
	_newsDomain "alta/cleanarch/businesses/news"
	"time"
)

type News struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func FromDomain(domain _newsDomain.Domain) News {
	return News{
		ID:        domain.ID,
		Title:     domain.Title,
		Content:   domain.Content,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

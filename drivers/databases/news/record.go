package news

import (
	_newsDomain "alta/cleanarch/businesses/news"
	"time"
)

type News struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

func fromDomain(domain *_newsDomain.Domain) *News {
	return &News{
		ID:        domain.ID,
		Title:     domain.Title,
		Content:   domain.Content,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *News) toDomain() _newsDomain.Domain {
	return _newsDomain.Domain{
		ID:        rec.ID,
		Title:     rec.Title,
		Content:   rec.Content,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

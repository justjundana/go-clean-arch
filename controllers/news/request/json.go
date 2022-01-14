package request

import (
	_newsDomain "alta/cleanarch/businesses/news"
)

type News struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (req *News) ToDomain() *_newsDomain.Domain {
	return &_newsDomain.Domain{
		Title:   req.Title,
		Content: req.Content,
	}
}

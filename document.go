package box

import (
	"errors"
	"fmt"
	"time"
)

type DocumentService struct {
	client *Client
}

type Document struct {
	Type      string    `json:"type"`
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type DocumentInput struct {
	URL        string `json:"url"`
	Name       string `json:"name"`
	Thumbnails string `json:"thumbnails"`
	NonSVG     bool   `json:"non_svg"`
}

func (s *DocumentService) NewURL(doc DocumentInput) (*Document, error) {
	if doc.URL == "" {
		return nil, errors.New("URL is required")
	}

	req, err := s.client.NewRequest("POST", "documents", doc)
	if err != nil {
		return nil, err
	}

	uResp := new(Document)
	_, err = s.client.Do(req, uResp)
	return uResp, err
}

func (s *DocumentService) FindOne(id string, fields string) (*Document, error) {
	u := fmt.Sprintf("/documents/%s", id)
	if fields != "" {
		u = fmt.Sprintf("%s?fields=%s", u, fields)
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	uResp := new(Document)
	_, err = s.client.Do(req, uResp)
	return uResp, err
}

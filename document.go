package box

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
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

	req, err := s.client.NewRequest("POST", "/1/documents", doc)
	if err != nil {
		return nil, err
	}

	uResp := new(Document)
	_, err = s.client.Do(req, uResp)
	return uResp, err
}

func (s *DocumentService) FindOne(id string, fields ...string) (*Document, error) {
	u := fmt.Sprintf("/1/documents/%s", id)
	if len(fields) > 0 {
		u = fmt.Sprintf("%s?fields=%s", u, strings.Join(fields, ","))
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	uResp := new(Document)
	_, err = s.client.Do(req, uResp)
	return uResp, err
}

func (s *DocumentService) GetThumbnail(id string, width, height int) (*http.Response, error) {
	u := fmt.Sprintf("/1/documents/%s/thumbnail?width=%d&height=%d", id, width, height)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	return resp, err
}

func (s *DocumentService) GetContent(id, extension string) (*http.Response, error) {
	u := fmt.Sprintf("/1/documents/%s/content.%s", id, extension)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	return resp, err
}

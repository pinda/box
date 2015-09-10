package box

import (
	"errors"
	"time"
)

type SessionService struct {
	client *Client
}

type Session struct {
	ID        string            `json:"id"`
	Type      string            `json:"type"`
	ExpiresAt time.Time         `json:"expires_at"`
	URLS      map[string]string `json:"urls"`
}

type SessionInput struct {
	DocumentID     string    `json:"document_id"`
	Duration       int       `json:"duration"`
	ExpiresAt      time.Time `json:"expires_at"`
	Downloadable   boolean   `json:"is_downloadable"`
	TextSelectable boolean   `json:"is_text_selectable"`
}

func (s *SessionService) New(session SessionInput) (*Session, error) {
	if session.DocumentID == "" {
		return nil, errors.New("Document ID is required")
	}

	req, err := s.client.NewRequest("POST", "sessions", session)
	if err != nil {
		return nil, err
	}

	uResp := new(Session)
	resp, err := s.client.Do(req, uResp)
	return &uResp, err
}

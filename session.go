package box

import (
	"errors"
	"time"
)

type SessionService struct {
	client *Client
}

type Session struct {
	ID        string              `json:"id"`
	Type      string              `json:"type"`
	ExpiresAt time.Time           `json:"expires_at"`
	URLs      map[string]string   `json:"urls"`
	Details   []map[string]string `json:"details"`
}

type SessionInput struct {
	DocumentID     string     `json:"document_id"`
	Duration       *int       `json:"duration,omitempty"`
	ExpiresAt      *time.Time `json:"expires_at,omitempty"`
	Downloadable   *bool      `json:"is_downloadable,omitempty"`
	TextSelectable *bool      `json:"is_text_selectable,omitempty"`
}

func (s *SessionService) New(session SessionInput) (*Session, error) {
	if session.DocumentID == "" {
		return nil, errors.New("Document ID is required")
	}

	req, err := s.client.NewRequest("POST", "/1/sessions", session)
	if err != nil {
		return nil, err
	}

	uResp := new(Session)
	_, err = s.client.Do(req, uResp)
	return uResp, err
}

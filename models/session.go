package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/emorydu/lenslocked/rand"
)

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session. When look up a session
	// this will be left empty, as we only store the hash of a session token
	// in our database, and we cannot reverse it into a raw token.
	Token     string
	TokenHash string
}

const (
	// MinBytesPerToken the minimum number of bytes to be used for each session token.
	MinBytesPerToken = 32
)

type SessionService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// each session token. If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used.
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// check
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	// Create the session token
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}
	// Store the session in our DB
	// 1. Try to update the user's session
	// 2. If err, create a new session
	row := ss.DB.QueryRow(`
		UPDATE sessions
		SET token_hash = $2
		WHERE user_id = $1 RETURNING id;`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if errors.Is(err, sql.ErrNoRows) {
		row = ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash) 
		VALUES ($1, $2) 
		RETURNING id;`, session.UserID, session.TokenHash)
		err = row.Scan(&session.ID)
	}
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// 1. Hash the session tokens
	hasher := ss.hash(token)

	// 2. Query for the session with that hash
	var user User
	row := ss.DB.QueryRow(`
		SELECT user_id FROM sessions 
		WHERE token_hash = $1;`, hasher)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	// 3. Using the UserID from the session, we need to query for that user
	row = ss.DB.QueryRow(`
		SELECT email, password_hash FROM users
		WHERE id = $1;`, user.ID)
	err = row.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	// 4. Return the user
	return &user, nil
}

func (ss *SessionService) hash(token string) string {
	hasher := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(hasher[:])
}

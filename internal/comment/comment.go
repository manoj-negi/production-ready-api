package comment

import (
	"context"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

var (
	ErrorFetchComment  = errors.New("some error while fetching")
	ErrorUpdateComment = errors.New("some error while updating")
	ErrorInsertComment = errors.New("error in insert")
	ErrorDeleteComment = errors.New("some error while deleting")
)

type Comment struct {
	ID     string `json:"id"`
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

type CommentStore interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	Ping(context.Context) error
}
type Service struct {
	Store CommentStore
}

func NewService(store CommentStore) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println("error in store")
		return Comment{}, ErrorFetchComment
	}
	return cmt, nil
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	cmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		log.Errorf("an error occurred adding the comment: %s", err.Error())
	}
	return cmt, nil
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(
	ctx context.Context, ID string, newComment Comment,
) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, newComment)
	if err != nil {
		log.Errorf("an error occurred updating the comment: %s", err.Error())
	}
	return cmt, nil
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	return s.Store.DeleteComment(ctx, ID)
}

// ReadyCheck - a function that tests we are functionally ready to serve requests
func (s *Service) ReadyCheck(ctx context.Context) error {
	log.Info("Checking readiness")
	return s.Store.Ping(ctx)
}

package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchComment  = errors.New("some error while fetching")
	ErrorUpdateComment = errors.New("some error while updating")
	ErrorInsertComment = errors.New("error in insert")
	ErrorDeleteComment = errors.New("some error while deleting")
)

type Comment struct {
	ID      string
	SLUG    string
	BODY    string
	AUTHOUR string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}
type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("--fetch all comments-")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println("error in store")
		return Comment{}, ErrorFetchComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrorUpdateComment
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrorDeleteComment
}

func (s *Service) InsertComment(ctx context.Context, cmt Comment) error {
	return ErrorInsertComment
}

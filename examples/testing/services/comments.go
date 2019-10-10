package services

type CommentsService struct{}

func NewComments() CommentsService {
	return CommentsService{}
}

func (c CommentsService) GetComments() ([]string, error) {
	return []string{
		"comment1",
		"comment2",
	}, nil
}

package services

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type commentsSuite struct {
	suite.Suite
	service CommentsService
}

func (s *commentsSuite) SetupSuite() {
	s.service = NewComments()
}

func (s *commentsSuite) Test_GetComment_Success() {
	expected := []string{
		"comment1",
		"comment2",
	}

	comments, err := s.service.GetComments()

	s.Require().NoError(err)
	s.Equal(expected, comments)
}

func TestComments(t *testing.T) {
	suite.Run(t, new(commentsSuite))
}

package controllers

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type commentsSuite struct {
	suite.Suite
	service *commentsServiceMock
}

func (s *commentsSuite) SetupSuite() {
	s.service = new(commentsServiceMock)
}

func (s *commentsSuite) TearDownSuite() {
	s.service.AssertExpectations(s.T())
}

func (s *commentsSuite) Test_GetComments_Success() {
	body := `["c1", "c2"]`
	req := httptest.NewRequest(
		http.MethodGet,
		"/comments",
		nil,
	)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetComments(s.service))
	s.service.
		On("GetComments").
		Return([]string{"c1", "c2"}, nil).
		Once()

	handler.ServeHTTP(w, req)

	s.Equal(200, w.Code)
	s.JSONEq(body, w.Body.String())
}

func TestComments(t *testing.T) {
	suite.Run(t, new(commentsSuite))
}

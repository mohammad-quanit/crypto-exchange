package programming

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupGin(IMock *MockInterface) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	SetRouterGroup(IMock, v1)

	return r
}

func TestPostUuidWithHyphen(t *testing.T) {
	mockInterface := MockInterface{}
	mockCall := mockInterface.On("NewUuid", false)
	mockCall.Return("1ce44be5-fe68-46f7-a153-51c1c91a4ae4")

	r := setupGin(&mockInterface)

	// Response recorder with all http status codes
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/programming/uuid", nil)

	// serving http request
	r.ServeHTTP(w, req)

	//asserting status code for success
	assert.Equal(t, w.Code, http.StatusOK)

	output := PostUuidOutput{}
	err := json.Unmarshal(w.Body.Bytes(), &output)

	assert.Nil(t, err)
	assert.Len(t, output.UUID, 36)
	assert.Contains(t, output.UUID, "-")

	mockInterface.AssertExpectations(t)
}

func TestPostUuidWithoutHyphen(t *testing.T) {
	mockInterface := MockInterface{}
	mockCall := mockInterface.On("NewUuid", true)
	mockCall.Return("1ce44be5fe6846f7a15351c1c91a4ae4")

	r := setupGin(&mockInterface)

	// Response recorder with all http status codes
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/programming/uuid?no-hyphens=true", nil)

	// serving http request
	r.ServeHTTP(w, req)

	//asserting status code for success
	assert.Equal(t, w.Code, http.StatusOK)

	output := PostUuidOutput{}
	err := json.Unmarshal(w.Body.Bytes(), &output)

	assert.Nil(t, err)
	assert.Len(t, output.UUID, 32)
	assert.NotContains(t, output.UUID, "-")

	mockInterface.AssertExpectations(t)
}

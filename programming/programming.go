package programming

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// The postUuidOutput struct defines the structure of the web service response.
type PostUuidOutput struct {
	UUID string `json:"uuid"`
}

// The SetRouterGroup function defines all the endpoints for the programming utilities
func SetRouterGroup(base *gin.RouterGroup) *gin.RouterGroup {
	programmingGroup := base.Group("/programming")
	{
		programmingGroup.POST("/uuid", postUuid())
		// programmingGroup.POST("/jwt", postJwtDebugger(p))
	}

	return programmingGroup
}

// The postUuid function, returns another function Once the server receives the POST /programming/uuid request
func postUuid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		noHyphensParamValue := ctx.Query("no-hyphens")
		isWithoutHyphens := noHyphensParamValue == "true"

		uuid := NewUuid(isWithoutHyphens)
		output := PostUuidOutput{UUID: uuid}

		ctx.JSON(http.StatusOK, output)
	}
}

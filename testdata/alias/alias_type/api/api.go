package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/swag/testdata/alias/alias_type/data"
)

/*// @Summary Get time as string
// @Description get time as string
// @ID time-as-string
// @Accept  json
// @Produce  json
// @Success 200 {object} data.StringAlias	"ok"
// @Router /testapi/time-as-string [get]
func GetTimeAsStringAlias(c *gin.Context) {
	var foo data.StringAlias = "test"
	log.Println(foo)
	//write your code
}*/

/*// @Summary Get time as time
// @Description get time as time
// @ID time-as-time
// @Accept  json
// @Produce  json
// @Success 200 {object} data.DateOnly	"ok"
// @Router /testapi/time-as-time [get]
func GetTimeAsTimeAlias(c *gin.Context) {
	var foo = data.DateOnly(time.Now())
	log.Println(foo)
	//write your code
}*/

// @Summary Get container with time and time alias
// @Description test container with time and time alias
// @ID time-as-time-container
// @Accept  json
// @Produce  json
// @Success 200 {object} data.TimeContainer	"ok"
// @Router /testapi/time-as-time-container [get]
func GetTimeAsTimeContainer(c *gin.Context) {
	now := time.Now()
	var foo = data.TimeContainer{
		Name:      "test",
		Timestamp: now,
		//CreatedAt: &now,
	}
	log.Println(foo)
	//write your code
}

// @Summary Get container with external alias type
// @Description test container with external alias type
// @ID external-alias-type-container
// @Accept  json
// @Produce  json
// @Success 200 {object} alias_external_types.SQLNullInt64Alias	"ok"
// @Router /testapi/external-alias-type-container [get]
func GetExternalAliasTypeContainer(c *gin.Context) {}

package v1

import (
	"github.com/gin-gonic/gin"
)

// gin-swagger 范例
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]

func GetTags(ctx *gin.Context) {
	//name := ctx.Query("name")
	//
	//maps := make(map[string]any)
	//data := make(map[string]any)
	//
	//if name != "" {
	//	maps["name"] = name
	//}
	//
	//var state int = -1
	//if arg := ctx.Query("state"); arg != "" {
	//	state = com.StrTo(arg).MustInt()
	//	maps["state"] = state
	//}
	//
	//code := e.SUCCESS
	//
	//data["lists"] = models.GetTags(util.GetPage(ctx), setting.AppSetting.PageSize, maps)
	//data["total"] = models.GetTagTotal(maps)
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": data,
	//})
}

// AddTag godoc
// @Summary     新增文章标签
// @Description 新增文章标签
// @Tags        tag
// @Accept      json
// @Produce     json
// @Param       name            query       string  true    "name of tag"
// @Param       state           query       int     false   "State"
// @Param       created_by      query       string  false   "CreatedBy"
// @Success     200             {string}    json    {"code": 200, "msg": "ok"}
// @Router      /api/v1/tags    [post]
func AddTag(ctx *gin.Context) {
	//name := ctx.Query("name")
	//state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()
	//createBy := ctx.Query("created_by")
	//
	//valid := validation.Validation{}
	//
	//valid.Required(name, "name").Message("name can't be empty")
	//valid.MaxSize(name, 100, "name").Message("name max length 100")
	//valid.Required(createBy, "create_by").Message("create user can't be empty")
	//valid.MaxSize(createBy, 100, "create_by").Message("create by max length 100")
	//valid.Range(state, 0, 1, "state").Message("invalid state")
	//
	//code := e.INVALID_PARAMS
	//if !valid.HasErrors() {
	//	if !models.ExistTagByName(name) {
	//		code = e.SUCCESS
	//		models.AddTag(name, state, createBy)
	//	} else {
	//		code = e.ERROR_EXIST_TAG
	//	}
	//} else {
	//	for _, err := range valid.Errors {
	//		log.Printf("err.key: %s, err.message: %s\n", err.Key, err.Message)
	//	}
	//}
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//})

}

func EditTag(ctx *gin.Context) {
	//id := com.StrTo(ctx.Param("id")).MustInt()
	//name := ctx.Query("name")
	//modifiedBy := ctx.Query("modified_by")
	//
	//valid := validation.Validation{}
	//
	//state := -1
	//if arg := ctx.Query("state"); arg != "" {
	//	state = com.StrTo(arg).MustInt()
	//	valid.Range(state, 0, 1, "state").Message("invalid state")
	//}
	//
	//valid.Required(id, "id").Message("ID can't be empty")
	//valid.Required(modifiedBy, "modified_by").Message("modify user can't be empty")
	//
	//code := e.INVALID_PARAMS
	//if !valid.HasErrors() {
	//	code = e.SUCCESS
	//
	//	if models.ExistTagByID(id) {
	//		data := make(map[string]any)
	//		data["modified_by"] = modifiedBy
	//		if name != "" {
	//			data["name"] = name
	//		}
	//		if state != -1 {
	//			data["state"] = state
	//		}
	//
	//		models.EditTag(id, data)
	//	} else {
	//		code = e.ERROR_NOT_EXIST_TAG
	//	}
	//}
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//})
}

func DeleteTag(ctx *gin.Context) {
	//id := com.StrTo(ctx.Param("id")).MustInt()
	//
	//valid := validation.Validation{}
	//valid.Min(id, 1, "id").Message("id must gt 0")
	//
	//code := e.SUCCESS
	//if !valid.HasErrors() {
	//	code = e.SUCCESS
	//	if models.ExistTagByID(id) {
	//		models.DeleteTag(id)
	//	} else {
	//		code = e.ERROR_NOT_EXIST_TAG
	//	}
	//}
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//})
}

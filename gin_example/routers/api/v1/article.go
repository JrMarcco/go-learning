package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_learning/gin_example/pkg/app"
	"go_learning/gin_example/pkg/e"
	"go_learning/gin_example/pkg/setting"
	"go_learning/gin_example/pkg/util"
	"go_learning/gin_example/service/article_service"
	"go_learning/gin_example/service/tag_service"
	"net/http"
)

// 省略验证部分

func GetArticle(ctx *gin.Context) {
	g := app.Gin{Ctx: ctx}

	id := com.StrTo(ctx.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("Invalid ID")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		g.Resp(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	as := article_service.Article{ID: id}
	exists, err := as.ExistByID()
	if err != nil {
		g.Resp(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	if !exists {
		g.Resp(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
	}

	article, err := as.Get()
	if err != nil {
		g.Resp(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	g.Resp(http.StatusOK, e.SUCCESS, article)
}

func GetArticles(ctx *gin.Context) {
	g := app.Gin{Ctx: ctx}

	state := -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagId := -1
	if arg := ctx.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
	}

	as := article_service.Article{
		TagID:    tagId,
		State:    state,
		PageNo:   util.GetPage(ctx),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := as.Count()
	if err != nil {
		g.Resp(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	articles, err := as.GetAll()
	if err != nil {
		g.Resp(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]any)

	data["lists"] = articles
	data["total"] = total

	g.Resp(http.StatusOK, e.SUCCESS, data)
}

func AddArticle(ctx *gin.Context) {

	var g = app.Gin{Ctx: ctx}

	tagId := com.StrTo(ctx.Query("tag_id")).MustInt()
	title := ctx.Query("title")
	desc := ctx.Query("desc")
	content := ctx.Query("content")
	createdBy := ctx.Query("created_by")
	state := com.StrTo(ctx.DefaultQuery("state", "0")).MustInt()

	ts := tag_service.Tag{ID: tagId}
	exists, err := ts.ExistByID()
	if err != nil {
		g.Resp(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	if !exists {
		g.Resp(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	as := article_service.Article{
		TagID:     tagId,
		Title:     title,
		Desc:      desc,
		Content:   content,
		State:     state,
		CreatedBy: createdBy,
	}

	if err := as.Add(); err != nil {
		g.Resp(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	g.Resp(http.StatusOK, e.SUCCESS, nil)
}

func EditArticle(ctx *gin.Context) {
	//id := com.StrTo(ctx.Param("id")).MustInt()
	//tagId := com.StrTo(ctx.Query("tag_id")).MustInt()
	//title := ctx.Query("title")
	//desc := ctx.Query("desc")
	//content := ctx.Query("content")
	//modifiedBy := ctx.Query("modified_by")
	//
	//state := -1
	//if arg := ctx.Query("state"); arg != "" {
	//	state = com.StrTo(arg).MustInt()
	//}
	//
	//code := e.SUCCESS
	//if models.ExistArticleByID(id) {
	//	if models.ExistTagByID(tagId) {
	//		data := make(map[string]any)
	//		if tagId > 0 {
	//			data["tag_id"] = tagId
	//		}
	//		if state != -1 {
	//			data["state"] = state
	//		}
	//		if title != "" {
	//			data["title"] = title
	//		}
	//		if desc != "" {
	//			data["desc"] = desc
	//		}
	//		if content != "" {
	//			data["content"] = content
	//		}
	//		data["modified_by"] = modifiedBy
	//
	//		models.EditArticle(id, data)
	//	} else {
	//		code = e.ERROR_NOT_EXIST_TAG
	//	}
	//} else {
	//	code = e.ERROR_NOT_EXIST_ARTICLE
	//}
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//})
}

func DeleteArticle(ctx *gin.Context) {
	//id := com.StrTo(ctx.Param("id")).MustInt()
	//
	//code := e.SUCCESS
	//
	//if models.ExistArticleByID(id) {
	//	models.DeleteArticle(id)
	//} else {
	//	code = e.ERROR_NOT_EXIST_ARTICLE
	//}
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//})
}

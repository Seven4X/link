package link

import (
	"github.com/Seven4X/link/web/app"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/api/messages"
	"github.com/Seven4X/link/web/library/config"
	"github.com/Seven4X/link/web/library/consts"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	svr = NewService()
)

/*
流控配置
[{
    "resource":"GET:/link/preview-token",
    "metricType":0,
    "tokenCalculateStrategy":0,
    "controlBehavior":0,
    "count": 3
}]

验证明令： ab -n1000 -c100 http://localhost:1323/link/preview-token
*/
func Router(e *echo.Echo) {
	g := e.Group("/link")
	g.POST("", createLink, mymw.JWT())
	g.GET("", listLink)
	g.POST("/actions/batch", batchImport, mymw.JWT())
	g.GET("/marks/hot", hotLink)
	g.GET("/marks/newest", newestLink)
	g.GET("/marks/mine", mineLink, mymw.JWT())
	g.GET("/marks/my", myAllPostLink, mymw.JWT())
	g.POST("/:lid/comment", postComment, mymw.JWT())
	g.GET("/:lid/comment", listComment)
	g.DELETE("/:lid/comment/:mid", deleteComment, mymw.JWT())
	g.GET("/preview-token", getPreviewToken)
}

// 解析URL中所有 a标签中的超链接保存
func batchImport(e echo.Context) error {
	url := e.Param("url")

	return e.JSON(http.StatusOK, map[string]interface{}{
		"url":  url,
		"size": 0,
	})

}

func getPreviewToken(e echo.Context) error {
	str := config.GetString(config.LinkPreviewToken)

	_ = e.HTML(http.StatusOK, str)
	return nil
}

func createLink(e echo.Context) error {
	req := new(CreateLinkRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	u := e.Get(consts.User)
	if u == nil {
		e.JSON(http.StatusOK, api.FailMsgId(messages.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*mymw.JwtCustomClaims)
	link := req.ToLink()
	link.CreateBy = claims.Id
	id, err := svr.Save(link)
	_ = e.JSON(http.StatusOK, api.Response(id, err))

	return nil
}

/*

 */
func listLink(e echo.Context) error {
	req := new(ListLinkRequest)
	e.Bind(req)
	if err := e.Validate(req); err != nil {
		return err
	}
	uid := app.GetUserId(e)
	req.UserId = uid
	res, err := svr.ListLink(req)
	e.JSON(http.StatusOK, api.Response(res, err))
	return nil
}
func hotLink(context echo.Context) error {
	return nil
}
func newestLink(context echo.Context) error {
	return nil
}
func mineLink(context echo.Context) error {
	return nil
}
func myAllPostLink(context echo.Context) error {
	return nil
}
func postComment(e echo.Context) error {
	req := new(NewCommentRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	u := e.Get(consts.User)
	if u == nil {
		e.JSON(http.StatusOK, api.FailMsgId(messages.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*mymw.JwtCustomClaims)
	req.CreateBy = claims.Id
	lid := e.Param("lid")
	if id, err := strconv.Atoi(lid); err != nil {
		e.JSON(http.StatusOK, api.FailMsgId(messages.GlobalParamWrong))
	} else {
		req.LinkId = id
	}

	id, err := svr.SaveComment(req)
	_ = e.JSON(http.StatusOK, api.Response(id, err))
	return nil
}
func listComment(context echo.Context) error {
	return nil
}
func deleteComment(context echo.Context) error {
	return nil
}
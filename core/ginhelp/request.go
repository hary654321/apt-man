package ginhelp

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetQueryParams(c *gin.Context) map[string]any {
	query := c.Request.URL.Query()
	var queryMap = make(map[string]any, len(query))
	for k := range query {
		if k == "limit" || k == "offset" || k == "order" {
			continue
		}
		if c.Query(k) == "" {
			continue
		}
		queryMap[k] = strings.Trim(c.Query(k), " ")
	}
	return queryMap
}

func GetPostFormParams(c *gin.Context) (map[string]any, error) {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		if !errors.Is(err, http.ErrNotMultipart) {
			return nil, err
		}
	}
	var postMap = make(map[string]any, len(c.Request.PostForm))
	for k, v := range c.Request.PostForm {
		if len(v) > 1 {
			postMap[k] = v
		} else if len(v) == 1 {
			postMap[k] = v[0]
		}
	}

	return postMap, nil
}

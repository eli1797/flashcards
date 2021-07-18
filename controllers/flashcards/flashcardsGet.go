package flashcards

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FlashcardsGet(c *gin.Context) ([]byte, int, error) {

	obj := make(map[string]string)
	obj["flashcards"] = "example"

	res, err := json.Marshal(obj)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}

package test

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MadMaxMR/Products-Restful/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{gin.Param{Key: "sku", Value: "FAL-8406280"}}

	controllers.DeleteProduct(c)

	assert.Equal(t, w.Code, 200)
	if w.Code != 200 {
		t.Errorf("Expected status 200, got %v", w.Code)
		t.Errorf("Error: %v", w.Body.String())
	}

	os.RemoveAll("public")
}

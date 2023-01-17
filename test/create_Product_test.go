package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/MadMaxMR/Products-Restful/controllers"
	"github.com/MadMaxMR/Products-Restful/database"
	"github.com/MadMaxMR/Products-Restful/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("sku", "FAL-8406280")
	_ = writer.WriteField("name", "500 Zapatilla Urbana Mujer")
	_ = writer.WriteField("brand", "New Balance")
	_ = writer.WriteField("size", "37")
	_ = writer.WriteField("price", "42990.10")

	img1Resp, err := os.Open("../test/imgTest.jpg")
	if err != nil {
		t.Errorf("Error abriendo: " + err.Error())
	}
	defer img1Resp.Close()
	part6, err := writer.CreateFormFile("image", "imgTest.jpg")
	if err != nil {
		t.Errorf("Error creando: " + err.Error())
	}
	_, err = io.Copy(part6, img1Resp)
	if err != nil {
		t.Errorf("Error copy: " + err.Error())
	}

	img2Resp, err := os.Open("../test/imgTest.jpg")
	if err != nil {
		t.Errorf("Error: " + err.Error())
	}
	defer img2Resp.Close()
	part7, err := writer.CreateFormFile("otherImages", "imgTest.jpg")
	if err != nil {
		t.Errorf("Error: " + err.Error())
	}
	_, err = io.Copy(part7, img2Resp)
	if err != nil {
		t.Errorf("Error copy: " + err.Error())
	}

	part8, err := writer.CreateFormFile("otherImages", filepath.Base("/../test/imgTest.jpg"))
	if err != nil {
		t.Errorf("Error: " + err.Error())
	}
	img3Resp, err := os.Open("../test/imgTest.jpg")
	if err != nil {
		t.Errorf("Error: " + err.Error())
	}
	io.Copy(part8, img3Resp)
	defer img3Resp.Close()

	writer.Close()

	c.Request, _ = http.NewRequest("POST", "/products", payload)

	c.Request.Header.Add("Content-Type", writer.FormDataContentType())

	controllers.CreateProduct(c)

	assert.Equal(t, c.Writer.Status(), 201)

	var result models.Product

	db := database.GetConnection()
	defer db.Close()

	db.First(&result, "SKU=", "FAL-8406280")
	assert.Equal(t, result.Name, "500 Zapatilla Urbana Mujer")
	assert.Equal(t, result.Brand, "New Balance")
	assert.Equal(t, result.Size, "37")
	assert.Equal(t, result.Price, "42990.10")
	assert.Equal(t, result.SKU, "FAL-8406280")

	if w.Code != 201 {
		t.Errorf("Expected status 201, got %v", w.Code)
		t.Errorf("Error: %v", w.Body.String())
		t.Log("FilePath ", img3Resp)

	}
}

package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/MadMaxMR/Products-Restful/controllers"
	"github.com/MadMaxMR/Products-Restful/database"
	"github.com/MadMaxMR/Products-Restful/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {

	os.MkdirAll("public/img/product/", os.ModePerm)
	os.MkdirAll("public/otherImages/", os.ModePerm)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	payload := new(bytes.Buffer)
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("sku", "FAL-8406280")
	_ = writer.WriteField("name", "500 Zapatilla Urbana Mujer")
	_ = writer.WriteField("brand", "New Balance")
	_ = writer.WriteField("size", "37")
	_ = writer.WriteField("price", "42990.10")

	//media part PRINCIPAL IMAGE

	img1Resp, err := os.Open("../test/imgTest.jpg")
	if err != nil {
		t.Errorf("Error: " + err.Error())
	}
	defer img1Resp.Close()

	part6, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": {`form-data;name ="image";filename="imgTest.jpeg"`},
		"Content-Type":        {"image/jpeg"},
	})
	if err != nil {
		t.Errorf("Error creando Part1: " + err.Error())
	}
	io.Copy(part6, img1Resp)

	//media part OTHER IMAGES
	img2Resp, err := os.Open("../test/imgTest.jpg")
	if err != nil {
		t.Errorf("Error: " + err.Error())
	}
	defer img2Resp.Close()
	part7, err := writer.CreateFormFile("otherImages", "imgTest.jpg")
	if err != nil {
		t.Errorf("Error: " + err.Error())
	}
	io.Copy(part7, img2Resp)

	//media part OTHER IMAGES
	img3Resp, err := os.Open("../test/imgTest.jpg")
	if err != nil {
		t.Errorf("Error: " + err.Error())
	}
	defer img3Resp.Close()
	part8, err := writer.CreateFormFile("otherImages", "imgTest.jpg")
	if err != nil {
		t.Errorf("Error: " + err.Error())
	}
	io.Copy(part8, img3Resp)

	writer.Close()

	c.Request, _ = http.NewRequest("POST", "/products", payload)

	c.Request.Header.Add("Content-Type", writer.FormDataContentType())

	controllers.CreateProduct(c)

	assert.Equal(t, c.Writer.Status(), 201)

	var result models.Product

	db := database.GetConnection()
	defer db.Close()
	db.Preload("OtherImages").First(&result, "SKU='FAL-8406280'")
	//db.First(&result, "SKU='FAL-8406280'")
	assert.Equal(t, result.SKU, "FAL-8406280")
	assert.Equal(t, result.Name, "500 Zapatilla Urbana Mujer")
	assert.Equal(t, result.Brand, "New Balance")
	assert.Equal(t, result.Size, "37")
	price, _ := strconv.ParseFloat("42990.10", 32)
	assert.Equal(t, result.Price, float32(price))

	var fileP = strings.Split(result.PrincipalImage, "/product/")[1]
	assert.FileExists(t, "public/img/product/"+fileP)

	for i := 0; i < len(result.OtherImages); i++ {
		var fileP = strings.Split(result.OtherImages[i].ImageURL, "/product/")[1]
		assert.Equal(t, result.OtherImages[i].ProductSKU, "FAL-8406280")
		assert.FileExists(t, "public/otherImages/"+fileP)
	}

	if w.Code != 201 {
		t.Errorf("Expected status 201, got %v", w.Code)
		t.Errorf("Error: %v", w.Body.String())
		//t.Log("Image de copy ",)
	}

}

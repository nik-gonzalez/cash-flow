package handlers

import (
	"bytes"
	"cash-flow-service/api"
	"cash-flow-service/mocks"
	"cash-flow-service/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandler_AddCashFlow(t *testing.T) {
	type fields struct {
		service api.IService
	}
	type args struct {
		c *gin.Context
		w *httptest.ResponseRecorder
	}
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	mService := mocks.IService{}
	mService.On("AddCashFlow", mock.Anything).Return(nil)

	wWithJson := httptest.NewRecorder()
	contextWithJson, _ := gin.CreateTestContext(wWithJson)
	contextWithJson.Request = &http.Request{
		Header: make(http.Header),
	}
	mockJson(contextWithJson, models.CashFlow{ID: "1"}, "POST")
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{"add cash-flow no json", fields{&mService}, args{context, w}, http.StatusBadRequest},
		{"add cash-flow", fields{&mService}, args{contextWithJson, wWithJson}, http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := Handler{
				service: tt.fields.service,
			}
			handler.AddCashFlow(tt.args.c)
			assert.Equal(t, tt.args.w.Result().StatusCode, tt.wantStatus)
		})
	}
}

func TestHandler_DeleteCashFlow(t *testing.T) {
	type fields struct {
		service api.IService
	}
	type args struct {
		c *gin.Context
	}
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}
	mService := mocks.IService{}
	mService.On("DeleteCashFlow", mock.Anything).Return(nil)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"delete cash-flow", fields{&mService}, args{context}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := Handler{
				service: tt.fields.service,
			}
			handler.DeleteCashFlow(tt.args.c)
		})
	}
}

func TestHandler_GetCashFlow(t *testing.T) {
	type fields struct {
		service api.IService
	}
	type args struct {
		c *gin.Context
	}
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	mService := mocks.IService{}
	mService.On("GetCashFlow", mock.Anything).Return(&models.CashFlow{}, nil)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"get cash-flow", fields{&mService}, args{context}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := Handler{
				service: tt.fields.service,
			}
			handler.GetCashFlow(tt.args.c)
		})
	}
}

func TestHandler_GetCashFlows(t *testing.T) {
	type fields struct {
		service api.IService
	}
	type args struct {
		c *gin.Context
	}
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	mService := mocks.IService{}
	mService.On("GetCashFlows", mock.Anything).Return([]models.CashFlow{{ID: "1"}, {ID: "2"}}, nil)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"get cash-flows", fields{&mService}, args{context}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := Handler{
				service: tt.fields.service,
			}
			handler.GetCashFlows(tt.args.c)
		})
	}
}

func TestHandler_UpdateCashFlows(t *testing.T) {
	type fields struct {
		service api.IService
	}
	type args struct {
		c *gin.Context
		w *httptest.ResponseRecorder
	}
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	mService := mocks.IService{}
	mService.On("UpdateCashFlow", mock.Anything).Return(nil)

	wWithJson := httptest.NewRecorder()
	contextWithJson, _ := gin.CreateTestContext(wWithJson)
	contextWithJson.Request = &http.Request{
		Header: make(http.Header),
	}
	mockJson(contextWithJson, models.CashFlow{ID: "1"}, "PUT")

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{"update cash-flow", fields{&mService}, args{context, w}, http.StatusBadRequest},
		{"update cash-flow with json", fields{&mService}, args{contextWithJson, wWithJson}, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := Handler{
				service: tt.fields.service,
			}
			handler.UpdateCashFlows(tt.args.c)
			assert.Equal(t, tt.args.w.Result().StatusCode, tt.wantStatus)
		})
	}
}
func mockJson(c *gin.Context /* the test context */, content interface{}, method string) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestNewWithService(t *testing.T) {
	type args struct {
		service api.IService
	}
	mService := mocks.IService{}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{"with new service test", args{&mService}, &Handler{&mService}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() =got %v, want %v", got, tt.want)
			}
		})
	}
}

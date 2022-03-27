package handlers

import (
	"cash-flow-service/api"
	"cash-flow-service/models"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"reflect"
	"testing"
)

type mockService struct{}

func (s mockService) GetCashFlows() ([]models.CashFlow, error) {
	return nil, nil
}
func (s mockService) AddCashFlow(cashFlow models.CashFlow) error {
	return nil
}
func (s mockService) UpdateCashFlow(cashFlow models.CashFlow) error {
	return nil
}
func (s mockService) GetCashFlow(id string) (*models.CashFlow, error) {
	return nil, nil
}
func (s mockService) DeleteCashFlow(id string) error {
	return nil
}
func TestHandler_AddCashFlow(t *testing.T) {
	type fields struct {
		service api.IService
	}
	type args struct {
		c *gin.Context
	}
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"add cash-flow", fields{mockService{}}, args{context}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := Handler{
				service: tt.fields.service,
			}
			handler.AddCashFlow(tt.args.c)
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
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"add cash-flow", fields{mockService{}}, args{context}},
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
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"add cash-flow", fields{mockService{}}, args{context}},
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
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"add cash-flow", fields{mockService{}}, args{context}},
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
	}
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"add cash-flow", fields{mockService{}}, args{context}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := Handler{
				service: tt.fields.service,
			}
			handler.UpdateCashFlows(tt.args.c)
		})
	}
}

func TestNewWithService(t *testing.T) {
	type args struct {
		service api.IService
	}
	tests := []struct {
		name string
		args args
		want Handler
	}{
		{"with new service test", args{mockService{}}, Handler{mockService{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

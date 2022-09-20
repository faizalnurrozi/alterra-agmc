package controllers

import (
	"github.com/faizalnurrozi/alterra-agmc/config"
	"github.com/faizalnurrozi/alterra-agmc/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBookController_GetBooks(t *testing.T) {

	/**
	 * Initiate config and echo framework
	 */

	config.InitDB()
	contract := config.Contract{DB: config.DB}
	e := echo.New()

	/**
	 * Declare commons struct
	 */

	type fields struct {
		Contract config.Contract
	}
	type args struct {
		ctx echo.Context
	}

	/**
	 * List of test case
	 */

	tests := []struct {
		name               string
		method             string
		fields             fields
		args               args
		statusCodeExpected int
		wantErr            bool
	}{
		{
			name:               "ok success",
			method:             http.MethodGet,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "nok invalid method",
			method:             http.MethodPost,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusMethodNotAllowed,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/v1/books")
			tt.args = args{ctx: ctx}

			uc := NewBookController(tt.fields.Contract)
			if assert.NoError(t, uc.GetBooks(ctx)) || assert.Error(t, uc.GetBooks(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

func TestBookController_GetBookByID(t *testing.T) {

	/**
	 * Initiate config and echo framework
	 */

	config.InitDB()
	contract := config.Contract{DB: config.DB}
	e := echo.New()

	/**
	 * Declare commons struct
	 */

	type fields struct {
		Contract config.Contract
	}
	type args struct {
		ctx echo.Context
	}

	/**
	 * List of test case
	 */

	tests := []struct {
		name               string
		method             string
		id                 string
		fields             fields
		args               args
		statusCodeExpected int
		wantErr            bool
	}{
		{
			name:               "ok success",
			method:             http.MethodGet,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "nok invalid method",
			method:             http.MethodPost,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusMethodNotAllowed,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/v1/books/" + tt.id)
			tt.args = args{ctx: ctx}

			uc := NewBookController(tt.fields.Contract)
			if assert.NoError(t, uc.GetBookByID(ctx)) || assert.Error(t, uc.GetBookByID(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

func TestBookController_Create(t *testing.T) {

	/**
	 * Initiate config and echo framework
	 */

	config.InitDB()
	contract := config.Contract{DB: config.DB}
	e := echo.New()

	/**
	 * Declare commons struct
	 */

	type fields struct {
		Contract config.Contract
	}
	type args struct {
		ctx echo.Context
	}

	/**
	 * List of test case
	 */

	tests := []struct {
		name               string
		jwt                bool
		method             string
		fields             fields
		args               args
		statusCodeExpected int
		wantErr            bool
	}{
		{
			name:               "ok success",
			jwt:                true,
			method:             http.MethodPost,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusCreated,
			wantErr:            false,
		},
		{
			name:               "nok without jwt POST",
			jwt:                false,
			method:             http.MethodPost,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            true,
		},
		{
			name:               "nok without jwt PUT",
			jwt:                false,
			method:             http.MethodPut,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            true,
		},
		{
			name:               "nok invalid method",
			jwt:                true,
			method:             http.MethodGet,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusMethodNotAllowed,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/v1/books")
			tt.args = args{ctx: ctx}

			if tt.jwt {
				tokenJWT, _ := middlewares.JWTMiddleware{}.CreateToken(1)
				req.Header.Set("Authorization", "Bearer "+tokenJWT)
			}

			uc := NewBookController(tt.fields.Contract)
			if assert.NoError(t, uc.Create(ctx)) || assert.Error(t, uc.Create(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

func TestBookController_Update(t *testing.T) {

	/**
	 * Initiate config and echo framework
	 */

	config.InitDB()
	contract := config.Contract{DB: config.DB}
	e := echo.New()

	/**
	 * Declare commons struct
	 */

	type fields struct {
		Contract config.Contract
	}
	type args struct {
		ctx echo.Context
	}

	/**
	 * List of test case
	 */

	tests := []struct {
		name               string
		jwt                bool
		method             string
		id                 string
		fields             fields
		args               args
		statusCodeExpected int
		wantErr            bool
	}{
		{
			name:               "ok success",
			jwt:                true,
			method:             http.MethodPut,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "nok without jwt PUT",
			jwt:                false,
			method:             http.MethodPut,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            true,
		},
		{
			name:               "nok without jwt GET",
			jwt:                false,
			method:             http.MethodGet,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            true,
		},
		{
			name:               "nok invalid method",
			jwt:                true,
			method:             http.MethodPost,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusMethodNotAllowed,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/v1/books/" + tt.id)
			tt.args = args{ctx: ctx}

			if tt.jwt {
				tokenJWT, _ := middlewares.JWTMiddleware{}.CreateToken(1)
				req.Header.Set("Authorization", "Bearer "+tokenJWT)
			}

			uc := NewBookController(tt.fields.Contract)
			if assert.NoError(t, uc.Update(ctx)) || assert.Error(t, uc.Update(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

func TestBookController_Delete(t *testing.T) {

	/**
	 * Initiate config and echo framework
	 */

	config.InitDB()
	contract := config.Contract{DB: config.DB}
	e := echo.New()

	/**
	 * Declare commons struct
	 */

	type fields struct {
		Contract config.Contract
	}
	type args struct {
		ctx echo.Context
	}

	/**
	 * List of test case
	 */

	tests := []struct {
		name               string
		jwt                bool
		method             string
		id                 string
		fields             fields
		args               args
		statusCodeExpected int
		wantErr            bool
	}{
		{
			name:               "ok success",
			jwt:                true,
			method:             http.MethodDelete,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "nok without jwt DELETE",
			jwt:                false,
			method:             http.MethodDelete,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            true,
		},
		{
			name:               "nok without jwt GET",
			jwt:                false,
			method:             http.MethodGet,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            true,
		},
		{
			name:               "nok invalid method",
			jwt:                true,
			method:             http.MethodPost,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusMethodNotAllowed,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/v1/books/" + tt.id)
			tt.args = args{ctx: ctx}

			if tt.jwt {
				tokenJWT, _ := middlewares.JWTMiddleware{}.CreateToken(1)
				req.Header.Set("Authorization", "Bearer "+tokenJWT)
			}

			uc := NewBookController(tt.fields.Contract)
			if assert.NoError(t, uc.Delete(ctx)) || assert.Error(t, uc.Delete(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

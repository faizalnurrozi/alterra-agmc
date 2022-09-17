package controllers

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/middlewares"
	"strings"
	"testing"
)

func TestUserController_GetUsers(t *testing.T) {

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
			method:             http.MethodGet,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "nok without jwt GET",
			jwt:                false,
			method:             http.MethodGet,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            false,
		},
		{
			name:               "nok without jwt POST",
			jwt:                false,
			method:             http.MethodPost,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            false,
		},
		{
			name:               "nok invalid method",
			jwt:                true,
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

			if tt.jwt {
				tokenJWT, _ := middlewares.JWTMiddleware{}.CreateToken(1)
				req.Header.Set("Authorization", "Bearer "+tokenJWT)
			}

			uc := NewUserController(tt.fields.Contract)
			if assert.NoError(t, uc.GetUsers(ctx)) || assert.Error(t, uc.GetUsers(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

func TestUserController_GetUserByID(t *testing.T) {

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
			method:             http.MethodGet,
			id:                 "2",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "nok without jwt GET",
			jwt:                false,
			method:             http.MethodGet,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            false,
		},
		{
			name:               "nok without jwt POST",
			jwt:                false,
			method:             http.MethodPost,
			id:                 "1",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            false,
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
		{
			name:               "nok invalid param :id",
			jwt:                true,
			method:             http.MethodGet,
			id:                 "id-7",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "nok data by id not found",
			jwt:                true,
			method:             http.MethodGet,
			id:                 "999",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/v1/books/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.id)
			tt.args = args{ctx: ctx}

			if tt.jwt {
				tokenJWT, _ := middlewares.JWTMiddleware{}.CreateToken(1)
				req.Header.Set("Authorization", "Bearer "+tokenJWT)
			}

			uc := NewUserController(tt.fields.Contract)
			if assert.NoError(t, uc.GetUserByID(ctx)) || assert.Error(t, uc.GetUserByID(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

func TestUserController_Create(t *testing.T) {

	/**
	 * Initiate config and echo framework
	 */

	config.InitDB()
	contract := config.Contract{DB: config.DB}
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

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
		requestBody        string
		fields             fields
		args               args
		statusCodeExpected int
		wantErr            bool
	}{
		{
			name:               "ok success",
			method:             http.MethodPost,
			requestBody:        `{"name":"Faizal nur rozi","email":"user@example.com","password":"faizal","gender":"MALE","nik":"14516347383","birth_date":"1995-12-30T13:35:32+07:00","married_status":true,"year_of_join":2014}`,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusCreated,
			wantErr:            false,
		},
		{
			name:               "nok wrong email format",
			method:             http.MethodPost,
			requestBody:        `{"name":"Faizal nur rozi","email":"user-example.com","password":"faizal","gender":"MALE","nik":"14516347383","birth_date":"1995-12-30T13:35:32+07:00","married_status":true,"year_of_join":2014}`,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            false,
		},
		{
			name:               "nok password is empty",
			method:             http.MethodPost,
			requestBody:        `{"name":"Faizal nur rozi","email":"user@example.com","gender":"MALE","nik":"14516347383","birth_date":"1995-12-30T13:35:32+07:00","married_status":true,"year_of_join":2014}`,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            false,
		},
		{
			name:               "nok without request body",
			method:             http.MethodPost,
			requestBody:        ``,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            false,
		},
		{
			name:               "nok invalid method",
			method:             http.MethodGet,
			requestBody:        ``,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusMethodNotAllowed,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", strings.NewReader(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/v1/books")
			tt.args = args{ctx: ctx}

			uc := NewUserController(tt.fields.Contract)
			if assert.NoError(t, uc.Create(ctx)) || assert.Error(t, uc.Create(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

func TestUserController_Update(t *testing.T) {

	/**
	 * Initiate config and echo framework
	 */

	config.InitDB()
	contract := config.Contract{DB: config.DB}
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

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
		requestBody        string
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
			requestBody:        `{"name":"Faizal nur rozi","email":"user@example.com","password":"user","gender":"MALE","nik":"14516347383","birth_date":"1995-12-30T13:35:32+07:00","married_status":true,"year_of_join":2014}`,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "nok wrong email format",
			jwt:                true,
			method:             http.MethodPut,
			id:                 "1",
			requestBody:        `{"name":"Faizal nur rozi","email":"user-example.com","password":"user","gender":"MALE","nik":"14516347383","birth_date":"1995-12-30T13:35:32+07:00","married_status":true,"year_of_join":2014}`,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            false,
		},
		{
			name:               "nok password is empty",
			jwt:                true,
			method:             http.MethodPut,
			id:                 "1",
			requestBody:        `{"name":"Faizal nur rozi","email":"user@example.com","gender":"MALE","nik":"14516347383","birth_date":"1995-12-30T13:35:32+07:00","married_status":true,"year_of_join":2014}`,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            false,
		},
		{
			name:               "nok without request body",
			jwt:                true,
			method:             http.MethodPut,
			id:                 "1",
			requestBody:        ``,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            false,
		},
		{
			name:               "nok invalid param :id",
			jwt:                true,
			method:             http.MethodPut,
			id:                 "id-2",
			requestBody:        `{"name":"Faizal nur rozi","email":"user@example.com","password":"user","gender":"MALE","nik":"14516347383","birth_date":"1995-12-30T13:35:32+07:00","married_status":true,"year_of_join":2014}`,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            false,
		},
		{
			name:               "nok invalid method",
			jwt:                true,
			method:             http.MethodGet,
			id:                 "1",
			requestBody:        ``,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusMethodNotAllowed,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", strings.NewReader(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/v1/books:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.id)
			tt.args = args{ctx: ctx}

			if tt.jwt {
				tokenJWT, _ := middlewares.JWTMiddleware{}.CreateToken(1)
				req.Header.Set("Authorization", "Bearer "+tokenJWT)
			}

			uc := NewUserController(tt.fields.Contract)
			if assert.NoError(t, uc.Update(ctx)) || assert.Error(t, uc.Update(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

func TestUserController_Delete(t *testing.T) {

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
			name:               "nok invalid param :id",
			jwt:                true,
			method:             http.MethodDelete,
			id:                 "id-2",
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            false,
		},
		{
			name:               "nok invalid method",
			jwt:                true,
			method:             http.MethodGet,
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
			ctx.SetPath("/v1/books:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.id)
			tt.args = args{ctx: ctx}

			if tt.jwt {
				tokenJWT, _ := middlewares.JWTMiddleware{}.CreateToken(1)
				req.Header.Set("Authorization", "Bearer "+tokenJWT)
			}

			uc := NewUserController(tt.fields.Contract)
			if assert.NoError(t, uc.Delete(ctx)) || assert.Error(t, uc.Delete(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

func TestUserController_Login(t *testing.T) {

	/**
	 * Initiate config and echo framework
	 */

	config.InitDB()
	contract := config.Contract{DB: config.DB}
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

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
		requestBody        string
		fields             fields
		args               args
		statusCodeExpected int
		wantErr            bool
	}{
		{
			name:               "ok success",
			method:             http.MethodPost,
			requestBody:        `{"email":"user2@example.com","password":"user2"}`,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "nok invalid email or password",
			method:             http.MethodPost,
			requestBody:        `{"email":"user1@example.com","password":"user1"}`,
			fields:             fields{Contract: contract},
			statusCodeExpected: http.StatusUnauthorized,
			wantErr:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", strings.NewReader(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/v1/books")
			tt.args = args{ctx: ctx}

			uc := NewUserController(tt.fields.Contract)
			if assert.NoError(t, uc.Login(ctx)) || assert.Error(t, uc.Login(ctx)) {
				assert.Equal(t, tt.statusCodeExpected, rec.Code)
			}
		})
	}
}

package handlers

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"GoMS/models"
	"GoMS/stores"
	"GoMS/stores/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestPing(t *testing.T) {

	t.Run("create", func(t *testing.T) {

		samplePing := models.Ping{ID: 1, CreatedBy: "127.0.0.1:1234", CreatedAt: time.Time{}}
		testCases := []struct {
			name               string
			pingStoreRes       *models.Ping
			pingStoreErr       error
			expectedStatusCode int
			expectedBody       string
		}{
			{"success", &samplePing, nil, http.StatusCreated, "{\"id\":1,\"createdBy\":\"127.0.0.1:1234\",\"createdAt\":\"0001-01-01T00:00:00Z\"}"},
			{"db failure", nil, errors.New("sample ping error"), http.StatusInternalServerError, "sample ping error\n"},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				mockPingStore := mocks.NewMockPingStore(ctrl)
				storeMap := stores.StoreMap{Ping: mockPingStore}
				handler := NewPingHandler(storeMap)
				recorder := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/", nil)
				req.RemoteAddr = samplePing.CreatedBy

				mockPingStore.EXPECT().Create(req.RemoteAddr).Return(testCase.pingStoreRes, testCase.pingStoreErr).Times(1)

				handler.Create(recorder, req)

				res := recorder.Result()
				defer req.Body.Close()

				if res.StatusCode != testCase.expectedStatusCode {
					t.Errorf("expected status code %v to equal %v", res.StatusCode, testCase.expectedStatusCode)
				}

				body, _ := ioutil.ReadAll(res.Body)

				if string(body) != testCase.expectedBody {
					t.Errorf("expected body %s to equal %s", string(body), testCase.expectedBody)
				}

			})
		}

	})

	t.Run("getAll", func(t *testing.T) {

		samplePing := models.Ping{ID: 1, CreatedBy: "127.0.0.1:1234", CreatedAt: time.Time{}}
		testCases := []struct {
			name               string
			pingStoreRes       []models.Ping
			pingStoreErr       error
			expectedStatusCode int
			expectedBody       string
		}{
			{"success", []models.Ping{samplePing}, nil, http.StatusOK, "[{\"id\":1,\"createdBy\":\"127.0.0.1:1234\",\"createdAt\":\"0001-01-01T00:00:00Z\"}]"},
			{"db failure", []models.Ping{}, errors.New("sample ping error"), http.StatusInternalServerError, "sample ping error\n"},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				mockPingStore := mocks.NewMockPingStore(ctrl)
				storeMap := stores.StoreMap{Ping: mockPingStore}
				handler := NewPingHandler(storeMap)
				recorder := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				req.RemoteAddr = samplePing.CreatedBy

				mockPingStore.EXPECT().GetAll().Return(testCase.pingStoreRes, testCase.pingStoreErr).Times(1)

				handler.GetAll(recorder, req)

				res := recorder.Result()
				defer req.Body.Close()

				if res.StatusCode != testCase.expectedStatusCode {
					t.Errorf("expected status code %v to equal %v", res.StatusCode, testCase.expectedStatusCode)
				}

				body, _ := ioutil.ReadAll(res.Body)

				if string(body) != testCase.expectedBody {
					t.Errorf("expected body %s to equal %s", string(body), testCase.expectedBody)
				}

			})
		}

	})

	t.Run("getById", func(t *testing.T) {

		samplePing := models.Ping{ID: 1, CreatedBy: "127.0.0.1:1234", CreatedAt: time.Time{}}

		t.Run("success", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPingStore := mocks.NewMockPingStore(ctrl)
			storeMap := stores.StoreMap{Ping: mockPingStore}
			handler := NewPingHandler(storeMap)
			recorder := httptest.NewRecorder()
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("pingID", "1")
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

			mockPingStore.EXPECT().GetById(1).Return(&samplePing, nil).Times(1)

			handler.GetByID(recorder, req)

			res := recorder.Result()
			defer req.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status code %v to equal 200", res.StatusCode)
			}

			body, _ := ioutil.ReadAll(res.Body)

			expectedBody := "{\"id\":1,\"createdBy\":\"127.0.0.1:1234\",\"createdAt\":\"0001-01-01T00:00:00Z\"}"

			if string(body) != expectedBody {
				t.Errorf("expected body %s to equal %s", string(body), expectedBody)
			}

		})

		t.Run("invalid url id (non-numerical)", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPingStore := mocks.NewMockPingStore(ctrl)
			storeMap := stores.StoreMap{Ping: mockPingStore}
			handler := NewPingHandler(storeMap)
			recorder := httptest.NewRecorder()
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("pingID", "abc")
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

			handler.GetByID(recorder, req)

			res := recorder.Result()
			defer req.Body.Close()

			if res.StatusCode != http.StatusBadRequest {
				t.Errorf("expected status code %v to equal 400", res.StatusCode)
			}

			body, _ := ioutil.ReadAll(res.Body)

			expectedBody := "Ping not found\n"

			if string(body) != expectedBody {
				t.Errorf("expected body %s to equal %s", string(body), expectedBody)
			}

		})

		t.Run("invalid url id (missing)", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPingStore := mocks.NewMockPingStore(ctrl)
			storeMap := stores.StoreMap{Ping: mockPingStore}
			handler := NewPingHandler(storeMap)
			recorder := httptest.NewRecorder()
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("pingID", "")
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

			handler.GetByID(recorder, req)

			res := recorder.Result()
			defer req.Body.Close()

			if res.StatusCode != http.StatusBadRequest {
				t.Errorf("expected status code %v to equal 400", res.StatusCode)
			}

			body, _ := ioutil.ReadAll(res.Body)

			expectedBody := "Ping not found\n"

			if string(body) != expectedBody {
				t.Errorf("expected body %s to equal %s", string(body), expectedBody)
			}

		})

		t.Run("valid id not found", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPingStore := mocks.NewMockPingStore(ctrl)
			storeMap := stores.StoreMap{Ping: mockPingStore}
			handler := NewPingHandler(storeMap)
			recorder := httptest.NewRecorder()
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("pingID", "5")
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

			mockPingStore.EXPECT().GetById(5).Return(nil, gorm.ErrRecordNotFound).Times(1)

			handler.GetByID(recorder, req)

			res := recorder.Result()
			defer req.Body.Close()

			if res.StatusCode != http.StatusNotFound {
				t.Errorf("expected status code %v to equal 404", res.StatusCode)
			}

			body, _ := ioutil.ReadAll(res.Body)

			expectedBody := "Ping not found\n"

			if string(body) != expectedBody {
				t.Errorf("expected body %s to equal %s", string(body), expectedBody)
			}

		})

		t.Run("db error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPingStore := mocks.NewMockPingStore(ctrl)
			storeMap := stores.StoreMap{Ping: mockPingStore}
			handler := NewPingHandler(storeMap)
			recorder := httptest.NewRecorder()
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("pingID", "5")
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

			mockPingStore.EXPECT().GetById(5).Return(nil, errors.New("sample error")).Times(1)

			handler.GetByID(recorder, req)

			res := recorder.Result()
			defer req.Body.Close()

			if res.StatusCode != http.StatusInternalServerError {
				t.Errorf("expected status code %v to equal 500", res.StatusCode)
			}

			body, _ := ioutil.ReadAll(res.Body)

			expectedBody := "sample error\n"

			if string(body) != expectedBody {
				t.Errorf("expected body %s to equal %s", string(body), expectedBody)
			}

		})

	})

}

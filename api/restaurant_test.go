package api

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/lakemanr/otamaq/db/mock"
	db "github.com/lakemanr/otamaq/db/sqlc"
	"github.com/lakemanr/otamaq/util"
	"github.com/stretchr/testify/require"
)

func TestCreateRestaurantApi(t *testing.T) {

	restaurant := createRandomRestaurant()

	arg := db.CreateRestaurantParams{
		OwnerID: restaurant.OwnerID,
		Name:    restaurant.Name,
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mock.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"owner_id": restaurant.OwnerID,
				"name":     restaurant.Name,
			},
			buildStubs: func(store *mock.MockStore) {
				store.EXPECT().CreateRestaurant(gomock.Any(), gomock.Eq(arg)).Times(1).Return(restaurant, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchRestaurant(t, recorder.Body, restaurant)
			},
		},
		{
			name: "InvalidName",
			body: gin.H{
				"owner_id": restaurant.OwnerID,
				"name":     "rest$$",
			},
			buildStubs: func(store *mock.MockStore) {
				store.EXPECT().CreateRestaurant(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"owner_id": restaurant.OwnerID,
				"name":     restaurant.Name,
			},
			buildStubs: func(store *mock.MockStore) {
				store.EXPECT().CreateRestaurant(gomock.Any(), gomock.Eq(arg)).Times(1).Return(db.Restaurant{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock store
	store := mock.NewMockStore(ctrl)

	// Create a server with the mock store
	server := NewServer(store)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.body)
			require.NoError(t, err)
			tc.buildStubs(store)
			request := httptest.NewRequest("POST", "/restaurants", bytes.NewReader(body))
			recorder := httptest.NewRecorder()
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func createRandomRestaurant() db.Restaurant {

	return db.Restaurant{
		ID:        util.RandomID(),
		OwnerID:   util.RandomID(),
		Name:      util.RandomRestaurantName(),
		CreatedAt: time.Now(),
	}
}

func requireBodyMatchRestaurant(t *testing.T, responseBody *bytes.Buffer, expectedRestaurant db.Restaurant) {
	var responseRestaurant db.Restaurant
	err := json.Unmarshal(responseBody.Bytes(), &responseRestaurant)
	require.NoError(t, err)
	require.Equal(t, responseRestaurant.ID, expectedRestaurant.ID)
	require.Equal(t, responseRestaurant.OwnerID, expectedRestaurant.OwnerID)
	require.Equal(t, responseRestaurant.Name, expectedRestaurant.Name)
	require.True(t, responseRestaurant.CreatedAt.Equal(expectedRestaurant.CreatedAt))
}

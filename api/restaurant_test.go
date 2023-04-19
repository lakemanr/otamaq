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

	testCases := []struct {
		name          string
		body          gin.H
		biuildStubs   func(store *mock.MockStore)
		checkResponce func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name": restaurant.Name,
			},
			biuildStubs: func(store *mock.MockStore) {
				store.EXPECT().CreateRestaurant(gomock.Any(), gomock.Eq(restaurant.Name)).Times(1).Return(restaurant, nil)
			},
			checkResponce: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchRestaurant(t, recorder.Body, restaurant)
			},
		},
		{
			name: "InvalidName",
			body: gin.H{
				"name": "",
			},
			biuildStubs: func(store *mock.MockStore) {
				store.EXPECT().CreateRestaurant(gomock.Any(), gomock.Any().String()).Times(0)
			},
			checkResponce: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalvError",
			body: gin.H{
				"name": restaurant.Name,
			},
			biuildStubs: func(store *mock.MockStore) {
				store.EXPECT().CreateRestaurant(gomock.Any(), gomock.Eq(restaurant.Name)).Times(1).Return(db.Restaurant{}, sql.ErrConnDone)
			},
			checkResponce: func(recorder *httptest.ResponseRecorder) {
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
			tc.biuildStubs(store)
			request := httptest.NewRequest("POST", "/restaurants", bytes.NewReader(body))
			recorder := httptest.NewRecorder()
			server.router.ServeHTTP(recorder, request)
			tc.checkResponce(recorder)
		})
	}
}

func createRandomRestaurant() db.Restaurant {
	return db.Restaurant{
		ID:        util.RandomID(),
		Name:      util.RandomRestaurantName(),
		CreatedAt: time.Now().Truncate(time.Second),
	}
}

func requireBodyMatchRestaurant(t *testing.T, responceBody *bytes.Buffer, expectedRestaurant db.Restaurant) {
	var responceRestaurant db.Restaurant
	err := json.Unmarshal(responceBody.Bytes(), &responceRestaurant)
	require.NoError(t, err)
	require.Equal(t, responceRestaurant, expectedRestaurant)
}

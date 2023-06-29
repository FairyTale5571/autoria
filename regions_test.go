package autoria

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_service_GetCities(t *testing.T) {
	type fields struct {
		apikey     string
		client     *http.Client
		mutex      sync.Mutex
		maxRetries int
		debug      bool
	}
	type args struct {
		stateID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    States
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetCities Vinytsa state",
			fields: fields{
				apikey:     os.Getenv("AUTORIA_API_KEY"),
				client:     http.DefaultClient,
				mutex:      sync.Mutex{},
				maxRetries: 0,
			},
			args: args{
				stateID: 1,
			},
			want: States{
				{Name: "Вінниця", Value: 1},
				{Name: "Жмеринка", Value: 27},
				{Name: "Козятин", Value: 30},
				{Name: "Крижопіль", Value: 31},
				{Name: "Липовець", Value: 32},
				{Name: "Літин", Value: 33},
				{Name: "Могилів-Подільський", Value: 34},
				{Name: "Муровані Курилівці", Value: 35},
				{Name: "Немирів", Value: 36},
				{Name: "Оратів", Value: 37},
				{Name: "Піщанка", Value: 38},
				{Name: "Погребище", Value: 39},
				{Name: "Теплик", Value: 40},
				{Name: "Тиврів", Value: 41},
				{Name: "Томашпіль", Value: 42},
				{Name: "Тростянець", Value: 43},
				{Name: "Тульчин", Value: 44},
				{Name: "Хмільник", Value: 45},
				{Name: "Чернівці", Value: 46},
				{Name: "Чечельник", Value: 47},
				{Name: "Шаргород", Value: 48},
				{Name: "Ямпіль", Value: 49},
				{Name: "Бар", Value: 597},
				{Name: "Бершадь", Value: 599},
				{Name: "Гайсин", Value: 602},
				{Name: "Іллінці", Value: 603},
				{Name: "Калинівка", Value: 604},
				{Name: "Гнівань", Value: 609},
				{Name: "Ладижин", Value: 644},
				{Name: "Якушинці", Value: 12680},
			},
			wantErr: assert.NoError,
		},
		{
			name: "GetCities error",
			fields: fields{
				apikey:     "123",
				client:     http.DefaultClient,
				mutex:      sync.Mutex{},
				maxRetries: 0,
			},
			args: args{
				stateID: 0,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey:     tt.fields.apikey,
				client:     tt.fields.client,
				mutex:      tt.fields.mutex,
				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetCities(tt.args.stateID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetCities(%v)", tt.args.stateID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetCities(%v)", tt.args.stateID)
		})
	}
}

func Test_service_GetStates(t *testing.T) {
	type fields struct {
		apikey     string
		client     *http.Client
		mutex      sync.Mutex
		maxRetries int
		debug      bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    States
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetStates",
			fields: fields{
				apikey:     os.Getenv("AUTORIA_API_KEY"),
				client:     http.DefaultClient,
				mutex:      sync.Mutex{},
				maxRetries: 0,
			},
			want: States{
				{Name: "Київська", Value: 10},
				{Name: "Вінницька", Value: 1},
				{Name: "Волинська", Value: 18},
				{Name: "Дніпропетровська", Value: 11},
				{Name: "Донецька", Value: 13},
				{Name: "Житомирська", Value: 2},
				{Name: "Закарпатська", Value: 22},
				{Name: "Запорізька", Value: 14},
				{Name: "Івано-Франківська", Value: 15},
				{Name: "Кіровоградська", Value: 16},
				{Name: "Луганська", Value: 17},
				{Name: "Львівська", Value: 5},
				{Name: "Миколаївська", Value: 19},
				{Name: "Одеська", Value: 12},
				{Name: "Полтавська", Value: 20},
				{Name: "Рівненська", Value: 9},
				{Name: "Сумська", Value: 8},
				{Name: "Тернопільська", Value: 3},
				{Name: "Харківська", Value: 7},
				{Name: "Херсонська", Value: 23},
				{Name: "Хмельницька", Value: 4},
				{Name: "Черкаська", Value: 24},
				{Name: "Чернівецька", Value: 25},
				{Name: "Чернігівська", Value: 6},
			},
			wantErr: assert.NoError,
		},
		{
			name: "GetStates error",
			fields: fields{
				apikey:     "123",
				client:     http.DefaultClient,
				mutex:      sync.Mutex{},
				maxRetries: 0,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey:     tt.fields.apikey,
				client:     tt.fields.client,
				mutex:      tt.fields.mutex,
				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetStates()
			if !tt.wantErr(t, err, fmt.Sprintf("GetStates()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetStates()")
		})
	}
}

func Test_service_GetCountries(t *testing.T) {
	type fields struct {
		apikey     string
		client     *http.Client
		mutex      sync.Mutex
		maxRetries int
		debug      bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    CountryTypes
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetCountries",
			fields: fields{
				apikey:     os.Getenv("AUTORIA_API_KEY"),
				client:     http.DefaultClient,
				mutex:      sync.Mutex{},
				maxRetries: 0,
			},
			want:    CountryTypes{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey:     tt.fields.apikey,
				client:     tt.fields.client,
				mutex:      tt.fields.mutex,
				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetCountries()
			if !tt.wantErr(t, err, fmt.Sprintf("GetCountries()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetCountries()")
		})
	}
}

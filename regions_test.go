package autoria

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_service_GetCities(t *testing.T) {
	type fields struct {
		client *http.Client
		apikey string
		debug  bool
	}
	type args struct {
		stateID int
	}
	tests := []struct {
		wantErr assert.ErrorAssertionFunc
		name    string
		fields  fields
		want    States
		args    args
	}{
		{
			name: "GetCities Vinytsa state",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: http.DefaultClient,
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
				apikey: "123",
				client: http.DefaultClient,
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
			s := &Service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,
				debug:  tt.fields.debug,
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
		client *http.Client
		apikey string

		debug bool
	}
	tests := []struct {
		wantErr assert.ErrorAssertionFunc
		name    string
		fields  fields
		want    States
	}{
		{
			name: "GetStates",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: http.DefaultClient,
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
				apikey: "123",
				client: http.DefaultClient,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				debug: tt.fields.debug,
			}
			got, err := s.GetStates()
			if !tt.wantErr(t, err, "GetStates()") {
				return
			}
			assert.Equalf(t, tt.want, got, "GetStates()")
		})
	}
}

func Test_service_GetCountries(t *testing.T) {
	type fields struct {
		client *http.Client
		apikey string

		debug bool
	}
	tests := []struct {
		wantErr assert.ErrorAssertionFunc
		name    string
		fields  fields
		want    CountryTypes
	}{
		{
			name: "GetCountries",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: http.DefaultClient,
			},
			want:    CountryTypes{Base{Name: "Австрія", Value: 40}, Base{Name: "Англія", Value: 826}, Base{Name: "Аргентина", Value: 32}, Base{Name: "Бельгія", Value: 56}, Base{Name: "Білорусь", Value: 112}, Base{Name: "Болгарія", Value: 100}, Base{Name: "Бразилія", Value: 76}, Base{Name: "Грузія", Value: 900}, Base{Name: "Данія", Value: 208}, Base{Name: "Естонія", Value: 233}, Base{Name: "Індія", Value: 356}, Base{Name: "Іран", Value: 364}, Base{Name: "Ірландія", Value: 901}, Base{Name: "Ісландія", Value: 903}, Base{Name: "Іспанія", Value: 724}, Base{Name: "Італія", Value: 380}, Base{Name: "Казахстан", Value: 398}, Base{Name: "Канада", Value: 124}, Base{Name: "Китай", Value: 158}, Base{Name: "Корея", Value: 408}, Base{Name: "Латвія", Value: 428}, Base{Name: "Литва", Value: 440}, Base{Name: "Люксембург", Value: 442}, Base{Name: "Малайзія", Value: 458}, Base{Name: "Молдова", Value: 498}, Base{Name: "Нідерланди", Value: 528}, Base{Name: "Німеччина", Value: 276}, Base{Name: "Норвегія", Value: 578}, Base{Name: "ОАЕ", Value: 902}, Base{Name: "Польша", Value: 616}, Base{Name: "Португалiя", Value: 620}, Base{Name: "Росія", Value: 643}, Base{Name: "Румунія", Value: 642}, Base{Name: "Сербія", Value: 688}, Base{Name: "Словаччина", Value: 703}, Base{Name: "Словенія", Value: 705}, Base{Name: "США", Value: 840}, Base{Name: "Туреччина", Value: 792}, Base{Name: "Угорщина", Value: 348}, Base{Name: "Узбекистан", Value: 860}, Base{Name: "Україна", Value: 804}, Base{Name: "Фінляндія", Value: 246}, Base{Name: "Франція", Value: 250}, Base{Name: "Чехія", Value: 203}, Base{Name: "Швейцарія", Value: 756}, Base{Name: "Швеція", Value: 752}, Base{Name: "Японія", Value: 392}},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				debug: tt.fields.debug,
			}
			got, err := s.GetCountries()
			if !tt.wantErr(t, err, "GetCountries()") {
				return
			}
			assert.Equalf(t, tt.want, got, "GetCountries()")
		})
	}
}

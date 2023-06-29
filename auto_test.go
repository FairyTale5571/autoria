package autoria

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_service_GetCategories(t *testing.T) {
	type fields struct {
		apikey     string
		client     *http.Client
		maxRetries int
	}
	tests := []struct {
		name    string
		fields  fields
		want    Categories
		wantErr bool
	}{
		{
			name: "Get Categories successfully",
			fields: fields{
				apikey:     os.Getenv("AUTORIA_API_KEY"),
				client:     &http.Client{},
				maxRetries: 0,
			},
			wantErr: false,
			want: Categories{
				{Name: "Легкові", Value: 1},
				{Name: "Мото", Value: 2},
				{Name: "Водний транспорт", Value: 3},
				{Name: "Спецтехніка", Value: 4},
				{Name: "Причепи", Value: 5},
				{Name: "Вантажівки", Value: 6},
				{Name: "Автобуси", Value: 7},
				{Name: "Автобудинки", Value: 8},
				{Name: "Повітряний транспорт", Value: 9},
				{Name: "Сільгосптехніка", Value: 10},
			},
		},
		{
			name: "Get Categories successfully",
			fields: fields{
				apikey:     "123",
				client:     &http.Client{},
				maxRetries: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey:     tt.fields.apikey,
				client:     tt.fields.client,
				maxRetries: tt.fields.maxRetries,
			}
			got, err := s.GetCategories()
			if tt.wantErr {
				if err == nil {
					t.Errorf("service.GetCategories() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			require.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_service_GetBodyStyles(t *testing.T) {
	type fields struct {
		apikey     string
		client     *http.Client
		maxRetries int
	}
	tests := []struct {
		name     string
		parentID int
		fields   fields
		want     []BaseWithParentID
		wantErr  bool
	}{
		{
			name: "Get Body styles successfully",
			fields: fields{
				apikey:     os.Getenv("AUTORIA_API_KEY"),
				client:     &http.Client{},
				maxRetries: 0,
			},
			want: []BaseWithParentID{
				{Base: Base{Name: "Седан", Value: 3}, ParentID: 0},
				{Base: Base{Name: "Позашляховик / Кросовер", Value: 5}, ParentID: 0},
				{Base: Base{Name: "Мінівен", Value: 8}, ParentID: 0},
				{Base: Base{Name: "Мікровен", Value: 449}, ParentID: 0},
				{Base: Base{Name: "Хетчбек", Value: 4}, ParentID: 0},
				{Base: Base{Name: "Універсал", Value: 2}, ParentID: 0},
				{Base: Base{Name: "Купе", Value: 6}, ParentID: 0},
				{Base: Base{Name: "Кабріолет", Value: 7}, ParentID: 0},
				{Base: Base{Name: "Пікап", Value: 9}, ParentID: 0},
				{Base: Base{Name: "Ліфтбек", Value: 307}, ParentID: 0},
				{Base: Base{Name: "Фастбек", Value: 448}, ParentID: 0},
				{Base: Base{Name: "Лімузин", Value: 252}, ParentID: 0},
				{Base: Base{Name: "Родстер", Value: 315}, ParentID: 0},
			},
			parentID: CategoryCars,
			wantErr:  false,
		},
		{
			name: "Unknown parentID",
			fields: fields{
				apikey:     os.Getenv("AUTORIA_API_KEY"),
				client:     &http.Client{},
				maxRetries: 0,
			},
			want:     []BaseWithParentID{},
			parentID: 9999999,
			wantErr:  false,
		},
		{
			name: "Get Body styles fail",
			fields: fields{
				apikey:     "123123123",
				client:     &http.Client{},
				maxRetries: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey:     tt.fields.apikey,
				client:     tt.fields.client,
				maxRetries: tt.fields.maxRetries,
			}
			got, err := s.GetBodyStyles(tt.parentID)
			if tt.wantErr {
				if err == nil {
					t.Errorf("service.GetCategories() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			require.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_service_GetBodyStylesGroups(t *testing.T) {
	type fields struct {
		apikey string
		client *http.Client

		maxRetries int
		debug      bool
	}
	tests := []struct {
		name     string
		parentID int
		fields   fields
		want     [][]BaseWithParentID
		wantErr  bool
	}{
		{
			name: "Get Body styles successfully for parentID 1",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 0,
				debug:      true,
			},
			want: [][]BaseWithParentID{
				{BaseWithParentID{Base: Base{Name: "Седан", Value: 3}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Позашляховик / Кросовер", Value: 5}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Мінівен", Value: 8}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Мікровен", Value: 449}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Хетчбек", Value: 4}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Універсал", Value: 2}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Купе", Value: 6}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Кабріолет", Value: 7}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Пікап", Value: 9}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Ліфтбек", Value: 307}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Фастбек", Value: 448}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Лімузин", Value: 252}, ParentID: 0}},
				{BaseWithParentID{Base: Base{Name: "Родстер", Value: 315}, ParentID: 0}},
			},
			parentID: CategoryCars,
			wantErr:  false,
		},
		{
			name: "Get Body styles successfully for parentID 2",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 0,
				debug:      true,
			},
			want: [][]BaseWithParentID{
				{
					BaseWithParentID{Base: Base{Name: "Інше", Value: 56}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Вантажні моторолери, мотоцикли, скутери, мопеди", Value: 429}, ParentID: 0},
				},
				{
					BaseWithParentID{Base: Base{Name: "Квадроцикли", Value: 35}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Всюдихід-амфібія", Value: 43}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Гольф-кар", Value: 44}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Картинг", Value: 45}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Квадроцикл дитячий", Value: 36}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Квадроцикл спортивний", Value: 39}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Квадроцикл утилітарний", Value: 41}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотовсюдиход", Value: 42}, ParentID: 0},
				},
				{
					BaseWithParentID{Base: Base{Name: "Міні мотоцикли", Value: 31}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Міні спорт", Value: 32}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Пітбайк", Value: 33}, ParentID: 0},
				},
				{
					BaseWithParentID{Base: Base{Name: "Мопеди", Value: 58}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Максі-скутер", Value: 12}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мокик", Value: 427}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Скутер / Мотороллер", Value: 11}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Скутер для інвалідів", Value: 426}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Скутер ретро", Value: 425}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Скутери з кабіною", Value: 430}, ParentID: 0},
				},
				{
					BaseWithParentID{Base: Base{Name: "Мотоцикли", Value: 13}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Боббер", Value: 421}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Кафе рейсер", Value: 423}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Багатоцільовий (All-round)", Value: 25}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Без обтікачів (Naked bike)", Value: 15}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл з коляскою", Value: 29}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Кастом", Value: 30}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Классік", Value: 14}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Кросс", Value: 19}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Круізер", Value: 24}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Позашляховий (Enduro)", Value: 21}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Спорт-туризм", Value: 17}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Супермото (Motard)", Value: 22}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Тріал", Value: 20}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Туризм", Value: 16}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотоцикл Чоппер", Value: 23}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Скремблер", Value: 422}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Спортбайк", Value: 18}, ParentID: 0},
				},
				{
					BaseWithParentID{Base: Base{Name: "Снігохід", Value: 46}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Гірські снігоходи", Value: 434}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Міні-розбірні снігоходи", Value: 432}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Мотобуксир", Value: 424}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Снігомопеди та снігоскутери", Value: 436}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Снігоходи для пполювання та рибалки", Value: 433}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Спортивні снігоходи", Value: 435}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Утилітарні снігоходи", Value: 431}, ParentID: 0},
				},
				{BaseWithParentID{Base: Base{Name: "Трайк", Value: 57}, ParentID: 0}},
				{
					BaseWithParentID{Base: Base{Name: "Трицикл", Value: 34}, ParentID: 0},
					BaseWithParentID{Base: Base{Name: "Вантажні трицикли", Value: 428}, ParentID: 0},
				},
			},
			parentID: CategoryMoto,
			wantErr:  false,
		},
		{
			name: "Unknown parentID",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 0,
			},
			want:     [][]BaseWithParentID(nil),
			parentID: 9999999,
			wantErr:  false,
		},
		{
			name: "Get Body groups styles fail",
			fields: fields{
				apikey: "123123123",
				client: &http.Client{},

				maxRetries: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetBodyStylesWithGroups(tt.parentID)
			if tt.wantErr {
				if err == nil {
					t.Errorf("service.GetCategories() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			require.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_service_GetMarksByCategory(t *testing.T) {
	type fields struct {
		apikey     string
		client     *http.Client
		maxRetries int
		debug      bool
	}
	type args struct {
		categoryID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Marks
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get marks by category cars",
			fields: fields{
				apikey:     os.Getenv("AUTORIA_API_KEY"),
				client:     &http.Client{},
				maxRetries: 0,
			},
			args: args{
				categoryID: CategoryCars,
			},
			want:    Marks{Base{Name: "Abarth", Value: 5166}, Base{Name: "Acura", Value: 98}, Base{Name: "Adler", Value: 2396}, Base{Name: "Aero", Value: 5165}, Base{Name: "AION", Value: 55172}, Base{Name: "Aixam", Value: 2}, Base{Name: "Alfa Romeo", Value: 3}, Base{Name: "Alpine", Value: 100}, Base{Name: "Altamarea", Value: 3988}, Base{Name: "AMC", Value: 5821}, Base{Name: "Anaig", Value: 5769}, Base{Name: "Armstrong Siddeley", Value: 5276}, Base{Name: "Aro", Value: 101}, Base{Name: "Artega", Value: 3105}, Base{Name: "Asia", Value: 4}, Base{Name: "AsiaStar", Value: 5793}, Base{Name: "Aston Martin", Value: 5}, Base{Name: "ATS Corsa", Value: 6160}, Base{Name: "Audi", Value: 6}, Base{Name: "Austin", Value: 7}, Base{Name: "Austin-Healey", Value: 4355}, Base{Name: "Autobianchi", Value: 102}, Base{Name: "Avatr", Value: 55330}, Base{Name: "BAIC", Value: 55051}, Base{Name: "Baojun", Value: 6228}, Base{Name: "Baoya", Value: 5245}, Base{Name: "Barkas (Баркас)", Value: 103}, Base{Name: "Baw", Value: 1540}, Base{Name: "Beijing", Value: 105}, Base{Name: "Bentley", Value: 8}, Base{Name: "Bertone", Value: 106}, Base{Name: "Bio Auto", Value: 3127}, Base{Name: "Blonell", Value: 108}, Base{Name: "BlueCar", Value: 55266}, Base{Name: "BMW", Value: 9}, Base{Name: "BMW-Alpina", Value: 99}, Base{Name: "Bollinger", Value: 6310}, Base{Name: "Borgward", Value: 5240}, Base{Name: "Brilliance", Value: 329}, Base{Name: "Bristol", Value: 10}, Base{Name: "Bugatti", Value: 109}, Base{Name: "Buick", Value: 110}, Base{Name: "BYD", Value: 386}, Base{Name: "Byton", Value: 6260}, Base{Name: "Cadillac", Value: 11}, Base{Name: "Caterham", Value: 3221}, Base{Name: "CATL", Value: 55284}, Base{Name: "Chana", Value: 407}, Base{Name: "Changan", Value: 1580}, Base{Name: "Changfeng", Value: 55143}, Base{Name: "Changhe", Value: 1567}, Base{Name: "Chanje", Value: 55171}, Base{Name: "Chery", Value: 190}, Base{Name: "Chevrolet", Value: 13}, Base{Name: "Chrysler", Value: 14}, Base{Name: "Citroen", Value: 15}, Base{Name: "Cupra", Value: 1451}, Base{Name: "Dacia", Value: 17}, Base{Name: "Dadi", Value: 198}, Base{Name: "Daechang", Value: 55242}, Base{Name: "Daewoo", Value: 18}, Base{Name: "DAF", Value: 115}, Base{Name: "DAF / VDL", Value: 1441}, Base{Name: "Dagger", Value: 3717}, Base{Name: "Daihatsu", Value: 19}, Base{Name: "Daimler", Value: 20}, Base{Name: "Datsun", Value: 4206}, Base{Name: "De Lorean", Value: 3071}, Base{Name: "Denza", Value: 55216}, Base{Name: "Derways", Value: 4328}, Base{Name: "Detroit Electric", Value: 4883}, Base{Name: "DFSK", Value: 4077}, Base{Name: "DKW", Value: 2243}, Base{Name: "Dodge", Value: 118}, Base{Name: "Dongfeng", Value: 308}, Base{Name: "Dr. Motor", Value: 5051}, Base{Name: "DS", Value: 4495}, Base{Name: "Eagle", Value: 120}, Base{Name: "Elwinn", Value: 5469}, Base{Name: "Enovate", Value: 55225}, Base{Name: "Ernst Auwarter", Value: 1413}, Base{Name: "Estrima", Value: 3921}, Base{Name: "FAW", Value: 121}, Base{Name: "FAW Bestune", Value: 55208}, Base{Name: "Ferrari", Value: 22}, Base{Name: "Fiat", Value: 23}, Base{Name: "Fiat-Abarth", Value: 4459}, Base{Name: "Fisker", Value: 3444}, Base{Name: "Ford", Value: 24}, Base{Name: "Fornasari", Value: 3104}, Base{Name: "Forthing", Value: 55370}, Base{Name: "FSO", Value: 25}, Base{Name: "FUQI", Value: 197}, Base{Name: "Gac", Value: 4506}, Base{Name: "GAC-Honda", Value: 55252}, Base{Name: "Geely", Value: 185}, Base{Name: "Genesis", Value: 2604}, Base{Name: "Geo", Value: 2790}, Base{Name: "GMC", Value: 123}, Base{Name: "Golf Car", Value: 4316}, Base{Name: "Gonow", Value: 183}, Base{Name: "Great Wall", Value: 124}, Base{Name: "Groz", Value: 1575}, Base{Name: "Hafei", Value: 191}, Base{Name: "Haima", Value: 3674}, Base{Name: "Hanomag", Value: 1784}, Base{Name: "Hansa", Value: 2053}, Base{Name: "Hanteng", Value: 55308}, Base{Name: "Haval", Value: 5456}, Base{Name: "Hawtai", Value: 5791}, Base{Name: "Hindustan", Value: 3411}, Base{Name: "Honda", Value: 28}, Base{Name: "Hong Xing", Value: 5572}, Base{Name: "Hongqi", Value: 55205}, Base{Name: "Horch", Value: 5624}, Base{Name: "Hozon", Value: 55108}, Base{Name: "Hozon Neta", Value: 55173}, Base{Name: "Huabei", Value: 2595}, Base{Name: "Huanghai", Value: 388}, Base{Name: "Human Horizons", Value: 55303}, Base{Name: "Humber", Value: 3002}, Base{Name: "Hummer", Value: 127}, Base{Name: "Humvee", Value: 4663}, Base{Name: "Hyundai", Value: 29}, Base{Name: "Infiniti", Value: 128}, Base{Name: "Innocenti", Value: 4273}, Base{Name: "Iran Khodro", Value: 3821}, Base{Name: "Isuzu", Value: 30}, Base{Name: "ItalCar", Value: 3757}, Base{Name: "Iveco", Value: 175}, Base{Name: "JAC", Value: 317}, Base{Name: "JAC-Volkswagen", Value: 55217}, Base{Name: "Jaguar", Value: 31}, Base{Name: "JCB", Value: 1590}, Base{Name: "Jeep", Value: 32}, Base{Name: "Jetour", Value: 55069}, Base{Name: "Jiangnan", Value: 335}, Base{Name: "Jinbei", Value: 2231}, Base{Name: "Jinbei Minibusus", Value: 4549}, Base{Name: "JMC", Value: 3018}, Base{Name: "Jonway", Value: 1574}, Base{Name: "Karosa", Value: 412}, Base{Name: "Karry", Value: 55200}, Base{Name: "Keyton", Value: 55282}, Base{Name: "Kia", Value: 33}, Base{Name: "King Long", Value: 2676}, Base{Name: "KingWoo", Value: 4606}, Base{Name: "Kirkham", Value: 4843}, Base{Name: "Koenigsegg", Value: 2643}, Base{Name: "Konecranes", Value: 4000}, Base{Name: "Lamborghini", Value: 35}, Base{Name: "Lancia", Value: 36}, Base{Name: "Land Rover", Value: 37}, Base{Name: "Landwind", Value: 406}, Base{Name: "LDV", Value: 134}, Base{Name: "Leap", Value: 55215}, Base{Name: "Leap Motor", Value: 55269}, Base{Name: "Leopaard", Value: 55213}, Base{Name: "Lesheng", Value: 55230}, Base{Name: "Letin", Value: 55375}, Base{Name: "Lexus", Value: 38}, Base{Name: "Lichi", Value: 5339}, Base{Name: "Lifan", Value: 334}, Base{Name: "Lincoln", Value: 135}, Base{Name: "Lingpao", Value: 55341}, Base{Name: "Link Tour", Value: 55197}, Base{Name: "Lotus", Value: 41}, Base{Name: "LTI", Value: 136}, Base{Name: "Lucid", Value: 6317}, Base{Name: "Luxgen", Value: 4269}, Base{Name: "MAN", Value: 177}, Base{Name: "Marshell", Value: 4064}, Base{Name: "Maruti", Value: 139}, Base{Name: "Maserati", Value: 45}, Base{Name: "Maxus", Value: 55270}, Base{Name: "Maybach", Value: 46}, Base{Name: "Mazda", Value: 47}, Base{Name: "McLaren", Value: 3101}, Base{Name: "MEGA", Value: 1904}, Base{Name: "Mercedes-Benz", Value: 48}, Base{Name: "Mercury", Value: 144}, Base{Name: "Merkur", Value: 3948}, Base{Name: "MG", Value: 49}, Base{Name: "Microcar", Value: 6055}, Base{Name: "Miles", Value: 4528}, Base{Name: "MINI", Value: 147}, Base{Name: "Mitsubishi", Value: 52}, Base{Name: "Mitsuoka", Value: 402}, Base{Name: "Mobility Ventures", Value: 4069}, Base{Name: "Morgan", Value: 53}, Base{Name: "Morris", Value: 54}, Base{Name: "MPM Motors", Value: 5039}, Base{Name: "MYBRO", Value: 5138}, Base{Name: "NEO", Value: 55114}, Base{Name: "Neta", Value: 55214}, Base{Name: "NIO", Value: 55080}, Base{Name: "Nissan", Value: 55}, Base{Name: "Norster", Value: 2489}, Base{Name: "Nysa (Ниса)", Value: 2045}, Base{Name: "Oldsmobile", Value: 148}, Base{Name: "Oltcit", Value: 2963}, Base{Name: "Opel", Value: 56}, Base{Name: "ORA", Value: 2974}, Base{Name: "Otosan", Value: 55313}, Base{Name: "Packard", Value: 3193}, Base{Name: "Pagani", Value: 2644}, Base{Name: "Peerless", Value: 1852}, Base{Name: "Peg-Perego", Value: 3446}, Base{Name: "Peterbilt", Value: 346}, Base{Name: "Peugeot", Value: 58}, Base{Name: "Pininfarina", Value: 3590}, Base{Name: "Pinzgauer", Value: 3215}, Base{Name: "Plymouth", Value: 181}, Base{Name: "Polestar", Value: 6131}, Base{Name: "Pontiac", Value: 149}, Base{Name: "Porsche", Value: 59}, Base{Name: "Praga Baby", Value: 2728}, Base{Name: "Proton", Value: 60}, Base{Name: "Qifeng", Value: 5340}, Base{Name: "Quicksilver", Value: 1332}, Base{Name: "Ram", Value: 4369}, Base{Name: "Ravon", Value: 4681}, Base{Name: "Raysince", Value: 55374}, Base{Name: "Renault", Value: 62}, Base{Name: "Rezvani", Value: 4466}, Base{Name: "Rich", Value: 55281}, Base{Name: "Rimac", Value: 3330}, Base{Name: "Rivian", Value: 6227}, Base{Name: "Robeta", Value: 5761}, Base{Name: "Roewe", Value: 55063}, Base{Name: "Rolls-Royce", Value: 63}, Base{Name: "Rover", Value: 64}, Base{Name: "Runhorse", Value: 55377}, Base{Name: "Saab", Value: 65}, Base{Name: "Saipa", Value: 3437}, Base{Name: "Saleen", Value: 196}, Base{Name: "Samand", Value: 192}, Base{Name: "Samson", Value: 3972}, Base{Name: "Samsung", Value: 325}, Base{Name: "Saturn", Value: 331}, Base{Name: "Sceo", Value: 2734}, Base{Name: "Scion", Value: 3268}, Base{Name: "SEAT", Value: 67}, Base{Name: "Secma", Value: 2492}, Base{Name: "Selena", Value: 1726}, Base{Name: "Shelby", Value: 3100}, Base{Name: "Shuanghuan", Value: 195}, Base{Name: "Sidetracker", Value: 4003}, Base{Name: "Sihao (Sol)", Value: 55218}, Base{Name: "Skoda", Value: 70}, Base{Name: "Skywell", Value: 55199}, Base{Name: "SMA", Value: 311}, Base{Name: "Smart", Value: 71}, Base{Name: "SouEast", Value: 194}, Base{Name: "South-East", Value: 55115}, Base{Name: "Soyat", Value: 3212}, Base{Name: "Spyker", Value: 411}, Base{Name: "SsangYong", Value: 73}, Base{Name: "Star", Value: 214}, Base{Name: "Studebaker", Value: 3228}, Base{Name: "Subaru", Value: 75}, Base{Name: "Suda", Value: 2879}, Base{Name: "Suda Hanen", Value: 55068}, Base{Name: "Sunbeam", Value: 385}, Base{Name: "Suzuki", Value: 76}, Base{Name: "T-King", Value: 5341}, Base{Name: "Talbot", Value: 77}, Base{Name: "Tarpan Honker", Value: 2497}, Base{Name: "TATA", Value: 78}, Base{Name: "Tatra", Value: 204}, Base{Name: "Tazzari", Value: 3922}, Base{Name: "Tesla", Value: 2233}, Base{Name: "Think", Value: 6092}, Base{Name: "Think Global", Value: 4237}, Base{Name: "Thunder Tiger", Value: 4033}, Base{Name: "Tianma", Value: 1578}, Base{Name: "Tiffany", Value: 55081}, Base{Name: "Tiger", Value: 2050}, Base{Name: "Tofas", Value: 2552}, Base{Name: "Toyota", Value: 79}, Base{Name: "Trabant", Value: 345}, Base{Name: "Triumph", Value: 80}, Base{Name: "TVR", Value: 81}, Base{Name: "Ultima", Value: 3017}, Base{Name: "Van Hool", Value: 206}, Base{Name: "Vantage", Value: 5873}, Base{Name: "Vauxhall", Value: 82}, Base{Name: "Venturi", Value: 83}, Base{Name: "Venucia", Value: 184}, Base{Name: "Vepr", Value: 3320}, Base{Name: "Volkswagen", Value: 84}, Base{Name: "Volvo", Value: 85}, Base{Name: "Voyah", Value: 55279}, Base{Name: "Wanderer", Value: 2021}, Base{Name: "Wanfeng", Value: 2403}, Base{Name: "Wartburg", Value: 339}, Base{Name: "Weilan", Value: 55113}, Base{Name: "Weltmeister", Value: 55088}, Base{Name: "Wiesmann", Value: 1992}, Base{Name: "Willys", Value: 1820}, Base{Name: "Wuling", Value: 2653}, Base{Name: "XEV", Value: 55337}, Base{Name: "Xiaolong", Value: 3452}, Base{Name: "Xinkai", Value: 338}, Base{Name: "Xpeng", Value: 107}, Base{Name: "Yema", Value: 55201}, Base{Name: "Yogomo", Value: 5285}, Base{Name: "Yugo", Value: 87}, Base{Name: "Zastava", Value: 2309}, Base{Name: "Zeekr", Value: 55280}, Base{Name: "Zhidou", Value: 182}, Base{Name: "Zimmer", Value: 2958}, Base{Name: "Zotye", Value: 3610}, Base{Name: "Zuk", Value: 3089}, Base{Name: "ZX", Value: 322}, Base{Name: "ЄРАЗ", Value: 170}, Base{Name: "Богдан", Value: 188}, Base{Name: "Бронто", Value: 3000}, Base{Name: "ВАЗ / Lada", Value: 88}, Base{Name: "ГАЗ", Value: 91}, Base{Name: "ГолАЗ", Value: 410}, Base{Name: "Жук", Value: 169}, Base{Name: "ЗАЗ", Value: 89}, Base{Name: "ЗИЛ", Value: 168}, Base{Name: "ЗИМ", Value: 1544}, Base{Name: "ЗИС", Value: 186}, Base{Name: "ИЖ", Value: 92}, Base{Name: "ЛуАЗ", Value: 189}, Base{Name: "Москвич/АЗЛК", Value: 94}, Base{Name: "Причеп", Value: 154}, Base{Name: "РАФ", Value: 327}, Base{Name: "Ретро автомобілі", Value: 199}, Base{Name: "Саморобний", Value: 2863}, Base{Name: "СеАЗ", Value: 96}, Base{Name: "СМЗ", Value: 2491}, Base{Name: "СПЭВ / SPEV", Value: 1351}, Base{Name: "ТагАЗ", Value: 4537}, Base{Name: "ТогАЗ", Value: 2638}, Base{Name: "Тренер", Value: 4038}, Base{Name: "УАЗ", Value: 93}, Base{Name: "Циклон", Value: 4021}, Base{Name: "Эстония", Value: 2307}},
			wantErr: assert.NoError,
		},
		{
			name: "Get marks by category moto",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 0,
			},
			args: args{
				categoryID: CategoryMoto,
			},
			want:    Marks{Base{Name: "ABM", Value: 3875}, Base{Name: "Access", Value: 1463}, Base{Name: "Acxa", Value: 2282}, Base{Name: "Adler", Value: 2396}, Base{Name: "Adly", Value: 1569}, Base{Name: "Aeon", Value: 1004}, Base{Name: "AGT", Value: 1770}, Base{Name: "Aie motor", Value: 1465}, Base{Name: "AIMA", Value: 4945}, Base{Name: "Akumoto", Value: 2762}, Base{Name: "Alesin", Value: 4103}, Base{Name: "Alfa", Value: 2393}, Base{Name: "Alfacore Seev", Value: 5350}, Base{Name: "Alfamoto", Value: 1558}, Base{Name: "Alpha", Value: 2108}, Base{Name: "Alphasports", Value: 1009}, Base{Name: "Amazing Scooter", Value: 5009}, Base{Name: "American ironhorse", Value: 1012}, Base{Name: "Apollo", Value: 4164}, Base{Name: "Aprilia", Value: 1014}, Base{Name: "Arctic cat", Value: 1581}, Base{Name: "Argo", Value: 389}, Base{Name: "Ariel", Value: 404}, Base{Name: "Armada Moto", Value: 3151}, Base{Name: "Arora", Value: 4968}, Base{Name: "Asia Wing", Value: 5857}, Base{Name: "Asia-moto", Value: 4567}, Base{Name: "Atala", Value: 2834}, Base{Name: "Atk", Value: 1015}, Base{Name: "atMoto", Value: 4364}, Base{Name: "ATV", Value: 2521}, Base{Name: "Auto Moto", Value: 2659}, Base{Name: "Avantis", Value: 4209}, Base{Name: "Azimut", Value: 1662}, Base{Name: "Baja", Value: 1741}, Base{Name: "Bajaj", Value: 3662}, Base{Name: "Baltmotors", Value: 4030}, Base{Name: "BAODIAO", Value: 4317}, Base{Name: "Baotian", Value: 2232}, Base{Name: "Barracuda", Value: 6085}, Base{Name: "Bashan", Value: 2431}, Base{Name: "Baumann", Value: 3976}, Base{Name: "Benelli", Value: 1020}, Base{Name: "Benyco", Value: 4464}, Base{Name: "Benzhou", Value: 5397}, Base{Name: "Beta", Value: 1021}, Base{Name: "Big Bear Choppers", Value: 4564}, Base{Name: "Big dog", Value: 2897}, Base{Name: "Big Sam G", Value: 4443}, Base{Name: "Bimota", Value: 3031}, Base{Name: "Birel", Value: 2637}, Base{Name: "Blata", Value: 2377}, Base{Name: "Bluera", Value: 4981}, Base{Name: "BMW", Value: 9}, Base{Name: "Board", Value: 5545}, Base{Name: "Bobcat", Value: 1618}, Base{Name: "Bombardier", Value: 4789}, Base{Name: "Bomber", Value: 1668}, Base{Name: "Bonez", Value: 3058}, Base{Name: "Boom Trikes", Value: 2574}, Base{Name: "Booster", Value: 3911}, Base{Name: "Boxter", Value: 6304}, Base{Name: "BPG-Werks", Value: 3615}, Base{Name: "Brammo", Value: 55136}, Base{Name: "Bravis", Value: 5431}, Base{Name: "Bravo", Value: 2987}, Base{Name: "BRIG", Value: 1665}, Base{Name: "Brixton Motorcycles", Value: 6086}, Base{Name: "BRP", Value: 1471}, Base{Name: "BSA", Value: 3761}, Base{Name: "BSE", Value: 55053}, Base{Name: "BT", Value: 4909}, Base{Name: "Buell", Value: 1031}, Base{Name: "Bugtrail", Value: 3506}, Base{Name: "BWK", Value: 2536}, Base{Name: "BWS", Value: 4800}, Base{Name: "Cagiva", Value: 1033}, Base{Name: "Campagna", Value: 3781}, Base{Name: "Cannondale", Value: 5018}, Base{Name: "CARGO", Value: 4327}, Base{Name: "Carroli", Value: 5852}, Base{Name: "Catran", Value: 2820}, Base{Name: "Cectek", Value: 2079}, Base{Name: "Celebrity", Value: 2228}, Base{Name: "Cezet (Чезет)", Value: 2002}, Base{Name: "Cezeta", Value: 5644}, Base{Name: "CFMOTO", Value: 1037}, Base{Name: "Challenger", Value: 2262}, Base{Name: "Charming", Value: 6284}, Base{Name: "Chenlon", Value: 4363}, Base{Name: "Chituma", Value: 3900}, Base{Name: "Chongqing Wonjan", Value: 3782}, Base{Name: "Citycoco", Value: 5568}, Base{Name: "Club Car", Value: 6207}, Base{Name: "Club Сar", Value: 2372}, Base{Name: "Cobra", Value: 2238}, Base{Name: "Comer", Value: 6267}, Base{Name: "Comman", Value: 3469}, Base{Name: "Commandor", Value: 2847}, Base{Name: "Corrado", Value: 1547}, Base{Name: "Corsa", Value: 5571}, Base{Name: "Cougar", Value: 3952}, Base{Name: "CPI", Value: 1043}, Base{Name: "Crafter", Value: 4099}, Base{Name: "CRG", Value: 2674}, Base{Name: "Crosser", Value: 3026}, Base{Name: "Cruiser", Value: 55316}, Base{Name: "Custom Culture Ukraine", Value: 5406}, Base{Name: "D-MONIAK", Value: 1692}, Base{Name: "Dabra", Value: 5699}, Base{Name: "Daelim", Value: 1047}, Base{Name: "Dafier", Value: 3653}, Base{Name: "DASY", Value: 3118}, Base{Name: "Daymak", Value: 55150}, Base{Name: "Dayun", Value: 3473}, Base{Name: "Defiant", Value: 1542}, Base{Name: "Delta", Value: 1560}, Base{Name: "Derbi", Value: 1938}, Base{Name: "Df motor", Value: 1054}, Base{Name: "DFL", Value: 3816}, Base{Name: "Diamo", Value: 1057}, Base{Name: "Dimen", Value: 55372}, Base{Name: "Dingo", Value: 6180}, Base{Name: "DINLI", Value: 2660}, Base{Name: "Dirtbike", Value: 5066}, Base{Name: "Discovery", Value: 2437}, Base{Name: "DKW", Value: 2243}, Base{Name: "Dongma", Value: 55340}, Base{Name: "Doohan", Value: 55100}, Base{Name: "Drift Trike", Value: 4907}, Base{Name: "Drive King", Value: 5908}, Base{Name: "Ducati", Value: 1060}, Base{Name: "E - ATV", Value: 3183}, Base{Name: "E-Cult", Value: 4872}, Base{Name: "E-Kross", Value: 5607}, Base{Name: "E-Madix", Value: 4809}, Base{Name: "E-mania", Value: 55084}, Base{Name: "E-Ton", Value: 4279}, Base{Name: "E-Twow", Value: 4602}, Base{Name: "E-Z-GO", Value: 2490}, Base{Name: "Eagle", Value: 120}, Base{Name: "East Dragon", Value: 2310}, Base{Name: "Eco-Extreme", Value: 3924}, Base{Name: "EcoMoto", Value: 5483}, Base{Name: "EF-TK", Value: 4641}, Base{Name: "Eglem", Value: 4952}, Base{Name: "Eglmoto", Value: 5490}, Base{Name: "EH Line", Value: 5472}, Base{Name: "Electric Scooter", Value: 5973}, Base{Name: "Electromoto", Value: 55363}, Base{Name: "Electrowin", Value: 5476}, Base{Name: "Elwinn", Value: 5469}, Base{Name: "Emco", Value: 6082}, Base{Name: "EmGo", Value: 55251}, Base{Name: "EMU Alltrack", Value: 2467}, Base{Name: "Energy Power", Value: 6264}, Base{Name: "Epella", Value: 5659}, Base{Name: "Eriskay", Value: 1548}, Base{Name: "ESO", Value: 6032}, Base{Name: "Espero", Value: 2558}, Base{Name: "Eurotex", Value: 4624}, Base{Name: "EvoBike", Value: 5595}, Base{Name: "Exdrive", Value: 5067}, Base{Name: "Explorer", Value: 4964}, Base{Name: "EZRaider", Value: 55376}, Base{Name: "Fada", Value: 1549}, Base{Name: "Falcon", Value: 1732}, Base{Name: "Fantic", Value: 2438}, Base{Name: "Farmer", Value: 3163}, Base{Name: "Favorit", Value: 3298}, Base{Name: "Feishen", Value: 3919}, Base{Name: "Flybo", Value: 1550}, Base{Name: "Flyer", Value: 5098}, Base{Name: "Forte", Value: 3164}, Base{Name: "Fortune", Value: 2759}, Base{Name: "Forward", Value: 2275}, Base{Name: "Fosti", Value: 1069}, Base{Name: "Foton", Value: 187}, Base{Name: "FoxWell", Value: 4210}, Base{Name: "Freedomotor", Value: 2321}, Base{Name: "FreeGo", Value: 5594}, Base{Name: "Fuego", Value: 173}, Base{Name: "Futong", Value: 1941}, Base{Name: "FYM", Value: 1573}, Base{Name: "G&A", Value: 1566}, Base{Name: "G-max", Value: 1561}, Base{Name: "GAMAX", Value: 4134}, Base{Name: "Garelli", Value: 1071}, Base{Name: "Gas Gas", Value: 1570}, Base{Name: "Geeli", Value: 5559}, Base{Name: "Geely", Value: 185}, Base{Name: "Genata", Value: 3844}, Base{Name: "Generic", Value: 3205}, Base{Name: "Gentoya", Value: 5897}, Base{Name: "Geon", Value: 1976}, Base{Name: "Gepard", Value: 2280}, Base{Name: "GerioMobil", Value: 2509}, Base{Name: "Gibbs", Value: 4058}, Base{Name: "Gilera", Value: 1077}, Base{Name: "Gladiator", Value: 4560}, Base{Name: "GM", Value: 3220}, Base{Name: "Godzilla", Value: 3828}, Base{Name: "Goes", Value: 1230}, Base{Name: "Gogoro", Value: 4301}, Base{Name: "Golf Car", Value: 4316}, Base{Name: "Goped", Value: 5690}, Base{Name: "Goricke", Value: 4988}, Base{Name: "Groz", Value: 1575}, Base{Name: "Guowei", Value: 6318}, Base{Name: "Guzzi", Value: 1128}, Base{Name: "Haibike", Value: 5410}, Base{Name: "Hamer", Value: 2621}, Base{Name: "Hammer", Value: 2432}, Base{Name: "Hammerhead", Value: 4527}, Base{Name: "Hanway", Value: 4151}, Base{Name: "Haobon", Value: 3872}, Base{Name: "Haojin", Value: 6012}, Base{Name: "Harley", Value: 6162}, Base{Name: "Harley-Davidson", Value: 1078}, Base{Name: "Harlo", Value: 4049}, Base{Name: "Havana", Value: 5666}, Base{Name: "Hecht", Value: 4264}, Base{Name: "Hercules", Value: 3655}, Base{Name: "Hero Electric", Value: 4224}, Base{Name: "Hero Majestic", Value: 4158}, Base{Name: "Hero Splendor", Value: 4365}, Base{Name: "Hisun", Value: 2468}, Base{Name: "Honda", Value: 28}, Base{Name: "Hongda", Value: 3997}, Base{Name: "Honling", Value: 2183}, Base{Name: "Honor", Value: 2944}, Base{Name: "Hornet", Value: 5542}, Base{Name: "Hors", Value: 2000}, Base{Name: "Horse", Value: 2424}, Base{Name: "HP Power", Value: 3775}, Base{Name: "Huanma", Value: 5689}, Base{Name: "Huatian", Value: 1577}, Base{Name: "Hubtex", Value: 3977}, Base{Name: "Hunter", Value: 2384}, Base{Name: "Huoniao", Value: 2683}, Base{Name: "Hupper", Value: 3686}, Base{Name: "Husaberg", Value: 2092}, Base{Name: "Husqvarna", Value: 1085}, Base{Name: "Hussar", Value: 2071}, Base{Name: "Hyosung", Value: 1086}, Base{Name: "IFA (ІФА)", Value: 1486}, Base{Name: "Imperial", Value: 2081}, Base{Name: "Indian", Value: 1087}, Base{Name: "Infinum", Value: 4798}, Base{Name: "Intrepid", Value: 4169}, Base{Name: "Invacare", Value: 5414}, Base{Name: "IO", Value: 6248}, Base{Name: "Irbis", Value: 3931}, Base{Name: "Iron Motorcycles", Value: 5094}, Base{Name: "Isuzu", Value: 30}, Base{Name: "Italjet", Value: 1089}, Base{Name: "iTango", Value: 6297}, Base{Name: "Jawa (ЯВА)", Value: 1487}, Base{Name: "Jawa (Ява)-cz", Value: 1091}, Base{Name: "Jazz", Value: 3177}, Base{Name: "JBW", Value: 2661}, Base{Name: "Jeek", Value: 5831}, Base{Name: "Jetstar", Value: 55055}, Base{Name: "Jialing", Value: 2705}, Base{Name: "Jianshe", Value: 1093}, Base{Name: "Jieda", Value: 2870}, Base{Name: "JimStar", Value: 2486}, Base{Name: "JINCHENG", Value: 2675}, Base{Name: "Jinding", Value: 3207}, Base{Name: "Jinling", Value: 4044}, Base{Name: "Jinlun", Value: 2799}, Base{Name: "Jinyee", Value: 4100}, Base{Name: "Jitao", Value: 3790}, Base{Name: "Jmstar", Value: 4146}, Base{Name: "Jnen", Value: 3138}, Base{Name: "John Deere", Value: 1583}, Base{Name: "Johnny Pag", Value: 4362}, Base{Name: "Jonway", Value: 1574}, Base{Name: "Joyner", Value: 2357}, Base{Name: "Jumbo", Value: 1896}, Base{Name: "Junak", Value: 4067}, Base{Name: "Juneng", Value: 4772}, Base{Name: "Kainuo", Value: 4411}, Base{Name: "Kaitong", Value: 4794}, Base{Name: "Kallio", Value: 5810}, Base{Name: "Kaminah", Value: 3898}, Base{Name: "Kangchao", Value: 3937}, Base{Name: "Kangda", Value: 4383}, Base{Name: "Kansas", Value: 4806}, Base{Name: "Kanuni", Value: 1100}, Base{Name: "Kavaki Motor", Value: 4724}, Base{Name: "Kawasaki", Value: 176}, Base{Name: "Kayo", Value: 2352}, Base{Name: "Kazuma", Value: 2753}, Base{Name: "Keeway", Value: 1101}, Base{Name: "Keine", Value: 5005}, Base{Name: "Kenrod", Value: 3474}, Base{Name: "Kentoya", Value: 4728}, Base{Name: "Kewesekl", Value: 1681}, Base{Name: "Kindroad", Value: 5957}, Base{Name: "Kinetik", Value: 4360}, Base{Name: "King Star", Value: 6291}, Base{Name: "KingSong", Value: 5638}, Base{Name: "Kingway", Value: 3475}, Base{Name: "Kinlon", Value: 3192}, Base{Name: "Kinroad", Value: 1499}, Base{Name: "Kioti", Value: 3429}, Base{Name: "Kosmic", Value: 5021}, Base{Name: "Kovi", Value: 5939}, Base{Name: "Kreidler", Value: 4776}, Base{Name: "Kross", Value: 2826}, Base{Name: "KTM", Value: 1104}, Base{Name: "Kv", Value: 3814}, Base{Name: "KXD moto", Value: 4115}, Base{Name: "Kymco", Value: 1105}, Base{Name: "Kymera", Value: 3260}, Base{Name: "Lambretta", Value: 1106}, Base{Name: "Land Tamer", Value: 5171}, Base{Name: "LandMax", Value: 55356}, Base{Name: "Lantana", Value: 2493}, Base{Name: "Laverda", Value: 1985}, Base{Name: "Leader", Value: 1628}, Base{Name: "Leike", Value: 4625}, Base{Name: "Leopard", Value: 2442}, Base{Name: "Lepton", Value: 5971}, Base{Name: "Lexmoto", Value: 4605}, Base{Name: "Liberty GMG", Value: 4098}, Base{Name: "Lifan", Value: 334}, Base{Name: "Ligier", Value: 55258}, Base{Name: "Like.Bike", Value: 5167}, Base{Name: "Lima", Value: 55174}, Base{Name: "Lingben", Value: 4990}, Base{Name: "Linhai", Value: 1109}, Base{Name: "Lml", Value: 1110}, Base{Name: "Loncin", Value: 2026}, Base{Name: "Longjia", Value: 4826}, Base{Name: "LVTong", Value: 6309}, Base{Name: "LZ", Value: 3923}, Base{Name: "MadAss", Value: 3909}, Base{Name: "Maja", Value: 5008}, Base{Name: "Makc", Value: 2315}, Base{Name: "Malaguti", Value: 1115}, Base{Name: "Malanca", Value: 1116}, Base{Name: "Mangosteen", Value: 55338}, Base{Name: "Mar-co", Value: 3989}, Base{Name: "Maraton", Value: 55145}, Base{Name: "Marsun", Value: 1551}, Base{Name: "Masai", Value: 4866}, Base{Name: "Mash", Value: 5538}, Base{Name: "Matador", Value: 4361}, Base{Name: "Max Trailer", Value: 3987}, Base{Name: "Maxxter", Value: 5488}, Base{Name: "MBK", Value: 2223}, Base{Name: "MegaStar", Value: 3783}, Base{Name: "Megelli", Value: 2075}, Base{Name: "Meiduo", Value: 4139}, Base{Name: "Melex", Value: 2057}, Base{Name: "Menila", Value: 2433}, Base{Name: "Metrakit", Value: 3084}, Base{Name: "Mikilon", Value: 3147}, Base{Name: "Milan", Value: 5019}, Base{Name: "MINI", Value: 147}, Base{Name: "MiniCross", Value: 55168}, Base{Name: "Miro", Value: 2496}, Base{Name: "Mista", Value: 3996}, Base{Name: "MKS", Value: 4759}, Base{Name: "Mondial", Value: 55345}, Base{Name: "Monster", Value: 6209}, Base{Name: "Montesa Honda", Value: 4615}, Base{Name: "Moto Aupa", Value: 4566}, Base{Name: "Moto Bellini", Value: 5036}, Base{Name: "Moto Guzzi", Value: 3123}, Base{Name: "MOTO MORINI", Value: 2590}, Base{Name: "Moto-Leader", Value: 4250}, Base{Name: "Motobecane", Value: 4217}, Base{Name: "Motobi", Value: 2226}, Base{Name: "Motoczysz", Value: 1132}, Base{Name: "MotoJet", Value: 2254}, Base{Name: "Motoland", Value: 4414}, Base{Name: "Motom", Value: 1134}, Base{Name: "Motorro", Value: 1506}, Base{Name: "Motowell", Value: 4431}, Base{Name: "Motron", Value: 4040}, Base{Name: "MSKart", Value: 3611}, Base{Name: "Musstang", Value: 1565}, Base{Name: "Mustang", Value: 2979}, Base{Name: "Mv agusta", Value: 1138}, Base{Name: "MYBRO", Value: 5138}, Base{Name: "MZ", Value: 1139}, Base{Name: "NanFang", Value: 4152}, Base{Name: "Navigator", Value: 1552}, Base{Name: "New Holland", Value: 1696}, Base{Name: "Nexus", Value: 4039}, Base{Name: "Ninebot", Value: 5704}, Base{Name: "Ninebot One", Value: 4419}, Base{Name: "Nitro", Value: 2951}, Base{Name: "NIU", Value: 5128}, Base{Name: "Norco", Value: 3784}, Base{Name: "NSU", Value: 2157}, Base{Name: "Nzita", Value: 5987}, Base{Name: "OlympMotors", Value: 6050}, Base{Name: "Omax", Value: 55318}, Base{Name: "Optima", Value: 4612}, Base{Name: "Orion", Value: 3591}, Base{Name: "Orix", Value: 55259}, Base{Name: "Oset", Value: 4223}, Base{Name: "Panda", Value: 3583}, Base{Name: "Pannon", Value: 4054}, Base{Name: "Pannonia", Value: 1149}, Base{Name: "Pard", Value: 55058}, Base{Name: "Patriot", Value: 3116}, Base{Name: "Patron", Value: 2794}, Base{Name: "Peda", Value: 1980}, Base{Name: "Peg-Perego", Value: 3446}, Base{Name: "Pegasus", Value: 3482}, Base{Name: "PERAVES", Value: 4290}, Base{Name: "Peripoli", Value: 5705}, Base{Name: "Peugeot", Value: 58}, Base{Name: "Pgo", Value: 1151}, Base{Name: "Phoenix", Value: 2625}, Base{Name: "Piaggio", Value: 313}, Base{Name: "Pioneer", Value: 4626}, Base{Name: "Pit bike", Value: 2789}, Base{Name: "Pitbull", Value: 4147}, Base{Name: "Piton", Value: 4654}, Base{Name: "PitsterPro", Value: 4274}, Base{Name: "Pocket", Value: 6210}, Base{Name: "Pocket bike", Value: 4738}, Base{Name: "Polaris", Value: 1231}, Base{Name: "Polini", Value: 4129}, Base{Name: "Pony", Value: 3584}, Base{Name: "Presto", Value: 5304}, Base{Name: "Pride", Value: 5132}, Base{Name: "Pronto", Value: 5896}, Base{Name: "Puch", Value: 2591}, Base{Name: "Pulse", Value: 5055}, Base{Name: "Q-tek", Value: 3873}, Base{Name: "Qianjiang", Value: 1939}, Base{Name: "Qingqi", Value: 1157}, Base{Name: "Qjiang", Value: 1754}, Base{Name: "Quad Bike", Value: 2606}, Base{Name: "Quadro", Value: 5584}, Base{Name: "Quadzilla", Value: 3876}, Base{Name: "Racer", Value: 3910}, Base{Name: "Racing", Value: 3493}, Base{Name: "Rage", Value: 3402}, Base{Name: "Raketa-Futong", Value: 2106}, Base{Name: "RAP", Value: 55378}, Base{Name: "Ratas Moto", Value: 4801}, Base{Name: "Rato", Value: 5888}, Base{Name: "Razor", Value: 3409}, Base{Name: "Reaper", Value: 2657}, Base{Name: "Rebel Master", Value: 55062}, Base{Name: "Regal", Value: 1595}, Base{Name: "Regal-Raptor", Value: 2009}, Base{Name: "Rewaco", Value: 2869}, Base{Name: "Rex", Value: 3832}, Base{Name: "Rexon", Value: 2027}, Base{Name: "Rialli", Value: 2829}, Base{Name: "RiderKart", Value: 2634}, Base{Name: "Rieju", Value: 2462}, Base{Name: "Rivero", Value: 4227}, Base{Name: "Rizzato", Value: 55127}, Base{Name: "RM", Value: 4280}, Base{Name: "Road Knight", Value: 4286}, Base{Name: "Road Legal", Value: 2962}, Base{Name: "Roadsign", Value: 5518}, Base{Name: "Rokon", Value: 4005}, Base{Name: "Rolektro", Value: 5119}, Base{Name: "Romet", Value: 1693}, Base{Name: "Rover", Value: 64}, Base{Name: "Roxon", Value: 5843}, Base{Name: "Royal Enfield", Value: 3443}, Base{Name: "Runhorse", Value: 55377}, Base{Name: "Runmaster Motor", Value: 5562}, Base{Name: "Sabur", Value: 1556}, Base{Name: "Sachs", Value: 1166}, Base{Name: "Sagitta", Value: 4255}, Base{Name: "Salardi", Value: 5551}, Base{Name: "Samada", Value: 4388}, Base{Name: "Sandi", Value: 5506}, Base{Name: "Sanyou", Value: 4287}, Base{Name: "SAXXX E-Roadster", Value: 55264}, Base{Name: "Scandic", Value: 4070}, Base{Name: "Scootmobiel", Value: 5928}, Base{Name: "Scorpa", Value: 3261}, Base{Name: "Scorpion", Value: 3582}, Base{Name: "Screamin", Value: 3962}, Base{Name: "SEEV", Value: 4629}, Base{Name: "Segway", Value: 1557}, Base{Name: "Senke", Value: 5272}, Base{Name: "Sensor", Value: 2832}, Base{Name: "SH", Value: 4795}, Base{Name: "Shark", Value: 5025}, Base{Name: "Shawoom", Value: 2085}, Base{Name: "Sherco", Value: 3087}, Base{Name: "Shineray", Value: 1572}, Base{Name: "Shinerey", Value: 3090}, Base{Name: "Shoprider", Value: 5909}, Base{Name: "Sigma Line", Value: 3262}, Base{Name: "Simson", Value: 1176}, Base{Name: "Sinski", Value: 2815}, Base{Name: "Skaut", Value: 5099}, Base{Name: "Skeeter", Value: 3998}, Base{Name: "SkyBike", Value: 4230}, Base{Name: "Skygo", Value: 3112}, Base{Name: "SkyMoto", Value: 1555}, Base{Name: "SkyTeam", Value: 4874}, Base{Name: "SM-MOTO", Value: 4520}, Base{Name: "Smart", Value: 71}, Base{Name: "SmartWay", Value: 4690}, Base{Name: "SMC", Value: 2746}, Base{Name: "Snow hawk", Value: 2567}, Base{Name: "SNOWMAX", Value: 4766}, Base{Name: "Sodikart", Value: 4176}, Base{Name: "Sonik", Value: 3241}, Base{Name: "Sontan", Value: 2366}, Base{Name: "Soul", Value: 1853}, Base{Name: "SP", Value: 3895}, Base{Name: "Spark", Value: 1554}, Base{Name: "Sparta", Value: 5935}, Base{Name: "Sparta Pharos", Value: 4887}, Base{Name: "Specialized", Value: 3787}, Base{Name: "Speed Gear", Value: 1553}, Base{Name: "Speed Shek", Value: 2186}, Base{Name: "Speedex", Value: 55098}, Base{Name: "Spider", Value: 2473}, Base{Name: "Spike", Value: 5984}, Base{Name: "Spike ZZ", Value: 4107}, Base{Name: "Sport Energy", Value: 2584}, Base{Name: "Spy", Value: 5484}, Base{Name: "Stalker", Value: 3266}, Base{Name: "Stels", Value: 2213}, Base{Name: "Sterling", Value: 2754}, Base{Name: "Stinger", Value: 1563}, Base{Name: "Stock", Value: 5609}, Base{Name: "Storm", Value: 2060}, Base{Name: "Stormbringer", Value: 4880}, Base{Name: "Sukida", Value: 5497}, Base{Name: "Sumoto", Value: 2798}, Base{Name: "Sun City", Value: 4878}, Base{Name: "Sunbeam", Value: 385}, Base{Name: "Sunra", Value: 5570}, Base{Name: "Sunrise", Value: 3406}, Base{Name: "Sunrise Medical", Value: 5593}, Base{Name: "Super", Value: 55293}, Base{Name: "Sur-Ron", Value: 5960}, Base{Name: "Suzuki", Value: 76}, Base{Name: "Swiss Hutless", Value: 2570}, Base{Name: "SWM", Value: 5513}, Base{Name: "Sym", Value: 2176}, Base{Name: "T3 Motion", Value: 5548}, Base{Name: "TaiLG", Value: 1990}, Base{Name: "Tank Vision", Value: 5573}, Base{Name: "Tante", Value: 4312}, Base{Name: "TARO", Value: 55184}, Base{Name: "Tauris", Value: 6041}, Base{Name: "Tayo", Value: 2465}, Base{Name: "TCS", Value: 3979}, Base{Name: "TDMC", Value: 3294}, Base{Name: "Tekken", Value: 4203}, Base{Name: "Terra Motors", Value: 3099}, Base{Name: "Terrot", Value: 4444}, Base{Name: "TGB", Value: 1775}, Base{Name: "Thunder Mountain", Value: 5494}, Base{Name: "Tiger", Value: 2050}, Base{Name: "Tiras", Value: 3055}, Base{Name: "Tisong", Value: 5257}, Base{Name: "Titan", Value: 3186}, Base{Name: "TM Racing", Value: 2409}, Base{Name: "Tms", Value: 5905}, Base{Name: "Tomos", Value: 2151}, Base{Name: "Tontse", Value: 5866}, Base{Name: "Tony Kart", Value: 2650}, Base{Name: "Top-Kart", Value: 6099}, Base{Name: "Tornado", Value: 2078}, Base{Name: "Toros", Value: 5621}, Base{Name: "Tour", Value: 55165}, Base{Name: "Tourer", Value: 4799}, Base{Name: "Triad", Value: 4453}, Base{Name: "Trike", Value: 4797}, Base{Name: "TRIKEtec", Value: 4894}, Base{Name: "TriRod", Value: 4580}, Base{Name: "Triton", Value: 3102}, Base{Name: "Triumph", Value: 80}, Base{Name: "Truva", Value: 2224}, Base{Name: "TRX Scooter", Value: 3001}, Base{Name: "TVS", Value: 55139}, Base{Name: "TVS2", Value: 55132}, Base{Name: "UABike", Value: 5191}, Base{Name: "UGBEST", Value: 4531}, Base{Name: "United Motors", Value: 3384}, Base{Name: "Unix", Value: 3837}, Base{Name: "VACCI Bike", Value: 4979}, Base{Name: "Valenti", Value: 4627}, Base{Name: "Vanderhall", Value: 6115}, Base{Name: "Vanguard", Value: 3800}, Base{Name: "Vapor", Value: 4562}, Base{Name: "Varan", Value: 2291}, Base{Name: "Vectrix", Value: 2750}, Base{Name: "VEGA", Value: 1998}, Base{Name: "Veken", Value: 3788}, Base{Name: "Venom", Value: 2512}, Base{Name: "Venta", Value: 2256}, Base{Name: "Vento", Value: 4987}, Base{Name: "Ventus", Value: 3951}, Base{Name: "Veola", Value: 5871}, Base{Name: "Vertemati", Value: 4384}, Base{Name: "Verucci", Value: 2949}, Base{Name: "Vespa", Value: 1199}, Base{Name: "Victoria", Value: 2245}, Base{Name: "Victory", Value: 2474}, Base{Name: "Viking", Value: 2858}, Base{Name: "Vimann", Value: 3959}, Base{Name: "Viper", Value: 1201}, Base{Name: "Voge", Value: 55131}, Base{Name: "Volta", Value: 3793}, Base{Name: "VZ Yachts", Value: 3990}, Base{Name: "Warrior", Value: 4133}, Base{Name: "Wels", Value: 4142}, Base{Name: "Wenling", Value: 2816}, Base{Name: "Werya", Value: 2850}, Base{Name: "Wexxtor", Value: 3154}, Base{Name: "Wheelman", Value: 6303}, Base{Name: "Wild Horses", Value: 4950}, Base{Name: "Windtech", Value: 5683}, Base{Name: "Winner", Value: 3539}, Base{Name: "WK BIKES", Value: 4650}, Base{Name: "Wmotion", Value: 5108}, Base{Name: "WOQU", Value: 5691}, Base{Name: "WSK", Value: 4292}, Base{Name: "Wuyi ZhengLong", Value: 2121}, Base{Name: "X-Man", Value: 3180}, Base{Name: "XGJAO", Value: 1814}, Base{Name: "Xiamen", Value: 4985}, Base{Name: "Xiaomi", Value: 5107}, Base{Name: "Xingye", Value: 5164}, Base{Name: "Xingyue", Value: 1762}, Base{Name: "Xinling", Value: 2313}, Base{Name: "XinYang", Value: 2636}, Base{Name: "Xispa", Value: 1211}, Base{Name: "Xmotos", Value: 1212}, Base{Name: "XYKD", Value: 4437}, Base{Name: "Yadea", Value: 5479}, Base{Name: "Yamaha", Value: 179}, Base{Name: "Yamasaki", Value: 1546}, Base{Name: "Yamati", Value: 3540}, Base{Name: "Yamoto", Value: 2290}, Base{Name: "YCF", Value: 4563}, Base{Name: "YiBen", Value: 2560}, Base{Name: "Yinxiang", Value: 4313}, Base{Name: "Yokomoto", Value: 55102}, Base{Name: "Yongkang", Value: 4338}, Base{Name: "Yugen", Value: 6152}, Base{Name: "YUKI", Value: 4796}, Base{Name: "Zabel", Value: 3798}, Base{Name: "Zealsun Prince", Value: 2314}, Base{Name: "Zemis", Value: 2911}, Base{Name: "Zenith", Value: 3789}, Base{Name: "Zero", Value: 4303}, Base{Name: "Zeus", Value: 6142}, Base{Name: "Zhejiang", Value: 2288}, Base{Name: "Zhongqi", Value: 5498}, Base{Name: "Zhongyu", Value: 3567}, Base{Name: "Zipp", Value: 4157}, Base{Name: "Znen", Value: 2671}, Base{Name: "Zonda", Value: 1439}, Base{Name: "Zonder", Value: 1576}, Base{Name: "Zongshen", Value: 1214}, Base{Name: "Zontes", Value: 4272}, Base{Name: "Zorro", Value: 4997}, Base{Name: "Ztech", Value: 5934}, Base{Name: "ZTR", Value: 5523}, Base{Name: "Zubr", Value: 2426}, Base{Name: "Zulrace", Value: 5910}, Base{Name: "Zumico", Value: 4777}, Base{Name: "Zundapp", Value: 1216}, Base{Name: "АКУ", Value: 4324}, Base{Name: "Алиса", Value: 5072}, Base{Name: "Багги", Value: 3239}, Base{Name: "Буран", Value: 2523}, Base{Name: "Верховина", Value: 1755}, Base{Name: "Восход", Value: 1204}, Base{Name: "Вятка", Value: 2317}, Base{Name: "Геркулес", Value: 6135}, Base{Name: "ДАЗ", Value: 3097}, Base{Name: "Днепр (КМЗ)", Value: 171}, Base{Name: "Дорожник", Value: 5906}, Base{Name: "ДТЗ", Value: 3351}, Base{Name: "Заря", Value: 2052}, Base{Name: "ЗиД", Value: 1535}, Base{Name: "ЗИМ", Value: 1544}, Base{Name: "ЗИФ", Value: 1959}, Base{Name: "ИЖ", Value: 92}, Base{Name: "ИМЗ (Урал*)", Value: 2370}, Base{Name: "Карпати", Value: 1854}, Base{Name: "Кастом", Value: 2764}, Base{Name: "Кельбаджары", Value: 4936}, Base{Name: "Ковровец", Value: 3427}, Base{Name: "Красный Октябрь", Value: 2244}, Base{Name: "ЛВЗ", Value: 3305}, Base{Name: "ЛМЗ", Value: 4649}, Base{Name: "Мінськ", Value: 1125}, Base{Name: "Махо", Value: 1999}, Base{Name: "Москва", Value: 2194}, Base{Name: "Муравей", Value: 1933}, Base{Name: "Партнер", Value: 5073}, Base{Name: "ПатриотМото", Value: 3445}, Base{Name: "Полесье", Value: 4892}, Base{Name: "Ракета-мото", Value: 3401}, Base{Name: "Рига", Value: 2218}, Base{Name: "РМЗ", Value: 3265}, Base{Name: "Рысь", Value: 2563}, Base{Name: "С-Мото", Value: 1537}, Base{Name: "Саморобний", Value: 2863}, Base{Name: "Тайга", Value: 1234}, Base{Name: "ТИЗ", Value: 2010}, Base{Name: "ТМЗ", Value: 1539}, Base{Name: "Трайк", Value: 3494}, Base{Name: "Тула", Value: 1579}, Base{Name: "УКРмото", Value: 3003}, Base{Name: "Урал", Value: 95}},
			wantErr: assert.NoError,
		},
		{
			name: "Get marks with invalid category",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 0,
			},
			args: args{
				categoryID: 9999999,
			},
			want:    Marks{},
			wantErr: assert.NoError,
		},
		{
			name: "Get marks by category fail",
			fields: fields{
				apikey: "123123123",
				client: &http.Client{},

				maxRetries: 0,
			},
			args: args{
				categoryID: CategoryMoto,
			},
			want:    Marks(nil),
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey:     tt.fields.apikey,
				client:     tt.fields.client,
				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetMarksByCategory(tt.args.categoryID)
			if !tt.wantErr(t, err, ErrInvalidAPIKey) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetMarksByCategory(%v)", tt.args.categoryID)
		})
	}
}

func Test_service_GetModelsByCategoryAndMarkID(t *testing.T) {
	type fields struct {
		apikey string
		client *http.Client

		maxRetries int
		debug      bool
	}
	type args struct {
		categoryID int
		markID     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Models
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get models by category and mark success",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 0,
			},
			args: args{
				categoryID: CategoryCars,
				markID:     6, // AUDI
			},
			want:    Models{Base{Name: "100", Value: 39}, Base{Name: "200", Value: 41}, Base{Name: "4000", Value: 61284}, Base{Name: "50", Value: 43274}, Base{Name: "5000", Value: 61285}, Base{Name: "60", Value: 63476}, Base{Name: "75", Value: 63517}, Base{Name: "80", Value: 43}, Base{Name: "90", Value: 44}, Base{Name: "A1", Value: 31914}, Base{Name: "A1 Allstreet", Value: 63470}, Base{Name: "A1 Citycarver", Value: 63480}, Base{Name: "A2", Value: 45}, Base{Name: "A3", Value: 46}, Base{Name: "A3 Sportback", Value: 59425}, Base{Name: "A4", Value: 47}, Base{Name: "A4 Allroad", Value: 3457}, Base{Name: "A5", Value: 2032}, Base{Name: "A5 Sportback", Value: 62310}, Base{Name: "A6", Value: 49}, Base{Name: "A6 Allroad", Value: 3460}, Base{Name: "A7 Sportback", Value: 32378}, Base{Name: "A8", Value: 51}, Base{Name: "Cabriolet", Value: 63516}, Base{Name: "Coupe", Value: 52863}, Base{Name: "e-tron", Value: 56328}, Base{Name: "e-tron GT", Value: 60267}, Base{Name: "e-tron S", Value: 48010}, Base{Name: "e-tron S Sportback", Value: 48011}, Base{Name: "e-tron Sportback", Value: 60665}, Base{Name: "Fox", Value: 63515}, Base{Name: "Front", Value: 61287}, Base{Name: "Q2", Value: 49591}, Base{Name: "Q2L e-tron", Value: 62950}, Base{Name: "Q3", Value: 35548}, Base{Name: "Q3 Sportback", Value: 62596}, Base{Name: "Q4", Value: 60289}, Base{Name: "Q4 Sportback", Value: 62601}, Base{Name: "Q5", Value: 3222}, Base{Name: "Q5 e-tron", Value: 63514}, Base{Name: "Q5 Sportback", Value: 62612}, Base{Name: "Q7", Value: 1943}, Base{Name: "Q8", Value: 54664}, Base{Name: "Q8 e-tron", Value: 64480}, Base{Name: "Q8 Sportback e-tron", Value: 64479}, Base{Name: "R8", Value: 2263}, Base{Name: "RS e-tron GT", Value: 62387}, Base{Name: "RS Q3", Value: 44182}, Base{Name: "RS Q3 Sportback", Value: 62609}, Base{Name: "RS Q8", Value: 61203}, Base{Name: "RS2", Value: 54303}, Base{Name: "RS3", Value: 35281}, Base{Name: "RS3 Sportback", Value: 63424}, Base{Name: "RS4", Value: 60}, Base{Name: "RS5", Value: 32945}, Base{Name: "RS5 Sportback", Value: 63513}, Base{Name: "RS6", Value: 58}, Base{Name: "RS7 Sportback", Value: 44161}, Base{Name: "S1", Value: 63481}, Base{Name: "S1 Sportback", Value: 63482}, Base{Name: "S2", Value: 30336}, Base{Name: "S3", Value: 61}, Base{Name: "S4", Value: 62}, Base{Name: "S5", Value: 2814}, Base{Name: "S5 Sportback", Value: 63479}, Base{Name: "S6", Value: 64}, Base{Name: "S7 Sportback", Value: 40988}, Base{Name: "S8", Value: 66}, Base{Name: "Sport Quattro", Value: 63471}, Base{Name: "SQ2", Value: 56346}, Base{Name: "SQ5", Value: 44183}, Base{Name: "SQ5 Sportback", Value: 32946}, Base{Name: "SQ7", Value: 49206}, Base{Name: "SQ8", Value: 59400}, Base{Name: "SQ8 e-tron", Value: 64482}, Base{Name: "SQ8 Sportback e-tron", Value: 64481}, Base{Name: "Super 90", Value: 63478}, Base{Name: "TT", Value: 1837}, Base{Name: "TT RS", Value: 33483}, Base{Name: "TT S", Value: 3452}, Base{Name: "V8", Value: 69}},
			wantErr: assert.NoError,
		},
		{
			name: "Get models with invalid mark",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 0,
			},
			args: args{
				categoryID: CategoryCars,
				markID:     99999999999,
			},
			want:    Models{},
			wantErr: assert.NoError,
		},
		{
			name: "Get models by category and mark fail",
			fields: fields{
				apikey: "123123123",
				client: &http.Client{},

				maxRetries: 0,
			},
			args: args{
				categoryID: CategoryCars,
				markID:     1,
			},
			want:    Models(nil),
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetModelsByCategoryAndMarkID(tt.args.categoryID, tt.args.markID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetModelsByCategoryAndMarkID(%v, %v)", tt.args.categoryID, tt.args.markID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetModelsByCategoryAndMarkID(%v, %v)", tt.args.categoryID, tt.args.markID)
		})
	}
}

func Test_service_GetModelsByCategoryAndMarkIDWithGroups(t *testing.T) {
	type fields struct {
		apikey string
		client *http.Client

		maxRetries int
		debug      bool
	}
	type args struct {
		categoryID int
		markID     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Models
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get models by category and mark success AUDI",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 0,
			},
			args: args{
				categoryID: CategoryCars,
				markID:     6, // AUDI
			},
			want:    Models{Base{Name: "100", Value: 39}, Base{Name: "200", Value: 41}, Base{Name: "4000", Value: 61284}, Base{Name: "50", Value: 43274}, Base{Name: "5000", Value: 61285}, Base{Name: "60", Value: 63476}, Base{Name: "75", Value: 63517}, Base{Name: "80", Value: 43}, Base{Name: "90", Value: 44}, Base{Name: "A1", Value: 31914}, Base{Name: "A1 Allstreet", Value: 63470}, Base{Name: "A1 Citycarver", Value: 63480}, Base{Name: "A2", Value: 45}, Base{Name: "A3", Value: 46}, Base{Name: "A3 Sportback", Value: 59425}, Base{Name: "A4", Value: 47}, Base{Name: "A4 Allroad", Value: 3457}, Base{Name: "A5", Value: 2032}, Base{Name: "A5 Sportback", Value: 62310}, Base{Name: "A6", Value: 49}, Base{Name: "A6 Allroad", Value: 3460}, Base{Name: "A7 Sportback", Value: 32378}, Base{Name: "A8", Value: 51}, Base{Name: "Cabriolet", Value: 63516}, Base{Name: "Coupe", Value: 52863}, Base{Name: "e-tron", Value: 56328}, Base{Name: "e-tron GT", Value: 60267}, Base{Name: "e-tron S", Value: 48010}, Base{Name: "e-tron S Sportback", Value: 48011}, Base{Name: "e-tron Sportback", Value: 60665}, Base{Name: "Fox", Value: 63515}, Base{Name: "Front", Value: 61287}, Base{Name: "Q2", Value: 49591}, Base{Name: "Q2L e-tron", Value: 62950}, Base{Name: "Q3", Value: 35548}, Base{Name: "Q3 Sportback", Value: 62596}, Base{Name: "Q4", Value: 60289}, Base{Name: "Q4 Sportback", Value: 62601}, Base{Name: "Q5", Value: 3222}, Base{Name: "Q5 e-tron", Value: 63514}, Base{Name: "Q5 Sportback", Value: 62612}, Base{Name: "Q7", Value: 1943}, Base{Name: "Q8", Value: 54664}, Base{Name: "Q8 e-tron", Value: 64480}, Base{Name: "Q8 Sportback e-tron", Value: 64479}, Base{Name: "R8", Value: 2263}, Base{Name: "RS e-tron GT", Value: 62387}, Base{Name: "RS Q3", Value: 44182}, Base{Name: "RS Q3 Sportback", Value: 62609}, Base{Name: "RS Q8", Value: 61203}, Base{Name: "RS2", Value: 54303}, Base{Name: "RS3", Value: 35281}, Base{Name: "RS3 Sportback", Value: 63424}, Base{Name: "RS4", Value: 60}, Base{Name: "RS5", Value: 32945}, Base{Name: "RS5 Sportback", Value: 63513}, Base{Name: "RS6", Value: 58}, Base{Name: "RS7 Sportback", Value: 44161}, Base{Name: "S1", Value: 63481}, Base{Name: "S1 Sportback", Value: 63482}, Base{Name: "S2", Value: 30336}, Base{Name: "S3", Value: 61}, Base{Name: "S4", Value: 62}, Base{Name: "S5", Value: 2814}, Base{Name: "S5 Sportback", Value: 63479}, Base{Name: "S6", Value: 64}, Base{Name: "S7 Sportback", Value: 40988}, Base{Name: "S8", Value: 66}, Base{Name: "Sport Quattro", Value: 63471}, Base{Name: "SQ2", Value: 56346}, Base{Name: "SQ5", Value: 44183}, Base{Name: "SQ5 Sportback", Value: 32946}, Base{Name: "SQ7", Value: 49206}, Base{Name: "SQ8", Value: 59400}, Base{Name: "SQ8 e-tron", Value: 64482}, Base{Name: "SQ8 Sportback e-tron", Value: 64481}, Base{Name: "Super 90", Value: 63478}, Base{Name: "TT", Value: 1837}, Base{Name: "TT RS", Value: 33483}, Base{Name: "TT S", Value: 3452}, Base{Name: "V8", Value: 69}},
			wantErr: assert.NoError,
		},
		{
			name: "Get models by category and mark success BMW",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 0,
			},
			args: args{
				categoryID: CategoryCars,
				markID:     9, // BMW
			},
			want:    Models{Base{Name: " iX1", Value: 63528}, Base{Name: "1 Series", Value: 2161}, Base{Name: "1500", Value: 63521}, Base{Name: "1502", Value: 48926}, Base{Name: "1602", Value: 47380}, Base{Name: "2000", Value: 48925}, Base{Name: "2002", Value: 47386}, Base{Name: "1600", Value: 63518}, Base{Name: "1600 GT", Value: 63526}, Base{Name: "1800", Value: 63519}, Base{Name: "1802", Value: 63520}, Base{Name: "2 Series", Value: 43023}, Base{Name: "2 Series Active Tourer", Value: 44594}, Base{Name: "2 Series Gran Coupe", Value: 60508}, Base{Name: "2 Series Gran Tourer", Value: 61149}, Base{Name: "3 Series", Value: 3219}, Base{Name: "3 Series Compact", Value: 63522}, Base{Name: "3 Series GT", Value: 43029}, Base{Name: "3200 CS", Value: 63527}, Base{Name: "4 Series", Value: 42495}, Base{Name: "4 Series Gran Coupe", Value: 44037}, Base{Name: "5 Series", Value: 2319}, Base{Name: "5 Series GT", Value: 44727}, Base{Name: "6 Series", Value: 3218}, Base{Name: "6 Series Gran Coupe", Value: 39420}, Base{Name: "6 Series GT", Value: 52144}, Base{Name: "7 Series", Value: 18490}, Base{Name: "700", Value: 63523}, Base{Name: "8 Series", Value: 94}, Base{Name: "8 Series Gran Coupe", Value: 59451}, Base{Name: "Dixi", Value: 33383}, Base{Name: "E3", Value: 63525}, Base{Name: "E9", Value: 63524}, Base{Name: "I3", Value: 44838}, Base{Name: "i3S", Value: 63123}, Base{Name: "i4", Value: 63124}, Base{Name: "i5", Value: 65411}, Base{Name: "i7", Value: 63477}, Base{Name: "i8", Value: 44537}, Base{Name: "Isetta", Value: 32380}, Base{Name: "iX", Value: 62701}, Base{Name: "iX3", Value: 61275}, Base{Name: "iX5", Value: 63529}, Base{Name: "M1", Value: 95}, Base{Name: "M2", Value: 44856}, Base{Name: "M3", Value: 3292}, Base{Name: "M4", Value: 44857}, Base{Name: "M5", Value: 3213}, Base{Name: "M6", Value: 3592}, Base{Name: "M6 Gran Coupe", Value: 59544}, Base{Name: "M8", Value: 59450}, Base{Name: "M8 Gran Coupe", Value: 60512}, Base{Name: "M3 Compact", Value: 63530}, Base{Name: "Neue Klasse", Value: 47379}, Base{Name: "X1", Value: 3597}, Base{Name: "X2", Value: 42029}, Base{Name: "X3", Value: 1866}, Base{Name: "X3 M", Value: 57774}, Base{Name: "X4", Value: 43735}, Base{Name: "X4 M", Value: 57775}, Base{Name: "X5", Value: 96}, Base{Name: "X5 M", Value: 3158}, Base{Name: "X6", Value: 2153}, Base{Name: "X6 M", Value: 3442}, Base{Name: "X7", Value: 49169}, Base{Name: "XM", Value: 64483}, Base{Name: "Z1", Value: 97}, Base{Name: "Z3", Value: 98}, Base{Name: "Z3 M", Value: 3223}, Base{Name: "Z4", Value: 99}, Base{Name: "Z4 M", Value: 59518}, Base{Name: "Z8", Value: 100}},
			wantErr: assert.NoError,
		},
		{
			name: "Get models with invalid mark",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 0,
			},
			args: args{
				categoryID: CategoryCars,
				markID:     99999999999,
			},
			want:    Models(nil),
			wantErr: assert.NoError,
		},
		{
			name: "Get models by category and mark fail",
			fields: fields{
				apikey: "123123123",
				client: &http.Client{},

				maxRetries: 0,
			},
			args: args{
				categoryID: CategoryCars,
				markID:     1,
			},
			want:    Models(nil),
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetModelsByCategoryAndMarkIDWithGroups(tt.args.categoryID, tt.args.markID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetModelsByCategoryAndMarkIDWithGroups(%v, %v)", tt.args.categoryID, tt.args.markID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetModelsByCategoryAndMarkIDWithGroups(%v, %v)", tt.args.categoryID, tt.args.markID)
		})
	}
}

func Test_service_GetGenerationsByModelID(t *testing.T) {
	type fields struct {
		apikey string
		client *http.Client

		maxRetries int
		debug      bool
	}
	type args struct {
		modelID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Generations
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetGenerationsByModelID success",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 3,
				debug:      false,
			},
			args: args{
				modelID: 1653,
			},
			want: []Generations{
				{
					ID:   1653,
					Name: "Caravelle",
					Generations: []Generation{
						{GenerationId: 16, Name: "III покоління (FL)/T5", YearFrom: 2009, YearTo: 2015, ModelId: 1653, Eng: "t5-fl"},
						{GenerationId: 17, Name: "IV покоління/T6", YearFrom: 2015, YearTo: 2019, ModelId: 1653, Eng: "t6"},
						{GenerationId: 8066, Name: "IV покоління (FL)/T6.1", YearFrom: 2019, YearTo: 0, ModelId: 1653, Eng: "t6-1"},
						{GenerationId: 13684, Name: "I покоління/T3", YearFrom: 1981, YearTo: 1992, ModelId: 1653, Eng: "1-pokolenie"},
						{GenerationId: 13686, Name: "II покоління/T4", YearFrom: 1990, YearTo: 1996, ModelId: 1653, Eng: "2-pokolenie"},
						{GenerationId: 13688, Name: "II покоління (FL)/T4", YearFrom: 1996, YearTo: 2003, ModelId: 1653, Eng: "2-pokolenie-fl"},
						{GenerationId: 13689, Name: "III покоління/T5", YearFrom: 2003, YearTo: 2009, ModelId: 1653, Eng: "3-pokolenie"},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "GetGenerationsByModelID error",
			fields: fields{
				apikey: "apikey",
				client: &http.Client{},

				maxRetries: 3,
			},
			args: args{
				modelID: 1653,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetGenerationsByModelID(tt.args.modelID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetGenerationsByModelID(%v)", tt.args.modelID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetGenerationsByModelID(%v)", tt.args.modelID)
		})
	}
}

func Test_service_GetModificationsByGenerationID(t *testing.T) {
	type fields struct {
		apikey string
		client *http.Client

		maxRetries int
		debug      bool
	}
	type args struct {
		generationID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Modifications
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetModificationsByGenerationID success",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 3,
			},
			args: args{
				generationID: 17,
			},
			want:    Modifications{Base{Name: "2.0 BiTDI DSG (180 к.с.)", Value: 115538}, Base{Name: "2.0 BiTDI DSG (180 к.с.) 4Motion", Value: 115539}, Base{Name: "2.0 TDI DSG (102 к.с.)", Value: 121490}, Base{Name: "2.0 TDI DSG (140 к.с.)", Value: 115537}, Base{Name: "2.0 TDI MT (102 к.с.)", Value: 115534}, Base{Name: "2.0 TDI MT (140 к.с.)", Value: 115535}, Base{Name: "2.0 TDI MT (140 к.с.) 4Motion", Value: 115536}, Base{Name: "2.0 TSI AT (150 к.с.)", Value: 130497}, Base{Name: "2.0 TSI AТ (140 к.с.)", Value: 131051}, Base{Name: "2.0 TSI DSG (204 к.с.) 4Motion", Value: 130575}, Base{Name: "2.0 TSI DSG6 (204 к.с.)", Value: 130498}, Base{Name: "2.0 TSI DSG7 (204 к.с.)", Value: 130502}, Base{Name: "2.0 TSI MТ (150 к.с.)", Value: 130501}, Base{Name: "Common Rail 2.0 TDI DSG (140 к.с.) 3000 L1H1", Value: 122254}, Base{Name: "Common Rail 2.0 TDI DSG (140 к.с.) 3400 L2H1", Value: 122253}, Base{Name: "Common Rail 2.0 TDI DSG (180 к.с.) 3400 L2H1", Value: 124228}, Base{Name: "Common Rail 2.0 TDI MT (140 к.с.) 3000 L1H1", Value: 122251}, Base{Name: "Common Rail 2.0 TDI MT (140 к.с.) 3400 L2H1", Value: 122252}, Base{Name: "LR 2.0 BiTDI DSG (180 к.с.)", Value: 115544}, Base{Name: "LR 2.0 BiTDI DSG (180 к.с.) 4Motion", Value: 115545}, Base{Name: "LR 2.0 TDI DSG (140 к.с.)", Value: 115543}, Base{Name: "LR 2.0 TDI MT (102 к.с.)", Value: 115540}, Base{Name: "LR 2.0 TDI MT (140 к.с.)", Value: 115541}, Base{Name: "LR 2.0 TDI MT (140 к.с.) 4Motion", Value: 115542}, Base{Name: "LR 2.0 TSI AТ (140 к.с.)", Value: 130499}, Base{Name: "LR 2.0 TSI DSG (204 к.с.) 4Motion", Value: 130574}, Base{Name: "LR 2.0 TSI МТ (150 к.с.)", Value: 130503}},
			wantErr: assert.NoError,
		},
		{
			name: "GetModificationsByGenerationID error",
			fields: fields{
				apikey: "apikey",
				client: &http.Client{},

				maxRetries: 3,
			},
			args: args{
				generationID: 17,
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetModificationsByGenerationID(tt.args.generationID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetModificationsByGenerationID(%v)", tt.args.generationID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetModificationsByGenerationID(%v)", tt.args.generationID)
		})
	}
}

func Test_service_GetDriverTypes(t *testing.T) {
	type fields struct {
		apikey string
		client *http.Client

		maxRetries int
		debug      bool
	}
	type args struct {
		categoryID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    DriverTypes
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetDriverTypes success",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 3,
			},
			args: args{
				categoryID: 1,
			},
			want:    DriverTypes{Base{Name: "Повний", Value: 1}, Base{Name: "Передній", Value: 2}, Base{Name: "Задній", Value: 3}},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetDriverTypes(tt.args.categoryID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetDriverTypes(%v)", tt.args.categoryID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetDriverTypes(%v)", tt.args.categoryID)
		})
	}
}

func Test_service_GetFuelTypes(t *testing.T) {
	type fields struct {
		apikey string
		client *http.Client

		maxRetries int
		debug      bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    FuelTypes
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetFuelTypes success",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 3,
			},
			want:    FuelTypes{Base{Name: "Бензин", Value: 1}, Base{Name: "Дизель", Value: 2}, Base{Name: "Газ", Value: 3}, Base{Name: "Газ / Бензин", Value: 4}, Base{Name: "Гібрид", Value: 5}, Base{Name: "Електро", Value: 6}, Base{Name: "Інше", Value: 7}, Base{Name: "Газ метан", Value: 8}, Base{Name: "Газ пропан-бутан", Value: 9}},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetFuelTypes()
			if !tt.wantErr(t, err, "GetFuelTypes()") {
				return
			}
			assert.Equalf(t, tt.want, got, "GetFuelTypes()")
		})
	}
}

func Test_service_GetGearboxTypes(t *testing.T) {
	type fields struct {
		apikey string
		client *http.Client

		maxRetries int
		debug      bool
	}
	type args struct {
		categoryID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    GearboxesTypes
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetGearboxTypes success",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 3,
			},
			args: args{
				categoryID: 1,
			},
			want:    GearboxesTypes{Base{Name: "Ручна / Механіка", Value: 1}, Base{Name: "Автомат", Value: 2}, Base{Name: "Типтронік", Value: 3}, Base{Name: "Робот", Value: 4}, Base{Name: "Варіатор", Value: 5}},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetGearboxTypes(tt.args.categoryID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetGearboxTypes(%v)", tt.args.categoryID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetGearboxTypes(%v)", tt.args.categoryID)
		})
	}
}

func Test_service_GetOptions(t *testing.T) {
	type fields struct {
		apikey string
		client *http.Client

		maxRetries int
		debug      bool
	}
	type args struct {
		categoryID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    OptionTypes
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetOptions success",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 3,
			},
			args: args{
				categoryID: 1,
			},
			want:    OptionTypes{Base{Name: "Антиблокувальна система (ABS)", Value: 217}, Base{Name: "Антипробуксовочна система (ASR)", Value: 606}, Base{Name: "Блокування замків задніх дверей", Value: 622}, Base{Name: "Датчик втоми водія", Value: 617}, Base{Name: "Датчик проникнення в салон (датчик об`єму)", Value: 623}, Base{Name: "Датчик тиску в шинах", Value: 620}, Base{Name: "Допомога при спуску", Value: 611}, Base{Name: "Допомога при старті в гору", Value: 609}, Base{Name: "Запобігання зіткнення", Value: 612}, Base{Name: "Імобілайзер", Value: 604}, Base{Name: "Контроль за смугою руху", Value: 616}, Base{Name: "Контроль сліпих зон", Value: 615}, Base{Name: "Нічне бачення", Value: 619}, Base{Name: "Розпізнавання дорожніх знаків", Value: 618}, Base{Name: "Розподіл гальмівних зусиль (BAS, EBD)", Value: 608}, Base{Name: "Сигналізація", Value: 303}, Base{Name: "Система кріплення IsoFix", Value: 621}, Base{Name: "Система стабілізації (ESP)", Value: 459}, Base{Name: "Стабілізація рульового управління (VSM)", Value: 607}, Base{Name: "Центральний замок", Value: 137}, Base{Name: "Адаптивний круїз", Value: 591}, Base{Name: "Бардачок з охолодженням", Value: 585}, Base{Name: "Бездротова зарядка для смартфону", Value: 546}, Base{Name: "Бортовий комп'ютер", Value: 188}, Base{Name: "Вибір режиму руху", Value: 610}, Base{Name: "Відкриття багажника без допомоги рук", Value: 582}, Base{Name: "Датчик дощу", Value: 255}, Base{Name: "Декоративне підсвічування салону", Value: 564}, Base{Name: "Декоративні накладки на педалі", Value: 565}, Base{Name: "Дистанційний запуск двигуна", Value: 595}, Base{Name: "Доводчик дверей", Value: 587}, Base{Name: "Електронна приладова панель", Value: 594}, Base{Name: "Електропривід дзеркал", Value: 583}, Base{Name: "Електропривід кришки багажника", Value: 629}, Base{Name: "Електрорегулювання керма", Value: 576}, Base{Name: "Електроскладання дзеркал", Value: 584}, Base{Name: "Запуск двигуна з кнопки", Value: 525}, Base{Name: "Кермо з пам'яттю положення", Value: 577}, Base{Name: "Круїз контроль", Value: 605}, Base{Name: "Люк", Value: 132}, Base{Name: "Мультифункціональне кермо", Value: 579}, Base{Name: "Обігрів керма", Value: 524}, Base{Name: "Обігрів лобового скла", Value: 589}, Base{Name: "Оздоблення керма шкірою", Value: 555}, Base{Name: "Оздоблення стелі чорного кольору", Value: 559}, Base{Name: "Оздоблення шкірою важеля КПП", Value: 556}, Base{Name: "Панорамний дах / Лобове скло", Value: 558}, Base{Name: "Передній центральний підлокітник", Value: 560}, Base{Name: "Підігрів дзеркал", Value: 443}, Base{Name: "Підкурювач і попільничка", Value: 588}, Base{Name: "Підрульові пелюстки перемикання передач", Value: 580}, Base{Name: "Проекційний дисплей", Value: 575}, Base{Name: "Регульований педальний вузол", Value: 586}, Base{Name: "Розетка 12V", Value: 568}, Base{Name: "Розетка 220V", Value: 567}, Base{Name: "Сидіння з масажем", Value: 581}, Base{Name: "Система доступу без ключа", Value: 572}, Base{Name: "Система «старт-стоп»", Value: 574}, Base{Name: "Складане заднє сидіння", Value: 548}, Base{Name: "Складний столик на спинках передніх сидінь", Value: 554}, Base{Name: "Сонцезахисна шторка на задньому склі", Value: 631}, Base{Name: "Сонцезахисні шторки в задніх дверях", Value: 563}, Base{Name: "Тоновані вікна", Value: 486}, Base{Name: "Третій задній підголівник", Value: 550}, Base{Name: "Третій ряд сидінь", Value: 551}, Base{Name: "Функція складання спинки сидіння пасажира", Value: 549}, Base{Name: "Холодильник", Value: 590}, Base{Name: "Швидка зарядка CHAdeMO", Value: 597}, Base{Name: "Android Auto", Value: 544}, Base{Name: "AUX", Value: 538}, Base{Name: "Bluetooth", Value: 539}, Base{Name: "CarPlay", Value: 545}, Base{Name: "MirrorLink", Value: 633}, Base{Name: "USB", Value: 540}, Base{Name: "Акустика", Value: 258}, Base{Name: "Аудіопідготовка", Value: 534}, Base{Name: "Голосове керування", Value: 543}, Base{Name: "Керування жестами", Value: 632}, Base{Name: "Мультимедіа система з LCD-екраном", Value: 536}, Base{Name: "Навігаційна система", Value: 355}, Base{Name: "Система мультимедіа для задніх пасажирів", Value: 541}, Base{Name: "Автономний обігрівач webasto", Value: 596}, Base{Name: "Багажник на дах", Value: 635}, Base{Name: "Газобалонне обладнання (ГБО)", Value: 246}, Base{Name: "Лебідка", Value: 636}, Base{Name: "Ліфтована підвіска", Value: 638}, Base{Name: "Пандус для людей з інвалідністю", Value: 526}, Base{Name: "Пневмопідвіска", Value: 592}, Base{Name: "Ручне керування для людей з інвалідністю", Value: 502}, Base{Name: "Шноркель", Value: 637}, Base{Name: "Авто в кредиті", Value: 630}, Base{Name: "Гаражне зберігання", Value: 477}, Base{Name: "Перша реєстрація", Value: 501}, Base{Name: "Перший власник", Value: 496}, Base{Name: "Сервісна книжка", Value: 484}, Base{Name: "Датчик світла", Value: 437}, Base{Name: "Денні ходові вогні", Value: 527}, Base{Name: "Омивач фар", Value: 441}, Base{Name: "Протитуманні фари", Value: 528}, Base{Name: "Система адаптивного освітлення", Value: 531}, Base{Name: "Система управління дальнім світлом", Value: 532}, Base{Name: "Броньований кузов", Value: 515}, Base{Name: "Довга база", Value: 552}, Base{Name: "Захист картера", Value: 569}, Base{Name: "Захист коробки", Value: 570}, Base{Name: "Кузов MAXI", Value: 553}, Base{Name: "Накладки на пороги", Value: 566}, Base{Name: "Фаркоп", Value: 571}, Base{Name: "Задня камера", Value: 602}, Base{Name: "Камера 360", Value: 603}, Base{Name: "Парктронік задній", Value: 598}, Base{Name: "Парктронік передній", Value: 192}, Base{Name: "Передня камера", Value: 601}, Base{Name: "Система автоматичного паркування", Value: 599}, Base{Name: "Бічні задні", Value: 626}, Base{Name: "Бічні передні", Value: 625}, Base{Name: "Віконні (шторки)", Value: 627}, Base{Name: "Водія", Value: 211}, Base{Name: "Колін водія", Value: 628}, Base{Name: "Пасажира", Value: 624}, Base{Name: "Центральна подушка між водієм та пасажиром", Value: 634}},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetOptions(tt.args.categoryID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetOptions(%v)", tt.args.categoryID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetOptions(%v)", tt.args.categoryID)
		})
	}
}

func Test_service_GetColors(t *testing.T) {
	type fields struct {
		apikey string
		client *http.Client

		maxRetries int
		debug      bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    ColorTypes
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetColors success",
			fields: fields{
				apikey: os.Getenv("AUTORIA_API_KEY"),
				client: &http.Client{},

				maxRetries: 3,
			},
			want:    ColorTypes{Base{Name: "Бежевий", Value: 1}, Base{Name: "Чорний", Value: 2}, Base{Name: "Синій", Value: 3}, Base{Name: "Коричневий", Value: 5}, Base{Name: "Зелений", Value: 7}, Base{Name: "Сірий", Value: 8}, Base{Name: "Помаранчевий", Value: 9}, Base{Name: "Фіолетовий", Value: 12}, Base{Name: "Червоний", Value: 13}, Base{Name: "Білий", Value: 15}, Base{Name: "Жовтий", Value: 16}},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				apikey: tt.fields.apikey,
				client: tt.fields.client,

				maxRetries: tt.fields.maxRetries,
				debug:      tt.fields.debug,
			}
			got, err := s.GetColors()
			if !tt.wantErr(t, err, "GetColors()") {
				return
			}
			assert.Equalf(t, tt.want, got, "GetColors()")
		})
	}
}

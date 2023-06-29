package autoria

const (
	// ErrorCodeApiKeyMissing	API ключ не було вказано. Детальніше про використання API ключа дивитись у розділі "Використання API Key" .
	ErrorCodeApiKeyMissing = "API_KEY_MISSING"
	// ErrorCodeApiKeyInvalid Було вказано неправильний ключ API. Переконайтеся, що було вказано коректний API ключ, або зареєструйтеся для отримання ключа.
	ErrorCodeApiKeyInvalid = "API_KEY_INVALID"
	// ErrorCodeApiKeyDisabled API ключ було вимкнено адміністратором. Будь ласка зв'яжіться з нами для допомоги.
	ErrorCodeApiKeyDisabled = "API_KEY_DISABLED"
	// ErrorCodeApiKeyUnauthorized API ключ не авторизований для доступу до цього сервісу. Будь ласка зв'яжіться з нами для допомоги.
	ErrorCodeApiKeyUnauthorized = "API_KEY_UNAUTHORIZED"
	// ErrorCodeApiKeyUnverified API ключ не підтверджено. Перевірте свою електронну пошту для підтвердження API ключа. Будь ласка зв'яжіться з нами для допомоги.
	ErrorCodeApiKeyUnverified = "API_KEY_UNVERIFIED"
	// ErrorCodeHttpsRequired Запити до цього API мають бути зроблені за HTTPS протоколом. Переконайтеся, будь ласка, що Ви використовуєте HTTPS протокол для Вашого запиту.
	ErrorCodeHttpsRequired = "HTTPS_REQUIRED"
	// ErrorCodeOverRateLimit Перевищено ліміт запитів для цього API ключа. Зв'яжіться з нами для отримання додаткової інформації про ліміти веб-сервісів .
	ErrorCodeOverRateLimit = "OVER_RATE_LIMIT"
	// ErrorCodeNotFound Не вдалося знайти API за цією URL-адресою. Перевірте свій URL.
	ErrorCodeNotFound = "NOT_FOUND"
)

const (
	// CategoryCars Легкові
	CategoryCars = iota + 1
	// CategoryMoto Мото
	CategoryMoto
	// CategoryWaterTransport Водний транспорт
	CategoryWaterTransport
	// CategorySpecialMachinery Спецтехніка
	CategorySpecialMachinery
	// CategoryTrailers Причепи
	CategoryTrailers
	// CategoryTrucks Вантажівки
	CategoryTrucks
	// CategoryBuses Автобуси
	CategoryBuses
	// CategoryAutoBuildings Автобудинки
	CategoryAutoBuildings
	// CategoryAirTransport Повітряний транспорт
	CategoryAirTransport
	// CategoryAgriculturalMachinery Сільгосптехніка
	CategoryAgriculturalMachinery
)

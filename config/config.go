package config

// Configuration интерфейс определяет методы для получения параметров
// конфигурации с поддержкой получения значений строк,
// целых чисел, чисел с плавающей запятой и логического значения.
// Существует также набор методов, позволяющих указать значение по
// умолчанию. Данные конфигурации допускают вложенные разделы
// конфигурации, которые можно получить с помощью метода GetSection.
type Configuration interface {
	GetString(name string) (configValue string, found bool)
	GetInt(name string) (configValue int, found bool)
	GetBool(name string) (configValue bool, found bool)
	GetFloat(name string) (configValue float64, found bool)
	GetStringDefault(name, defVal string) (configValue string)
	GetIntDefault(name string, defVal int) (configValue int)
	GetBoolDefault(name string, defVal bool) (configValue bool)
	GetFloatDefault(name string, defVal float64) (configValue float64)
	GetSection(sectionName string) (section Configuration, found bool)
}

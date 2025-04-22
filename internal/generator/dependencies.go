package generator

var featureDependencies = map[string][]string{
	"postgres": {
		"gorm.io/gorm",
		"gorm.io/driver/postgres",
	},
	"viper": {
		"github.com/spf13/viper",
	},
	"redis": {
		"github.com/redis/go-redis/v9",
	},
	"jwt": {
		"github.com/golang-jwt/jwt/v5",
	},
}

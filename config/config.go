package config

// Config adalah root struct yang merepresentasikan seluruh isi config.yaml
type Config struct {
	Env          string       `mapstructure:"env"`
	Port         int          `mapstructure:"port"`
	JWT          JWTConfig    `mapstructure:"jwt"`
	DB           DBConfig     `mapstructure:"db"`
	DefaultValue DefaultValue `mapstructure:"default-value"`
}

type JWTConfig struct {
	Expire    int    `mapstructure:"expire"`
	SecretKey string `mapstructure:"secret-key"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Interval int    `mapstructure:"interval"`
	IdleConn int    `mapstructure:"idle-conn"`
	OpenConn int    `mapstructure:"open-conn"`
}

type DefaultValue struct {
	Timezone  string          `mapstructure:"timezone"`
	Breakfast BreakfastConfig `mapstructure:"breakfast"`
	Checkout  CheckoutConfig  `mapstructure:"checkout"`
}

type BreakfastConfig struct {
	Time     string `mapstructure:"time"`
	Location string `mapstructure:"location"`
}

type CheckoutConfig struct {
	Time       string `mapstructure:"time"`
	LatePolicy string `mapstructure:"late-policy"`
}

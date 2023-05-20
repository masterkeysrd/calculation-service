package jwt

type Config struct {
	SecretKey       string `mapstructure:"secretKey"`
	AccessTokenTTL  int    `mapstructure:"accessTokenTTL"`
	RefreshTokenTTL int    `mapstructure:"refreshTokenTTL"`
	Issuer          string `mapstructure:"issuer"`
	Audience        string `mapstructure:"audience"`
}

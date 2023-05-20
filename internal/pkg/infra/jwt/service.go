package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/dig"
)

type Claims struct {
	jwt.RegisteredClaims
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Service interface {
	GenerateTokens(userId string) (*Tokens, error)
	ValidateToken(c *gin.Context) error
}

type ServiceParams struct {
	dig.In
	Config *Config
}

type service struct {
	config *Config
}

func NewJwtService(params ServiceParams) Service {
	return &service{
		config: params.Config,
	}
}

func (s *service) GenerateTokens(userId string) (*Tokens, error) {
	id := uuid.New().String()
	issuedAt := time.Now().Add(-5 * time.Second)

	accessToken, err := s.generateAccessToken(id, userId, issuedAt)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(id, userId, issuedAt)

	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *service) generateAccessToken(id string, userId string, issuedAt time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	expirationTime := time.Now().Add(s.getAccessTokenTTL())
	token.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.config.Issuer,
			Subject:   userId,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			Audience:  jwt.ClaimStrings{s.config.Audience},
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			ID:        id,
		},
	}

	return token.SignedString([]byte(s.config.SecretKey))
}

func (s *service) generateRefreshToken(id string, userId string, issuedAt time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	expirationTime := time.Now().Add(s.getRefreshTokenTTL())
	token.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.config.Issuer,
			Subject:   userId,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			Audience:  jwt.ClaimStrings{s.config.Audience},
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			ID:        id,
		},
	}

	return token.SignedString([]byte(s.config.SecretKey))
}

func (s *service) getAccessTokenTTL() time.Duration {
	return time.Duration(s.config.AccessTokenTTL) * time.Second
}

func (s *service) getRefreshTokenTTL() time.Duration {
	return time.Duration(s.config.RefreshTokenTTL) * time.Second
}

func (s *service) ValidateToken(c *gin.Context) error {
	tokenString := ExtractToken(c)

	if len(tokenString) == 0 {
		return ErrTokenIsNotProvided
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}

		return []byte(s.config.SecretKey), nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrorInvalidToken
	}

	userId, err := strconv.ParseUint(claims["sub"].(string), 10, 64)
	if err != nil {
		return errors.New("invalid user id")
	}

	c.Set("tokenId", claims["jti"])
	c.Set("userId", userId)

	return nil
}

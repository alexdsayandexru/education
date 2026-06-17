module auth

go 1.23.0

replace jwt => ../jwt

require (
	github.com/google/uuid v1.6.0
	jwt v0.0.0-00010101000000-000000000000
)

require github.com/golang-jwt/jwt/v5 v5.2.1 // indirect

package helpers

import (
    "time"
    "errors"
    "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(GetEnv("JWT_SECRET", "DEFAULT_SECRET_KEY"))

func GenerateJWT(userID uint, email string) (string, error) {
    claims := jwt.MapClaims{}
    
    claims["userID"] = userID
    claims["email"] = email
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil {
        if verr, ok := err.(*jwt.ValidationError); ok && verr.Errors == jwt.ValidationErrorExpired {
            return nil, errors.New("토큰이 만료되었습니다.")
        }
        return nil, errors.New("토큰이 유효하지 않습니다.")
    }

    if !token.Valid {
        return nil, errors.New("토큰 형식이 올바르지 않습니다.")
    }

    return token, nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
    token, err := ValidateJWT(tokenString)
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, errors.New("토큰에서 클레임을 추출하는 데 실패했습니다.")
    }

    return claims, nil
}

func ExtractUserIDAndEmail(claims jwt.MapClaims) (uint, string, error) {
    userID, ok := claims["userID"].(float64)
    if !ok {
        return 0, "", errors.New("클레임에 userID 정보가 포함되어 있지 않습니다.")
    }

    email, ok := claims["email"].(string)
    if !ok {
        return 0, "", errors.New("클레임에 email 정보가 포함되어 있지 않습니다.")
    }

    return uint(userID), email, nil
}

func GetUserIDEndEmailFromToken(tokenString string) (uint, string, error) {
    claims, err := GetClaimsFromToken(tokenString)
    if err != nil {
        return 0, "", err
    }

    return ExtractUserIDAndEmail(claims)
}

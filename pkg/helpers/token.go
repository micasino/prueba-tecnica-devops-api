package helper

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

// SignedDetails
type SignedDetails struct {
	Username   string
	Email      string
	First_name string
	Last_name  string
	ID         uint
	jwt.StandardClaims
}

var SECRET_KEY string = "key_validate"           //os.Getenv("SECRET_KEY")
var JWT_EXPIRATION_IN_MILISECOND int64 = 3600000 //60s // (3600000 --1h) //os.Getenv("SECRET_KEY")

// GenerateAllTokens generates both teh detailed token and refresh token
func GenerateAllTokens(username string, email string, firstName string, lastName string, userType string, id uint) (signedToken string, signedRefreshToken string, expiresInMilSec int64, err error) {

	var timeDurationInSecound = JWT_EXPIRATION_IN_MILISECOND / 1000
	claims := &SignedDetails{
		Username:   username,
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		ID:         id,
		//User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(timeDurationInSecound)).Unix(), //divide el tiempo entre 1000 por el valor en miliseconds
		},
	}

	refreshClaims := &SignedDetails{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, JWT_EXPIRATION_IN_MILISECOND, err
}

//ValidateToken validates the jwt token
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

func ExtractTokenMetadata(token *jwt.Token) map[string]interface{} {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims
	}
	return nil
}

//UpdateAllTokens renews the user tokens when they login
// func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	var updateObj primitive.D

// 	updateObj = append(updateObj, bson.E{"token", signedToken})
// 	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

// 	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

// 	upsert := true
// 	filter := bson.M{"user_id": userId}
// 	opt := options.UpdateOptions{
// 		Upsert: &upsert,
// 	}

// 	_, err := userCollection.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.D{
// 			{"$set", updateObj},
// 		},
// 		&opt,
// 	)
// 	defer cancel()

// 	if err != nil {
// 		log.Panic(err)
// 		return
// 	}

// 	return
// }

package config

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/redis/go-redis/v9"

	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/redis/go-redis.v9"
)

func AzCredentialToRedis(credential azcore.TokenCredential) func(context.Context) (string, string, error) {
	return func(ctx context.Context) (string, string, error) {
		tk, err := credential.GetToken(ctx, policy.TokenRequestOptions{
			Scopes: []string{"https://redis.azure.com/.default"},
		})
		if err != nil {
			return "", "", err
		}

		parts := strings.Split(tk.Token, ".") // the token is a JWT; get the principal's object ID from its payload
		if len(parts) != 3 {
			return "", "", errors.New("token must have 3 parts")
		}
		payload, err := base64.RawURLEncoding.DecodeString(parts[1])
		if err != nil {
			return "", "", fmt.Errorf("failed decoding payload: %s", err)
		}

		claims := struct {
			OID string `json:"oid"`
		}{}
		err = json.Unmarshal(payload, &claims)
		if err != nil {
			return "", "", fmt.Errorf("couldn't unmarshal payload: %s", err)
		}

		if claims.OID == "" {
			return "", "", errors.New("missing object ID claim")
		}
		return claims.OID, tk.Token, nil
	}
}

func InitRedis() redis.UniversalClient {
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Printf("failed getting Azure credential: %s", err)
		return nil
	}

	redisDbUrl := os.Getenv("REDIS_DATABASE_URL")

	opt, err := redis.ParseURL(redisDbUrl)
	if err != nil {
		fmt.Println("rdb error: ", err.Error())
		os.Exit(1)
	}

	return redistrace.NewClient(&redis.Options{
		Addr:                       opt.Addr,
		Password:                   opt.Password,
		DB:                         0, // use default DB,
		TLSConfig:                  &tls.Config{MinVersion: tls.VersionTLS12},
		CredentialsProviderContext: AzCredentialToRedis(credential),
	})

}

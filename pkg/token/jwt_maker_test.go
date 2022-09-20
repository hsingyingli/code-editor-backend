package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	secretkey := "werwioejofisjdofj2iojroijrowkjefojfoiwejfojweifjwjiefjwoeijfiwjefjowiejfiowefjwioejfwioj"
	maker, err := NewJWTMaker(secretkey)

	require.NoError(t, err)

	username := "aaron"
	duration := time.Minute

	createAt := time.Now()
	expiredAt := createAt.Add(duration)

	token, _, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, createAt, payload.CreateAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	secretkey := "werwioejofisjdofj2iojroijrowkjefojfoiwejfojweifjwjiefjwoeijfiwjefjowiejfiowefjwioejfwioj"
	maker, err := NewJWTMaker(secretkey)

	require.NoError(t, err)

	username := "aaron"
	duration := time.Minute
	token, _, err := maker.CreateToken(username, -duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)

}

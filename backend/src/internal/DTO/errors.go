package DTO

import "errors"

var (
	ErrUserNotFound      = errors.New("пользователь не найден")
	ErrUserAlreadyExists = errors.New("пользователь с таким email уже существует")
	ErrInvalidPassword   = errors.New("неверный пароль")
	ErrInvalidToken      = errors.New("невалидный или просроченный JWT токен")
	ErrNoToken           = errors.New("токен авторизации отсутствует")
)

package ntpconsumer

import (
	ntp "github.com/beevik/ntp"
	"time"
)

// GetTime берет данные с нтп сервера ntp1.vniiftri.ru (Рандомный нтп сервер) и возращает время.
// Если произошла ошибка возвращается ошибка.
func GetTime() (time.Time, error) {
	return ntp.Time("ntp1.vniiftri.ru")
}

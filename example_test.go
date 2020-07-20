package s4lenium_test

import (
	"../s4lenium"
	"fmt"
	"github.com/xela07ax/toolsXela/chLogger"
	"github.com/xela07ax/toolsXela/tp"
	"path/filepath"
	"testing"
	"time"
)


func TestLocal(t *testing.T) {
	// 1 Создаем логер
	FullLogPath := filepath.Join("/","log")
	err := tp.CheckMkdir(FullLogPath)
	tp.FckText(fmt.Sprintf("Ошибка временной папки: %s",FullLogPath),err)

	// 2 Запускаем логер
	d2 := 300* time.Millisecond
	logEr := chLogger.NewChLoger(FullLogPath,&d2)
	logEr.RunMinion()
	logEr.ChInLog <- [4]string{"Welcome","nil",fmt.Sprintf("Вас приветствует \"Бот S4len Контроллер\" v2.1 (150720) ")}

	// 3 Создадим экземпляр
	s := s4lenium.NewS4len(logEr.ChInLog)
	wd, service, err := s.InitializeLocal("D:\\Programs\\chromeDriver\\chromedriver_83.exe")
	tp.FckText(fmt.Sprintf("Создадим экземпляр %s",FullLogPath),err)
	// Обязательно!
	defer service.Stop()
	defer wd.Quit()

	// Выполняем тесты
	tp.FckText(fmt.Sprintf("Ошибка выполнения теста: %s","1"),s.Test1(wd))
}
func TestRemote(t *testing.T) {
	// 1 Создаем логер
	FullLogPath := filepath.Join("/","log")
	err := tp.CheckMkdir(FullLogPath)
	tp.FckText(fmt.Sprintf("Ошибка временной папки: %s",FullLogPath),err)

	// 2 Запускаем логер
	d2 := 300* time.Millisecond
	logEr := chLogger.NewChLoger(FullLogPath,&d2)
	logEr.RunMinion()
	logEr.ChInLog <- [4]string{"Welcome","nil",fmt.Sprintf("Вас приветствует \"Бот S4len Контроллер\" v2.1 (150720) ")}

	// 3 Создадим экземпляр
	s := s4lenium.NewS4len(logEr.ChInLog)
	wd, err := s.InitializeRemote("http://192.168.99.106:4444/wd/hub")
	tp.FckText(fmt.Sprintf("Создадим экземпляр %s",FullLogPath),err)
	// Обязательно!
	defer wd.Quit()

	// Выполняем тесты
	tp.FckText(fmt.Sprintf("Ошибка выполнения теста: %s","1"),s.Test1(wd))
}
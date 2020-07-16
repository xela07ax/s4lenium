package Docs_test

import (
	"github.com/xela07ax/s4lenium"
	"github.com/xela07ax/toolsXela/chLogger"
	"github.com/xela07ax/toolsXela/tp"
	"fmt"
	"path/filepath"
	"time"
)

func main()  {
	// Создаем логер
	FullLogPath := filepath.Join("/","log")
	err := tp.CheckMkdir(FullLogPath)
	tp.FckText(fmt.Sprintf("Ошибка временной папки: %s",FullLogPath),err)

	// Запускаем логер
	d2 := 300* time.Millisecond
	logEr := chLogger.NewChLoger(FullLogPath,&d2)
	logEr.RunMinion()
	logEr.ChInLog <- [4]string{"Welcome","nil",fmt.Sprintf("Вас приветствует \"Бот S4len Контроллер\" v2.1 (150720) ")}

	// Создадим экземпляр
	s := s4lenium.NewS4len(logEr.ChInLog)
	wd, service, err := s.Initialize("D:\\Programs\\chromeDriver\\chromedriver_83.exe")
	tp.FckText(fmt.Sprintf("Создадим экземпляр %s",FullLogPath),err)
	// Обдязательно!
	defer service.Stop()
	defer wd.Quit()

	// Выполняем тесты
	tp.FckText(fmt.Sprintf("Ошибка выполнения теста: %s","1"),s.Test1(wd))

}


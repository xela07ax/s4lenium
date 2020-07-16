# Selenium Go package Alpha from tebeka/selenium
Пакет создан для уменьшения количества повторений кода инициализации селениума, путем его переиспользования.
Порядок менять нельзя!
### Использование:
```go
import (
	"github.com/xela07ax/s4lenium"
)
// Создадим экземпляр
	selen := s4lenium.NewS4len(logEr.ChInLog)
	wd, service, err := selen.Initialize("D:\\Programs\\chromeDriver\\chromedriver_83.exe")
	tp.FckText("Создадим экземпляр селениума",err)
	// Обязательно!
	defer service.Stop()
	defer wd.Quit()
    selen.Test1(wd)
```

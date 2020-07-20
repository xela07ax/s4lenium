# Selenium Go package Alpha from tebeka/selenium
Пакет создан для уменьшения количества повторений кода инициализации селениума, путем его переиспользования.
Порядок менять нельзя!
### Использование:
```go
import (
	"github.com/xela07ax/s4lenium"
)
// Создадим экземпляр
	selen := s4lenium.NewS4len(p.Loger)
	wd, err := selen.InitializeRemote(p.Config.SeleniumServer)
	tp.FckText(fmt.Sprintf("Создадим экземпляр селениума"),err)
	// Обязательно!
	//defer service.Stop()
	defer wd.Quit()
	selen.Test1(wd)
```

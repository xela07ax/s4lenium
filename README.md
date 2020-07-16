# Selenium Go package from tebekka
Usage

```go
// Создадим экземпляр
	selen := s4lenium.NewS4len(logEr.ChInLog)
	wd, service, err := selen.Initialize("D:\\Programs\\chromeDriver\\chromedriver_83.exe")
	tp.FckText("Создадим экземпляр селениума",err)
	// Обязательно!
	defer service.Stop()
	defer wd.Quit()
    selen.Test1(wd)
```

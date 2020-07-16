# Selenium Go package from tebekka
Usage

```go
// Создадим экземпляр
	s := s4lenium.NewS4len(logEr.ChInLog)
	wd, service, err := s.Initialize("D:\\Programs\\chromeDriver\\chromedriver_83.exe")
	tp.FckText(fmt.Sprintf("Создадим экземпляр %s",FullLogPath),err)
	// Обязательно!
	defer service.Stop()
	defer wd.Quit()
```

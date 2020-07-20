package s4lenium


import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

type Selen struct {
	NameFunc string
	Wd selenium.WebDriver
	Service *selenium.Service
	Loger chan <- [4]string
}
func NewS4len(Loger chan <- [4]string) *Selen{
	return &Selen{
		NameFunc: "S4len",
		Loger: Loger,
	}
}

const (
	//设置常量 分别设置chromedriver.exe的地址和本地调用端口
	//seleniumPath = `resources/chromedriver.exe`
	port         = 9515
	regionNum = "101"
	fio = "Иванов Иван Иванович"
	phone = "84950000000"
	email = "ivanov@ivan.ivanovich"
	shortDescription = "Тест скороти заведения заявки"
	detailedDescription = "Подробный текст: Тест скороти заведения заявки"
	dateTicket = "10.07.2020 22:14:04"
	filePath = "C:\\windows-version.txt"
)
func (s *Selen) InitializeLocal(selenPath string) (selenium.WebDriver,*selenium.Service, error) {
	//1.开启selenium服务
	//设置selium服务的选项,设置为空。根据需要设置。
	var sPath = "http://127.0.0.1:9515/wd/hub"
		ops := []selenium.ServiceOption{}
		service, err := selenium.NewChromeDriverService(selenPath, port, ops...)
		if err != nil {
			return nil, service, fmt.Errorf("Error starting the ChromeDriver server: %v", err)
		}
		//延迟关闭服务
		//defer service.Stop()
		wd, err := s.InitializeRemote(sPath)
		if err != nil {
			service.Stop()
			return nil,nil, fmt.Errorf("Error DefaultURLPrefix the ChromeDriver server: %v", err)
		}
		return wd, service, err
}
func (s *Selen) InitializeRemote(sPath string) (selenium.WebDriver,error) {
	caps := selenium.Capabilities{
		"browserName": "chrome",
		"enableVNC": true,
	}
	//调用浏览器urlPrefix: 测试参考：DefaultURLPrefix = "http://127.0.0.1:4444/wd/hub", "http://172.31.50.74:4444/wd/hub"

	wd, err := selenium.NewRemote(caps, sPath)
	//if err != nil {
	//	//service.Stop()
	//	return wd, fmt.Errorf("Error DefaultURLPrefix the ChromeDriver server: %v", err)
	//}
	//defer wd.Quit()
	return wd, err
}

func (s *Selen) Test1(wd selenium.WebDriver) error  {
	if err := wd.Get("http://192.168.99.106:3036/"); err != nil {
		return err
	}
	GetElement(selenium.ByID,"qf_region", wd).SendKeys(regionNum)
	GetElement(selenium.ByID,"qf_detailedDescription", wd).SendKeys(detailedDescription)
	GetElement(selenium.ByID,"qf_fio", wd).SendKeys(fio)
	GetElement(selenium.ByID,"qf_phone", wd).SendKeys(phone)
	GetElement(selenium.ByID,"qf_dateTicket", wd).SendKeys(dateTicket)
	GetElement(selenium.ByID,"qf_filePath", wd).SendKeys(filePath)
	GetElement(selenium.ByID,"qf_email", wd).SendKeys(email)
	GetElement(selenium.ByID,"qf_shortDescription", wd).SendKeys(shortDescription)
	time.Sleep(1*time.Minute)
	return nil
}

func GetElement(bySelector string, selector string, wd selenium.WebDriver) selenium.WebElement  {
	we, err := GetElementWithErr(bySelector,selector,wd)
	if err != nil {
		panic(err)
	}
	return we
}

func GetElementWithErr(bySelector string, selector string, wd selenium.WebDriver) (selenium.WebElement,error)  {
	var c selenium.Condition= func(wb selenium.WebDriver)(bool,error){
		//when i delete the line, it can wait
		fmt.Printf("Ищем:%s\n",selector)
		var err error
		_,err = wb.FindElement(bySelector, selector)
		if err != nil {
			return false,nil
		}
		return true,err
	}
	var err error
	err = wd.WaitWithTimeoutAndInterval(c,20*time.Second,1*time.Second)
	if err != nil {
		return nil,err
	}
	var we selenium.WebElement
	we,err = wd.FindElement(bySelector, selector)
	return we,err
}

func FindSetID(id string,val string, wd selenium.WebDriver) {
	//Find Baidu input box id
	we, err := wd.FindElement(selenium.ByID, id)
	if err != nil {
		panic(err)
	}
	//Send '' to input box
	err = we.SendKeys(val)
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)
}
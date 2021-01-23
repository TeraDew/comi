package common

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/mitchellh/go-homedir"
	"github.com/yumenaka/comi/locale"
	"image"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type ServerConfig struct  {
	OpenBrowser        bool   `json:"-"` //不要解析这个字段
	DisableLAN         bool   `json:"-"` //不要解析这个字段
	Template           string `json:"template"`
	Auth               string `json:"-"` //不要解析这个字段
	PrintAllIP         bool   `json:"-"` //不要解析这个字段
	Port               int
	ConfigPath         string `json:"-"` //不要解析这个字段
	CheckImageInServer bool
	LogToFile          bool   `json:"-"` //不要解析这个字段
	LogFilePath        string `json:"-"` //不要解析这个字段
	LogFileName        string `json:"-"` //不要解析这个字段
	MaxDepth           int    `json:"-"` //不要解析这个字段
	MinImageNum        int
	ServerHost         string
	EnableWebpServer    bool
	WebpConfig          WebPServerConfig `json:"-"` //不要解析这个字段
	EnableFrpcServer    bool
	FrpConfig           FrpClientConfig `json:"-"` //不要解析这个字段
	ZipFilenameEncoding string          `json:"-"` //不要解析这个字段
}

//通过路径名或执行文件名，来设置默认网页模板这个参数
func (config *ServerConfig) SetTemplateByName(FileName string){
	//如果执行文件名包含 comi或multi，设定为多页漫画模式
	if strings.Contains(FileName, "comi") || strings.Contains(FileName, "multi")  || strings.Contains(FileName, "多页"){
		config.Template= "multi"
		fmt.Println(locale.GetString("multi_page_template"))
	}
	//如果执行文件名包含 single，设定为 single 漫画模式
	if strings.Contains(FileName, "single")|| strings.Contains(FileName, "单页"){
		config.Template ="single"
		fmt.Println(locale.GetString("single_page_template"))
	}
	//如果执行文件名包含 sketch或croquis，设定为速写参考模式
	if strings.Contains(FileName, "sketch") || strings.Contains(FileName, "croquis")|| strings.Contains(FileName, "速写"){
		config.Template ="sketch"
		fmt.Println(locale.GetString("sketch_page_template"))
	}
	//如果用goland调试
	if strings.Contains(FileName, "build"){
		config.Template ="sketch"
		fmt.Println(locale.GetString("sketch_page_template"))
	}
}

var Config = ServerConfig{
	OpenBrowser:         true,
	DisableLAN:          false,
	Template:            "multi", //multi、single、random etc.
	Port:                1234,
	CheckImageInServer:  false,
	LogToFile:           false,
	MaxDepth:            2,
	MinImageNum:         3,
	ZipFilenameEncoding: "",
	WebpConfig: WebPServerConfig{
		WebpCommand:  "webp-server",
		HOST:         "127.0.0.1",
		PORT:         "3333",
		ImgPath:      "",
		QUALITY:      70,
		AllowedTypes: []string{"jpg", "png", "jpeg", "bmp"},
		ExhaustPath:  "",
	},
	EnableFrpcServer: false,
	FrpConfig: FrpClientConfig{
		FrpcCommand:      "frpc",
		ServerAddr:       "localhost", //server_addr
		ServerPort:       7000,        //server_port
		Token:            "&&%%!2356",
		FrpType:          "tcp",
		RemotePort:       -1, //remote_port
		RandomRemotePort: true,
		//AdminAddr:   "127.0.0.1",
		//AdminPort:   "12340",
		//AdminUser:   "",
		//AdminPwd :   "",
	},
	ServerHost: "",
}

var ReadingBook Book
var BookList []Book
var (
	//ReadFileName           string
	TempDir    string
	PictureDir string
	//PrintVersion    bool
	Version         string = "v0.2.4"
	SupportPicType         = [...]string{".png", ".jpg", ".jpeg", "bmp", ".gif", ".webp"}
	SupportFileType        = [...]string{
		".zip",
		".tar",
		".rar",
		".tar.gz",
		".tgz",
		".tar.bz2",
		".tbz2",
		".tar.xz",
		".txz",
		".tar.lz4",
		".tlz4",
		".tar.sz",
		".tsz",
		".bz2",
		".gz",
		".lz4",
		".sz",
		".xz"}
)

func init() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
	}
	Config.LogFilePath = home
	Config.LogFileName = "comigo.ini"
}

type Book struct {
	Name            string      `json:"name"`
	Author          string      `json:"author"`
	Title           string      `json:"title"`
	FilePath        string      `json:"-"` //不要解析这个字段
	PageNum         int         `json:"page_num"`
	PageInfo        []ImageInfo `json:"pages"`
	FileType        string      `json:"file_type"`
	FileSize        int64       `json:"file_size"`
	Modified        time.Time   `json:"modified_time"`
	UUID            string      `json:"uuid"`
	IsFolder        bool        `json:"is_folder"`
	ExtractNum      int         `json:"extract_num"`
	ExtractComplete bool        `json:"extract_complete"`
	ReadPercent     float64     `json:"read_percent"`
}

type ImageInfo struct {
	Height        int    `json:"height"`
	Width         int    `json:"width"`
	UrlPath       string `json:"url"`
	LocalPath     string `json:"-"` //不要解析这个字段
	InArchiveName string `json:"-"` //不要解析这个字段
	ImgType       string `json:"class"`
}

//一些绑定到Book结构体的方法
func (b *Book) SetArchiveBookName(name string) {
	post := strings.LastIndex(name, "/") //Unix路径分隔符
	if post == -1 {
		post = strings.LastIndex(name, "\\") //windows分隔符
	}
	if post != -1 {
		//name = string([]rune(name)[post:]) //为了防止中文字符被错误截断，先转换成rune，再转回来?
		name = name[post:]
		name = strings.ReplaceAll(name, "\\", "")
		name = strings.ReplaceAll(name, "/", "")
	}
	b.Name = name
}

func (b *Book) SetImageFolderBookName(name string) {
	b.Name = name
}

func (b *Book) SetPageNum() {
	//页数，目前只支持漫画
	b.PageNum = len(b.PageInfo)
}

func (b *Book) SetFilePath(path string) {
	b.FilePath = path
}

func (b *Book) GetName() string { //绑定到Book结构体的方法
	return b.Name
}

func (b *Book) GetPicNum() int {
	var PicNum = 0
	for _, p := range b.PageInfo {
		if checkPicExt(p.UrlPath) {
			PicNum++
		}
	}
	return PicNum
}

//服务器端分析单双页
func (b *Book) ScanAllImage() {
	log.Println(locale.GetString("check_image_start"))
	for i := 0; i < len(b.PageInfo); i++ { //此处不能用range，因为需要修改
		SetImageType(&b.PageInfo[i])
	}
	log.Println(locale.GetString("check_image_completed"))
}

//并发分析
func (b *Book) ScanAllImageGo() {
	//var wg sync.WaitGroup
	log.Println(locale.GetString("check_image_start"))
	res := make(chan string)
	count := 0
	extractNum := 0
	Percent := 0
	tempPercent := 0
	for i := 0; i < len(b.PageInfo); i++ { //此处不能用range，因为需要修改
		//wg.Add(1)
		count++
		//并发处理，提升图片分析速度
		go func(i int) {
			//defer wg.Done()
			SetImageType(&b.PageInfo[i])
			res <- fmt.Sprintf("Finished %d", i)
		}(i)
	}
	//wg.Wait()
	for i := 0; i < count; i++ {
		extractNum++
		if b.PageNum != 0 {
			Percent = int((float32(extractNum) / float32(b.PageNum)) * 100)
			if tempPercent != Percent {
				if (Percent%20) == 0 || Percent == 10 {
					fmt.Println(strconv.Itoa(Percent) + "% ")
				}
			}
			tempPercent = Percent
		}
		//fmt.Println(<-res)
		<-res
	}
	log.Println(locale.GetString("check_image_completed"))
}

func SetImageType(p *ImageInfo) {
	err := p.GetImageSize()
	//log.Println(locale.GetString("check_image_ing"), p.LocalPath)
	if err != nil {
		log.Println(locale.GetString("check_image_error") + err.Error())
	}
	if p.Width == 0 && p.Height == 0 {
		p.ImgType = "UnKnow"
		return
	}
	if p.Width > p.Height {
		p.ImgType = "Horizontal"
	} else {
		p.ImgType = "Vertical"
	}
}

//获取图片分辨率
func (i *ImageInfo) GetImageSize() (err error) {
	var img image.Image
	img, err = imaging.Open(i.LocalPath)
	if err != nil {
		log.Printf(locale.GetString("check_image_error")+" %v\n", err)
	} else {
		i.Width = img.Bounds().Dx()
		i.Height = img.Bounds().Dy()
	}
	return err
}

//中断处理：程序被中断的时候，清理临时文件
func SetupCloseHander() {
	c := make(chan os.Signal, 2)
	//SIGHUP（挂起）, SIGINT（中断）或 SIGTERM（终止）默认会使得程序退出。
	//1、SIGHUP 信号在用户终端连接(正常或非正常)结束时发出。
	//2、syscall.SIGINT 和 os.Interrupt 是同义词,按下 CTRL+C 时发出。
	//3、SIGTERM（终止）:kill终止进程,允许程序处理问题后退出。
	//4.syscall.SIGHUP,终端控制进程结束(终端连接断开)
	//5、syscall.SIGQUIT，CTRL+\ 退出
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	go func() {
		<-c
		fmt.Println("\r"+locale.GetString("start_clear_file"))
		deleteTempFiles()
		os.Exit(0)
	}()
}
func InitReadingBook() (err error) {
	//准备解压，设置图片文件夹
	if ReadingBook.IsFolder {
		PictureDir = ReadingBook.FilePath
		ReadingBook.ExtractComplete = true
		ReadingBook.ExtractNum = ReadingBook.PageNum
	} else {
		err = SetTempDir()
		if err != nil {
			fmt.Println(locale.GetString("temp_folder_error"), err)
			return err
		}
		PictureDir = TempDir
		err = ExtractArchive(&ReadingBook)
		if err != nil {
			fmt.Println(locale.GetString("file_not_found"))
			return err
		}
		ReadingBook.SetArchiveBookName(ReadingBook.FilePath) //设置书名
	}
	//服务器分析图片分辨率
	if Config.CheckImageInServer {
		ReadingBook.ScanAllImageGo() //扫描所有图片，取得分辨率信息，使用了协程
	}
	return err
}

//设置临时文件夹，退出时会被清理
func SetTempDir() (err error) {
	if TempDir != "" {
		return err
	}
	TempDir, err = ioutil.TempDir("", "comic_cache_A8cG")
	if err != nil {
		println(locale.GetString("temp_folder_create_error"))
	} else {
		fmt.Println(locale.GetString("temp_folder_path") + TempDir)
	}
	return err
}

func deleteTempFiles() {
	fmt.Println(locale.GetString("clear_temp_file_start"))
	if strings.Contains(TempDir, "comic_cache_A8cG") { //判断文件夹前缀，避免删错文件
		err := os.RemoveAll(TempDir)
		if err != nil {
			fmt.Println(locale.GetString("clear_temp_file_error") + TempDir)
		} else {
			fmt.Println(locale.GetString("clear_temp_file_completed") + TempDir)
		}
	}
	deleteOldTempFiles()
}

//根据权限，清理老文件可能失败
func deleteOldTempFiles() {
	tempDirUpperFolder := TempDir
	post := strings.LastIndex(TempDir, "/") //Unix风格的路径分隔符
	if post == -1 {
		post = strings.LastIndex(TempDir, "\\") //windows风格的分隔符
	}
	if post != -1 {
		tempDirUpperFolder = string([]rune(TempDir)[:post]) //为了防止中文字符被错误截断，先转换成rune，再转回来
		fmt.Println(locale.GetString("temp_folder_path"), tempDirUpperFolder)
	}
	files, err := ioutil.ReadDir(tempDirUpperFolder)
	if err != nil {
		fmt.Println(err)
	}
	for _, fi := range files {
		if fi.IsDir() {
			oldTempDir := tempDirUpperFolder + "/" + fi.Name()
			if strings.Contains(oldTempDir, "comic_cache_A8cG") { //判断文件夹前缀，避免删错文件
				err := os.RemoveAll(oldTempDir)
				if err != nil {
					fmt.Println(locale.GetString("clear_temp_file_error") + oldTempDir)
				} else {
					fmt.Println(locale.GetString("clear_temp_file_completed") + oldTempDir)
				}
			}
		}
	}
}

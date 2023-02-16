package v1

import (
	"bufio"
	"errors"
	"fiber-nuzn-api/controllers"
	"fiber-nuzn-api/pkg/utils"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UploadController struct {
	controllers.Base
}

func NewUploadController() *UploadController {
	return &UploadController{}
}

func (t *UploadController) SaveUpload(ctx *fiber.Ctx) error {
	// 接收文件file
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	// 获取文件后缀
	extName := path.Ext(file.Filename)
	// 拼接文件路径
	err, pathDir := Mkdir(extName, "")
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	// 保存文件
	if err := ctx.SaveFile(file, pathDir); err != nil {
		return ctx.JSON(t.Fail(err))
	}
	return ctx.JSON(t.Ok(pathDir))
}

// MergeFile 合并切片
func (t *UploadController) MergeFile(ctx *fiber.Ctx) error {
	// hash值（区分当前文件是那个的，也可以用uuid，nanoid，等）
	fileId := ctx.FormValue("fileId")
	fileIndex := ctx.FormValue("fileIndex")
	fileName := ctx.FormValue("fileName")
	// 获取文件后缀
	extName := path.Ext(fileName)
	atom, err := strconv.Atoi(fileIndex)
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	p := "static/upload/" + fileId + extName
	newFile, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0766)
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	index := 0
	for {
		if atom == index {
			break
		}
		// 文件流 存的路径 ：public/blob_1lV4DJWf2qs8MdQojPMwb_1
		filePath := "static/" + "blob" + "_" + fileId + "_" + strconv.Itoa(index)
		f, _ := os.Open(filePath)
		r := bufio.NewReader(f)
		data := make([]byte, 1024, 1024)
		for {
			total, err := r.Read(data)
			if err == io.EOF {
				f.Close()
				os.Remove(filePath)
				break
			}
			_, err = newFile.Write(data[:total])
		}
		index++
	}
	defer func() {
		newFile.Close()
	}()
	return ctx.JSON(t.Ok(map[string]interface{}{
		"msg":  "合并成功",
		"切片序号": p,
	}))
}

// ChunkFile 上传切片
func (t *UploadController) ChunkFile(ctx *fiber.Ctx) error {
	// 文件名名称
	//fileName := ctx.FormValue("fileName")
	// hash值（区分当前文件是那个的，也可以用uuid，nanoid，等）
	fileId := ctx.FormValue("fileId")
	fileIndex := ctx.FormValue("fileIndex")
	// 接收文件的file 分片
	file, err := ctx.FormFile("fileChunk")
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	// 文件流 存的路径 ：public/blob_1lV4DJWf2qs8MdQojPMwb_1
	filePath := "static/" + file.Filename + "_" + fileId + "_" + fileIndex
	// 转成file
	upFile, _ := file.Open()
	// 创建文件
	fileBool, err := createFile(filePath, upFile)
	if !fileBool {
		return ctx.JSON(t.Fail(err))
	}
	return ctx.JSON(t.Ok(map[string]interface{}{
		"msg":  "上传成功",
		"切片序号": fileIndex,
	}))
}

// 创建文件
func createFile(filePath string, upFile multipart.File) (bool, error) {
	fileBool, err := fileExists(filePath)
	if fileBool && err == nil {
		return true, errors.New("文件以存在")
	} else {
		newFile, err := os.Create(filePath)
		data := make([]byte, 1024, 1024)
		for {
			total, err := upFile.Read(data)
			if err == io.EOF {
				break
			}
			_, err = newFile.Write(data[:total])
			if err != nil {
				return false, errors.New("文件上传失败")
			}
		}
		defer newFile.Close()
		if err != nil {
			return false, errors.New("创建文件失败")
		}
	}
	return true, nil
}

// 判断文件或文件夹是否存在
func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

/*
效验文件后缀
extName：文件后缀
extMap：效验后缀
返回值：布尔值
*/

func allExtMap(extMap []string, extName string) bool {
	allowExtMap := make(map[string]bool)
	for _, val := range extMap {
		allowExtMap[val] = true
	}
	// 判断excel上传是否合法
	if _, ok := allowExtMap[extName]; !ok {

		return false
	}
	return true
}

/**
创建文件夹/文件名
extName：文件后缀
route：设置特定目录后缀
返回值：bool,路径
*/

func Mkdir(extName, route string) (error, string) {
	// 组成 文件路径
	dir := "static/upload/" + utils.GetDay() + route
	// 创建文件路径
	if err := os.MkdirAll(dir, 0666); err != nil {
		return err, ""
	}
	//生成文件名称   144325235235.png
	fileUnixName := strconv.FormatInt(utils.GetUnixNano(), 10)
	saveDir := path.Join(dir, fileUnixName+extName)
	return nil, saveDir
}

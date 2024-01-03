package utils

import (
	"archive/zip"
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	"zrDispatch/common/log"
	"zrDispatch/core/slog"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// 换行的数据
func ReadLineData(userDict string) (users []string, err error) {
	file, err := os.Open(userDict)
	if err != nil {
		log.Error("router.Run error", zap.Error(err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		user := strings.TrimSpace(scanner.Text())
		if user != "" {
			users = append(users, user)
		}
	}
	return users, err
}

func Write(path, str string) {
	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR+os.O_TRUNC, 0764)
	if err != nil {
		log.Error("router.Run error", zap.Error(err))
	}

	//jsonBuf := append([]byte(result),[]byte("\r\n")...)
	f.WriteString(str)
}
func WriteAppend(path, str string) {
	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR+os.O_APPEND, 0764)
	if err != nil {
		log.Error("router.Run error", zap.Error(err))
	}

	//jsonBuf := append([]byte(result),[]byte("\r\n")...)
	str += "\n"
	f.WriteString(str)
}

func WriteAppendHh(path, str string) {
	f, err := os.OpenFile(path, os.O_CREATE+os.O_RDWR+os.O_APPEND, 0764)
	if err != nil {
		log.Error("router.Run error", zap.Error(err))
	}

	//jsonBuf := append([]byte(result),[]byte("\r\n")...)
	str += "\n"
	f.WriteString(str)
}

func Read(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Error("router.Run error", zap.Error(err))
	}
	return string(content)
}

func ReadLast(path string) string {
	cmd := exec.Command("tail", "-1", path)

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Println(slog.DEBUG, string(out))
		slog.Println(slog.DEBUG, err)
	}

	return string(out)
}

// 压缩为zip格式
// source为要压缩的文件或文件夹, 绝对路径和相对路径都可以
// target是目标文件
// filter是过滤正则(Golang 的 包 path.Match)
func ZipFile(source, target, timeStr string) error {
	var err error
	if isAbs := filepath.IsAbs(source); !isAbs {
		source, err = filepath.Abs(source) // 将传入路径直接转化为绝对路径
		if err != nil {
			return errors.WithStack(err)
		}
	}
	//创建zip包文件
	zipfile, err := os.Create(target)
	if err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		if err := zipfile.Close(); err != nil {
			log.Error("router.Run error", zap.Error(err))
		}
	}()

	//创建zip.Writer
	zw := zip.NewWriter(zipfile)

	defer func() {
		if err := zw.Close(); err != nil {
			log.Error("router.Run error", zap.Error(err))
		}
	}()

	info, err := os.Stat(source)
	if err != nil {
		return errors.WithStack(err)
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return errors.WithStack(err)
		}

		//将遍历到的路径与pattern进行匹配
		//println(filter, info.Name())
		//ism, err := filepath.Match(filter, info.Name())
		//
		//if err != nil {
		//	return errors.WithStack(err)
		//}
		////如果匹配就忽略
		//if ism {
		//	return nil
		//}

		GetTime()
		t, _ := time.Parse("2006-01-02 15:04:05", timeStr)
		if timeStr != "" && info.ModTime().Before(t) {
			return nil
		}
		println(info.Name())
		//创建文件头
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return errors.WithStack(err)
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		//写入文件头信息
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return errors.WithStack(err)
		}

		if info.IsDir() {
			return nil
		}
		//写入文件内容
		file, err := os.Open(path)
		if err != nil {
			return errors.WithStack(err)
		}

		defer func() {
			if err := file.Close(); err != nil {
				log.Error("router.Run error", zap.Error(err))
			}
		}()
		_, err = io.Copy(writer, file)

		return errors.WithStack(err)
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// 获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

// 获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(root string) (files []string, err error) {

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	return files, err
}

func GetFiles(root, name string) (files []string, err error) {

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if GetStrACount(name, path) > 0 {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

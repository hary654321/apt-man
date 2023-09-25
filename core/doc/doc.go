package doc

import (
	"log"
	"os"
	"time"
	"zrDispatch/common/utils"
	"zrDispatch/core/utils/define"

	// "strconv"

	"github.com/nguyenthenguyen/docx"
)

type biaozhi struct {
	yeMeiTime  string // 页眉的时间编号
	neiRouTime string // 内容的时间格式

	sjmc   string   // 事件名称
	sjjb   string   // 事件级别（漏洞、风险、应急）
	ipaddr string   // 告警对象（应急）
	gjlj   []string // 告警路径（应急、风险）
	tms    string   // 告警数量（应急）
	gjxx   string   // 告警描述（漏洞、风险）
	czjy   []string // 处置建议(漏洞、风险)
	sjly   string   // 事件类型(应急)
	llmc   string   // 漏洞名称(漏洞)
	yxbb   []string // 影响版本(漏洞)
	ywxt   string   // 业务系统（应急、风险）
	ywdw   string   // 业务单位（漏洞、应急、风险）
	fxms   string   // 风险验证描述（风险）
	kyfx   string   // 可疑分析（可疑）

	czjyStr   string // 写入excel的处置建立
	lujingStr string // 写入excel中的路径
	lcw       string // 另存为文件的名称
}

var (
	// Baogao 要打开的报告类型
	Baogao string
	// SaveFile 要保存的报告类型
	SaveFile string
	dataList biaozhi
	// qiantou这个比较重要。
	/*
		我不知道什么原因，这两个参数我已经第二次修改了，可能是因为docx的问题，所以建议如果你要动程序之前，最好先自己测试一下docx的这两个参数对不对。直接解压docx看就可以了。
	*/
	qiantou string = `<w:p><w:pPr><w:rPr><w:rFonts w:hint="eastAsia" w:ascii="仿宋_GB2312" w:hAnsi="仿宋_GB2312" w:eastAsia="仿宋_GB2312" w:cs="仿宋_GB2312"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="zh-CN"/></w:rPr></w:pPr><w:r><w:rPr><w:rFonts w:hint="eastAsia" w:ascii="仿宋_GB2312" w:hAnsi="仿宋_GB2312" w:eastAsia="仿宋_GB2312" w:cs="仿宋_GB2312"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="zh-CN"/></w:rPr><w:t>`
	jiewei  string = `</w:t></w:r><w:bookmarkStart w:id="0" w:name="_GoBack"/><w:bookmarkEnd w:id="0"/></w:p>`
)

// 1.判断输出的是什么通告
// func tonggao() {
// 	cl, err := excelize.OpenFile("素材.xlsx")
// 	if err != nil {
// 		log.Println("打开素材文件错误！")
// 		return
// 	}
// 	SaveFile, _ = cl.GetCellValue("需要填写的资料", "B1")
// 	switch SaveFile {
// 	case "可疑行为":
// 		Baogao = "安全可疑行为通告模板.docx"
// 	case "风险预警":
// 		Baogao = "安全风险预警通告模板.docx"
// 	case "应急预警":
// 		Baogao = "安全事件应急通告模板.docx"
// 	case "漏洞预警":
// 		Baogao = "安全漏洞预警通告模板.docx"
// 	default:
// 		log.Println("请检查(素材.xlsx)文件（需要填写的资料）表中报告类型是否正确！")
// 	}

// 	// 获取相关的数据
// 	dataList.sjmc, _ = cl.GetCellValue("需要填写的资料", "B2")   // 风险预警：获取事件名称
// 	dataList.sjjb, _ = cl.GetCellValue("需要填写的资料", "B3")   // 所有预警：获取事件级别
// 	dataList.ipaddr, _ = cl.GetCellValue("需要填写的资料", "B4") // 应急预警：告警对象
// 	gjlj, _ := cl.GetCellValue("需要填写的资料", "B5")
// 	dataList.gjlj = strings.Split(gjlj, "\n")           // 告警路径
// 	dataList.tms, _ = cl.GetCellValue("需要填写的资料", "B6")  // 告警数量
// 	dataList.gjxx, _ = cl.GetCellValue("需要填写的资料", "B7") // 告警描述
// 	dataList.czjyStr, _ = cl.GetCellValue("需要填写的资料", "B8")
// 	dataList.czjy = strings.Split(dataList.czjyStr, "\n") // 所有预警：获取处置建议
// 	dataList.sjly, _ = cl.GetCellValue("需要填写的资料", "B9")   // 应急预警：事件类型
// 	dataList.llmc, _ = cl.GetCellValue("需要填写的资料", "B10")  // 漏洞预警：漏洞名称
// 	yxbb, _ := cl.GetCellValue("需要填写的资料", "B11")
// 	dataList.yxbb = strings.Split(yxbb, "\n")            // 漏洞预警：影响版本
// 	dataList.ywxt, _ = cl.GetCellValue("需要填写的资料", "B12") // 风险预警：业务系统名称
// 	dataList.ywdw, _ = cl.GetCellValue("需要填写的资料", "B13") // 风险预警：业务单位名称
// 	dataList.fxms, _ = cl.GetCellValue("需要填写的资料", "B14") // 风险预警：风险验证描述
// 	dataList.kyfx, _ = cl.GetCellValue("需要填写的资料", "B15") // 可疑行为：可疑分析

// }

// 2.处理并输出报告
func ExportDoc(task *define.DetailTask, data map[string]interface{}) {
	// 打开一个已有格式的文档，这个是要打开的文档路径。
	filesDir := "report.docx"
	// 添加页眉
	r, err := docx.ReadDocxFile(filesDir)
	if err != nil {
		log.Printf("打开word文档错误，错误信息: %s", err)
	}
	log.Println("打开报告完成")
	defer r.Close()
	docx1 := r.Editable()

	yeMeiTime := time.Now().Format("200601021504")
	docx1.ReplaceHeader("页眉日期", yeMeiTime)

	//任务信息
	docx1.Replace("task.name", task.Name, -1)
	docx1.Replace("task.ip", task.Ip, -1)
	docx1.Replace("task.port", task.Port, -1)
	docx1.Replace("task.create_time", utils.GetInterfaceToString(task.Ctime), -1)
	docx1.Replace("task.update_time", utils.GetInterfaceToString(task.Utime), -1)

	//汇总信息
	docx1.Replace("ip_count", utils.GetInterfaceToString(data["ip_count"]), -1)
	docx1.Replace("live_ip_count", utils.GetInterfaceToString(data["live_ip_count"]), -1)
	docx1.Replace("port_count", utils.GetInterfaceToString(data["port_count"]), -1)
	docx1.Replace("live_port", utils.GetInterfaceToString(data["live_port"]), -1)
	docx1.Replace("ip_count", utils.GetInterfaceToString(data["ip_count"]), -1)
	docx1.Replace("match_ip_count", utils.GetInterfaceToString(data["match_ip_count"]), -1)

	sp := "./report/" + task.Name + ".docx"
	err = docx1.WriteToFile(sp)
	if err != nil {
		os.Mkdir("./report", os.ModePerm)
		log.Println("当前目录中没有report报告文件夹，现在创建文件夹成功！！！")
		docx1.WriteToFile(sp)
		log.Println("保存docx文件成功！！！")
	}
	log.Println("保存docx文件成功！！！")
}

// func exc() {
// 	f, err := excelize.OpenFile("预警通报汇总表.xlsx")
// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	//告警函编号、业务单位、业务系统、告警状态、漏洞名称、告警类型、告警级别、告警对象、告警描述、应急处置方法、发生时间、结束时间、告警根治处置方法建议、安全处置联系人、支撑人员、修复方法描述、修复时间、修复验证描述、修复验证人、修复验证时间

// 	// 非应急预警的格式
// 	exStr := []string{dataList.yeMeiTime, dataList.ywdw, dataList.ywxt, "未修复", dataList.llmc, dataList.事件名称, dataList.sjjb, dataList.ipaddr, dataList.gjxx, "", dataList.neiRouTime, "", dataList.czjyStr, "", "xxx", "", "", "", dataList.ywdw, SaveFile}

// 	// 检查A列一千行，判断是否有值，如果有值就不写。没有值就整行填写记录
// 	for n1 := 1; n1 < 1000; n1++ {
// 		s1 := strconv.Itoa(n1)
// 		lie := "A" + s1
// 		// 获取当前表，固定表格的值
// 		cell, _ := f.GetCellValue("Sheet1", lie)
// 		if cell == "" { // 判断是否为空
// 			// 为空就提交单元格编号并跳出循环
// 			f.SetSheetRow("Sheet1", lie, &exStr)
// 			// celSave = lie
// 			break
// 		}
// 	}
// 	// 保存文件
// 	f.Save()
// 	log.Println("通告内容已经写入“预警通报汇总表.xlsx”表中！！！")
// }

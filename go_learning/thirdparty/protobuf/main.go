package main

import (
	"fmt"
	// "io/ioutil"
	// jsonStruct "go_learning/thirdparty/protobuf/report_struct/json_struct"
	protobufStruct "go_learning/thirdparty/protobuf/report_struct/protobuf_struct"
	"go_learning/thirdparty/protobuf/spm"
	"os"

	// "time"

	// "encoding/json"

	"github.com/golang/protobuf/proto"
	// "github.com/pquerna/ffjson/ffjson"
)

// -----------------------------------
// protobuf data len: 578
// proto.Marshal cost: 815.045051ms
// proto.Unmarshal cost: 1.110265959s
// -----------------------------------
// ffjson data len: 818
// ffjson.Marshal cost: 2.46182064s
// ffjson.Unmarshal cost: 3.417351879s
// -----------------------------------
// json data len: 818
// json.Marshal cost: 8.623781071s
// json.Unmarshal cost: 9.524305675s
// -----------------------------------

func main() {
	spmDemo()
	// protoBufDemo()
	// jsonDemo()
	// protobufWriteFile()
}

func spmDemo() {
	pb := &spm.ReportInputParams{
		Events: []*spm.Event{
			{
				UserId:     "",
				ClientTime: 1606752000000,
				Session:    "",
				EventId:    999,
				EventType:  "VIEW",
				Utm:        "",
				Spm:        "MiShop_A.Home.transformer.0",
				Spmref:     "MiShop_A.Home.transformer.0",
				Scm:        "m1.p1.a1.av1.c1.p1.ext",
				Scmref:     "m1.p1.a1.av1.c1.p1.ext",
				Uri:        "https%253a%252f%252fwww.mi.com%252fbuy%252fdetail%253fproduct_id%253d12609",
				Uriref:     "https%253a%252f%252fwww.mi.com%252fbuy%252fdetail%253fproduct_id%253d12609",
				Params:     "",
			},
			{
				UserId:     "",
				ClientTime: 1606752000000,
				Session:    "",
				EventId:    998,
				EventType:  "VIEW",
				Utm:        "",
				Spm:        "MiShop_A.Home.transformer.0",
				Spmref:     "MiShop_A.Home.transformer.0",
				Scm:        "m1.p1.a1.av1.c1.p1.ext",
				Scmref:     "m1.p1.a1.av1.c1.p1.ext",
				Uri:        "https%253a%252f%252fwww.mi.com%252fbuy%252fdetail%253fproduct_id%253d12609",
				Uriref:     "https%253a%252f%252fwww.mi.com%252fbuy%252fdetail%253fproduct_id%253d12609",
				Params:     "",
			},
		},
		UserAgent: "ua%253dxxxxx%255cu0001app_channel%253dxxx%255cu0001app_version%253d10",
	}
	
	pbByte, err := proto.Marshal(pb)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("data:%s, err:%s\n", pbByte, err)
	book := spm.ReportInputParams{}
	if err := proto.Unmarshal(pbByte, &book); err != nil {
	}


	fmt.Printf("data2:%+v\n", &book)

}

// func protoBufDemo() {
// 	var data []byte

// 	params := protobufStruct.ReportInputParams{
// 		Data: []*protobufStruct.ReportItem{
// 			&protobufStruct.ReportItem {
// 				UserId     : "TMP1601481600000000000001",
// 				ClientTime : 1601481600000,
// 				Session    : "1234123_asdfwerfasd_123412341234",
// 				EventId    : 1,
// 				EventType  : "TOUCH",
// 				Event      : &protobufStruct.Event{
// 					Utm    : "BiMiHome_A.HomeShop.back.0",
// 					Scm    : "BiMiHome_A.HomeShop.back.0",
// 					Spm    : "BiMiHome_A.HomeShop.back.0",
// 					Spmref : "BiMiHome_A.HomeShop.back.0",
// 					Uri    : "https://xiaomi.f.mioffice.cN/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
// 					Uriref : "https://xiaomi.f.mioffice.cn/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
// 					Params : "key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value",
// 				},
// 				UserAgent  : "ua=value&os_name=value&os_version=value&network=value&ip=value&app_name=value&app_version=value&app_channel=value&model=value&cgi=value",
// 			},
// 		},
// 	}

// 	t1 := time.Now()
// 	for i := 0; i < 1000000; i++ {
// 		data, _ = proto.Marshal(&params)
// 	}
// 	t2 := time.Now()
// 	fmt.Println("----------------------------------")
// 	fmt.Printf("protobuf data len: %d\n", len(data))
// 	fmt.Printf("proto.Marshal cost: %s\n", t2.Sub(t1))

// 	params2 := protobufStruct.ReportInputParams{}
// 	t1 = time.Now()
// 	for i := 0; i < 1000000; i++ {
// 		proto.Unmarshal(data, &params2)
// 	}
// 	t2 = time.Now()
// 	fmt.Printf("proto.Unmarshal cost: %s\n", t2.Sub(t1))
// }

// func jsonDemo() {
// 	var data []byte

// 	params := jsonStruct.ReportInputParams{
// 		Data: []jsonStruct.ReportItem{
// 			jsonStruct.ReportItem {
// 				UserId     : "TMP1601481600000000000001",
// 				ClientTime : 1601481600000,
// 				Session    : "1234123_asdfwerfasd_123412341234",
// 				EventId    : 1,
// 				EventType  : "TOUCH",
// 				Event      : jsonStruct.Event{
// 					Utm    : "BiMiHome_A.HomeShop.back.0",
// 					Scm    : "BiMiHome_A.HomeShop.back.0",
// 					Spm    : "BiMiHome_A.HomeShop.back.0",
// 					Spmref : "BiMiHome_A.HomeShop.back.0",
// 					Uri    : "https://xiaomi.f.mioffice.cN/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
// 					Uriref : "https://xiaomi.f.mioffice.cn/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
// 					Params : "key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value",
// 				},
// 				UserAgent  : "ua=value&os_name=value&os_version=value&network=value&ip=value&app_name=value&app_version=value&app_channel=value&model=value&cgi=value",
// 			},
// 		},
// 	}

// 	t1 := time.Now()
// 	for i := 0; i < 1000000; i++ {
// 		data, _ = ffjson.Marshal(&params)
// 	}
// 	t2 := time.Now()
// 	fmt.Println("----------------------------------")
// 	fmt.Printf("ffjson data len: %d\n", len(data))
// 	fmt.Printf("ffjson.Marshal cost: %s\n", t2.Sub(t1))

// 	params2 := jsonStruct.ReportInputParams{}
// 	t1 = time.Now()
// 	for i := 0; i < 1000000; i++ {
// 		ffjson.Unmarshal(data, &params2)
// 	}
// 	t2 = time.Now()
// 	fmt.Printf("ffjson.Unmarshal cost: %s\n", t2.Sub(t1))
	
// 	//////////////////////////////////////////////////////////////////////

// 	t1 = time.Now()
// 	for i := 0; i < 1000000; i++ {
// 		data, _ = json.Marshal(&params)
// 	}
// 	t2 = time.Now()
// 	fmt.Println("----------------------------------")
// 	fmt.Printf("json data len: %d\n", len(data))
// 	fmt.Printf("json.Marshal cost: %s\n", t2.Sub(t1))
// 	t1 = time.Now()
// 	for i := 0; i < 1000000; i++ {
// 		json.Unmarshal(data, &params2)
// 	}
// 	t2 = time.Now()
// 	fmt.Printf("json.Unmarshal cost: %s\n", t2.Sub(t1))
// }


func protobufWriteFile() {
	data := protobufStruct.ReportInputParams{
		UserId     : "AW0wGfyZyLyTglEwiQjtMw==",
		ClientTime : 1607574892929,
		Session    : "1607574891",
		EventType  : "VISIBLE",
		Event      : []*protobufStruct.Event{
			{
				EventId : 1234124,
				Utm     : "asdfaf",
				Scm     : "asd;lfal;sdfj;asdf",
				Spm     : "MiShop_I.Home.tab.1",
				Spmref  : "https://www.baidu.com/s?wd=asdfasdfadsf",
				Uri     : "https://www.baidu.com/s?wd=asdfasdfadsf",
				Uriref  : "https://www.baidu.com/s?wd=asdfasdfadsf",
				Params  : "tab%3D%25E9%25A6%2596%25E9%25A1%25B5",
			},
		},
		UserAgent  : "ua%3DMozilla%252F5.0%2520%2528iPhone%253B%2520CPU%2520iPhone%2520OS%252014_2%2520like%2520Mac%2520OS%2520X%2529%2520AppleWebKit%252F605.1.15%2520%2528KHTML%252C%2520like%2520Gecko%2529%2520Mobile%252F15E148%252FXiaoMi%252FMiuiBrowser%252F4.3%252FShop%252Fios%252Fx86_64%252F5.1.99%2520APP%252Fcom.xiaomi.mishop.dailybuild%2520APPV%252F5.1.98%2520iosPassportSDK%252F3.9.5%2520iOS%252F14.2%5Cu0001device_id%3D4377CE1030FC50FBAD113FC9A57DEAE8",
	}
	for len(data.Event) < 30 {
		data.Event = append(data.Event, data.Event[0])
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		fmt.Printf("proto.Marshal failed, err: %s", err)
		return 
	}
	file, err := os.OpenFile("./report_protobuf", os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open file failed, err: %s", err)
		return 
	}
	nBytes, err := file.Write(dataBytes)
	if err != nil {
		fmt.Printf("Write file failed, err: %s", err)
		return 
	}
	fmt.Printf("Write %d bytes.", nBytes)
}

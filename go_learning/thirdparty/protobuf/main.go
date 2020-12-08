package main

import (
	"fmt"
	jsonStruct "go_learning/thirdparty/protobuf/report_struct/json_struct"
	protobufStruct "go_learning/thirdparty/protobuf/report_struct/protobuf_struct"
	"os"
	"time"

	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/pquerna/ffjson/ffjson"
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
	// protoBufDemo()
	// jsonDemo()
	protobufWriteFile()
}

func protoBufDemo() {
	var data []byte

	params := protobufStruct.ReportInputParams{
		Data: []*protobufStruct.ReportItem{
			&protobufStruct.ReportItem {
				UserId     : "TMP1601481600000000000001",
				ClientTime : 1601481600000,
				Session    : "1234123_asdfwerfasd_123412341234",
				EventId    : 1,
				EventType  : "TOUCH",
				Event      : &protobufStruct.Event{
					Utm    : "BiMiHome_A.HomeShop.back.0",
					Scm    : "BiMiHome_A.HomeShop.back.0",
					Spm    : "BiMiHome_A.HomeShop.back.0",
					Spmref : "BiMiHome_A.HomeShop.back.0",
					Uri    : "https://xiaomi.f.mioffice.cN/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
					Uriref : "https://xiaomi.f.mioffice.cn/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
					Params : "key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value",
				},
				UserAgent  : "ua=value&os_name=value&os_version=value&network=value&ip=value&app_name=value&app_version=value&app_channel=value&model=value&cgi=value",
			},
		},
	}

	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
		data, _ = proto.Marshal(&params)
	}
	t2 := time.Now()
	fmt.Println("----------------------------------")
	fmt.Printf("protobuf data len: %d\n", len(data))
	fmt.Printf("proto.Marshal cost: %s\n", t2.Sub(t1))

	params2 := protobufStruct.ReportInputParams{}
	t1 = time.Now()
	for i := 0; i < 1000000; i++ {
		proto.Unmarshal(data, &params2)
	}
	t2 = time.Now()
	fmt.Printf("proto.Unmarshal cost: %s\n", t2.Sub(t1))
}

func jsonDemo() {
	var data []byte

	params := jsonStruct.ReportInputParams{
		Data: []jsonStruct.ReportItem{
			jsonStruct.ReportItem {
				UserId     : "TMP1601481600000000000001",
				ClientTime : 1601481600000,
				Session    : "1234123_asdfwerfasd_123412341234",
				EventId    : 1,
				EventType  : "TOUCH",
				Event      : jsonStruct.Event{
					Utm    : "BiMiHome_A.HomeShop.back.0",
					Scm    : "BiMiHome_A.HomeShop.back.0",
					Spm    : "BiMiHome_A.HomeShop.back.0",
					Spmref : "BiMiHome_A.HomeShop.back.0",
					Uri    : "https://xiaomi.f.mioffice.cN/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
					Uriref : "https://xiaomi.f.mioffice.cn/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
					Params : "key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value",
				},
				UserAgent  : "ua=value&os_name=value&os_version=value&network=value&ip=value&app_name=value&app_version=value&app_channel=value&model=value&cgi=value",
			},
		},
	}

	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
		data, _ = ffjson.Marshal(&params)
	}
	t2 := time.Now()
	fmt.Println("----------------------------------")
	fmt.Printf("ffjson data len: %d\n", len(data))
	fmt.Printf("ffjson.Marshal cost: %s\n", t2.Sub(t1))

	params2 := jsonStruct.ReportInputParams{}
	t1 = time.Now()
	for i := 0; i < 1000000; i++ {
		ffjson.Unmarshal(data, &params2)
	}
	t2 = time.Now()
	fmt.Printf("ffjson.Unmarshal cost: %s\n", t2.Sub(t1))
	
	//////////////////////////////////////////////////////////////////////

	t1 = time.Now()
	for i := 0; i < 1000000; i++ {
		data, _ = json.Marshal(&params)
	}
	t2 = time.Now()
	fmt.Println("----------------------------------")
	fmt.Printf("json data len: %d\n", len(data))
	fmt.Printf("json.Marshal cost: %s\n", t2.Sub(t1))
	t1 = time.Now()
	for i := 0; i < 1000000; i++ {
		json.Unmarshal(data, &params2)
	}
	t2 = time.Now()
	fmt.Printf("json.Unmarshal cost: %s\n", t2.Sub(t1))
}


func protobufWriteFile() {
	data := protobufStruct.ReportInputParams{
		Data: []*protobufStruct.ReportItem{
			&protobufStruct.ReportItem {
				UserId     : "TMP1601481600000000000001",
				ClientTime : 1601481600000,
				Session    : "1234123_asdfwerfasd_123412341234",
				EventId    : 1,
				EventType  : "TOUCH",
				Event      : &protobufStruct.Event{
					Utm    : "BiMiHome_A.HomeShop.back.0",
					Scm    : "BiMiHome_A.HomeShop.back.0",
					Spm    : "BiMiHome_A.HomeShop.back.0",
					Spmref : "BiMiHome_A.HomeShop.back.0",
					Uri    : "https://xiaomi.f.mioffice.cN/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
					Uriref : "https://xiaomi.f.mioffice.cn/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
					Params : "key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value",
				},
				UserAgent  : "ua=value&os_name=value&os_version=value&network=value&ip=value&app_name=value&app_version=value&app_channel=value&model=value&cgi=value",
			},
			&protobufStruct.ReportItem {
				UserId     : "TMP1601481600000000000001",
				ClientTime : 1601481600000,
				Session    : "1234123_asdfwerfasd_123412341234",
				EventId    : 1,
				EventType  : "TOUCH",
				Event      : &protobufStruct.Event{
					Utm    : "BiMiHome_A.HomeShop.back.0",
					Scm    : "BiMiHome_A.HomeShop.back.0",
					Spm    : "BiMiHome_A.HomeShop.back.0",
					Spmref : "BiMiHome_A.HomeShop.back.0",
					Uri    : "https://xiaomi.f.mioffice.cN/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
					Uriref : "https://xiaomi.f.mioffice.cn/docs/dock4Jm5lAGXGILv45dsFrYyrX1?sidebarOpen=1",
					Params : "key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value&key=value",
				},
				UserAgent  : "ua=value&os_name=value&os_version=value&network=value&ip=value&app_name=value&app_version=value&app_channel=value&model=value&cgi=value",
			},
		},
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

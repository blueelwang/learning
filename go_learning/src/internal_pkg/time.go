package libs

import (
	"fmt"
	"time"
)

func TimeDemo() {
	// 时间输出，以 2006-01-02 15:04:05 作为格式
	// 2006 代表年，01:月（两位），02:日（两位）， 15:小时（两位，24小时表示）， 04:分钟（两位）， 05:秒（两位）
	// 1:月，2:日， 3:小时， 4:分钟， 5:秒
	t := time.Now()
	fmt.Printf("%+v\n", t)							// 2020-06-08 19:47:02.263471 +0800 CST m=+0.000140413
	fmt.Println(t.Unix())							// 1591616822
	fmt.Println(t.Year())							// 2020
	fmt.Println(t.Month())							// June
	fmt.Println(int(t.Month()))						// 6
	fmt.Println(t.Day())							// 8
	fmt.Println(t.Hour())							// 19
	fmt.Println(t.Minute())							// 47
	fmt.Println(t.Second())							// 2
	fmt.Println(t.Weekday())						// Monday
	fmt.Println(int(t.Weekday()))					// 1
	fmt.Println(t.Format("2006-01-02 15:04:05"))	// 2020-06-08 19:47:02
	fmt.Println(t.Format("2006-1-2 3:4:5"))			// 2020-6-8 7:47:2

	// 通过字符串解析时间，解析出错返回err
	t, err := time.ParseInLocation("2006-01-02", "2020-02-27", time.Local)
	fmt.Println(t.Unix());	// 1582732800
	fmt.Println(t.AddDate(-1, 0, 0).Format("2006-01-02"), err)	// 2019-02-27 <nil>

	t, err = time.ParseInLocation("2006-01-02", "2020-02-30", time.Local)
	fmt.Println(t.Format("2006-01-02"), err)	// 0001-01-01 parsing time "2020-02-30": day out of range

	// 年，月，日 加减计算，如果结果不存在，会向后顺延，err为nil，如：2019-02-29，实际表示的是 2019-03-01
	t, err = time.ParseInLocation("2006-01-02", "2020-02-29", time.Local)
	fmt.Println(t.AddDate(-1, 0, 0).Format("2006-01-02"), err)	// 2019-03-01 <nil>

	t, _ = time.ParseInLocation("2006-01-02", "2020-03-01", time.Local)
	fmt.Println(t.AddDate(0, 0, -1).Format("2006-01-02"))	// 2020-02-29
	fmt.Println(t.AddDate(-1, 0, -1).Format("2006-01-02"))	// 2019-02-28
	
	t, _ = time.ParseInLocation("2006-01-02", "2020-03-31", time.Local)
	fmt.Println(t.AddDate(0, -1, 0).Format("2006-01-02"))	// 2030-03-02
	
}
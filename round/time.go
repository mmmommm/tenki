package round
import(
	"time"
)

var layout = "2006-01-02 15:04:05"
func timeToString(t time.Time) string {
    str := t.Format(layout)
    return str
}
var minutesList = [...]string{" 06:00:00", " 12:00:00", " 15:00:00", " 18:00:00"}

//現在時刻を取得して翌日の6時、12時、15時、18時のデータを表示できるようにする
func Tommorow() []string {
	tNow := time.Now()
	//翌日の同時刻を文字列で取得
	tTomorrow := timeToString(tNow.AddDate(0, 0, 1))
	//文字列から年月日だけを抽出
	picDate := tTomorrow[:10]
	var forecastMinutes []string
	for _, v := range minutesList {
		forecastMinutes = append(forecastMinutes, picDate + v)
	}
	return forecastMinutes
}
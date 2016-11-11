package time

import "fmt"
import "time"

// human readable format
func Size(bytes uint64) string {
	switch {
	case bytes < 1024:
		return fmt.Sprintf("%d B", bytes)
	case bytes < 1024*1024:
		return fmt.Sprintf("%.2f K", float64(bytes)/1024)
	case bytes < 1024*1024*1024:
		return fmt.Sprintf("%.2f M", float64(bytes)/1024/1024)
	default:
		return fmt.Sprintf("%.2f G", float64(bytes)/1024/1024/1024)
	}
}

// short string format
func Time(d time.Duration) string {
	u := uint64(d)
	if u < uint64(time.Second) {
		switch {
		case u == 0:
			return "0"
		case u < uint64(time.Microsecond):
			return fmt.Sprintf("%.2f ns", float64(u))
		case u < uint64(time.Millisecond):
			return fmt.Sprintf("%.2f μs", float64(u)/1000)
		default:
			return fmt.Sprintf("%.2f ms", float64(u)/1000/1000)
		}
	} else {
		switch {
		case u < uint64(time.Minute):
			return fmt.Sprintf("%.2f s", float64(u)/1000/1000/1000)
		case u < uint64(time.Hour):
			return fmt.Sprintf("%.2f m", float64(u)/1000/1000/1000/60)
		default:
			return fmt.Sprintf("%.2f h", float64(u)/1000/1000/1000/60/60)
		}
	}
}

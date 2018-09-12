package timeCheck
import "time"


func ISZE()bool{
	hour := time.Now().Hour()
	if hour>0 && hour<6{
		return true
	}
	return false
}
func ISAM()bool{
	hour := time.Now().Hour()
	if hour>=6 && hour<12{
		return true
	}
	return false
}


func ISPM()bool{
	hour := time.Now().Hour()
	if hour>=12 && hour<18{
		return true
	}
	return false
}


func ISEVE()bool{
	hour := time.Now().Hour()
	if hour >= 18 && hour<24{
		return true
	}
	return false
}
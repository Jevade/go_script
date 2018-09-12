package tz

type TZ int
const(
   ZH=TZ(8) 
   UTC=TZ(0) 
)
var tzmap map[TZ]string = map[TZ]string{ZH:"China beijing time",UTC:"Universal Greenwith time"}

func(T TZ)String()string{
    return tzmap[T]
}


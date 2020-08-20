package main
import (
    "net/http"
    "os"
    "strconv"
    "fmt"
    "io/ioutil"
    "math"
    //"bytes"
    //"encoding/json"
    "strings"
    "time"
)

func main(){
    ZERO_POS_LNG := 116.4748838
    ZERO_POS_LAT := 40.1327017
    R := 0.20
	if len(os.Args) > 1 {
        if os.Args[1]=="-h"{
            fmt.Println("请输入参数: 经度 纬度 半径")
            fmt.Println("Usage:")
            fmt.Println("    1:./insertplanstatus.exe  lng  lat  R  ")
            fmt.Println("    2:example: ./insertplanstatus.exe  116.4748838 40.1327017 0.2 ")
            return
        }else{
               if s, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
                  ZERO_POS_LNG = s
                }
               if s, err := strconv.ParseFloat(os.Args[2], 64); err == nil {
                  ZERO_POS_LAT = s
                }
               if s, err := strconv.ParseFloat(os.Args[3], 64); err == nil {
                  R = s
                }
        }
    }

    //song := make(map[string]string)
	//song["username"] = "******"
	//song["password"] = "******"
	//bytesData, _ := json.Marshal(song)
    //data := bytes.NewBuffer([]byte(bytesData))
    //contentTypeJson := "application/json"
    //fmt.Println(data)

    angle := 0
    for;;{
        time.Sleep(200 * time.Millisecond)
        angle++
        angle%=360
        lng := ZERO_POS_LNG + R* math.Sin(math.Pi/360.0*float64(angle))
        lat := ZERO_POS_LAT + R* math.Cos(math.Pi/360.0*float64(angle))
        fmt.Println(angle,lng,lat)
        now := time.Now()      // current local time
        //timestamp := now.UnixNano()
        timestamp := now.Unix()

        para := fmt.Sprintf("regno=NO12345&&planid=1&&time=%d&&ht=1&&alt=1&&height=1&&lng=%.7f&&lat=%.7f&&spd=1&&head=1&&pit=1&&roll=1&&side_slip=1&&flap_pos=2&&act=1&&mode=1&&arm=1&&air=1&&bat_status=1&&bat_work_time=1&&bat_surplus=1&&bat_voltage=1&&bat_current=1&&bat_temp=1&&hyd_pressure=1&&hyd_surplus=1&&hyd_con_rate=1&&hyd_temp=1&&datatra_status=1&&datatra_db=1&&pictra_status=1&&pictra_db=1&&equip_5G_status=1&&equip_5G_db=1&&lidar_status=1&&airdata_mod_status=1&&INS_status=1&&GPS_recept_status=1&&radio_altimeter_status=1&&actuator1_status=1&&actuator2_status=1&&actuator3_status=1&&actuator4_status=1&&actuator5_status=1&&temp=1&&wind_spd=1&&wind_dirc=1&&tc=1",timestamp,lng,lat)
        data := strings.NewReader(para)
        url :=  "http://123.57.22.225:8000/api/post_plane_status/"
        contentTypeForm := "application/x-www-form-urlencoded"
        resp, err :=http.Post(url, contentTypeForm, data)
        if err !=nil{
            fmt.Println(err)
        }
        defer resp.Body.Close()
        body,err := ioutil.ReadAll(resp.Body)
        if err !=nil{
            fmt.Println(err)
        }
        fmt.Println(string(body))
    }

}

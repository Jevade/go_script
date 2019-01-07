package sd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

//HealthCheck is check health
// @Summary Check node health
// @Description check node health
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200
// @Router /sd/health [get]
func HealthCheck(c *gin.Context) {
	log.Info("visit health")
	message := fmt.Sprintf("the status is %s", "ok")
	c.JSON(http.StatusOK, message)
}

//InfoCheck will return vps info
func InfoCheck(c *gin.Context) {
	a, _ := load.Avg()
	u, _ := disk.Usage("/")
	m, _ := mem.VirtualMemory()
	cpus, _ := cpu.Percent(time.Second, false)
	data := make(map[string]interface{})
	data["l1"] = a.Load1
	data["l5"] = a.Load5
	data["l15"] = a.Load15
	data["busyCPU"] = cpus[0]
	data["usedDisk"] = float32(u.Used) / MB
	data["totalDisk"] = float32(u.Total) / MB
	data["usedMEM"] = float32(m.Used) / MB
	data["totalMEM"] = float32(m.Total) / MB
	status := http.StatusOK
	c.JSON(status, data)
}

//DiskCheck is to check disk
// @Summary Check disk health
// @Description check disk health
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 429 {object} httputil.HTTPError "{"Code":429,"Message":"Source busy"}"
// @Failure 500 {object} httputil.HTTPError "{"Code":500,"Message":"Source used up"}"
// @Router /sd/disk [get]
func DiskCheck(c *gin.Context) {
	log.Info("visit disk")
	u, _ := disk.Usage("/")
	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "ok"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	message := fmt.Sprintf("%s - Free Space:%dMB (%dGB)/%dMB (%dGB)|Used:  %d", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.JSON(status, message)
}

//CPUCheck is to cpu Load
// @Summary Check cpu health
// @Description check load health
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 429 {object} httputil.HTTPError "{"Code":429,"Message":"Source busy"}"
// @Failure 500 {object} httputil.HTTPError "{"Code":500,"Message":"Source used up"}"
// @Router /sd/cpu [get]
func CPUCheck(c *gin.Context) {
	cpus, _ := cpu.Percent(time.Second, false)
	cores, _ := cpu.Counts(false)
	status := http.StatusOK
	text := "ok"
	if cpus[0] >= 0.9 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if cpus[0] >= 0.8 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	message := fmt.Sprintf("%s - CPU busy in one second: %.2f| Cores: %d", text, cpus[0], cores)
	c.JSON(status, message)
}

//LoadCheck is to check Load
// @Summary Check load health
// @Description check load health
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 429 {object} httputil.HTTPError "{"Code":429,"Message":"Source busy"}"
// @Failure 500 {object} httputil.HTTPError "{"Code":500,"Message":"Source used up"}"
// @Router /sd/load [get]
func LoadCheck(c *gin.Context) {
	log.Info("visit load")
	cores, _ := cpu.Counts(false)
	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "ok"
	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f ,| Cores: %d", text, l1, l5, l15, cores)
	c.JSON(status, message)

}

//RAMCheck is to check ram
// @Summary Check ram health
// @Description check ram health
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 429 {object} httputil.HTTPError "{"Code":429,"Message":"Source busy"}"
// @Failure 500 {object} httputil.HTTPError "{"Code":500,"Message":"Source used up"}"
// @Router /sd/ram [get]
func RAMCheck(c *gin.Context) {
	log.Info("visit ram")
	u, _ := mem.VirtualMemory()
	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "ok"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}
	message := fmt.Sprintf("%s - Free Memory:%dMB (%dGB)/%dMB (%dGB)|Used:  %d", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.JSON(status, message)
}

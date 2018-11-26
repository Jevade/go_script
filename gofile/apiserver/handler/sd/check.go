package sd

import (
	"fmt"
	"net/http"

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

//CPUCheck is to check CPU
// @Summary Check cpu health
// @Description check cpu health
// @Tags sd
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 429 {object} httputil.HTTPError "{"Code":429,"Message":"Source busy"}"
// @Failure 500 {object} httputil.HTTPError "{"Code":500,"Message":"Source used up"}"
// @Router /sd/cpu [get]
func CPUCheck(c *gin.Context) {
	log.Info("visit cpu")
	test, _ := cpu.Counts(true)
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
	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f ,test:%d| Cores: %d", text, l1, l5, l15, cores, test)
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

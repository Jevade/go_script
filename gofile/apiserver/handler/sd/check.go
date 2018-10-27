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
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

//HealthCheck is check health
func HealthCheck(c *gin.Context) {
	log.Info("visit health")
	message := "ok"
	c.String(http.StatusOK, message)
}

//DiskCheck is to check disk
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
	message := fmt.Sprintf("%s - Free Space:%dMB (%dGB)/%dMB (%dGB)|Used:  %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, message)
}

//CPUCheck is to check CPU
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
	c.String(status, message)

}

//RAMCheck is to check ram
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
	message := fmt.Sprintf("%s - Free Memory:%dMB (%dGB)/%dMB (%dGB)|Used:  %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, message)
}

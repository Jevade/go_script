package transploar

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"

	"../parse"
)

type ploar struct {
	radius float64
	angle  float64
}

type cartesian struct {
	x float64
	y float64
}

const result = "Polar:radius=%.02f angle=%.02f degrees --Cartesian:x=%.02f y=%.02f\n"
const prompt = "Enter a radius and an angle (in degrees), e.g., 12 90, " + "or %s to quit.\n"

func init() {
	if runtime.GOOS == "windows" {
		fmt.Fprintf(os.Stdout, prompt, "Ctrl+Z,Enter")
	} else {
		fmt.Fprintf(os.Stdout, prompt, "Ctrl+D")
	}
}

// P2C func as trans
func P2C() {
	p := make(chan ploar)
	defer close(p)
	c := transP2C(p)
	defer close(c)
	interact(p, c)
}

func transP2C(p chan ploar) chan cartesian {
	var c = make(chan cartesian)
	var car cartesian
	go func() {
		for {
			plo := <-p
			angle := plo.angle * math.Pi / 180.0
			car.x = plo.radius * math.Cos(angle)
			car.y = plo.radius * math.Sin(angle)
			c <- car
		}
	}()
	return c
}

func interact(p chan ploar, c chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please input R and angle")
		line, err := reader.ReadString('\n')
		fmt.Println("the goten line is:", line)
		if err != nil {
			fmt.Printf("err:%s,Please try again", err)
			continue
		}
		if numbers := strings.Fields(line); len(numbers) == 2 {
			numsli, err := myparse.Fields2number(numbers)
			if err != nil {
				fmt.Println("trans to num field")
				continue
			}
			radius, angle := float64(numsli[0]), float64(numsli[1])
			p <- ploar{radius, angle}
			coord := <-c
			fmt.Printf(result, radius, angle, coord.x, coord.y)

		} else {
			fmt.Println("Please input R and angle")
		}

	}
}

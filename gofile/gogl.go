package main

import (
	"log"
	"runtime"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"
	// "log"
)
const(
	width = 500
	height = 500
)
func main(){
	runtime.LockOSThread()
	window:=initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()
	for !window.ShouldClose(){
		draw(window,program)
	}
}

func initGlfw()*glfw.Window{
	if err := glfw.Init();err != nil{
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable,glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor,4)
	glfw.WindowHint(glfw.ContextVersionMinor,1)
	glfw.WindowHint(glfw.OpenGLProfile,glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible,glfw.True)
	window , err := glfw.CreateWindow(width,height,"life game",nil,nil)
	if err != nil{
		panic(err)
	}
	window.MakeContextCurrent()//binding the window to our current thread.
	return window
}

func initOpenGL()uint32{
	if err:=gl.Init();err!=nil{
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
	prog := gl.CreateProgram()
	return prog
}

func draw(window *glfw.Window, pragram uint32){
	gl.Clear(gl.COLOR_BUFFER_BIT|gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(pragram)
	glfw.PollEvents()
	window.SwapBuffers()
}

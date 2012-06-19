package main

import (
	"fmt"
	"github.com/genbattle/Go2D/sdl"
	gl "github.com/chsc/gogl/gl33"
	"os"
)

func Quit(code int) {
	fmt.Println("Exiting with code", code)
	sdl.Quit()
	os.Exit(code)
}

func main() {
	fmt.Println("OpenGL Test!!!")
	// Init SDL
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		fmt.Println("Init error:", err.Error())
		Quit(1)
	}
	// Init OpenGL
	sdl.GL_SetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
	sdl.GL_SetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3)
	sdl.GL_SetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	sdl.GL_SetAttribute(sdl.GL_DEPTH_SIZE, 24)
	
	// Create SDL Window
	window, err := sdl.CreateWindow("OpenGL Test Nick", sdl.WINDOWPOS_CENTRED, sdl.WINDOWPOS_CENTRED, 512, 512, sdl.WINDOW_OPENGL | sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Error while creating window:", err.Error())
		Quit(2)
	}
	
	context := sdl.GL_CreateContext(window)
	
	err = sdl.GL_SetSwapInterval(1)
	if err != nil {
		fmt.Println("Error while setting swap interval:", err.Error())
		Quit(3)
	}
	
	//show red
	gl.ClearColor(1.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	sdl.GL_SwapWindow(window)
	sdl.Delay(2000)
	//show green
	gl.ClearColor(0.0, 1.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	sdl.GL_SwapWindow(window)
	sdl.Delay(2000)
	//show blue
	gl.ClearColor(0.0, 0.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	sdl.GL_SwapWindow(window)
	sdl.Delay(2000)
	
	sdl.GL_DeleteContext(context)
	sdl.DestroyWindow(window)
	Quit(0)	
}

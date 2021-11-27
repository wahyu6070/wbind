package main

import (
    "fmt"
    "os"
    "os/exec"
    "log"
)
func print(IN string){
	fmt.Println(IN)
	}
func IS_NOT_DIR_EXIST(INPUT string) {
	_, err := os.Stat(INPUT)
	if os.IsNotExist(err) {
    	err := os.MkdirAll(INPUT, 0755)
    	if err != nil {
        	log.Fatal(err)
    	}
    	}
	}
func MOUNT_ALL(){
	cmd := exec.Command("mount", "-o", "rw,remount", "/")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
		}
    fmt.Println(string(stdout))
	
	}
func MOUNT_RO(){
	IS_NOT_DIR_EXIST(OUTPUT)
	cmd := exec.Command("mount", "-o", "rw,remount", INPUT, OUTPUT)
	stdout, err := cmd.Output()
	fmt.Println("- Mounting RO <", INPUT,"> To <", OUTPUT,">")
	if err != nil {
		fmt.Println(err.Error())
		return
		}
    fmt.Println(string(stdout))
	
	}
func MOUNT_BIND(INPUT string, OUTPUT string) {
	IS_NOT_DIR_EXIST(OUTPUT)
	cmd := exec.Command("mount", "-o", "bind", INPUT, OUTPUT)
	stdout, err := cmd.Output()
	fmt.Println("- Binding <", INPUT,"> To <", OUTPUT,">")
	if err != nil {
		fmt.Println(err.Error())
		return
		}
    fmt.Println(string(stdout))
    
	}
func main() {
	print("     Wbind by wahyu6070 ")
	print(" ")
	MOUNT_ALL()
	LIST_BIND := map[string]string{
		"/data/media/0" : "/home/sdcard",
		"/system" : "/home/root/system",
		"/vendor" : "/home/root/vendor",
		"/product" : "/home/root/product",
		"/system_ext" : "/home/root/system_ext",
		"/data/app" : "/home/app",
		"/data/data" : "/home/data",
		"/data/data/com.termux/files/usr/bin" : "/system/xbin",
		"/data/data/com.termux/files/home" : "/home/thome",
		"/data/adb/modules" : "/home/mmodules",
		}
    for P, W := range LIST_BIND {
    	_, err := os.Stat(P)
		if !os.IsNotExist(err) {
    		MOUNT_BIND(P, W)
    	}
    	
    }
}
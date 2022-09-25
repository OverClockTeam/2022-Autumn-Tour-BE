package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func VideoSplit(filename string, start string, end string) error {
	flgst := strings.Split(start, ".")
	flgen := strings.Split(end, ".")
	start = verifyInput(start)
	end = verifyInput(end)
	realfile := path.Join(filename)
	cmdArguments := []string{"-i", realfile, "-ss", start, "-to", end,
		"-async", "1", "-y", path.Join(strings.Split(filename, ".")[0] + "-" + flgst[0] + "-" + flgen[0] + "-videocut" + path.Ext(filename))}
	//fmt.Println(cmdArguments)
	cmd := exec.Command("C:/Users/fuyik/Workspace/ffmpeg/bin/ffmpeg", cmdArguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	log.Printf("分割成功")
	return err
}

func verifyInput(param string) string {
	paramSplit := strings.Split(param, ".")
	paramInt, err := strconv.Atoi(paramSplit[0])
	if err != nil {
		log.Printf(err.Error())
		return ""
	}
	min := paramInt / 60
	minString := strconv.Itoa(min)
	if min < 10 {
		minString = "0" + minString
	}
	paramInt = paramInt % 60
	secString := strconv.Itoa(paramInt)
	if paramInt < 10 {
		secString = "0" + secString
	}
	rlt := "00:" + minString + ":" + secString + "." + paramSplit[1]
	return rlt
}

package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

/*
ffmpeg  -y -i 1.mp4 \
	-r 20 -vb 400k \
	-acodec aac -strict experimental -ac 2 -ar 8000 -ab 24k \
	-vf scale=352:288:force_original_aspect_ratio=decrease,pad=352:288:x=(352-iw)/2:y=(288-ih)/2:color=yellow \
	1.3gp

*/
func main() {
	//args := "-i /Users/pavel/Downloads/1.mov -acodec copy -vcodec copy -f flv rtmp://aaa/bbb"
	//cmd := exec.Command("ffmpeg", strings.Split(args, " ")...)
	args := "-v error -select_streams v:0 -show_entries stream=width,height -of csv=s=x:p=0 /Users/pavel/Downloads/1.mov"
	cmd := exec.Command("ffprobe", strings.Split(args, " ")...)

	output, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	errScanner := bufio.NewScanner(stderr)
	errScanner.Split(bufio.ScanWords)
	for errScanner.Scan() {
		fmt.Print(errScanner.Text() + " ")
	}

	outScanner := bufio.NewScanner(output)
	for outScanner.Scan() {
		m := outScanner.Text()
		fmt.Println(m)
		// matchPackets, _ := regexp.MatchString("Packets", m)
		// matchMinimum, _ := regexp.MatchString("Minimum", m)

		// if matchPackets {
		// 		fmt.Println("Ping statistics for 127.0.0.1")
		// 		fmt.Println(m)
		// }

		// if matchMinimum {
		// 		fmt.Println("Approximate round trip times in milli-seconds:")
		// 		fmt.Println(m)
		// }
	}

	cmd.Wait()
}

package jfile

import (
	"fmt"
	"os"
	"os/exec"
)

func ConvertVideoToH264(inputPath, outputPath string) error {
	fmt.Println("Convert video to h264")
	cmd := exec.Command("ffmpeg",
		"-y", // <- force overwrite
		"-i", inputPath,
		"-c:v", "libx264",
		"-preset", "slow",
		"-crf", "28",
		"-c:a", "aac", "-b:a", "96k",
		"-movflags", "+faststart",
		outputPath,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

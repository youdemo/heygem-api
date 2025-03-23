package test

import (
	"bytes"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"testing"
)

func TestFfmpeg(t *testing.T) {
	// 输入视频文件路径
	inputFile := "http://127.0.0.1:8000/temp/20250321002510240.mp4"
	//// 输出音频文件路径
	//outputAudio := "output_audio.mp3"
	//// 输出视频文件路径（无音频）
	//outputVideo := "output_video.mp4"
	//
	//// 分离音频
	//err := ffmpeg.Input(inputFile).
	//	Output(outputAudio, ffmpeg.KwArgs{"q:a": 0, "map": "a"}).
	//	Run()
	//if err != nil {
	//	log.Fatalf("分离音频失败: %s", err)
	//}
	//log.Println("音频分离成功")
	//
	//// 分离视频
	//err = ffmpeg.Input(inputFile).
	//	Output(outputVideo, ffmpeg.KwArgs{"an": "", "vcodec": "copy"}).
	//	Run()
	//if err != nil {
	//	log.Fatalf("分离视频失败: %s", err)
	//}
	//log.Println("视频分离成功")
	// 使用 ffmpeg-go 获取视频信息
	info, err := ffmpeg.ProbeReader(bytes.NewReader(g.Client().GetBytes(gctx.New(), inputFile)))
	if err != nil {
		log.Fatalf("获取视频信息失败: %s", err)
	}
	duration := gjson.New(info).Get("streams.0.duration").Float64()
	// 打印视频信息
	fmt.Println("视频信息时长:", duration)
}

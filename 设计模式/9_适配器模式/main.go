package main

import "fmt"

/**
* @Author : awen
* @Date : 2020/2/24 12:03 下午
 */
/* 适配器模式 */

// 参考：
//	 https://blog.csdn.net/qibin0506/article/details/50598359

// 定义：
//   将一个类的接口转换成客户希望的另外一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作

/* 声音适配的例子 */

type IPlayer interface { // 1 播放音乐接口
	PlayMusic()
}

// 音乐播放器类
type MusicPlayer struct {
	Src string // 音乐路径
}
// 实现Player接口，调用此方法达到播放音乐音乐效果
func (mp *MusicPlayer) PlayMusic() {
	fmt.Println("play music: " + mp.Src)
}

// 现在我们想要实现播放游戏声音功能
type GameSoundPlayer struct {
	Src string
}
func (gp *GameSoundPlayer) PlaySound() {
	fmt.Println("play music: " + gp.Src)
}

/* 适配器 */
type GameSoundAdapter struct {
	SoundPlayer GameSoundPlayer
}
// 适配器 实现PlayMusic()
func (gsa *GameSoundAdapter) PlayMusic() {
	gsa.SoundPlayer.PlaySound() // 调用嵌套子类的 PlaySound()
}

func main() {
	//musicPlayer := &MusicPlayer{Src: "music1.mp3"}
	soundPlayer := GameSoundPlayer{Src:"music2.mp3"}
	soundAdapter := &GameSoundAdapter{SoundPlayer:soundPlayer}
	play(soundAdapter)
}

func play(player IPlayer)  {
	player.PlayMusic()
}

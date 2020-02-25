package main

import (
	"fmt"
	"strings"
)

/**
* @Author : awen
* @Date : 2020/2/25 12:18 上午
 */

/*代理模式 https://studygolang.com/articles/7193
	定义：
		给某一个对象提供一个代理，并由代理对象控制对原对象的引用。
		代理模式的英文叫做Proxy或Surrogate，它是一种对象结构型模式。

	结构：
		Client 就是上面那个段子中的你， 你是行为的主导者。
		Subject 是代理人和被代理的抽象接口
		RealSubject 被代理的对象， 也就是上面的妹子
		Proxy 代理者， 对应上面的妹子室友
	例子：
		比如你喜欢一个妹子， 不好意思跟人家开口， 这时候你可能就通过她舍友来表达你的爱慕了
*/

// 被代理的抽象接口
type Git interface {
	Clone(url string) bool
}
// 被代理的真实对象
type Github struct {}
func (g Github) Clone(url string) bool {
	if strings.HasPrefix(url,"https") {
		fmt.Println("clone from "+url)
		return true
	}
	fmt.Println("failed to clone from " + url)
	return false
}

// 代理者
type GitBash struct{
	GitCmd Git
}
func (p GitBash) Clone(url string) bool {
	return p.GitCmd.Clone(url)
}

// 使用者
type Coder struct{}
func (p Coder) GetCode(url string) {
	gitBash := GetGit(1)
	if gitBash.Clone(url) {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}
func GetGit(t int) Git {
	if t == 1 {
		return GitBash{GitCmd: Github{}}
	}
	return nil // 可能还有其他的git源
}

func main() {
	var coder Coder = Coder{}
	coder.GetCode("https://github.com/qibin0506/go-designpattern")
	coder.GetCode("http://github.com/qibin0506/go-designpattern")
}

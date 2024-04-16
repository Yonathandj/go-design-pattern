package design_pattern

import "fmt"

/*
	Observer adalah design pattern dimana terdapat 2 entitas yang saling berhubungan yaitu subject dan satu atau lebih observer atau pelanggan. Dimana entity subject dapat mendaftarkan atau menghapus observer yang berlangganan kepadanya. Subject mempunyai method untuk memberitahu observer apabila state di subject berubah.

	Design pattern observer pada intinya adalah hubungan antara subject dengan para observer yang mengamatinya dimana apabila terjadi perubahan kondisi atau state pada subject maka subject akan memanggil method milik observer untuk memberitahu observer bahwa state di subject telah berubah sehingga observer dapat melakukan tindakan terhadap perubahan tersebut.
*/

// Observer
type ISubscriber interface {
	Update(message string)
	Subscribe(channel IYoutubeChannel)
	UnSubscribe(channel IYoutubeChannel)
}
type Subscriber struct {
	Name string
	YoutubeChannels []IYoutubeChannel
}
func (s *Subscriber) Update(message string) {
	fmt.Println("Hello " + s.Name + ", " + message)
}
func (s *Subscriber) Subscribe(channel IYoutubeChannel) {
	s.YoutubeChannels = append(s.YoutubeChannels, channel)
	channel.Register(s);
}
func (s *Subscriber) UnSubscribe(channel IYoutubeChannel) {
	for i, youtubeChannel := range s.YoutubeChannels {
		if youtubeChannel == channel {
			s.YoutubeChannels = append(s.YoutubeChannels[:i], s.YoutubeChannels[:i+1]...)
		}
	}
	channel.UnRegister(s)
}

// Subject
type IYoutubeChannel interface {
	Register(subscriber ISubscriber)
	UnRegister(subscriber ISubscriber)
	NotifyLive()
}
type YoutubeChannel struct {
	Name string
	Subscribers []ISubscriber
}
func (yc *YoutubeChannel) Register(subscriber ISubscriber)  {
	yc.Subscribers = append(yc.Subscribers, subscriber)
}
func (yc *YoutubeChannel) UnRegister(subscriber ISubscriber)  {
	for i, v := range yc.Subscribers {
		if v == subscriber {
			yc.Subscribers = append(yc.Subscribers[:i], yc.Subscribers[i + 1:]...)
			break
		}
	}
}
func (yc *YoutubeChannel) NotifyLive() {
	for _, subscriber := range yc.Subscribers {
		subscriber.Update(yc.Name + " has live now")
	}
}
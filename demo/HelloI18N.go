package demo

import "fmt"

type I18NMessage struct {
  msgmap map[string]string
}

func (m *I18NMessage) AddMsg(lang, msg string){
  if m.msgmap == nil {
    m.msgmap = make(map[string]string)
  }
  m.msgmap[lang] = msg
}

func (m *I18NMessage) GetLen() int{
  return len(m.msgmap)
}

func (m *I18NMessage) Clear(){
  m.msgmap = nil
}

func (m *I18NMessage) Say(lang string) {
  var msg = m.msgmap[lang]
  if len(msg) > 0 {
    fmt.Println(msg)
  } else {
    panic("no language specified")
  }
}

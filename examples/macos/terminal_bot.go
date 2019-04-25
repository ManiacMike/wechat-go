package main

import (
	"github.com/songtianyi/rrframework/logs"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/cleaner"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/faceplusplus"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/forwarder"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/gifer"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/joker"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/laosj"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/replier"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/revoker"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/switcher"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/system"
	"github.com/ManiacMike/wechat-go/plugins/wxweb/youdao"
	"github.com/ManiacMike/wechat-go/wxweb"
	"time"
)

func main() {
	// create session
	session, err := wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}
	// load plugins for this session
	faceplusplus.Register(session)
	replier.Register(session)
	switcher.Register(session)
	gifer.Register(session)
	cleaner.Register(session)
	laosj.Register(session)
	joker.Register(session)
	revoker.Register(session)
	forwarder.Register(session)
	system.Register(session)
	youdao.Register(session)

	// enable by type example
	if err := session.HandlerRegister.EnableByType(wxweb.MSG_SYS); err != nil {
		logs.Error(err)
		return
	}

	for {
		if err := session.LoginAndServe(false); err != nil {
			logs.Error("session exit, %s", err)
			for i := 0; i < 3; i++ {
				logs.Info("trying re-login with cache")
				if err := session.LoginAndServe(true); err != nil {
					logs.Error("re-login error or session down, %s", err)
				}
				time.Sleep(3 * time.Second)
			}
			if session, err = wxweb.CreateSession(nil, session.HandlerRegister, wxweb.TERMINAL_MODE); err != nil {
				logs.Error("create new sesion failed, %s", err)
				break
			}
		} else {
			logs.Info("closed by user")
			break
		}
	}
}

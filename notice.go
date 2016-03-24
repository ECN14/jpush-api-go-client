package jpushclient

import (
	"fmt"
)

type Notice struct {
	Alert    string          `json:"alert,omitempty"`
	Android  *AndroidNotice  `json:"android,omitempty"`
	IOS      *IOSNotice      `json:"ios,omitempty"`
	WINPhone *WinPhoneNotice `json:"winphone,omitempty"`
}

type AndroidNotice struct {
	Alert     string                 `json:"alert"`
	Title     string                 `json:"title,omitempty"`
	BuilderId int                    `json:"builder_id,omitempty"`
	Extras    map[string]interface{} `json:"extras,omitempty"`
}

type IOSNotice struct {
	Alert            string                 `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            interface{}            `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"Content-available,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
}

type WinPhoneNotice struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

func (this *Notice) SetAlert(alert string) {
	this.Alert = alert
}

func (this *Notice) SetAndroidNotice(n *AndroidNotice) {
	this.Android = n
}

func (this *Notice) SetIOSNotice(n *IOSNotice) {
	this.IOS = n
}

func (this *Notice) SetWinPhoneNotice(n *WinPhoneNotice) {
	this.WINPhone = n
}

func (this *Notice) SetIOSSound(sound string) error {
	if this.IOS == nil {
		return fmt.Errorf("IOS is nil pointer")
	}
	this.IOS.Sound = sound
	return nil
}

func (this *Notice) SetIOSBadge(Badge interface{}) error {
	if this.IOS == nil {
		return fmt.Errorf("IOS is nil pointer")
	}
	this.IOS.Badge = Badge
	return nil
}

func (this *Notice) SetIOSExtras(Extras map[string]interface{}) error {
	if this.IOS == nil {
		return fmt.Errorf("IOS is nil pointer")
	}
	this.IOS.Extras = Extras
	return nil
}

func (this *Notice) SetAndroidExtras(Extras map[string]interface{}) error {
	if this.Android == nil {
		return fmt.Errorf("Android is nil pointer")
	}
	this.Android.Extras = Extras
	return nil
}

func (this *Notice) SetWINPhoneExtras(Extras map[string]interface{}) error {
	if this.WINPhone == nil {
		return fmt.Errorf("WINPhone is nil pointer")
	}
	this.WINPhone.Extras = Extras
	return nil
}

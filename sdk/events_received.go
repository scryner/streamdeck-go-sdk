package sdk

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	EventDidReceiveSettings      = "didReceiveSettings"
	EventDidReceiveGlobalSetting = "didReceiveGlobalSettings"
	EventKeyDown                 = "keyDown"
	EventKeyUp                   = "keyUp"
	EventTouchTap                = "touchTap"
	EventDialPress               = "dialPress" // deprecated
	EventDialUp                  = "dialUp"
	EventDialDown                = "dialDown"
	EventDialRotate              = "dialRotate"
	EventWillAppear              = "willAppear"
	EventWillDisappear           = "willDisappear"
	EventTitleParametersDidChang = "titleParametersDidChange"
	EventDeviceDidConnect        = "deviceDidConnect"
	EventDeviceDidDisconnect     = "deviceDidDisconnect"
	EventApplicationDidLaunch    = "applicationDidLaunch"
	EventApplicationDidTerminate = "applicationDidTerminate"
	EventSystemDidWakeUp         = "systemDidWakeUp"
	EventPropertyInspectorDidApp = "propertyInspectorDidAppear"
	EventPropertyInspectorDidDis = "propertyInspectorDidDisappear"
	EventSendToPlugin            = "sendToPlugin"
	EventSendToPropertyInspector = "sendToPropertyInspector"
)

type EventEnvelope struct {
	Path  string
	Event interface{}
}

func (g *EventEnvelope) UnmarshalJSON(data []byte) error {
	var aux struct {
		Event string `json:"event"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch aux.Event {
	// case EventDidReceiveSettings:
	// case EventDidReceiveGlobalSetting:
	// case EventKeyDown:
	// case EventKeyUp:
	case EventTouchTap:
		event := &TouchTapEvent{}

		if err := json.Unmarshal(data, event); err != nil {
			return err
		}

		g.Path = fmt.Sprintf("%s/%s", event.Action, event.Event)

		g.Event = event

	case EventDialPress:
		event := &DialPressEvent{}

		if err := json.Unmarshal(data, event); err != nil {
			return err
		}

		g.Path = fmt.Sprintf("%s/%s", event.Action, event.Event)

		g.Event = event

	case EventDialUp:
		event := &DialEvent{}

		if err := json.Unmarshal(data, event); err != nil {
			return err
		}

		g.Path = fmt.Sprintf("%s/%s", event.Action, event.Event)

		g.Event = event

	case EventDialDown:
		event := &DialEvent{}

		if err := json.Unmarshal(data, event); err != nil {
			return err
		}

		g.Path = fmt.Sprintf("%s/%s", event.Action, event.Event)

		g.Event = event

	case EventDialRotate:
		event := &DialRotateEvent{}

		if err := json.Unmarshal(data, event); err != nil {
			return err
		}

		g.Path = fmt.Sprintf("%s/%s", event.Action, event.Event)

		g.Event = event

	case EventWillAppear:
		event := &WillAppearEvent{}

		if err := json.Unmarshal(data, event); err != nil {
			return err
		}

		g.Path = fmt.Sprintf("%s/%s", event.Action, event.Event)

		g.Event = event

	case EventWillDisappear:
		event := &WillDisappearEvent{}

		if err := json.Unmarshal(data, event); err != nil {
			return err
		}

		g.Path = fmt.Sprintf("%s/%s", event.Action, event.Event)

		g.Event = event

	case EventDeviceDidConnect:
		event := &DeviceDidConnectEvent{}

		if err := json.Unmarshal(data, event); err != nil {
			return err
		}

		g.Path = fmt.Sprintf("%s", event.Event)

		g.Event = event

	case EventDeviceDidDisconnect:
		event := &DeviceDidDisconnectEvent{}

		if err := json.Unmarshal(data, event); err != nil {
			return err
		}

		g.Path = fmt.Sprintf("%s", event.Event)

		g.Event = event

	// case EventTitleParametersDidChang:
	// case EventApplicationDidLaunch:
	// case EventApplicationDidTerminate:
	// case EventSystemDidWakeUp:
	// case EventPropertyInspectorDidApp:
	// case EventPropertyInspectorDidDis:
	// case EventSendToPlugin:
	// case EventSendToPropertyInspector:
	default:
		log.Printf("WARNING: %q is unimplemented!", aux.Event)
	}

	return nil
}

type TouchTapEvent struct {
	Action  string `json:"action,omitempty"`
	Event   string `json:"event,omitempty"`
	Context string `json:"context,omitempty"`
	Device  string `json:"device,omitempty"`
	Payload struct {
		Settings    interface{} `json:"settings,omitempty"`
		Coordinates struct {
			Column int `json:"column,omitempty"`
			Row    int `json:"row,omitempty"`
		} `json:"coordinates,omitempty"`
		TapPos []int `json:"tapPos,omitempty"`
		Hold   bool  `json:"hold,omitempty"`
	} `json:"payload,omitempty"`
}

type DialPressEvent struct {
	Action  string `json:"action,omitempty"`
	Event   string `json:"event,omitempty"`
	Context string `json:"context,omitempty"`
	Device  string `json:"device,omitempty"`
	Payload struct {
		Settings    interface{} `json:"settings,omitempty"`
		Coordinates struct {
			Column int `json:"column,omitempty"`
			Row    int `json:"row,omitempty"`
		} `json:"coordinates,omitempty"`
		Pressed bool `json:"pressed,omitempty"`
	} `json:"payload,omitempty"`
}

type DialEvent struct {
	Action  string `json:"action,omitempty"`
	Event   string `json:"event,omitempty"`
	Context string `json:"context,omitempty"`
	Device  string `json:"device,omitempty"`
	Payload struct {
		Controller  string      `json:"controller,omitempty"`
		Settings    interface{} `json:"settings,omitempty"`
		Coordinates struct {
			Column int `json:"column,omitempty"`
			Row    int `json:"row,omitempty"`
		} `json:"coordinates,omitempty"`
	} `json:"payload,omitempty"`
}

type DialRotateEvent struct {
	Action  string `json:"action,omitempty"`
	Event   string `json:"event,omitempty"`
	Context string `json:"context,omitempty"`
	Device  string `json:"device,omitempty"`
	Payload struct {
		Settings    interface{} `json:"settings,omitempty"`
		Coordinates struct {
			Column int `json:"column,omitempty"`
			Row    int `json:"row,omitempty"`
		} `json:"coordinates,omitempty"`
		Ticks   int  `json:"ticks,omitempty"`
		Pressed bool `json:"pressed,omitempty"`
	} `json:"payload,omitempty"`
}

type WillAppearEvent struct {
	Action  string `json:"action,omitempty"`
	Event   string `json:"event,omitempty"`
	Context string `json:"context,omitempty"`
	Device  string `json:"device,omitempty"`
	Payload struct {
		Controller  string      `json:"controller,omitempty"`
		Settings    interface{} `json:"settings,omitempty"`
		Coordinates struct {
			Column int `json:"column,omitempty"`
			Row    int `json:"row,omitempty"`
		} `json:"coordinates,omitempty"`
		State           int  `json:"state,omitempty"`
		IsInMultiAction bool `json:"isInMultiAction,omitempty"`
	} `json:"payload,omitempty"`
}

type WillDisappearEvent struct {
	Action  string `json:"action,omitempty"`
	Event   string `json:"event,omitempty"`
	Context string `json:"context,omitempty"`
	Device  string `json:"device,omitempty"`
	Payload struct {
		Controller  string      `json:"controller,omitempty"`
		Settings    interface{} `json:"settings,omitempty"`
		Coordinates struct {
			Column int `json:"column,omitempty"`
			Row    int `json:"row,omitempty"`
		} `json:"coordinates,omitempty"`
		State           int  `json:"state,omitempty"`
		IsInMultiAction bool `json:"isInMultiAction,omitempty"`
	} `json:"payload,omitempty"`
}

type DeviceDidConnectEvent struct {
	Event      string `json:"event,omitempty"`
	Device     string `json:"device,omitempty"`
	DeviceInfo struct {
		Name string `json:"name,omitempty"`
		Type int    `json:"type,omitempty"`
		Size struct {
			Columns int `json:"columns,omitempty"`
			Rows    int `json:"rows,omitempty"`
		} `json:"size,omitempty"`
	} `json:"deviceInfo,omitempty"`
}

type DeviceDidDisconnectEvent struct {
	Event  string `json:"event,omitempty"`
	Device string `json:"device,omitempty"`
}

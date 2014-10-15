// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"github.com/chanxuehong/wechat/message/passive/request"
	"net/http"
)

// 明文模式
// 消息（事件）分路器, 可以根据实际业务来调整顺序!
func msgDispatch(w http.ResponseWriter, r *http.Request, msg *request.Request, rawXMLMsg []byte, timestamp int64, agent Agent) {
	switch msg.MsgType {
	case request.MSG_TYPE_TEXT:
		agent.ServeTextMsg(w, r, msg.Text(), rawXMLMsg, timestamp)

	case request.MSG_TYPE_EVENT:
		switch msg.Event {
		case request.EVENT_TYPE_LOCATION:
			agent.ServeLocationEvent(w, r, msg.LocationEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_CLICK:
			agent.ServeMenuClickEvent(w, r, msg.MenuClickEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_VIEW:
			agent.ServeMenuViewEvent(w, r, msg.MenuViewEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_SCANCODE_PUSH:
			agent.ServeMenuScanCodePushEvent(w, r, msg.MenuScanCodePushEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_SCANCODE_WAITMSG:
			agent.ServeMenuScanCodeWaitMsgEvent(w, r, msg.MenuScanCodeWaitMsgEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_PIC_SYSPHOTO:
			agent.ServeMenuPicSysPhotoEvent(w, r, msg.MenuPicSysPhotoEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_PIC_PHOTO_OR_ALBUM:
			agent.ServeMenuPicPhotoOrAlbumEvent(w, r, msg.MenuPicPhotoOrAlbumEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_PIC_WEIXIN:
			agent.ServeMenuPicWeixinEvent(w, r, msg.MenuPicWeixinEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_LOCATION_SELECT:
			agent.ServeMenuLocationSelectEvent(w, r, msg.MenuLocationSelectEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_TEMPLATESENDJOBFINISH:
			agent.ServeTemplateSendJobFinishEvent(w, r, msg.TemplateSendJobFinishEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_MASSSENDJOBFINISH:
			agent.ServeMassSendJobFinishEvent(w, r, msg.MassSendJobFinishEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_MERCHANTORDER:
			agent.ServeMerchantOrderEvent(w, r, msg.MerchantOrderEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_SUBSCRIBE:
			if msg.Ticket == "" { // 普通订阅
				agent.ServeSubscribeEvent(w, r, msg.SubscribeEvent(), rawXMLMsg, timestamp)
			} else { // 扫描二维码订阅
				agent.ServeSubscribeByScanEvent(w, r, msg.SubscribeByScanEvent(), rawXMLMsg, timestamp)
			}

		case request.EVENT_TYPE_UNSUBSCRIBE:
			agent.ServeUnsubscribeEvent(w, r, msg.UnsubscribeEvent(), rawXMLMsg, timestamp)

		case request.EVENT_TYPE_SCAN:
			agent.ServeScanEvent(w, r, msg.ScanEvent(), rawXMLMsg, timestamp)

		default: // unknown event type
			agent.ServeUnknownMsg(w, r, rawXMLMsg, timestamp)
		}

	case request.MSG_TYPE_LINK:
		agent.ServeLinkMsg(w, r, msg.Link(), rawXMLMsg, timestamp)

	case request.MSG_TYPE_VOICE:
		agent.ServeVoiceMsg(w, r, msg.Voice(), rawXMLMsg, timestamp)

	case request.MSG_TYPE_LOCATION:
		agent.ServeLocationMsg(w, r, msg.Location(), rawXMLMsg, timestamp)

	case request.MSG_TYPE_IMAGE:
		agent.ServeImageMsg(w, r, msg.Image(), rawXMLMsg, timestamp)

	case request.MSG_TYPE_VIDEO:
		agent.ServeVideoMsg(w, r, msg.Video(), rawXMLMsg, timestamp)

	default: // unknown message type
		agent.ServeUnknownMsg(w, r, rawXMLMsg, timestamp)
	}
}

// 兼容模式, 安全模式
// 消息（事件）分路器, 可以根据实际业务来调整顺序!
func encryptedMsgDispatch(w http.ResponseWriter, r *http.Request, msg *request.Request, rawXMLMsg []byte, timestamp int64, nonce, encryptType string, AESKey, random []byte, agent Agent) {
	switch msg.MsgType {
	case request.MSG_TYPE_TEXT:
		agent.ServeEncryptedTextMsg(w, r, msg.Text(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

	case request.MSG_TYPE_EVENT:
		switch msg.Event {
		case request.EVENT_TYPE_LOCATION:
			agent.ServeEncryptedLocationEvent(w, r, msg.LocationEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_CLICK:
			agent.ServeEncryptedMenuClickEvent(w, r, msg.MenuClickEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_VIEW:
			agent.ServeEncryptedMenuViewEvent(w, r, msg.MenuViewEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_SCANCODE_PUSH:
			agent.ServeEncryptedMenuScanCodePushEvent(w, r, msg.MenuScanCodePushEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_SCANCODE_WAITMSG:
			agent.ServeEncryptedMenuScanCodeWaitMsgEvent(w, r, msg.MenuScanCodeWaitMsgEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_PIC_SYSPHOTO:
			agent.ServeEncryptedMenuPicSysPhotoEvent(w, r, msg.MenuPicSysPhotoEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_PIC_PHOTO_OR_ALBUM:
			agent.ServeEncryptedMenuPicPhotoOrAlbumEvent(w, r, msg.MenuPicPhotoOrAlbumEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_PIC_WEIXIN:
			agent.ServeEncryptedMenuPicWeixinEvent(w, r, msg.MenuPicWeixinEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_LOCATION_SELECT:
			agent.ServeEncryptedMenuLocationSelectEvent(w, r, msg.MenuLocationSelectEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_TEMPLATESENDJOBFINISH:
			agent.ServeEncryptedTemplateSendJobFinishEvent(w, r, msg.TemplateSendJobFinishEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_MASSSENDJOBFINISH:
			agent.ServeEncryptedMassSendJobFinishEvent(w, r, msg.MassSendJobFinishEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_MERCHANTORDER:
			agent.ServeEncryptedMerchantOrderEvent(w, r, msg.MerchantOrderEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_SUBSCRIBE:
			if msg.Ticket == "" { // 普通订阅
				agent.ServeEncryptedSubscribeEvent(w, r, msg.SubscribeEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)
			} else { // 扫描二维码订阅
				agent.ServeEncryptedSubscribeByScanEvent(w, r, msg.SubscribeByScanEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)
			}

		case request.EVENT_TYPE_UNSUBSCRIBE:
			agent.ServeEncryptedUnsubscribeEvent(w, r, msg.UnsubscribeEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		case request.EVENT_TYPE_SCAN:
			agent.ServeEncryptedScanEvent(w, r, msg.ScanEvent(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

		default: // unknown event type
			agent.ServeEncryptedUnknownMsg(w, r, rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)
		}

	case request.MSG_TYPE_LINK:
		agent.ServeEncryptedLinkMsg(w, r, msg.Link(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

	case request.MSG_TYPE_VOICE:
		agent.ServeEncryptedVoiceMsg(w, r, msg.Voice(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

	case request.MSG_TYPE_LOCATION:
		agent.ServeEncryptedLocationMsg(w, r, msg.Location(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

	case request.MSG_TYPE_IMAGE:
		agent.ServeEncryptedImageMsg(w, r, msg.Image(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

	case request.MSG_TYPE_VIDEO:
		agent.ServeEncryptedVideoMsg(w, r, msg.Video(), rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)

	default: // unknown message type
		agent.ServeEncryptedUnknownMsg(w, r, rawXMLMsg, timestamp, nonce, encryptType, AESKey, random)
	}
}
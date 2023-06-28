package API

import (
	"github.com/botscommunity/icqgo/keyboard"
	"github.com/botscommunity/icqgo/update"
)

type ChatID string
type UserID string

// AnswerCallback
type QueryID string
type Text string
type URL string
type ShowAlert bool

// SendMessage
type ReplyMessageID string
type ReplyMessagesID []string
type ForwardChatID string
type ForwardMessageID string
type ForwardMessagesID []string
type InlineKeyboard *keyboard.InlineKeyboard
type Format update.Format
type ParseMode string

// SendFile
type Caption string

// DeleteMessages
type MessageID string
type MessagesID []string

// BanDeleteMessages
type DeleteMessages bool

// ResolvePending
type Approve bool
type Everyone bool

// ChatControl
type Title string
type About string

type Actions []Action
type Action string

// GetFile
type FileID string

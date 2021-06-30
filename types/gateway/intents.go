package gateway

import (
	"github.com/kislball/goocord/utils"
)

const (
	IntentGuilds = 1 << 0
	IntentGuildMembers = 1 << 1
	IntentGuildBans = 1 << 2
	IntentGuildEmojis = 1 << 3
	IntentGuildIntegrations = 1 << 4
	IntentGuildWebhooks = 1 << 5
	IntentGuildInvites = 1 << 6
	IntentGuildVoiceStates = 1 << 7
	IntentGuildPresences = 1 << 8
	IntentGuildMessages = 1 << 9
	IntentGuildMessageReactions = 1 << 10
	IntentGuildMessageTyping = 1 << 11
	IntentDirectMessages = 1 << 12
	IntentDirectMessageReactions = 1 << 13
	IntentDirectMessageTyping = 1 << 14
)

var UnprivilegedIntents = utils.Flags{Flags: 0}
var AllIntents = utils.Flags{Flags: 0}

func init() {
	arr := []int{
		IntentGuilds,
		IntentGuildBans,
		IntentGuildEmojis,
		IntentGuildIntegrations,
		IntentGuildWebhooks,
		IntentGuildInvites,
		IntentGuildVoiceStates,
		IntentGuildMessages,
		IntentGuildMessageReactions,
		IntentGuildMessageTyping,
		IntentDirectMessages,
		IntentDirectMessages,
		IntentDirectMessageReactions,
		IntentDirectMessageTyping,
	}
	UnprivilegedIntents.Add(arr...)

	AllIntents.Add(IntentGuildMembers, IntentGuildPresences, UnprivilegedIntents.Flags)
}

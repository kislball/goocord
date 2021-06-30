package gateway

// https://discord.com/developers/docs/topics/gateway#payloads-gateway-payload-structure
type Payload struct {
	Opcode   int         `json:"opcode"`
	Data     interface{} `json:"d,omitempty"`
	Sequence int         `json:"s,omitempty"`
	Event    string      `json:"t,omitempty"`
}

const (
	// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
	ActivityFlagInstance = 1 << 0
	// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
	ActivityFlagJoin = 1 << 1
	// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
	ActivityFlagSpectate = 1 << 2
	// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
	ActivityFlagJoinRequest = 1 << 3
	// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
	ActivityFlagSync = 1 << 4
	// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
	ActivityFlagPlay = 1 << 5
)

// https://discord.com/developers/docs/topics/gateway#update-presence-gateway-presence-update-structure
type UpdatePresence struct {
	Since      *int64     `json:"since,omitempty"`
	Status     string     `json:"status"`
	AFK        bool       `json:"afk"`
	Activities []Activity `json:"activities"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-structure
type Activity struct {
	Name          string              `json:"name"`
	Type          int                 `json:"type"`
	URL           *string             `json:"url,omitempty"`
	CreatedAt     int64               `json:"created_at"`
	ApplicationID *string             `json:"application_id,omitempty"`
	Details       *string             `json:"details,omitempty"`
	State         *string             `json:"state,omitempty"`
	Instance      *bool               `json:"instance,omitempty"`
	Flags         *int                `json:"flags,omitempty"`
	Timestamps    *ActivityTimestamps `json:"timestamps,omitempty"`
	Emoji         *ActivityEmoji      `json:"emoji,omitempty"`
	Party         *ActivityParty      `json:"party,omitempty"`
	Assets        *ActivityAssets     `json:"assets,omitempty"`
	Secrets       *ActivitySecrets    `json:"secrets,omitempty"`
	Buttons       *[]ActivityButtons  `json:"buttons,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-timestamps
type ActivityTimestamps struct {
	Start *int `json:"start,omitempty"`
	End   *int `json:"end,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-emoji
type ActivityEmoji struct {
	Name     string  `json:"name"`
	Id       *string `json:"id,omitempty"`
	Animated *bool   `json:"animated,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-party
type ActivityParty struct {
	Id   *string `json:"id,omitempty"`
	Size *[2]int `json:"size,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-assets
type ActivityAssets struct {
	LargeImage *string `json:"large_image,omitempty"`
	LargeText  *string `json:"large_text,omitempty"`
	SmallImage *string `json:"small_image,omitempty"`
	SmallText  *string `json:"small_text,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-secrets
type ActivitySecrets struct {
	Join     *string `json:"join,omitempty"`
	Spectate *string `json:"spectate,omitempty"`
	Match    *string `json:"match,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-buttons
type ActivityButtons struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type UpdatePresencePayload struct {
	Payload
	Data UpdatePresence `json:"d"`
}

// https://discord.com/developers/docs/topics/gateway#identify-identify-structure
type Identify struct {
	Token          string          `json:"token"`
	Compress       *bool           `json:"compress,omitempty"`
	LargeThreshold *int            `json:"large_threshold,omitempty"`
	Shard          *[2]int         `json:"shard,omitempty"`
	Presence       *UpdatePresence `json:"presence,omitempty"`
	Intents        int             `json:"intents"`
}

// https://discord.com/developers/docs/topics/gateway#identify-identify-connection-properties
type IdentifyProperties struct {
	OS      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

type IdentifyPayload struct {
	Payload
	Data Identify `json:"d"`
}

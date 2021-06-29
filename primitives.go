package goocord

// Primitives are structures, which are sent to/by Discord.
// All primitives must be suffixed with Primitive, e.g.
// GuildMemberPrimitive

// GatewayPresenceUpdatePrimitive is sent by the client to indicate a presence or status update.
type GatewayPresenceUpdatePrimitive struct {
	// Unix time (in milliseconds) of when the client went idle, or null if the client is not idle
	Since int64 `json:"since,omitempty"`
	// The user's new status
	Status string `json:"status"`
	// Whether or not the client is afk
	AFK bool `json:"afk"`
	// The user's activities
	Activities []ActivityPrimitive `json:"activities"`
}

type ActivityPrimitive struct {
	Name       string                      `json:"name"`
	Type       string                      `json:"type"`
	URL        string                      `json:"url,omitempty"`
	CreatedAt  int64                       `json:"created_at"`
	Timestamps ActivityTimestampsPrimitive `json:"timestamps,omitempty"`
	Details    string                      `json:"details,omitempty"`
	State      string                      `json:"state,omitempty"`
	Instance   bool                        `json:"instance,omitempty"`
	Flags      int                         `json:"flags,omitempty"`
	Party      ActivityPartyPrimitive      `json:"party,omitempty"`
	Assets     ActivityAssetsPrimitive     `json:"assets,omitempty"`
	Secrets    ActivitySecretsPrimitive    `json:"secrets,omitempty"`
	Buttons    []ActivityButtonPrimitive   `json:"buttons,omitempty"`
}

type ActivityTimestampsPrimitive struct {
	End   int64 `json:"end,omitempty"`
	Start int64 `json:"start,omitempty"`
}

type ActivityPartyPrimitive struct {
	Id   string `json:"id,omitempty"`
	Size [2]int `json:"size,omitempty"`
}

type ActivityAssetsPrimitive struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type ActivitySecretsPrimitive struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

type ActivityButtonPrimitive struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

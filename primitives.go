package goocord

// Primitives are structures, which are sent to/by Discord.
// All primitives must be suffixed with Primitive, e.g.
// GuildMemberPrimitive

// GatewayPresenceUpdatePrimitive is sent by the client to indicate a presence or status update.
type GatewayPresenceUpdatePrimitive struct {
	Since      int64               `json:"since,omitempty"` // Unix time (in milliseconds) of when the client went idle, or null if the client is not idle
	Status     string              `json:"status"`          // The user's new status
	AFK        bool                `json:"afk"`             // Whether or not the client is afk
	Activities []ActivityPrimitive `json:"activities"`      // The user's activities
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
}

type ActivityTimestampsPrimitive struct {
	End   int64 `json:"end,omitempty"`
	Start int64 `json:"start,omitempty"`
}

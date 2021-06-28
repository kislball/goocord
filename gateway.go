package goocord

import (
	"time"
)

type GatewayPresenceUpdate struct {
	GatewayPresenceUpdatePrimitive
	Since time.Time
	Activities []Activity
}

func (g *GatewayPresenceUpdate) FromPrimitive(prim interface{}) {
	p := prim.(*GatewayPresenceUpdatePrimitive)
	g.Since = time.Unix(p.Since * 1000, 0)
	g.Status = p.Status
	g.AFK = p.AFK

	var activities []Activity
	for _, v := range g.Activities {
		act := Activity{}
		act.FromPrimitive(v)
		activities = append(activities, v)
	}

	g.Activities = activities
}

func (g *GatewayPresenceUpdate) ToPrimitive() interface{} {
	return nil
}

type Activity struct {
	ActivityPrimitive
	CreatedAt time.Time
	Flags Flags
}

func (a *Activity) FromPrimitive(prim interface{}) {
	p := prim.(*ActivityPrimitive)
	a.Name = p.Name
	a.Type = p.Type
	a.URL = p.URL
	a.CreatedAt = time.Unix(p.CreatedAt * 1000, 0)
	a.Details = p.Details
	a.State = p.State
	a.Instance = p.Instance
	a.Flags = Flags{p.Flags}
}

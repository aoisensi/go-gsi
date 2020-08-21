package dota2

type Data struct {
	Provider  *Provider  `json:"provider"`
	Map       *Map       `json:"map"`
	Player    *Player    `json:"player"`
	Hero      *Hero      `json:"hero"`
	Abilities *Abilities `json:"abilities"`
	Items     *Items     `json:"items"`
}

type Provider struct {
	Name    string
	AppID   int
	Version int
	//Timestamp TODO
}

type Map struct {
	Name                string `json:"name"`
	MatchID             string `json:"matchid"`
	GameTime            int    `json:"game_time"`
	ClockTime           int    `json:"clock_time"`
	Daytime             bool   `json:"daytime"`
	NightStalkerNight   bool   `json:"nightstalker_night"`
	GameState           string `json:"game_state"`
	Paused              bool   `json:"paused"`
	WinTeam             string `json:"win_team"`
	CustomGameName      string `json:"customgamename"`
	WardPurcaseCooldown int    `json:"ward_purcase_cooldown"`
}

type Player struct {
	SteamID            string   `json:"steamid"`
	Name               string   `json:"name"`
	Activity           string   `json:"activity"`
	Kills              int      `json:"kills"`
	Deaths             int      `json:"deaths"`
	Assists            int      `json:"assists"`
	LastHits           int      `json:"last_hits"`
	Denies             int      `json:"denies"`
	KillStreak         int      `json:"kill_streak"`
	CommandsIssued     int      `json:"commands_issued"`
	KillList           struct{} `json:"kill_list"` // TODO
	TeamName           string   `json:"team_name"`
	Gold               int      `json:"gold"`
	GoldReliable       int      `json:"gold_reliable"`
	GoldUnreliable     int      `json:"gold_unreliable"`
	GoldFromHeroKills  int      `json:"gold_from_hero_kills"`
	GoldFromCreepKills int      `json:"gold_from_creep_kills"`
	GoldFromIncome     int      `json:"gold_from_income"`
	GoldFromShared     int      `json:"gold_from_Shared"`
	GPM                int      `json:"gpm"`
	XPM                int      `json:"xpm"`
}

type Hero struct {
	XPos            int    `json:"xpos"`
	YPos            int    `json:"ypos"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Level           int    `json:"level"`
	Alive           bool   `json:"alive"`
	RespawnSeconds  int    `json:"respawn_seconds"`
	BuybackCost     int    `json:"buyback_cost"`
	BuybackCooldown int    `json:"buyback_cooldown"`
	Health          int    `json:"health"`
	MaxHealth       int    `json:"max_health"`
	HealthPercent   int    `json:"health_percent"`
	Mana            int    `json:"mana"`
	MaxMana         int    `json:"max_mana"`
	ManaPercent     int    `json:"mana_percent"`
	Silenced        bool   `json:"silenced"`
	Stunned         bool   `json:"stunned"`
	Disarmed        bool   `json:"disarmed"`
	MagicImmune     bool   `json:"magicimmune"`
	Hexed           bool   `json:"hexed"`
	Muted           bool   `json:"muted"`
	Break           bool   `json:"break"`
	Smoked          bool   `json:"smoked"`
	HasDebuff       bool   `json:"has_debuff"`
	Talent1         bool   `json:"talent_1"`
	Talent2         bool   `json:"talent_2"`
	Talent3         bool   `json:"talent_3"`
	Talent4         bool   `json:"talent_4"`
	Talent5         bool   `json:"talent_5"`
	Talent6         bool   `json:"talent_6"`
	Talent7         bool   `json:"talent_7"`
	Talent8         bool   `json:"talent_8"`
}

type Abilities struct {
	Ability0 Ability `json:"ability0"`
	Ability1 Ability `json:"ability1"`
	Ability2 Ability `json:"ability2"`
	Ability3 Ability `json:"ability3"`
	Ability4 Ability `json:"ability4"`
	Ability5 Ability `json:"ability5"`
	Ability6 Ability `json:"ability6"`
	Ability7 Ability `json:"ability7"`
	Ability8 Ability `json:"ability8"`
}

type Ability struct {
	Name          string `json:"name"`
	Level         int    `json:"level"`
	CanCast       bool   `json:"can_cast"`
	Passive       bool   `json:"passive"`
	AbilityActive bool   `json:"ability_active"`
	Cooldown      int    `json:"cooldown"`
	Ultimate      bool   `json:"ultimate"`
}

type Items struct {
	Slot0  Item `json:"slot0"`
	Slot1  Item `json:"slot1"`
	Slot2  Item `json:"slot2"`
	Slot3  Item `json:"slot3"`
	Slot4  Item `json:"slot4"`
	Slot5  Item `json:"slot5"`
	Slot6  Item `json:"slot6"`
	Slot7  Item `json:"slot7"`
	Slot8  Item `json:"slot8"`
	Stash0 Item `json:"stash0"`
	Stash1 Item `json:"stash1"`
	Stash2 Item `json:"stash2"`
	Stash3 Item `json:"stash3"`
	Stash4 Item `json:"stash4"`
	Stash5 Item `json:"stash5"`
}

type Item struct {
	Name string `json:"name"`
}

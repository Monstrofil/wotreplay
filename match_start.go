package wotreplay

// MatchStart is block data about "Match Start"
type MatchStart struct {
	ClientVersionFromXML string                `json:"clientVersionFromXml"`
	ClientVersionFromExe string                `json:"clientVersionFromExe"`
	DateTime             string                `json:"dateTime"`
	MapDisplayName       string                `json:"mapDisplayName"`
	MapName              string                `json:"mapName"`
	GameplayID           string                `json:"gameplayID"`
	HasMods              bool                  `json:"hasMods"`
	RegionCode           string                `json:"regionCode"`
	PlayerID             int                   `json:"playerID"`
	PlayerName           string                `json:"playerName"`
	PlayerVehicle        string                `json:"playerVehicle"`
	BattleType           battleType            `json:"battleType"`
	Vehicles             map[vehicleID]vehicle `json:"vehicles"`
	ServerName           string                `json:"serverName"`
	ServerSettings       struct {
		Roaming             []interface{} `json:"roaming"`
		SpgRedesignFeatures struct {
			StunEnabled           bool `json:"stunEnabled"`
			MarkTargetAreaEnabled bool `json:"markTargetAreaEnabled"`
		} `json:"spgRedesignFeatures"`
		IsPotapovQuestEnabled bool `json:"isPotapovQuestEnabled"`
	} `json:"serverSettings"`
}

type battleType int

// http://wiki.wargaming.net/en/Game_Modes
func (x battleType) String() string {
	return []string{
		0:  "???",
		1:  "Random",
		2:  "Training",
		3:  "Tank Company",
		4:  "Special Battle",
		5:  "Clan Wars",
		6:  "???",
		7:  "7vs7",
		8:  "???",
		9:  "???",
		10: "Strongholds",
	}[int(x)]
}
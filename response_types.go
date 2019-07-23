package gocrapi

type Member struct {
	tag      string
	name     string
	trophies int
}

/* Sample clan structure from server:
{
  "state": "string",
  "warEndTime": "string",
  "clan": {
    "tag": "string",
    "name": "string",
    "badgeId": 0,
    "clanScore": 0,
    "participants": 0,
    "battlesPlayed": 0,
    "wins": 0,
    "crowns": 0
  },
  "participants": [
    {
      "tag": "string",
      "name": "string",
      "cardsEarned": 0,
      "battlesPlayed": 0,
      "wins": 0
    }
  ]
}
*/

type WarParticipant struct {
	Tag           string
	Name          string
	CardsEarned   int
	BattlesPlayed int
	Wins          int
}

type WarClan struct {
	Tag           string
	Name          string
	badgeID       int
	ClanScore     int
	Participants  int
	BattlesPlayed int
	Wins          int
	Crowns        int
}

type ClanCurrentWar struct {
	State        string
	WarEndTime   string
	Participants []WarParticipant
}

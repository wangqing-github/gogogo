package card

import (
	"encoding/json"
	"gogogo/center/redis"
)

var cacheCardMap map[string]*[]Card
var cacheConfigCardMap map[int32]*ConfigCard

type Card struct {
	id     int32
	power  int32
	userId string
	level  int32
	start  int32
}

type ConfigCard struct {
	id    int32
	power int32
	level int32
	start int32
}

func (card *Card) updateLevel(level int32) {
	card.level += level
}

func initCard(ids []string) {
	for _, userId := range ids {
		value := redis.GetByKey(userId)
		var cardInfo []Card
		err := json.Unmarshal([]byte(value), &cardInfo)
		if err != nil {
			continue
		}
		cacheCardMap[userId] = &cardInfo
	}
}

func getByUserId(id string) *[]Card {
	cardMap := cacheCardMap[id]
	if cardMap == nil {
		return nil
	}
	return cardMap
}

func getByUidAndCid(id string, cardId int32) *Card {
	cardMap := cacheCardMap[id]
	if cardMap == nil {
		return nil
	}
	c := getCardByCardId(cardId, cardMap)
	return c
}

func levelUp(userId string, cardId, num int32) {
	cardMap := cacheCardMap[userId]
	if cardMap == nil {
		return
	}
	c := getCardByCardId(cardId, cardMap)
	if c != nil {
		c.updateLevel(num)
	}
}

func addCard(cardId int32, userId string) {
	config := cacheConfigCardMap[cardId]
	if config == nil {
		return
	}
	cardMap := cacheCardMap[userId]
	if cardMap == nil {
		cards := make([]Card, 0)
		cacheCardMap[userId] = &cards
		cardMap = &cards
	}
	card := getCardByCardId(cardId, cardMap)
	if card != nil {
		return
	}
	card = &Card{
		id:    cardId,
		start: config.start,
		level: config.level,
	}
	cards := append(*cardMap, *card)
	cacheCardMap[userId] = &cards
}

func getCardByCardId(cardId int32, cardMap *[]Card) *Card {
	for _, c := range *cardMap {
		if c.id == cardId {
			return &c
		}
	}
	return nil
}

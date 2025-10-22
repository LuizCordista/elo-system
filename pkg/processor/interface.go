package processor

import "github.com/luizcordista/elo-system/pkg/model"

type MatchProcessor interface {
	ProcessMatchResult(match model.MatchResult) []model.PlayerMMRChange
}

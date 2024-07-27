package FightingGame

import (
	"errors"
	"fmt"
	"strings"
)

/*
PREDEFINED ERRORS
*/
var ErrorTitleMissing = errors.New("title is missing")
var ErrorCharacterMinimum = errors.New("roster must have 2 characters at minimum")

/*
*
FIGHTING GAME
*/
type FightingGame struct {
	Roster  []Character
	Title   string
	Version int
}

type Character struct {
	Name string
}

func (fg *FightingGame) ToString() string {
	var info strings.Builder
	info.WriteString("============================================\n")
	info.WriteString(fmt.Sprintf("Title: %s Version: %d\n", fg.Title, fg.Version))
	info.WriteString("============================================\n")
	for index, char := range fg.Roster {
		info.WriteString(fmt.Sprintf("%d: %s\n", index, char.Name))
	}
	info.WriteString("\n")
	return info.String()
}

/**
 * BUILDER
 */
type IBuilder interface {
	SetTitle(title string) IBuilder
	SetVersion(version int) IBuilder
	AddCharacter(character Character) IBuilder
	Reset() IBuilder
	Clone() IBuilder
	Build() (FightingGame, error)
}

type Builder struct {
	fightingGame FightingGame
}

func (builder *Builder) Reset() IBuilder {
	builder.fightingGame = FightingGame{}
	return builder
}

func (builder *Builder) Clone() IBuilder {
	// Create a new builder instance
	newBuilder := &Builder{
		fightingGame: FightingGame{
			Title: builder.fightingGame.Title,
			// Make a deep copy of the Roster slice to ensure the clone is independent
			Roster: append([]Character(nil), builder.fightingGame.Roster...),
		},
	}
	return newBuilder
}

func (builder *Builder) SetTitle(title string) IBuilder {
	builder.fightingGame.Title = title
	return builder
}

func (builder *Builder) SetVersion(version int) IBuilder {
	builder.fightingGame.Version = version
	return builder
}

func (builder *Builder) AddCharacter(character Character) IBuilder {
	builder.fightingGame.Roster = append(builder.fightingGame.Roster, character)
	return builder
}

func (builder *Builder) Build() (FightingGame, error) {
	if builder.fightingGame.Title == "" {
		return FightingGame{}, ErrorTitleMissing
	}
	if len(builder.fightingGame.Roster) < 2 {
		return FightingGame{}, ErrorCharacterMinimum
	}
	return builder.fightingGame, nil
}

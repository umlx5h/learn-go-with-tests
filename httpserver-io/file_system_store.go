package main

import (
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, error) {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins, nil
	}

	return 0, fmt.Errorf("no such player")
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{
			Name: name,
			Wins: 1,
		})
	}

	f.database.Seek(0, 0)
	league.Save(f.database)
}

package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"sync"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
	mu       sync.Mutex
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player file %q: %w", file.Name(), err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %q: %w", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from %q: %w", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})

	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, error) {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins, nil
	}

	return 0, fmt.Errorf("no such player")
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{
			Name: name,
			Wins: 1,
		})
	}

	f.database.Encode(f.league)
}

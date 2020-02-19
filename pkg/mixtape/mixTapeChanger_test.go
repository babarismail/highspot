package mixtape

import (
	"testing"

	myjson "../json"
)

func TestAddPlayList(t *testing.T) {
	outputFile := "../../data/output.json"
	err := HandleChangesToMixTape("../../data/mixtape_sample1.json", "../../data/changes_add.json", outputFile)
	if err != nil {
		t.Errorf("unexpected error occurred applying changes. %v", err)
	}
	updatedMixTape, err := myjson.ReadMixTape(outputFile)
	if len(updatedMixTape.PlayLists) != 3 {
		t.Errorf("expecting %d playlists but got %d", 3, len(updatedMixTape.PlayLists))
	}
}

func TestRemovePlayLists(t *testing.T) {
	outputFile := "../../data/output.json"
	err := HandleChangesToMixTape("../../data/mixtape_sample1.json", "../../data/changes_remove_playlists.json", outputFile)
	if err != nil {
		t.Errorf("unexpected error occurred applying changes. %v", err)
	}
	updatedMixTape, err := myjson.ReadMixTape(outputFile)
	if len(updatedMixTape.PlayLists) != 0 {
		t.Errorf("expecting %d playlists but got %d", 0, len(updatedMixTape.PlayLists))
	}
}

func TestUpdateSongs(t *testing.T) {
	outputFile := "../../data/output.json"
	err := HandleChangesToMixTape("../../data/mixtape_sample1.json", "../../data/changes_update.json", outputFile)
	if err != nil {
		t.Errorf("unexpected error occurred applying changes. %v", err)
	}
	updatedMixTape, err := myjson.ReadMixTape(outputFile)
	for _, playlist := range updatedMixTape.PlayLists {
		if playlist.ID == "1" {
			if len(playlist.SongIDs) != 5 {
				t.Errorf("expecting %d songs in playlist with ID %s but got %d", 5, playlist.ID, len(playlist.SongIDs))
			}
		}
	}
}

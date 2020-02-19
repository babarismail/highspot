package json

import "testing"

func TestReadMixTapeSuccess(t *testing.T) {
	mixTape, err := ReadMixTape("../../data/mixtape_sample1.json")
	if len(mixTape.Users) != 1 {
		t.Errorf("Unexpected user count")
	}

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestReadMixTapeNoFile(t *testing.T) {
	_, err := ReadMixTape("../../data/nofile.json")
	if err == nil {
		t.Errorf("expecting error but didn't get one")
	}
}

func TestReadChangesSuccess(t *testing.T) {
	changes, err := ReadChanges("../../data/changes_sample1.json")

	if len(changes.NewPlayLists) != 2 {
		t.Errorf("unexpect number of new playlists in the changes: %d", len(changes.NewPlayLists))
	}
	if len(changes.RemovePlayLists) != 2 {
		t.Errorf("unexpect number of playlists to delete in the changes: %d", len(changes.RemovePlayLists))
	}
	if len(changes.UpdatePlayLists) != 2 {
		t.Errorf("unexpect number of playlists to update in the changes: %d", len(changes.UpdatePlayLists))
	}

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestReadChangesNoFile(t *testing.T) {
	_, err := ReadChanges("../../data/nofile.json")
	if err == nil {
		t.Errorf("expecting error but didn't get one")
	}
}

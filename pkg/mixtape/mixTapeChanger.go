package mixtape

import (
	"fmt"

	myjson "../json"
	"../types"
)

//adds new playlists. throws an error if a playlist doesn't contain a song
func addNewPlayLists(mixTape *types.MixTape, newLists []types.PlayList) (*types.MixTape, error) {
	for _, playListToAdd := range newLists {
		if len(playListToAdd.SongIDs) == 0 {
			return mixTape, fmt.Errorf("Number of songs in a new playlist cannot be zero")
		}
	}
	mixTape.PlayLists = append(mixTape.PlayLists, newLists...)
	return mixTape, nil
}

// removes specified playlists from the mixtape
// using a set to optimize removal time. depending on requirements, we can optimize either for space or time
func removePlayLists(mixTape *types.MixTape, playListsToRemove []string) *types.MixTape {
	newPlayLists := make([]types.PlayList, 0)
	playListsToRemoveSet := make(map[string]bool)
	for _, playlistID := range playListsToRemove {
		playListsToRemoveSet[playlistID] = true
	}
	for _, playlist := range mixTape.PlayLists {
		if !playListsToRemoveSet[playlist.ID] {
			newPlayLists = append(newPlayLists, playlist)
		}
	}
	mixTape.PlayLists = newPlayLists
	return mixTape
}

//updates existing playlist by adding existing songs to it
func updatePlayLists(mixTape *types.MixTape, playListsToUpdate []types.PlayList) *types.MixTape {
	playListsToUpdateSet := make(map[string]types.PlayList)
	for _, playlist := range playListsToUpdate {
		playListsToUpdateSet[playlist.ID] = playlist
	}
	for i := 0; i < len(mixTape.PlayLists); i++ {
		if playListToUpdate, exists := playListsToUpdateSet[mixTape.PlayLists[i].ID]; exists {
			mixTape.PlayLists[i].SongIDs = append(mixTape.PlayLists[i].SongIDs, playListToUpdate.SongIDs...)
		}
	}
	return mixTape
}

//applies all 3 change types currently allowed
func applyChanges(mixTape types.MixTape, changes types.Changes) (*types.MixTape, error) {
	updatedMixTape, err := addNewPlayLists(&mixTape, changes.NewPlayLists)
	if err != nil {
		return updatedMixTape, err
	}
	updatedMixTape = removePlayLists(updatedMixTape, changes.RemovePlayLists)
	updatedMixTape = updatePlayLists(updatedMixTape, changes.UpdatePlayLists)
	return updatedMixTape, nil
}

//HandleChangesToMixTape controller method calling various other methods to apply user provided changes to mix tape
func HandleChangesToMixTape(mixTapeFilePath string, changesFilePath string, outputFilePath string) error {
	mixTape, err := myjson.ReadMixTape(mixTapeFilePath)
	if err != nil {
		return err
	}
	changes, err := myjson.ReadChanges(changesFilePath)
	if err != nil {
		return err
	}

	updatedMixTape, err := applyChanges(mixTape, changes)

	if err != nil {
		return err
	}

	err = myjson.WriteOutputToFile(outputFilePath, *updatedMixTape)
	return err
}

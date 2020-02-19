package types

//MixTape mixtape.json will be deserialized into this struct.
//output.json will also be serialized from this struct
type MixTape struct {
	Users     []User     `json:"users"`
	PlayLists []PlayList `json:"playlists"`
	Songs     []Song     `json:"songs"`
}

//Changes represents the changes that will be required to be made
type Changes struct {
	NewPlayLists    []PlayList `json:"new_playlists"`
	RemovePlayLists []string   `json:"remove_playlists"`
	//PlayList ID will be used to identify the playlist to which the songs will be appended
	UpdatePlayLists []PlayList `json:"update_playlists"`
}

//User represents a single user in MixTape
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//PlayList represents a single playlist in MixTape
type PlayList struct {
	ID      string   `json:"id"`
	UserID  string   `json:"user_id,omitempty"`
	SongIDs []string `json:"song_ids"`
}

//Song represents a single song in MixTape
type Song struct {
	ID     string `json:"id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

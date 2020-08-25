package main

import "fmt"

//Playlist represents linked list
type Playlist struct {
	Name             string
	Head             *Song
	CurrentlyPlaying *Song
}

//Song represents linked list entity
type Song struct {
	Artist string
	Title  string
	Next   *Song
}

//CreatePlaylist returns new list
func CreatePlaylist(name string) *Playlist {
	pl := new(Playlist)
	pl.Name = name
	return pl
}

//AddSong adds new element to the list
func (pl *Playlist) AddSong(artist, title string) {
	song := new(Song)
	song.Artist = artist
	song.Title = title
	if pl.Head == nil {
		pl.Head = song
		pl.CurrentlyPlaying = song
		return
	}
	current := pl.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = song
}

//ListSongs lists all elements of the list
func (pl *Playlist) ListSongs() []*Song {
	songs := make([]*Song, 0)
	current := pl.Head
	for current.Next != nil {
		songs = append(songs, current)
		current = current.Next
	}
	return songs
}

//PlayNext switches the element, returns error if list is empty or it's finished
func (pl *Playlist) PlayNext() error {
	if pl.Head == nil {
		return fmt.Errorf("Playlist is empty")
	}
	if pl.CurrentlyPlaying == nil {
		pl.CurrentlyPlaying = pl.Head
		return nil
	}
	if pl.CurrentlyPlaying.Next == nil {
		pl.CurrentlyPlaying = pl.Head
		return fmt.Errorf("Playlist is finished, starting over")
	}
	pl.CurrentlyPlaying = pl.CurrentlyPlaying.Next
	return nil
}

func main() {
	pl := CreatePlaylist("My fav")
	pl.AddSong("Lil Wayne", "Funeral")
	pl.AddSong("Logic", "Everyday")
	pl.AddSong("RSAC", "NBA")
	pl.AddSong("Foals", "Spanish Sahara")
	songs := pl.ListSongs()
	fmt.Println("listing all songs")
	for i := 0; i < len(songs); i++ {
		fmt.Printf("\n Artist is %v, Song is %v\n", songs[i].Artist, songs[i].Title)
	}
	fmt.Printf("\nCurretnly playing is %v\n", pl.CurrentlyPlaying)
	fmt.Println("Next song")
	pl.PlayNext()
	fmt.Printf("\n CurrentlyPlaying is %v\n", pl.CurrentlyPlaying)
	fmt.Println("Going to the end")
	pl.PlayNext()
	pl.PlayNext()
	fmt.Printf("\n The last song is %v", pl.CurrentlyPlaying)
	err := pl.PlayNext()
	if err != nil {
		fmt.Printf("\nGot error as expected, error is %v", err)
	}
	fmt.Printf("\nCurrent after last is %v\n", pl.CurrentlyPlaying)

}

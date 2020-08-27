package main

import (
	"fmt"
	"log"
)

type song struct {
	artist   string
	title    string
	previous *song
	next     *song
}

type playlist struct {
	name    string
	head    *song
	tail    *song
	current *song
}

func createPlaylist(name string) *playlist {
	p := &playlist{
		name: name,
	}
	return p
}

func (p *playlist) addSong(artist, title string) {
	s := &song{
		artist: artist,
		title:  title,
	}
	if p.head == nil {
		p.head = s
		p.tail = s
		return
	}
	curNode := p.tail
	curNode.next = s
	s.previous = curNode
	s.next = p.head
	p.head.previous = s
	p.tail = s
}

func (p *playlist) showAllSongs() ([]*song, error) {
	curNode := p.head
	allSongs := make([]*song, 0)
	if curNode == nil {
		return nil, fmt.Errorf("playlist is emptry")
	}
	for e := curNode; e != nil; e = e.next {
		allSongs = append(allSongs, e)
	}
	return allSongs, nil
}

func (p *playlist) startPlaying() error {
	if p.head == nil {
		return fmt.Errorf("playlist is empty")
	}
	p.current = p.head
	return nil
}

func (p *playlist) nextSong() {
	cur := p.current
	p.current = cur.next
}
func (p *playlist) previousSong() {
	cur := p.current
	p.current = cur.previous
}

func main() {
	pl := createPlaylist("myplaylist")
	pl.addSong("Lil Wayne", "MAMMAMIA")
	pl.addSong("Intelligence", "August")
	pl.addSong("Tash Sultana", "Notion")
	err := pl.startPlaying()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\nStarted to play %v", pl.current)
	pl.nextSong()
	fmt.Printf("\nPlaying next song %v", pl.current)
	pl.previousSong()
	fmt.Printf("\nReturned to previous song %v", pl.current)
	pl.nextSong()
	pl.nextSong()
	fmt.Printf("\nPlaying last song in playlist %v", pl.current)
	pl.nextSong()
	fmt.Printf("\nStarting over with song %v", pl.current)
	pl.previousSong()
	fmt.Printf("\nPrevious song for first is last %v\n", pl.current)
}

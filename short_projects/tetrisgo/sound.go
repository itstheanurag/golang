package main

import (
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/generators"
	"github.com/gopxl/beep/v2/speaker"
)

var (
	audioRate  = beep.SampleRate(44100)
	audioReady bool
)

func initAudio() {
	speaker.Init(audioRate, audioRate.N(time.Second/10))
	audioReady = true
}

func closeAudio() {
	if audioReady {
		speaker.Close()
	}
}

func playTone(freq float64, duration time.Duration) {
	if !audioReady {
		return
	}
	sine, err := generators.SineTone(audioRate, freq)
	if err != nil {
		return
	}
	stream := &effects.Volume{
		Streamer: beep.Take(audioRate.N(duration), sine),
		Base:     2,
		Volume:   -3,
		Silent:   false,
	}
	speaker.Play(stream)
}

func playLineClearSound(cleared int) {
	if cleared <= 0 {
		return
	}
	go func() {
		notes := []struct {
			freq     float64
			duration time.Duration
			gap      time.Duration
		}{
			{523.25, 70 * time.Millisecond, 75 * time.Millisecond},  // C5
			{659.25, 70 * time.Millisecond, 75 * time.Millisecond},  // E5
			{783.99, 70 * time.Millisecond, 75 * time.Millisecond},  // G5
			{1046.50, 110 * time.Millisecond, 0},                    // C6
		}

		switch cleared {
		case 1:
			playTone(440, 90*time.Millisecond)
		case 2:
			playTone(notes[0].freq, notes[0].duration)
			time.Sleep(notes[0].gap)
			playTone(notes[1].freq, notes[1].duration)
		case 3:
			for i := 0; i < 3; i++ {
				playTone(notes[i].freq, notes[i].duration)
				time.Sleep(notes[i].gap)
			}
		default:
			for i := 0; i < 4; i++ {
				playTone(notes[i].freq, notes[i].duration)
				if notes[i].gap > 0 {
					time.Sleep(notes[i].gap)
				}
			}
		}
	}()
}
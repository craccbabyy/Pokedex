package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"net/http"
)

func init() {
	supportedCommands["inspect"] = cliCommand{
		name:   "inspect <pokemon_name>",
		desc:   "Inspect a Pokemon you have caught",
		callbk: commandInspect,
	}
}

func commandInspect(cfg *config, args ...string) error {
	// this command is different, we are only accessing the cache and therefore has no associated method in pokeapi
	// but DOES make API calls with a helper function, so perhaps this design is a bit wonky (move to internal/)

	//check for exactly 1 argument
	if len(args) != 1 {
		return errors.New("please provide a Pokemon name")
	}
	name := args[0]
	//check the cache for the pokemon name in users caught pokemon
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon yet!")
	}

	// if we made it here, the pokemon was already caught, so return it's details

	// adding functionality to PRINT SPRITE
	url := pokemon.Sprites.Other.OfficialArtwork.FrontDefault
	if url == "" {
		url = pokemon.Sprites.FrontDefault
	}
	if url != "" {
		err := PrintSprite(url)
		if err != nil {
			return err
		}
	}

	// text details
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  -%v\n", typeInfo.Type.Name)
	}
	return nil
}

/* desired output structure :
Pokedex > inspect pidgey
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying
*/

// helper function to render sprite (old school technique)
func PrintSprite(url string) error { // need to feed it the correct URL from the JSON data struct
	if url == "" { // string empty
		return fmt.Errorf("url string cannot be empty")
	}

	// download the png file - linked from JSON data in struct
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	bounds := img.Bounds()
	srcWidth := bounds.Max.X
	srcHeight := bounds.Max.Y

	// Set a safe layout target (32x32 characters looks crisp and won't line-wrap)
	targetWidth := 32
	targetHeight := 16

	// Use a bytes.Buffer to assemble the output in memory.
	// This makes it render instantly instead of stuttering row-by-row.
	var buf bytes.Buffer

	// Increment by 1 row of characters (which handles 2 rows of pixels)
	for tY := 0; tY < targetHeight; tY++ {
		for tX := 0; tX < targetWidth; tX++ {
			srcX := bounds.Min.X + (tX * srcWidth / targetWidth)
			srcY1 := bounds.Min.Y + ((tY * 2) * srcHeight / (targetHeight * 2))
			srcY2 := bounds.Min.Y + ((tY*2 + 1) * srcHeight / (targetHeight * 2))

			r1, g1, b1, a1 := img.At(srcX, srcY1).RGBA()
			r2, g2, b2, a2 := img.At(srcX, srcY2).RGBA()

			// Safely shift 16-bit color map down to 8-bit standard
			r1_8, g1_8, b1_8 := r1>>8, g1>>8, b1>>8
			r2_8, g2_8, b2_8 := r2>>8, g2>>8, b2>>8

			topAlpha := a1 >> 8
			botAlpha := a2 >> 8

			// Strict background transparency thresholding
			if topAlpha < 20 && botAlpha < 20 {
				buf.WriteString(" ")
				continue
			}

			// Only bottom pixel is transparent (Draw upper block ▀)
			if topAlpha >= 20 && botAlpha < 20 {
				buf.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm▀\033[0m", r1_8, g1_8, b1_8))
				continue
			}

			// Only top pixel is transparent (Draw lower block ▄)
			if topAlpha < 20 && botAlpha >= 20 {
				buf.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm▄\033[0m", r2_8, g2_8, b2_8))
				continue
			}

			// Both are active pixels - layer foreground and background styling cleanly
			buf.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm▄\033[0m", r2_8, g2_8, b2_8, r1_8, g1_8, b1_8))
		}
		// Reset line styles and append a newline
		buf.WriteString("\n")
	}

	// Flush the buffer to the screen all at once
	fmt.Print(buf.String())
	return nil
}

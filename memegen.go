package main

/*
//	FILE: memegen.go
//
//  PURPOSE: Meme generator for .mark command
//
//  AUTHORS:
//    PROGRAMMER: switch
//    21/04/11 Implementation
//
// ----------------------------------------------------------------------------
*/

/*
//  Method: GenerateMeme
//
//  PURPOSE: Generate Meme image, store meme image (path - db and local file)
//	         Usage: .mark add <text> if image is supplied this method will be called.
//
//  ARGS: image text, image
*/
func GenerateMeme() string {
	return ""
}

/* NOTES

api: https://imgflip.com/api

use this to combine text with image and get back the data to store in db the path and local file of the image.

changes needed:
  database reformat - id, text, (optional) image path
  randomMarkQuote - revise so that if there is an image supplied show that over the text
  listMarkQuote - revise so that only text is shown if there is any. otherwise dont list.
  others... need to look into cases
*/

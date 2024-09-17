# goinvaders

This is yet another Go version of the famous arcade game of the '70s
Only one level is playable, but I will continue to add features.

The main purpose I wrote it is because I am learning Raylib in Go and want to experiment different features.
The code was writtem by following a YouTube video called [C++ Space Invaders Tutorial with raylib - Beginner Tutorial (OOP)](https://youtu.be/TGo3Oxdpr5o) by Nick Koumaris
I strongly suggest you to watch his video, which I followed step by step translating his C++ code into Go.

I am addind more features while learning. For instance I modified Nick's handling of textures by introducing an Atlas file generated on my Mac with [TexturePacker](https://www.codeandweb.com/texturepacker)
and then unmarshaling the xml to a structure to process textures (see file atlas.go)

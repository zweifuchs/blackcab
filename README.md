# Black CAB
Explore the very basic Cellular Automata.

This programm calculates images from basic Cellular Automata. It can calculate the first 255 basic forms of CA described by Stephen Wolfram. 
[https://en.wikipedia.org/wiki/Elementary_cellular_automaton](https://en.wikipedia.org/wiki/Elementary_cellular_automaton) 

## Table of Contents

## Getting Started
Compile and run. 

### Prerequisites

You need to have GO installed on your machine. Right now no pre-builds are provided. However, this programm only uses stdlib. Just compile the code and you are ready to go. 

A generate.sh file is added as an example to mass generate images. 

### Examples

Basic options:

```
-pop 500 -generations 500 -rule 30
```
creates a 500x500 Image with Rule 30
```
-pop 500 -generations 500 -rule 73 -genesis 001000100
```
for a different starting block

```
-rnd -file xyz.png
```
for randomized start and a different filename to save the result

## Notes


## Built With

* [GO](https://golang.org/) - Golang

## Notes



## Contributing

Feel free to fork and spread the love. Also, PRs are welcome too.

* **Alfred Eichenseher** - *Initial work* - [Me on Github](https://github.com/zweifuchs)

## Contact
contact me at :
alfred . eichenseher |AT| googlemail . com 
https://www.alfred-eichenseher.de

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* A big hello to all the happy coders out there.
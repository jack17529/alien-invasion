# alien-invasion

Most random and most optimized simulation of alien invasion problem. The code's time complexity is flexible in order to get the best performance for simulation.

## Installation

Tested on Ubuntu. No prerequisites are required except Golang.

1. Download the project and place it in your GOPATH.
2. Use `make` command to install custom packages, execute tests and build the binary file.
3. Run using `./main -map=(map input path) -aliens=(no. of aliens)`

## Assumptions

1. The map is an undirected map, meaning if foo is north of bar, then bar is south of foo. Both the connections need to be present in the map or my application will not accept it. Also aleins look down upon the world of X thus they want to conqer it, hence they won't follow traffic rules :)

2. If a map has redundant or incomplete connections they are also not accept. Check out `storage\WrongMaps` which I used for testing my application. Also check out correct sample maps in `storage`.

3. Any positive number of aliens with the upper bound of int can come in the map. Only `(number of cities)*2` aliens can come in the world of X at a time. So if the aliens are more they never come to the world of X as it is already destroyed by the aliens that first came.

4. Aliens don't move until all aliens have landed in world of X. They don't want to lose the ownership of the current city.

5. Aliens fight as soon as they enter city having another alien. This happens when they come to the world of X as well as when they move.

6. Only one alien move at a time, but it can recieve many consecutive chances to move at random.

7. The algorithm is works in a few seconds always, but if the map supplied is too big it can be overwhelming for the best algorithm known to man. So please make changes according to your machine efficiency according to the `lv` parameter in Line 70 in `cmd/alien-invasion/main.go` file.

8. `4*(no. of cities)<<10^8` and also `4*(no. of cities with 1 alien after they have come)<<10^8` for the simulation to complete properly on a normal machine.

## Tests

Storage folder contains many sample maps to test the code. Few of them were taken from the net randomly.
Use `make test` to run the test cases.
Although the application is extensivly tested on all the edge cases that came in my mind for durability. I wrote trivial test cases so that any developer watching can understand my project, but the storage covers all the edge cases maps.

Still excited, want to generate more random test cases?
I got you covered.
1. Just generate a grid with dimension Height x Width of cities connected to all the neighbouring cities.
2. Select a random city and delete one of the random neighbours road.
3. Execute step 2 random number of times while making sure that the graph is fully connected.
4. And Voila, you have a random map of X.

One can easily write C++ code for this.

## TODO

1. Make a dockerfile to make the installation robust on all platforms.
2. Writing high quality tests.

## Contributing

Good criticism will be highly appreciated. Feel free to report issues, submit PR or ask questions.

## Author

Shivam Sharma(jack17529)

## MIT Licence

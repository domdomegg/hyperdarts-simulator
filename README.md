# 🎯 HyperDarts simulator

This is a simulator based off [Greg Egan's tweet](https://twitter.com/gregeganSF/status/1160461092973211648) and [this Numberphile + 3blue1brown video](https://www.youtube.com/watch?v=6_yU9eJ0NxA).

The original puzzle is (credit to Greg Egan):

> In the game of HyperDarts, the bull’s eye starts out with a diameter of 1, but each dart that lands reduces it to the length of the chord the dart lies on (or 0, if the dart is outside the bull’s eye).
> 
> Suppose a (very unskilled) player is equally likely to hit *any* point in the blue square. They get 1 point just for playing, then 1 point every time they hit the current bull’s eye.
> 
> If they throw an infinite number of darts, what is their expected score?

## 🔧 Implementation

This implementation is written in Go, as I've been trying to learn it as a language and this seemed like a sensible project at the time to try. As I worked to improve the efficiency of the program it did become more and more trivial though.

I've also extended the original problem from the square distribution to a circle distribution, which has expected value `e`, which is quite a neat result. It's fairly easy to add new custom distributions and run them against the simulator too.

To run the script from the command line run `go run simulator.go`

It's fairly fast, and on my (fairly modest) laptop simulates 12 million games per second.
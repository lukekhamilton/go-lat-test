# Latitude code exercise

Please find enclosed the latitude code exercise based around the getMaxProfit function (main.go) which takes in an slice of int's and returns the profile from trading the possible lowest and highest buy/sell signels.

All code is writen in golang

I have taken a TTD approach to this problem and setup the test cases (main_test.go) to support easlier added new sample data tables for quickly added extra data sets.

I have also wrap the project up in a Dockerfile for quick and easy running of the project with the only dependency then being docker.


## Usage

To buld the docker image and run the test, just simple run:

`make`

To run the test locally without docker if you already have a golang environment setup simple run:

`make test`

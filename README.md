## Jack Henry code review

A sample go project for purposes of team review

assuming that usage of the go package "github.com/briandowns/openweathermap" is not allowed

### Requirements

For developer guidance, this project was developed using an MacOS M1 and Visual Studio Code and the "dev containers" extension (V0.348.0) by load devcontainer.json.  Orbstack was used instead of Docker-Desktop for docker ops.  (brew install orbstack)

"Make" was installed on the local machine not related to dev containers. If not, use homebrew to install Make.

"Bruno" is a Postman like REST testing application that can be installed using homebrew:
(brew install bruno).  You can find the Bruno json queries file in the /bruno directory.

You will need to select the JackHenry environment configuration and also load your APIKEY into this
environment config as well.

### Testing

Make sure you change the value of the APIKEY const in the main.go file before running tests. If not, you will receive 401 error.

When loading the devcontainers.json file, make sure you load the project in the container and rebuild the container to validate your docker dev environment is correct.

execute "make test" from the terminal command prompt to check the state of the project.

### Dev notes

No coverage or other testing rituals are included in this project.


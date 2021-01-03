# septa
Septa golang client and command line executable for getting the next regional rail train to arrive.  This project uses https://www3.septa.org/api/

## Build

[![Go Report Card](https://goreportcard.com/badge/github.com/dherbst/septa?style=flat-square)](https://goreportcard.com/report/github.com/dherbst/septa)
![golint](https://github.com/dherbst/septa/workflows/golint/badge.svg)

To build the command, clone this repo, then type `make`.  The septa executable will be copied to `$HOME/bin`.   Prerequisite is that you have `docker` installed, and `$HOME/bin` is on your `$PATH`.

    git clone https://github.com/dherbst/septa
    cd septa
    make

You can create a bash alias to make looking up the next to arrive for your favorite stations:

    alias nta="septa next Narberth \"Suburban Station\""

Alternatively, you can use the `go install` command which should install the `septa` binary to `$HOME/go/bin`:

    git clone https://github.com/dherbst/septa
    cd septa
    make install-local

### Download the binary
You can download the latest osx binary `septa` at https://github.com/dherbst/septa/releases.  If you need a windows version please open an issue.

## Commands

### Next to Arrive - `septa next <src> <dest>`
Returns the next to arrive between the `src` and `dest` stations.

    $ septa next Narberth "Suburban Station"

    Paoli/Thorndale Train 5340 Departs  4:57PM Arrives  5:16PM On time
    Paoli/Thorndale Train 1542 Departs  5:27PM Arrives  5:45PM On time
    Paoli/Thorndale Train 5344 Departs  5:57PM Arrives  6:16PM On time
    Paoli/Thorndale Train 5348 Departs  6:57PM Arrives  7:16PM On time
    Paoli/Thorndale Train 5352 Departs  7:57PM Arrives  8:16PM On time


### List of stations - `septa stations`
Returns the list of stations that are used as arguments in the `next` command.

    $ septa stations

    9th St
    30th Street Station
    49th St
    Airport Terminal A
    Airport Terminal B
    ...


### Alerts - `septa alerts [route]`
Returns the alerts for all routes if no `route` specified, or just for the `route` specified.

    $ septa alerts generic

    Generic All SEPTA service will operate on a Sunday schedule on Friday, December 25 (Christmas Day). All SEPTA passengers are required to wear a face mask or covering. Regular service schedules are in effect, however, the Cynwyd and Chestnut Hill West Lines are suspended. The 12th & Filbert doors at Jefferson Station and the 16th & JFK doors at Suburban Station will be open during the following hours: 6:30 a.m. to 9:30 a.m. and 3:30 p.m. to 6:30 p.m. (Monday-Friday only). During all other hours these doors will be locked. We appreciate your patience during this time. With the launch of the SEPTA Key Card Travel Wallet feature for Regional Rail, the sale of paper tickets, single trip, and 10 trip strips ended <a href="http://septa.org/key/updates/paper-ticket-sales-ending.html">October 2, 2020</a>. Paper tickets will continue to be accepted through the valid date stamped on the back and valid for 180 days from the purchase date.

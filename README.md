# septa
Septa command line for getting the next to arrive.

## Commands

### septa next <src> <dest>
Returns the next to arrive between the src and dest stations.

   $ septa next Narberth "Suburban Station"

   Paoli/Thorndale Train 5340 Departs  4:57PM Arrives  5:16PM On time
   Paoli/Thorndale Train 1542 Departs  5:27PM Arrives  5:45PM On time
   Paoli/Thorndale Train 5344 Departs  5:57PM Arrives  6:16PM On time
   Paoli/Thorndale Train 5348 Departs  6:57PM Arrives  7:16PM On time
   Paoli/Thorndale Train 5352 Departs  7:57PM Arrives  8:16PM On time


### septa stations
TODO: will list the stations.


### Alerts
curl https://www3.septa.org/hackathon/Alerts/?req1=rr_route_pao | python -m json.tool

```json
[
    {
        "route_id": "rr_route_pao",
        "route_name": "
/Thorndale",
        "mode": "Regional Rail",
        "isadvisory": "Yes",
        "isdetour": "N",
        "isalert": "N",
        "issuppend": "N",
        "last_updated": "Oct 25 2019  9:03PM",
        "isSnow": "N",
        "description": "Between Center City, Bryn Mawr, Malvern, Paoli & Thorndale"
    }
]
```

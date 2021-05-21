# filteringMatches

This projects uses GraphQL for matches query/filter. This application runs on port 8080.

> **schema.graphqls** contains the graphql schema

> **graph** contains the graphql server 

> **store** contains the repository

Repository.FindMatches contains the implemntation for the filtering 


Query Sample

Url: https://test-filtering-matches.herokuapp.com
```
query Match($input:InputFilter!) {
  matches(input: $input) {
    displayName
    city {
      lat
      lon
    }
  }
}
```

Variables

```
{
    "input": {
        "hasPhoto": true,
        "inContact": true,
        "favourite": false,
        "compatibilityScore": {
            "from": 10,
            "to": 99
        },
        "age": {
            "from": 10,
            "to": 99
        },
        "height": {
            "from": 10,
            "to": 990
        },
        "distanceInKm": {
            "lat": 51.509865,
            "lng": -0.118092,
            "from": 10,
            "to": 99
        }
    }
}
```

To Run project, 
1. Run *go mod init* 
2. Run *go run server.go*

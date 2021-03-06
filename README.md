Frontend Assessment
===============================
This repository holds the code and instructions for a Frontend Assessment.

# Assignment

## Dux
Recently the market for rubber duckies has seen a major uptake. We want to get in on this squeaky action
by providing an application where users can easily buy and sell rubber duckies.

Our UX designer has done extensive user research and came up with a wire frame for the application. It's stored as 
`wireframe.png` and can be found in the root of this repository.

Our backend engineers have started work on an API that should allow for building a SPA that provides our users with the
functionality our UX designer has envisioned. It too resides in this repository, inside the `api` directory.

Building the SPA is up to you!

### Prerequisites
The API is written in Go and as such you will need to [install Go](https://golang.org/doc/install) on your system. For 
local development you will probably want to run your SPA on a different port as the API. However, once in production 
the two will serve from the same host name.

To run the API execute the following command from the `api` directory:
```
go run main.go
```

To run its tests issue the following command from the `api` directory:
```
go test ./...
```

### Features
Below you'll find the features we want in the SPA, as well as their descriptions. Currently there is only one user 
account, a test account, identified by the ID `RGFya3dpbmcgRHVjawo=`. It's data is available through the API.

#### Ducks on offer
The market is moving fast and we want our users to be able to find the best deals at any given time. The API endpoint 
`/api/offers/` provides a constant, real time stream of rubber ducks being offered for sale. We want to present this 
stream to our users. You may assume no one else will buy an offer within 10 seconds. You should consider it sold after 
10 seconds. The API endpoint is intended to be used via a long polling mechanism.

#### Money to spend
Our users need to know whether they have enough money to make a trade. You can get their current balance by calling the
API endpoint `/api/users/{userId}/balance/`. You may assume the user only spends money through this application. Other 
possible drains on budget do not have to be accounted for.

#### Making a trade
When a user has enough funds, and wants to make a trade, we should facilitate this. The API endpoint `/api/trades/` can 
be used to execute trades.

### Bonus
Congrats on making it this far! If you're up for a challenge, and you feel like impressing us, you may take a stab at 
the features below.

#### Authentication and authorization
We'll never be able to put this app into production if there's no security in place. Can you provide a basic 
authentication and authorization setup?

#### Go wild
What killer feature did our UX designer miss? What functionality will launch this app into orbit? Go wild and impress 
us with your creativity and ingenuity.

### Hints
To err is human and your colleagues who developed the API are human.

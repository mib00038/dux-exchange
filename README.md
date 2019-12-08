Katana Labs Frontend Assessment
===============================
This repository holds the code and instructions for the Katana Labs Frontend Assessment. This assessment is meant to 
give us insights into your capabilities as a frontend developer.

Once you receive the assessment you will have 7 days to complete the assignment outlined below. Before the 8th day you 
should hand in your solution in a zip file, by email, to the person who initially gave you the assignment.

After your solution has been reviewed a meeting will be planned where your solution will be discussed. Attention will 
be given to the architecture you created and the functionality you delivered.

If you need a bit of assistance, or some hints, you may reach out to the person who gave you the assignment. 
Alternatively you could use a [rubber duckie](https://en.wikipedia.org/wiki/Rubber_duck_debugging) ;-).

To make sure we can use this assessment for multiple candidates, we kindly request you not to share this repository or 
it's contents with anyone.

We wish you good luck and are keen to see your solution!

# Assignment

## Dux
Recently the market for rubber duckies has seen a major uptake. We as Katana Labs want to get in on this squeaky action
by providing an application where users can easily buy and sell rubber duckies.

Our UX designer has done extensive user research and came up with a wire frame for the application. It's stored as 
`wireframe.png` and can be found in the root of this repository.

Our backend engineers have started work on an API that should allow for building a SPA that provides our users with the
functionality our UX designer has envisioned. It too resides in this repository, inside the `api` directory.

Building the SPA is up to you!

### Features
Below you'll find the features we want in the SPA, as well as their descriptions. Currently there is only one user 
account, a test account, identified by the ID `RGFya3dpbmcgRHVjawo=`. It's data is available through the API.

#### Ducks on offer
The market is moving fast and we want our users to be able to find the best deals at any given time. The API endpoint 
`/api/offers/` provides a constant, real time stream of rubber ducks being offered for sale. We want to present this 
stream to our users. You may assume no one else will buy an offer within 10 seconds. You should consider it sold after 
10 seconds.

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

# Design

## TABLES

- Transaction
- Source
- Destination
- User
- Account
- Card
- Mobile

TRANSACTION
HAS: 1 to 1 USER
: 1 to 1 DESTINATION
User
HAS: 1 to many CARD
HAS: 1 to many MOBILE

SOURCE
HAS: 1 to 1 ACCOUNT
: 1 to 1 USER

DESTINATION
HAS: 1 to 1 ACCOUNT
: 1 to 1 USER

CARD
FK: USERID

TRANSACTION TYPE
int

PAYEE ID
int

working key,
isAlive
shouldnot be traced
at all

## Queries

User:

- getAll()
- getFailedCount()
- getSucceededCount()
- getFailedAmount()
- getSucceededAmount()
- getMostUsedService()
- getLeastUsedService()
- getTotalSpending()
- getCards()
- getMobile()

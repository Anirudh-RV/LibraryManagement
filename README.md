# Library Management

**Technologies used**

1. Golang for running GO API's to connect the application with the Database
2. MongoDB as Database

# Features available

1. POST, PUT, GET requests for Member, Books and Issuance collection
2. Have JWT set up for the API calls where a token has to be added in the header to access the backend.
3. The collections have been setup with unique keys.
4. Syntax for CRUD REST Calls:

   1. POST
      <fqdn/IP>/<entity-name>

   2. GET
      For one ID:
      <fqdn/IP>/<entity-name>/<entity-id>

      For All data:
      <fqdn/IP>/<entity-name>

   3. PUT
      <fqdn/IP>/<entity-name>

5. The methods need to have the JWT Token which gets generated when the app is run
6. The POST/PUT API accepts JSON in the following format:
   {
   "members": [
   {
   "mem_id": 1,
   "mem_name": "Joe",
   "mem_phone": "+91 7799228833",
   "mem_email": "Joe@mama.com"
   },
   {
   "mem_id": 2,
   "mem_name": "Bideen",
   "mem_phone": "+91 2082e8008",
   "mem_email": "Bideen@yaya.com"
   }
   ]
   }
7. The GET API responds with a JSON array with the relevant data.

# Steps to run the application

1. To setup the database relations, given I have chosen to go with non-relational db for familiarity of building APIs, use:

DB:
use LibraryAdmin

member:
db.createCollection("Member");
db.Member.createIndex( { "mem_id" : 1 }, { unique : true });

book:
db.createCollection("Book");
db.Book.createIndex( { "book_id" : 1 }, { unique : true });

issuance:
db.createCollection("Issuance");
db.Issuance.createIndex( { "issuance_id" : 1 }, { unique : true });

collection:
db.createCollection("Collection");
db.Issuance.createIndex( { "collection_id" : 1 }, { unique : true });

category:
db.createCollection("Category");
db.Issuance.createIndex( { "cat_id" : 1 }, { unique : true });

membership:
db.createCollection("Membership");
db.Issuance.createIndex( { "membership_id" : 1 }, { unique : true });

2. To run the go API, go to the Api directory and run: go run main.go

3. I have also added the Dockerfiles to the directories to be containerized for deployment.

# Code rundowm

1. Api's folder has the Go Api related code
   a. HandleBooks handles the CRUD Apis for Books collection
   b. HandleBooks handles the CRUD Apis for Books collection
   c. HandleBooks handles the CRUD Apis for Books collection
   d. HandleJWT handles the authorization of the API and making it secure
   e. main.go handles the routes and which functions need to be called
   f. TestRestMethods should be filled with test suites for the API calls
2. webapp contains the reactJS app that needs to be implemented

# NOTES

1. I would've liked to check more into Swagger and how it will help in building APIs, but given the time constraint, I chose to build it from scratch.

2. The DB is use is non-relational and we can emulate the relations between the collections through processing the queries from the API end.

3. I had set up a ReactJS environment to run the UI but will need to be improved upon with more time.

# Enhancements needed

1. Need to add UI for which members have books pending for return on a given day

- Use ReactJS for a single page application
- Use axios for calls
- Make sure we use async await to handle the axios calls

2. Add a constants file to store all the constant values for better code maintainance
3. Use constant files inside each of the Handle module
4. Validation checks before insertion of the Issuance entity to check if the book and member Ids exist
5. Add Tests to the APIs
6. Check Swagger and see how it might help build our APIs

# SAMPLE DATA

MEMBERS:
{
"members": [
{
"mem_id": 1,
"mem_name": "mem1",
"mem_phone": "+91 7799228833",
"mem_email": "mem1@email.com"
},
{
"mem_id": 2,
"mem_name": "mem2",
"mem_phone": "+91 12332133",
"mem_email": "mem2@email.com"
},
{
"mem_id": 3,
"mem_name": "mem3",
"mem_phone": "+91 131353131",
"mem_email": "mem3@email.com"
},
{
"mem_id": 4,
"mem_name": "mem4",
"mem_phone": "+91 513561361",
"mem_email": "mem4@email.com"
}
,{
"mem_id": 5,
"mem_name": "mem5",
"mem_phone": "+91 16136611",
"mem_email": "mem5@email.com"
}
]
}

BOOKS:

{
"books": [
{
"book_id": 1,
"book_name": "Book1",
"book_cat_id": "book_cat_id1",
"book_collection_id": "book_collection_id1",
"book_launch_date": "2017-11-01T00:00:00",
"book_publisher": "book_publisher1"
},
{
"book_id": 2,
"book_name": "Book2",
"book_cat_id": "book_cat_id2",
"book_collection_id": "book_collection_id2",
"book_launch_date": "2019-11-01T00:00:00",
"book_publisher": "book_publisher1"
}
,
{
"book_id": 3,
"book_name": "Book3",
"book_cat_id": "book_cat_id3",
"book_collection_id": "book_collection_id1",
"book_launch_date": "2010-07-08T00:00:00",
"book_publisher": "book_publisher2"
}
,
{
"book_id": 4,
"book_name": "Book4",
"book_cat_id": "book_cat_id4",
"book_collection_id": "book_collection_id1",
"book_launch_date": "2017-12-10T00:00:00",
"book_publisher": "book_publisher4"
}
,
{
"book_id": 5,
"book_name": "Book5",
"book_cat_id": "book_cat_id5",
"book_collection_id": "book_collection_id1",
"book_launch_date": "2017-11-01T00:00:00",
"book_publisher": "book_publisher6"
}
]
}

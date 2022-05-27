# PaulCamper offline challenge

## Application process

To provide your solution in the most suitable way please use following process:
* fork a git repository
* iterate on the solution and commit changes step by step (one commit per logical change set)
with descriptive one line comments
* open a pull request (merge request) to this repo in a 72 hours
(all commits withing 72 hours after challenge has been shared with you are considered as an accepted solution)

You will find 3 parts of the task below in this Readme file. Parts differs in complexity.
They can be treated as separate subtasks and implemented one by one starting from top to bottom.
Partial solutions also accepted.



## Domain

We are working on the service which uses external 3rd-party translation provider API.
We have translation provider API interface (TranslatorAPI) and implementation (translatorStub) for testing purposes.
We know that translation API calls:
* can take long time
* sometimes fails
* costs us money
* result may change sometimes but business decided we can cache them for at least N hours


## Task

* Create a service that will properly handle external translation API with:
1. retry requests N times with exponential back off before failing with an error
2. cache request results in the storage to avoid charges for the same queries (simplest in-memory storage is enough)
3. deduplicate simultaneous queries for the same parameters to avoid charges for same query burst
* Cover new functionality with tests. This is to the same extent important as a solution code itself.


## Source code

translator.go and main.go should not be modified. Please use service.go and any new files for the solution.



## Added by Yogesh


## Implementation of Logic 
I have started the functionality by reading the values from a config file. We have two values which needs to be configured, instead of hard coding those values in code, I have used json file to configured it.

Then I have used a map as a in-memory storage to store catch. I have initialized a catch object in begining.

I have used three structure for catch management. 
1. To store value 
2. To store key
3. To store catched map and mutex

Whenever main module invoking Translate method, 
I am first checking the key (which is composed of from Lang, To Lang and data) in catch, if key is there in catch then I am sending catched value from catch without hiting third party api. 

If key is not available in catch then I am invoking third party api and storing the value in catch for future reference.

While invoking the api I also implemented the logic of maximum try if api fails to return any values.

I built the logic to expire the catched element after elapsing of configured period of time.

Instead of using any third party catch management, I tried to development own catch management using map data structure and go routine.


## Further improvements
1. We can delete the catch if not accessed by the application recently. We can discuss ageing. Instead of providing blanket expiration criteria, we can check which key is not accessed in long time and we can delete that key only.
2. We can implement log functionality
    We can capture and write all errors in one log file. Errors from any operation or from any event, function
3. We can add more test cases here as I considered only major scenarios.
    
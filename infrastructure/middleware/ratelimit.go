package middleware

/* 
In this step your goal is to implement the token bucket algorithm for rate limiting. The token bucket algorithm works like this:

There is a ‘bucket’ that has capacity for N tokens. Usually this is a bucket per user or IP address.
Every time period a new token is added to the bucket, if the bucket is full the token is discarded.
When a request arrives and the bucket contains tokens, the request is handled and a token is removed from the bucket.
When a request arrives and the bucket is empty, the request is declined.
For this step, implement this strategy such that the bucket is per IP address, has a capacity of 10 tokens with new tokens added at a rate of 1 token per second.

When a request is rejected you should return the HTTP status code of 429 - Too Many Requests.

Once you have implemented that you can use Postman to test it. There is a blog post that introduces the performance testing abilities of Postman and explains how to set it up here.

I configured a test to hit the limited API endpoint with 10 Virtual Users, as you can see that results in no errors initially, (the bucket had ten tokens, then after a second 90% of the requests fail as there are 10 users trying to access the API, but only one token being added per second.
*/


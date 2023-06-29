# about myBA-backend
myBA is the backend of the app were you can administrate your business dates for your customers and the services that you offer as task for your employees. This app was building with golang language, implements a clean architecture with different layers.
In domain layer, there are all entities, repositories interfaces and the respective logic business (use cases).
In infrastrcuture layer, there are the files that implements the databse, messaging, and http route and other configurations. 
finally, of a implicitly way, I use the singleton design pattern, taking advantage of the use of pointers in the Go language. By injecting the reference to the memory space for database instances, Rabbit, among others, is possible.

# run
this app has a Dockerfile, there you should change the values according to your configurations, you can run the app in route ./cmd/api there you can find the main file and follow the flux of the application, create the image and run the app or run in your local machine configuring the environment variables as .env.config says.

# libraries used
The project use wire to inject dependencies, gin to http request configurations, pq to connect to postgreSQL database and go-rabbitmq library to publish messages in a queue. 

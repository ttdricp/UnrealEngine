# UnrealEngine
Payment Microservice part
The principle of operation of your microservice "Payment service" can be described as follows:

Initialization:

A microservice to consume and listen on the bulk port (in the case of "localhost:50051") for incoming gRPC requests.
Connection to the main server:

The microservice creates a connection to the main server, which provides the payment processing function. This happens by calling grpc.Dial and passing the server address along with the security setting (in the case of using grpc.WithInsecure() to detect transport security).
Payment Processing:

When the microservice receives a request to process a payment from a client, it creates a client connection created to the main server using the previously connected one.
A PaymentRequest object is generated with information about the payment, including the user ID (UserId) and the payment amount (Amount).
The microservice then raised the ProcessPayment method on the client, passing the created request and the execution context.
The main search query and the search for the PaymentResponse object, determining the status of the payment (Status), the payment identifier (PaymentId) and the user identifier (UserId).
The microservice receives a response from the main server and displays information about payments in a log using log.Println.
Payment Processing Simulation:

After processing the payment, the microservice calls the SimulationPaymentProcessing function to simulate payment processing. It takes on a random boolean value, which turns out to be a successful or unsuccessful payment. The probability of successful payment is 80%.
Depending on the successful payment result, the microservice generates information about the payment status using the getPaymentStatus function.
Storage of payment data in the database:

A microservice for lifting the storePaymentInDatabase function, which simply prints information about payments in a log using log.Printf.
Publishing payment data to RabbitMQ:

The publishPaymentToRabbitMQ microservice function for publishing information about payments in turn to RabbitMQ.
A connection is created to RabbitMQ using amqp.Dial.
A channel (ch) is created using conn.Channel().
Declared as a queue using ch.QueueDeclare.
The billing data is in turn linked using ch.Publish messages, where the billing information refers to the quality of the content.
Completion of work:

Upon completion of payment processing, the microservice displays a message about the end of work in the log.
Thus, the microservice receives requests for payment processing from clients via gRPC, passes requests to the main server, quickly receives responses, simulates payment processing, saves information about payments in the database and takes it into account in the RabbitMQ queue for further processing by other services or problems.

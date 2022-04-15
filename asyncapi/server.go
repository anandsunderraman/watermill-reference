package asyncapi

// GetKafkaBroker return the AMQP URI.
func GetKafkaBroker() string {
	//this must be passed in or created by the app based on the bindings
	return "localhost:9092"
}

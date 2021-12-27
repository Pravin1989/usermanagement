package models

const (
	// TransactionIDHeaderKey is the key of the HTTP/Kafka request header used to transfer TransactionID across service boundry
	TransactionIDHeaderKey = "X-Request-Id"

	CorrelationIDHeaderKey = "X-Correlation-Id"
	// CorrelationIDHeaderKey is the key of the HTTP/Kafka response header used to  transfer TransactionID across service boundry
)

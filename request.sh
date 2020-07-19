curl --header "Content-Type: application/json" \
--header "Source-Type: game" \
--request POST \
--data '{"state": "win", "amount": "10.15", "transactionId": "34567"}' \
localhost:5000
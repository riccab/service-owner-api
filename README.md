# service_owner_api

## AUTHOR

Ricca Brunolli (ricca.brunolli@gmail.com)

## Description

Allows users to send an http request for information regarding owners of products and/or services

```none
curl localhost:9858/serviceowner/api/v1/product/name/<name_of_service>
curl localhost:9858/serviceowner/api/v1/product/owner/<owner_of_service>
curl localhost:9858/serviceowner/api/v1/product/handle/<slack_handle>
curl localhost:9858/serviceowner/api/v1/product/email/<owners_email>
```

curling the above address will return the service owner's contact details for the requested product

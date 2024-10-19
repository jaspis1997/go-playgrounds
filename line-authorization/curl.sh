curl -v -X POST https://api.line.me/v2/bot/message/push \
-H 'Content-Type: application/json' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInZlciI6InYzIiwidHlwIjoiSldUIn0.eyJzdWIiOjIwMDYwMjIwMzgsImlhdCI6MTcyMjk1MDMzNzQxNSwiZXhwIjoxNzIyOTUxMjM3NDE1fQ.mTJQzi8_QCoIUPRMr9Mdv_w-zHdbx4LAcyTMPvsJimA' \
-H 'X-Line-Retry-Key: 123e4567-e89b-12d3-a456-426614174000' \
-d '{
    "to": "U97dad58774e23f2c3c3e89ba622f0318",
    "messages":[
        {
            "type":"text",
            "text":"Hello, world1"
        }
    ]
}'






curl -v -X POST https://api.line.me/oauth2/v2.1/token \
-H 'Content-Type: application/x-www-form-urlencoded' \
-d 'grant_type=authorization_code' \
-d 'code=1upNLtF7zgn8ay2QVD1y' \
--data-urlencode 'redirect_uri=http://localhost:8009/callback?token=example_token' \
-d 'client_id=2006022039' \
-d 'client_secret=aca6ddcdcb74086440643601cea6d7b9' \
-d 'code_verifier=wJKN8qz5t8SSI9lMFhBB6qwNkQBkuPZoCxzRhwLRUo1'




curl -v -X GET https://api.line.me/v2/profile \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.ZIO1nrqBzYmz9GiuWa3qedQr6Xlwvx73LtGOQRQE_gJH5ATnvPgohfLnWsDT3f887g_yUSojcZsFc3AmMpfvh8PlN1AQxmEjTvpG_-1cVrnktxJ3oCkHxZCp-Ozp9DDDjpHzXl93NMaEwNXh6lEhU339s9nAcU2QZ7Fgvhz03lI.udm107DCbQSxdpTcHv-A_VbgPjvmTUF-pbBzxmZzbzk'


curl -v -X POST https://api.line.me/oauth2/v3/token \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=client_credentials' \
--data-urlencode 'client_id=2006022038' \
--data-urlencode 'client_secret=68210c2c8958eaeb396586ee9ddb01e8'
# CONTACT API

## About

Sends contact request email to businesses from their frontends which I wrote.

## Flow

1. Receives post body (map of strings), and potentially image(s).
2. Checks the sending host
3. Gets the schema for the expected format of the received data from the database based off the host.
4.
  a. If no host found return 403.
  b. If schema found continue
5. Process images into s3 zip archive
6. Store received data in DB

7. Send "MESSAGE_READY" event with: the ID of the data in NoSQL, and the path of the image archive

-- Next Service --

8. Recveives event from queue
9. Retrieves the message and the template to use from the database
10. combines the data into the email template
11. Sends the email
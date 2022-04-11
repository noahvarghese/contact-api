# contact-api

Serverless API made to send templated emails for unauthenticated contact forms.

## Application Flow

1. Receives post body (JSON)
2. Checks that hostname was passed in the body, as well as a data object
3. If no hostname set, or body is empty return 400
4. Checks database for hostname
5. If hostname is not recognized return 403
6. Checks the database for a template associated with the hostname
7. If no template found return 404
8. Retrieves list of fields from database that are used in the template
9. Parses data object and checks against the fields retrieved from the database
10. If a mandatory field is missing return 400
11. Store stringified JSON in database in case email fails
12. Bind data to template
13. Send email
14. a. If failed return 500
    b. else set message sent in database and return 201

## Requirements

- MYSQL database configured
- AWS CLI already configured with Administrator permission
- [Docker installed](https://www.docker.com/community-edition)
- [Golang](https://golang.org)
- SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Development

Utilizes GitHub Actions for CI/CD

1. Fork repository
2. Execute ./storage/database/\_schema.sql against your database instance
3. Following the sample data found in ./storage/database/data.sql create the desired hosts,fields, and templates for your supported websites
4. Run SQL commands found in ./storage/database/\_users.sql REPLACING 'tmp_password' WITH THE PASSWORD OF YOUR CHOICE
5. Setup all environment variables as GitHub repository secrets (Replace DB_USER and DB_PWD with the user you just created)
6. Create .env with environment variables (Replace DB_USER and DB_PWD with the user you just created)

## Deployment

Adding/modifying templates and hosts is done through SQL and is seperate from the application code.

You can find your API Gateway Endpoint URL in the output values displayed after deployment.

1. Clone repository
2. Execute ./storage/database/\_schema.sql against your database instance
3. Following the sample data found in ./storage/database/data.sql create the desired hosts,fields, and templates for your supported websites
4. Run SQL commands found in ./storage/database/\_users.sql REPLACING 'tmp_password' WITH THE PASSWORD OF YOUR CHOICE
5. Create .env with environment variables (Replace DB_USER and DB_PWD with the user you just created)
6. Execute

   ```shell
   ./bin/deploy.sh
   ```

7. Go to the lambda environment and configure the application environment variables

## Environment Variables

Application type variables are needed for your app to run locally and in prod.
Build type variables are used during the build/deployment process and are not needed in your lambda environment.

| name                   | description                                                                           | type        |
| ---------------------- | ------------------------------------------------------------------------------------- | ----------- |
| AWS_ACCESS_KEY_ID      | The ID of the AWS IAM user                                                            | build       |
| AWS_SECRECT_ACCESS_KEY | AWS secret for the user                                                               | build       |
| DB_NAME                | The name of the database                                                              | application |
| DB_PORT                | The port to access the database on                                                    | application |
| DB_PWD                 | The password for the database user                                                    | application |
| DB_URL                 | The IP address or URL to access the database on                                       | application |
| DB_USER                | The username to log in to the database as                                             | application |
| S3_BUCKET              | The bucket to upload the build to                                                     | build       |
| S3_PREFIX              | If using a single s3 bucket for multiple application artifacts, this is the subfolder | build       |
| SMTP_PWD               | The password for your email server                                                    | application |
| SMTP_PORT              | The port to access the email server                                                   | application |
| SMTP_URL               | The URL or IP address of the email server                                             | application |
| SMTP_USER              | The user to login to the email server                                                 | application |

## Resources

https://pkg.go.dev/html/template
https://aws.amazon.com/premiumsupport/knowledge-center/custom-headers-api-gateway-lambda/
https://aws.amazon.com/blogs/compute/using-github-actions-to-deploy-serverless-applications/
https://github.com/aws/aws-lambda-go/blob/main/events/README_ApiGatewayEvent.md

## Setup process

### Installing dependencies & building the target

In this example we use the built-in `sam build` to automatically download all the dependencies and package our build target.  
Read more about [SAM Build here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli-command-reference-sam-build.html)

The `sam build` command is wrapped inside of the `Makefile`. To execute this simply run

```shell
make
```

### Local development

**Invoking function locally through local API Gateway**

```bash
./scripts/local.sh
```

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000`

### Packaging and deployment

Requires environemnt variables set in .env
Uses SAM CLI to build and deploy

```shell
./bin/deployment.sh
```

### Testing

Requires that the environment variables be loaded into memory.
Uses the `testing` package that is built-in in Golang.

```shell
./bin/test.sh
```

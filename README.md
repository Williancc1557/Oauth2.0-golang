# API de autenticação com Oauth2.0

Hello, this API was created to authenticate the user with good security and efficiency. I used to create this backend the clean architecture and I tried to cover all important parts of the code with tests.

First, I'll start talking about, what is Oauth2.0 and how I implemented it.

## Oauth2.0

Let's start by talking about how I set up the Oauth2.0 logic in this application. It consists of a way of authenticating with an application using the **OpenId**.

In Oauth2.0 we have 2 tokens, the `access-token` and the `refresh-token`. These tokens are important for building our system.

- `access-token`: It is associated with the user's ID and is usually temporary for a short period of time, however, we can use the decode on this token and obtain the user's ID. Once you have the user's ID, you can collect the information from the database that is related to that ID.

- `refresh-token`: This token is usually saved in the database with the user ID of the person who owns it. It's so that when the access-token becomes invalid, a new access-token can be generated without the user having to go to the login page again. In other words, as long as the refresh-token is valid, we will always have a valid access-token.

- `OpenId`: This is the way to store the user identification and use it in other applications. In this API I use OpenId in my JWT token as the claim sub, the subject is like the title of the token, in other words, the principal data of the token.

## `/api/auth/sign-up` **(POST)**

I created this route to register the user in the database, this route will save the user by email and password and then return the `accessToken` and `refreshToken` respectively. See below an example of how to request this route:

**BODY:**

```json
{
  "email": "test@example.com",
  "password": "valid_password"
}
```

**Response:**

```json
{
  "expiresIn": 300,
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiI2MzAxODU1YzQyMmJkY2Y1M2NjMzQ2YTUiLCJzdWIiOiJjbGllbnQiLCJpYXQiOjE2NjEwNDQwNjAsImV4cCI6MTY2MTA0NDM2MH0.CV_vO_lq0TBz3t7fW_9S1nUFDVpNXOV214_jSURpmbE",
  "refreshToken": "2615de11c12c7bcc3d74f9196"
}
```

## `/api/auth/sign-in` **(POST)**

This route will be used when the user no longer has a valid refresh token, you can direct them to the login screen. Returning and updating your refresh token. To do this, you must provide a valid account with email and password.

For example:

**BODY:**

```json
{
  "email": "test@example.com",
  "password": "valid_password"
}
```

**Response:**

```json
{
  "refreshToken": "sb9910f04e2cbafa604a69e1b"
}
```

But, as you can see, there's no access token in this response, but as I said before, we can get a new access token using this refresh token in the next route that I will talk about.

## `/api/auth/refresh-token` **(GET)**

This route was created with the purpose to use the refresh token, it will return an valid access token with the expires time of the access token.

**HEADER:**

```json
"refreshtoken": "ba9910f04e2cbafa604a69e1b"
```

**Response:**

```json
{
  "expiresIn": 300,
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiI2MmY4NWE4M2IwMDY0YzExODk0M2JlNzYiLCJzdWIiOiJjbGllbnQiLCJpYXQiOjE2NjExMTU5MzAsImV4cCI6MTY2MTExNjIzMH0.HeKmXNam6ds0X_xKskPtbjF68JHeod9TRrA9s_9kWms"
}
```

## Environment variables

This API needs some additional information to be able to turn on there are three environment variables, these variables will be used to connect with the PostgreSQL database by URL, write the secret of the JWT token, and put the port that the API will turn on with.

To be able to set variables, create a file named `.env` outside `src`, I'll demonstrate a variable below for example:

```env
  POSTGRE_URL=postgres://myuser:mypassword@localhost:5432/mydatabase?
  TOKEN_SECRET=tokenSecretExample
  PORT=1234
```

## Tests with testing

In this application I tried to cover all important parts of code with unit tests, like controllers, repositories, utils, and data. I did it because to refactor the code will be a simple task and will guarantee more security and efficiency when the developer change the code without fear of change the business rule.

To mock dependencies I used the following libraries:

- github.com/DATA-DOG/go-sqlmock
- github.com/golang/mock

The first library I used to mock sql queries and put an fake result for theses queries as you can see in the code below:

```go
t.Run("Success", func(t *testing.T) {
  repo, mock, db := setupMocks(t)
  defer db.Close()

  email := "test@example.com"
  query := regexp.QuoteMeta("SELECT * FROM users WHERE email = $1")

  rows := sqlmock.NewRows([]string{"id", "email", "password", "refresh_token"}).
    AddRow(1, email, "fake_hashed_password", "fake_refresh_token")

  mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)

  account, err := repo.Get(email)
  require.NoError(t, err)
  require.NotNil(t, account)
  require.Equal(t, email, account.Email)
})
```

And the next library `golang/mock` is used to mock interfaces, these interfaces are used as dependencies of the controller, data, and repositories. See below a simple example of mocking a dependency:

```go
t.Run("InvalidEmailCredentials", func(t *testing.T) {
  signInController, _, mockGetAccountByEmail, _, ctrl := setupMocks(t)
  defer ctrl.Finish()

  mockGetAccountByEmail.EXPECT().Get(email).Return(nil, errors.New("fake-error"))

  httpRequest := createHttpRequest(t, email, password)
  httpResponse := signInController.Handle(*httpRequest)

  verifyHttpResponse(t, httpResponse, http.StatusBadRequest, "invalid credentials")
})
```

In this case, the fake dependency will be `mockGetAccountByEmail`, as you can see, while I'm returning an error, if the email is not the same as the passed, an error will occur in the test.

### Commands

These are the commands you can use to run tests, or make a coverage of them.

- `make test`: This command will execute all tests in the `tests` file;

- `make coverage`: This test will generate a coverage of some specifics parts of the code, these important parts are: controllers, data, infra, utils.

## System Layers

We can also see a lot of layers that I created to maintain a good code separation and their responsibilities. Below I will talk about the most important layers used in this API.

- `data`: This layer was created to communicate the controller and repositories, but not directly;

- `domain`: This layer is the core of business rules, all dependencies used by controllers are declared here as an interface;

- `infra`: Is used to communicate with the database using repositories methods;

- `presentation`: This layer is used to create the controllers of the application;

- `utils`: I used this layer to put dependencies that I can use in all layers of my project.

# Tasks

## Creation of email validation

[ ] - Create any token and set it to last at least 5 minutes, then create a route to validate this token and verify the user's account.

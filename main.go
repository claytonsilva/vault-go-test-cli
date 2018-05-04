// create a vault client
client, err := api.NewClient(&api.Config{Address: url, HttpClient: httpClient})
if err != nil {
    panic(err)
}

// to pass the password
options := map[string]interface{}{
   "password": password,        
}

// the login path
// this is configurable, change userpass to ldap etc
path := fmt.Sprintf("auth/userpass/login/%s", username)

// PUT call to get a token
secret, err := client.Logical().Write(path, options)

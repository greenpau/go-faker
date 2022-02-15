# go-faker

Generate fake user identities for testing purposes, e.g. names, phone
numbers, addresses, etc.

Create a single fake identity with the password `foobar` and email
domain `contoso.com`:

```bash
faker --password foobar --domain=contoso.com user
```

Create 10 different user with the same attributes:

```bash
faker --password foobar --domain=contoso.com --count=10 user
```

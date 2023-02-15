# Pvectl

`pvectl` is a cli tool for managing Proxmox Virtual Environment clusters

## Installation
TODO

## Configuration
`pvectl` can be configured by using a file named `pvectl.yaml` in the current directory. Here's an example of the file:
```yaml
endpoint: https://pve.homelab.local:8006/api2/json
auth:
  method: apitoken
  tokenid: user@pam!token
  secret: your-api-secret
```

### Authentication
`pvectl` supports two authentication methods: `login` and `apitoken`.

```yaml
auth:
  method: login
  username: your-username
  password: your-password
```

```yaml
auth:
  method: apitoken
  tokenid: user@realm!token
  secret: your-api-secret
```

### Allow insecure connection
If using self-signed certificates, use:
```yaml
allowInsecureConnection: true # (default: false)
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

## License
[Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)

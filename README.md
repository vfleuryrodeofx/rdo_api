# RDO API
A simple API server in Go to fetch data from RDO

# Example

```bash
curl -X POST http://localhost:8080/login   -H "Content-Type: application/json"   -d '{"appname":"rdo_publish_pipeline","password":"rodeo!"}'
```


# TODO : 

- [ ] Add a proper Token generator using JWT
- [ ] Hashing password
- [ ] A simple py script to mimic a Python Client making requests to get data
- [ ] More routes !

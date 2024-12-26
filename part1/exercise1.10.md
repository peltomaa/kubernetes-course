```bash
$ go run worker/main.go
write to file 2024-12-26T11:05:52+02:00: 14a07d88-97fb-493e-8347-e69c16719f58
write to file 2024-12-26T11:05:57+02:00: 14a07d88-97fb-493e-8347-e69c16719f58
write to file 2024-12-26T11:06:02+02:00: 14a07d88-97fb-493e-8347-e69c16719f58
$ go run web/main.go
Starting server
Server started in port 3000
```

## another terminal window
```bash
$ curl http://localhost:3000
2024-12-26T11:06:02+02:00: 14a07d88-97fb-493e-8347-e69c16719f58

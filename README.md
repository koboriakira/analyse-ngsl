# analyse-ngsl

## Docker

Create image.

```bash
docker build -t go-image .
```

Run for Windows.

```Command Prompt
docker run -it --rm --name analyse-ngsl -v C:\git\analyse-ngsl:/go go-image
```

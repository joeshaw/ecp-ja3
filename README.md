# Compute@Edge service for displaying JA3 fingerprints

Like it says on the tin.  This prints the JA3 string and the MD5ed
fingerprint.

```
$ curl https://ja3.joeshaw.org
771,4866-4867-4865-49196-49200-159-52393-52392-52394-49195-49199-158-49188-49192-107-49187-49191-103-49162-49172-57-49161-49171-51-157-156-61-60-53-47-255,0-11-10-13172-16-22-23-49-13-43-45-51-21,29-23-30-25-24,0-1-2

f436b9416f37d134cadd04886327d3e8
```

## Development / deployment

`fastly compute build` to compile

`fastly compute serve` to run locally, though it doesn't have TLS so it's not very useful.

`fastly compute deploy` to deploy the service.

If you want to deploy this to your own service, you will need to unset `service_id` in `fastly.toml`.

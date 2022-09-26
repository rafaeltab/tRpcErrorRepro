in pnpm-lock.yaml
add quotes around the tarball string = no errors
no quotes around the tarball string = errors:

```
panic: yaml: line 20849: did not find expected ',' or '}'

goroutine 1 [running]:
main.DecodePnpmLockfile({0xc000180000, 0xb444a, 0xb444b})
        <url to root>/hello.go:30 +0x78
main.main()
        <url to root>/hello.go:17 +0x76
```

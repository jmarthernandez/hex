```
func main() {
	rom, err := ioutil.ReadFile("path/to/rom.ch8")
	if err != nil {
        panic(err)
    }
	hex.Hexdump(rom, 4)
}
```

outputs
```
00000000 12 25 53 50
00000004 41 43 45 20
00000008 49 4e 56 41
0000000c 44 45 52 53
00000010 20 30 2e 39
00000014 31 20 42 79
```
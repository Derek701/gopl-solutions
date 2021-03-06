package main

// Compact const declarations for kilobyte through yottabyte.
const (
	KB = 1000    // kilobyte  = 1000
	MB = KB * KB // megabyte  = 1000²
	GB = MB * KB // gigabyte  = 1000³
	TB = GB * KB // terabyte  = 1000⁴ (exceeds 1 << 32)
	PB = TB * KB // petabyte  = 1000⁵
	EB = PB * KB // exabyte   = 1000⁶
	ZB = EB * KB // zettabyte = 1000⁷ (exceeds 1 << 64)
	YB = ZB * KB // yottabyte = 1000⁸
)

func main() {
	// Do nothing.
}

// Code snippet of compact const declarations for kB through YB.

const (
	kB = 1000    // kilobyte  = 1000
	MB = kB * kB // megabyte  = 1000²
	GB = MB * kB // gigabyte  = 1000³
	TB = GB * kB // terabyte  = 1000⁴ (exceeds 1 << 32)
	PB = TB * kB // petabyte  = 1000⁵
	EB = PB * kB // exabyte   = 1000⁶
	ZB = EB * kB // zettabyte = 1000⁷ (exceeds 1 << 64)
	YB = ZB * kB // yottabyte = 1000⁸
)

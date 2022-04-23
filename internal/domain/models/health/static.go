package health

const (
	Alloc         = "alloc"
	BuckHashSys   = "buckhashsys"
	Frees         = "frees"
	GCCPUFraction = "gccpufraction"
	GCSys         = "gcsys"
	HeapAlloc     = "heapalloc"
	HeapIdle      = "heapidle"
	HeapInuse     = "heapinuse"
	HeapObjects   = "heapobjects"
	HeapReleased  = "heapreleased"
	HeapSys       = "heapsys"
	LastGC        = "lastgc"
	Lookups       = "lookups"
	MCacheInuse   = "mcacheinuse"
	MCacheSys     = "mcachesys"
	MSpanInuse    = "mspaninuse"
	MSpanSys      = "mspansys"
	Mallocs       = "mallocs"
	NextGC        = "nextgs"
	NumForcedGC   = "numforcedgc"
	NumGC         = "numgc"
	OtherSys      = "othersys"
	PauseTotalNs  = "pausetotalns"
	StackInuse    = "stackinuse"
	StackSys      = "stacksys"
	Sys           = "sys"
	TotalAlloc    = "totalalloc"
	PollCount     = "pollcount"
	RandomValue   = "randomvalue"
)

const (
	TypeGauge   = "gauge"
	TypeCounter = "counter"
)

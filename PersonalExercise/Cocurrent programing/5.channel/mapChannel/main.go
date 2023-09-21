package main

func main() {
	var mapchan chan map[string]string
	mapchan = make(chan map[string]string, 21)
	m1 := make(map[string]string, 10)
	m1["adddress1"] = "tianjing"
	m1["adddress2"] = "shangjing"
	m2 := make(map[string]string, 20)
	m2["man1"] = "zs"
	m2["man2"] = "as"
	mapchan <- m1
	mapchan <- m2
}

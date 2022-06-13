package others

import (
	"bytes"
	"fmt"
)


func sliceDemo() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')
    
	dir1 := path[:sepIndex] // 这种方式会
	// dir1 := path[:sepIndex:sepIndex] 这种不会导致共享内存
	dir2 := path[sepIndex+1:]

	// dir1 and dir2 sharing memory
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB
	
	dir1 = append(dir1,"suffix"...)

	// 因为cap足够，所以append后数据扩展到dir2
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB
}
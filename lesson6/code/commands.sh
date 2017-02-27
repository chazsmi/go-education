// START OMIT
go test -run=NONE -bench=. ./... > old.txt
// END OMIT
// START SECOND OMIT
go test -run=NONE -bench=. ./... > new.txt
benchcmp old.txt new.txt
// END SECOND OMIT
// START THIRD OMIT
go get github.com/chazsmi/profiling-example
profiling-example
// END THIRD OMIT
// START FOURTH OMIT
go tool pprof http://localhost:8080/debug/pprof/heap
// END FOURTH OMIT
// START FIFTH OMIT
go tool pprof http://localhost:8080/debug/pprof/profile
// END FIFTH OMIT
// START SIXTH OMIT
curl localhost:8080
You have [eggs apples bananas]
// END SIXTH OMIT
// START SEVENTH OMIT
(pprof) png > graph.png
// END SEVENTH OMIT


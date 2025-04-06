go test ./... -v -cover
	github.com/brotigen23/goph-keeper/server/cmd/server		coverage: 0.0% of statements
?   	github.com/brotigen23/goph-keeper/server/internal/dto	[no test files]
?   	github.com/brotigen23/goph-keeper/server/internal/middleware	[no test files]
?   	github.com/brotigen23/goph-keeper/server/internal/model	[no test files]
?   	github.com/brotigen23/goph-keeper/server/internal/repository	[no test files]
?   	github.com/brotigen23/goph-keeper/server/internal/repository/mock	[no test files]
	github.com/brotigen23/goph-keeper/server/internal/handler		coverage: 0.0% of statements
	github.com/brotigen23/goph-keeper/server/internal/app		coverage: 0.0% of statements
=== RUN   TestConfigFromEnv
--- PASS: TestConfigFromEnv (0.00s)
PASS
coverage: 0.0% of statements
ok  	github.com/brotigen23/goph-keeper/server/internal/config	(cached)	coverage: 0.0% of statements
	github.com/brotigen23/goph-keeper/server/pkg/crypt		coverage: 0.0% of statements
	github.com/brotigen23/goph-keeper/server/pkg/migration		coverage: 0.0% of statements
	github.com/brotigen23/goph-keeper/server/internal/service		coverage: 0.0% of statements
	github.com/brotigen23/goph-keeper/server/pkg/pgErrors		coverage: 0.0% of statements
	github.com/brotigen23/goph-keeper/server/internal/server		coverage: 0.0% of statements
=== RUN   TestCreateMetadata
=== RUN   TestCreateMetadata/Test_OK
--- PASS: TestCreateMetadata (0.00s)
    --- PASS: TestCreateMetadata/Test_OK (0.00s)
=== RUN   TestGetMetadataByID
=== RUN   TestGetMetadataByID/Test_OK
--- PASS: TestGetMetadataByID (0.00s)
    --- PASS: TestGetMetadataByID/Test_OK (0.00s)
=== RUN   TestGetMetadataByRowID
=== RUN   TestGetMetadataByRowID/Test_OK
time=2025-04-06T18:18:45.368+03:00 level=ERROR msg="Error occured" error="Query: could not match actual sql: \"SELECT id, data, created_at, updated_at FROM metadata WHERE table_name = $1 AND row_id = $2\" with expected regexp \"SELECT id, data, created_at, updated_at FROM metadata WHERE table_name = ? AND row_id = ?\"" file=/home/user/Coding/go/goph-keeper/server/internal/repository/postgres/metadata.go line=131
    metadata_test.go:235: 
        	Error Trace:	/home/user/Coding/go/goph-keeper/server/internal/repository/postgres/metadata_test.go:235
        	Error:      	Not equal: 
        	            	expected: <nil>(<nil>)
        	            	actual  : *errors.errorString(&errors.errorString{s:"Query: could not match actual sql: \"SELECT id, data, created_at, updated_at FROM metadata WHERE table_name = $1 AND row_id = $2\" with expected regexp \"SELECT id, data, created_at, updated_at FROM metadata WHERE table_name = ? AND row_id = ?\""})
        	Test:       	TestGetMetadataByRowID/Test_OK
--- FAIL: TestGetMetadataByRowID (0.00s)
    --- FAIL: TestGetMetadataByRowID/Test_OK (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x709a7a]

goroutine 30 [running]:
testing.tRunner.func1.2({0x7494e0, 0xa44f50})
	/usr/lib/golang/src/testing/testing.go:1632 +0x230
testing.tRunner.func1()
	/usr/lib/golang/src/testing/testing.go:1635 +0x35e
panic({0x7494e0?, 0xa44f50?})
	/usr/lib/golang/src/runtime/panic.go:791 +0x132
github.com/brotigen23/goph-keeper/server/internal/repository/postgres.TestGetMetadataByRowID.func1(0xc0001e31e0)
	/home/user/Coding/go/goph-keeper/server/internal/repository/postgres/metadata_test.go:237 +0x51a
testing.tRunner(0xc0001e31e0, 0xc00019d7d0)
	/usr/lib/golang/src/testing/testing.go:1690 +0xf4
created by testing.(*T).Run in goroutine 28
	/usr/lib/golang/src/testing/testing.go:1743 +0x390
FAIL	github.com/brotigen23/goph-keeper/server/internal/repository/postgres	0.006s
=== RUN   TestLogger
--- PASS: TestLogger (0.00s)
PASS
coverage: 60.0% of statements
ok  	github.com/brotigen23/goph-keeper/server/pkg/logger	0.002s	coverage: 60.0% of statements
FAIL

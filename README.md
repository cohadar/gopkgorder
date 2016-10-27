## Go standard library package order, from least to most dependent.
Hover to see dependency list. Number in curly braces is just a number of direct imports. Each row is one level in topsorted graph.

[errors{0}](https://golang.org/pkg/errors "nil")
[sort{0}](https://golang.org/pkg/sort "nil")
[unicode{0}](https://golang.org/pkg/unicode "nil")
[unsafe{0}](https://golang.org/pkg/unsafe "nil")

[container{1}](https://golang.org/pkg/container "sort")

[strconv{3}](https://golang.org/pkg/strconv "errors, math, unicode/utf8")

[sync{3}](https://golang.org/pkg/sync "internal/race, runtime, unsafe")

[reflect{5}](https://golang.org/pkg/reflect "math, runtime, strconv, sync, unsafe")
[syscall{5}](https://golang.org/pkg/syscall "errors, internal/race, runtime, sync, unsafe")

[bytes{4}](https://golang.org/pkg/bytes "errors, io, unicode, unicode/utf8")
[strings{4}](https://golang.org/pkg/strings "errors, io, unicode, unicode/utf8")
[time{4}](https://golang.org/pkg/time "errors, runtime, sync, syscall")

[bufio{4}](https://golang.org/pkg/bufio "bytes, errors, io, unicode/utf8")
[hash{2}](https://golang.org/pkg/hash "io, sync")

[fmt{8}](https://golang.org/pkg/fmt "errors, io, math, os, reflect, strconv, sync, unicode/utf8")
[path{6}](https://golang.org/pkg/path "errors, os, runtime, sort, strings, unicode/utf8")
[regexp{8}](https://golang.org/pkg/regexp "bytes, io, sort, strconv, strings, sync, unicode, unicode/utf8")

[context{5}](https://golang.org/pkg/context "errors, fmt, reflect, sync, time")
[flag{8}](https://golang.org/pkg/flag "errors, fmt, io, os, reflect, sort, strconv, time")
[index{5}](https://golang.org/pkg/index "bytes, encoding/binary, io, regexp, sort")
[io{8}](https://golang.org/pkg/io "bytes, errors, os, path/filepath, sort, strconv, sync, time")
[math{9}](https://golang.org/pkg/math "bytes, encoding/binary, errors, fmt, io, strconv, strings, sync, unsafe")
[runtime{11}](https://golang.org/pkg/runtime "bufio, bytes, fmt, io, os, sort, strings, sync, text/tabwriter, time, unsafe")
[vendor{11}](https://golang.org/pkg/vendor "bytes, errors, fmt, io, os, runtime, strings, sync, syscall, unicode/utf8, unsafe")

[archive{19}](https://golang.org/pkg/archive "bufio, bytes, compress/flate, encoding/binary, errors, fmt, hash, hash/crc32, io, io/ioutil, math, os, path, sort, strconv, strings, sync, syscall, time")
[compress{13}](https://golang.org/pkg/compress "bufio, encoding/binary, errors, fmt, hash, hash/adler32, hash/crc32, io, math, sort, strconv, sync, time")
[database{10}](https://golang.org/pkg/database "errors, fmt, io, reflect, runtime, sort, strconv, sync, sync/atomic, time")
[encoding{19}](https://golang.org/pkg/encoding "bufio, bytes, errors, fmt, io, math, math/big, os, reflect, runtime, sort, strconv, strings, sync, sync/atomic, time, unicode, unicode/utf16, unicode/utf8")
[os{14}](https://golang.org/pkg/os "bytes, context, errors, fmt, io, path/filepath, runtime, strconv, strings, sync, sync/atomic, syscall, time, unsafe")
[testing{20}](https://golang.org/pkg/testing "bytes, errors, flag, fmt, io, log, math, math/rand, os, reflect, runtime, runtime/debug, runtime/pprof, runtime/trace, sort, strconv, strings, sync, sync/atomic, time")
[text{16}](https://golang.org/pkg/text "bytes, errors, fmt, io, io/ioutil, net/url, os, path/filepath, reflect, runtime, sort, strconv, strings, sync, unicode, unicode/utf8")

[debug{12}](https://golang.org/pkg/debug "bytes, compress/zlib, encoding/binary, errors, fmt, io, os, path, sort, strconv, strings, sync")
[html{13}](https://golang.org/pkg/html "bytes, encoding/json, fmt, io, io/ioutil, path/filepath, reflect, strings, sync, text/template, text/template/parse, unicode, unicode/utf8")
[image{11}](https://golang.org/pkg/image "bufio, bytes, compress/lzw, compress/zlib, encoding/binary, errors, fmt, hash, hash/crc32, io, strconv")
[internal{16}](https://golang.org/pkg/internal "bufio, bytes, flag, fmt, io, math/rand, os, os/exec, path/filepath, runtime, sort, strconv, strings, sync, testing, unsafe")
[log{9}](https://golang.org/pkg/log "errors, fmt, io, net, os, runtime, strings, sync, time")

[mime{15}](https://golang.org/pkg/mime "bufio, bytes, crypto/rand, encoding/base64, errors, fmt, io, io/ioutil, net/textproto, os, sort, strings, sync, unicode, unicode/utf8")

[crypto{25}](https://golang.org/pkg/crypto "bufio, bytes, container/list, encoding/asn1, encoding/binary, encoding/hex, encoding/pem, errors, fmt, hash, io, io/ioutil, math/big, net, os, os/exec, runtime, strconv, strings, sync, sync/atomic, syscall, time, unicode/utf8, unsafe")

[expvar{12}](https://golang.org/pkg/expvar "bytes, encoding/json, fmt, log, math, net/http, os, runtime, sort, strconv, sync, sync/atomic")
[go{50}](https://golang.org/pkg/go "archive/zip, bufio, bytes, container/heap, debug/elf, encoding/binary, encoding/gob, encoding/json, encoding/xml, errors, expvar, flag, fmt, html, html/template, index/suffixarray, io, io/ioutil, log, math, math/big, net, net/http, net/http/httptest, net/http/httputil, net/http/pprof, net/url, os, os/exec, path, path/filepath, reflect, regexp, runtime, runtime/debug, runtime/pprof, runtime/trace, sort, strconv, strings, sync, sync/atomic, syscall, text/scanner, text/tabwriter, text/template, time, unicode, unicode/utf8, unsafe")

[cmd{67}](https://golang.org/pkg/cmd "bufio, bytes, compress/gzip, container/heap, crypto/md5, crypto/sha1, crypto/sha256, crypto/tls, debug/dwarf, debug/elf, debug/gosym, debug/macho, debug/pe, debug/plan9obj, encoding/base64, encoding/binary, encoding/hex, encoding/json, encoding/xml, errors, flag, fmt, go/ast, go/build, go/constant, go/doc, go/format, go/importer, go/parser, go/printer, go/scanner, go/token, go/types, html, html/template, internal/singleflight, internal/trace, io, io/ioutil, log, math, math/rand, net, net/http, net/url, os, os/exec, os/signal, path, path/filepath, reflect, regexp, runtime, runtime/debug, runtime/pprof, sort, strconv, strings, sync, syscall, text/scanner, text/tabwriter, text/template, time, unicode, unicode/utf8, unsafe")
[net{47}](https://golang.org/pkg/net "bufio, bytes, compress/gzip, container/list, context, crypto/hmac, crypto/md5, crypto/tls, encoding/base64, encoding/binary, encoding/gob, encoding/json, errors, flag, fmt, html/template, internal/nettrace, internal/singleflight, io, io/ioutil, log, math, math/rand, mime, mime/multipart, os, os/exec, path, path/filepath, reflect, regexp, runtime, runtime/pprof, runtime/trace, sort, strconv, strings, sync, sync/atomic, syscall, time, unicode, unicode/utf8, unsafe, vendor/golang_org/x/net/http2/hpack, vendor/golang_org/x/net/lex/httplex, vendor/golang_org/x/net/route")

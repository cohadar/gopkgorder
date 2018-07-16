## Go standard library package order, from least to most dependent.
Hover to see dependency list. Number in curly braces is just a number of direct imports. Each row is one level in topsorted graph.
Go Version: go1.10.2

[errors{0}](https://golang.org/pkg/errors "nil")
[unicode{0}](https://golang.org/pkg/unicode "nil")
[unsafe{0}](https://golang.org/pkg/unsafe "nil")

[strconv{3}](https://golang.org/pkg/strconv "errors, math, unicode/utf8")

[sync{3}](https://golang.org/pkg/sync "internal/race, runtime, unsafe")

[plugin{3}](https://golang.org/pkg/plugin "errors, sync, unsafe")
[reflect{7}](https://golang.org/pkg/reflect "math, runtime, strconv, sync, unicode, unicode/utf8, unsafe")
[syscall{5}](https://golang.org/pkg/syscall "errors, internal/race, runtime, sync, unsafe")

[bytes{5}](https://golang.org/pkg/bytes "errors, internal/cpu, io, unicode, unicode/utf8")
[sort{1}](https://golang.org/pkg/sort "reflect")
[strings{6}](https://golang.org/pkg/strings "errors, internal/cpu, io, unicode, unicode/utf8, unsafe")
[time{4}](https://golang.org/pkg/time "errors, runtime, sync, syscall")

[bufio{4}](https://golang.org/pkg/bufio "bytes, errors, io, unicode/utf8")
[container{1}](https://golang.org/pkg/container "sort")
[hash{5}](https://golang.org/pkg/hash "errors, internal/cpu, io, sync, unsafe")

[regexp{8}](https://golang.org/pkg/regexp "bytes, io, sort, strconv, strings, sync, unicode, unicode/utf8")

[fmt{8}](https://golang.org/pkg/fmt "errors, io, math, os, reflect, strconv, sync, unicode/utf8")
[index{5}](https://golang.org/pkg/index "bytes, encoding/binary, io, regexp, sort")
[path{6}](https://golang.org/pkg/path "errors, os, runtime, sort, strings, unicode/utf8")

[context{5}](https://golang.org/pkg/context "errors, fmt, reflect, sync, time")
[flag{9}](https://golang.org/pkg/flag "errors, fmt, io, os, reflect, sort, strconv, strings, time")
[io{9}](https://golang.org/pkg/io "bytes, errors, os, path/filepath, sort, strconv, sync, sync/atomic, time")
[math{10}](https://golang.org/pkg/math "bytes, encoding/binary, errors, fmt, internal/cpu, io, strconv, strings, sync, unsafe")

[archive{23}](https://golang.org/pkg/archive "bufio, bytes, compress/flate, encoding/binary, errors, fmt, hash, hash/crc32, io, io/ioutil, math, os, os/user, path, reflect, runtime, sort, strconv, strings, sync, syscall, time, unicode/utf8")
[compress{14}](https://golang.org/pkg/compress "bufio, encoding/binary, errors, fmt, hash, hash/adler32, hash/crc32, io, math, math/bits, sort, strconv, sync, time")
[encoding{20}](https://golang.org/pkg/encoding "bufio, bytes, errors, fmt, io, math, math/big, math/bits, os, reflect, runtime, sort, strconv, strings, sync, sync/atomic, time, unicode, unicode/utf16, unicode/utf8")
[os{16}](https://golang.org/pkg/os "bytes, context, errors, fmt, internal/poll, internal/testlog, io, path/filepath, runtime, strconv, strings, sync, sync/atomic, syscall, time, unsafe")
[text{16}](https://golang.org/pkg/text "bytes, errors, fmt, io, io/ioutil, net/url, os, path/filepath, reflect, runtime, sort, strconv, strings, sync, unicode, unicode/utf8")

[database{13}](https://golang.org/pkg/database "context, errors, fmt, io, reflect, runtime, sort, strconv, sync, sync/atomic, time, unicode, unicode/utf8")
[debug{12}](https://golang.org/pkg/debug "bytes, compress/zlib, encoding/binary, errors, fmt, io, os, path, sort, strconv, strings, sync")
[html{13}](https://golang.org/pkg/html "bytes, encoding/json, fmt, io, io/ioutil, path/filepath, reflect, strings, sync, text/template, text/template/parse, unicode, unicode/utf8")
[image{11}](https://golang.org/pkg/image "bufio, bytes, compress/lzw, compress/zlib, encoding/binary, errors, fmt, hash, hash/crc32, io, strconv")
[internal{20}](https://golang.org/pkg/internal "bufio, bytes, errors, flag, fmt, io, math/rand, os, os/exec, path/filepath, runtime, sort, strconv, strings, sync, sync/atomic, syscall, testing, time, unsafe")
[log{9}](https://golang.org/pkg/log "errors, fmt, io, net, os, runtime, strings, sync, time")
[runtime{19}](https://golang.org/pkg/runtime "bufio, bytes, compress/gzip, context, encoding/binary, errors, fmt, io, io/ioutil, math, os, regexp, sort, strconv, strings, sync, text/tabwriter, time, unsafe")

[mime{15}](https://golang.org/pkg/mime "bufio, bytes, crypto/rand, encoding/base64, errors, fmt, io, io/ioutil, net/textproto, os, sort, strings, sync, unicode, unicode/utf8")
[testing{24}](https://golang.org/pkg/testing "bufio, bytes, errors, flag, fmt, internal/race, internal/testlog, io, log, math, math/rand, os, reflect, regexp, runtime, runtime/debug, runtime/pprof, runtime/trace, sort, strconv, strings, sync, sync/atomic, time")

[crypto{35}](https://golang.org/pkg/crypto "bufio, bytes, container/list, encoding/asn1, encoding/binary, encoding/hex, encoding/pem, errors, fmt, hash, internal/cpu, io, io/ioutil, math, math/big, net, net/url, os, os/exec, os/user, path/filepath, reflect, runtime, strconv, strings, sync, sync/atomic, syscall, time, unicode/utf8, unsafe, vendor/golang_org/x/crypto/chacha20poly1305, vendor/golang_org/x/crypto/cryptobyte, vendor/golang_org/x/crypto/cryptobyte/asn1, vendor/golang_org/x/crypto/curve25519")
[vendor{28}](https://golang.org/pkg/vendor "bytes, container/list, crypto/cipher, crypto/subtle, encoding/asn1, encoding/binary, errors, fmt, io, io/ioutil, log, math, math/big, math/rand, net, net/url, os, reflect, runtime, sort, strconv, strings, sync, syscall, testing, time, unicode/utf8, unsafe")

[go{87}](https://golang.org/pkg/go "archive/tar, archive/zip, bufio, bytes, compress/bzip2, compress/flate, compress/gzip, compress/zlib, container/heap, context, crypto, crypto/aes, crypto/cipher, crypto/des, crypto/dsa, crypto/ecdsa, crypto/elliptic, crypto/hmac, crypto/md5, crypto/rand, crypto/rc4, crypto/rsa, crypto/sha1, crypto/sha256, crypto/sha512, crypto/subtle, crypto/tls, crypto/x509, crypto/x509/pkix, debug/elf, encoding/asn1, encoding/base64, encoding/binary, encoding/gob, encoding/hex, encoding/json, encoding/pem, encoding/xml, errors, expvar, flag, fmt, hash, html, html/template, image, image/jpeg, index/suffixarray, io, io/ioutil, log, math, math/big, math/bits, math/rand, net, net/http, net/http/httptest, net/http/httputil, net/http/pprof, net/textproto, net/url, os, os/exec, os/user, path, path/filepath, reflect, regexp, runtime, runtime/debug, runtime/pprof, runtime/trace, sort, strconv, strings, sync, sync/atomic, syscall, text/scanner, text/tabwriter, text/template, time, unicode, unicode/utf16, unicode/utf8, unsafe")

[expvar{12}](https://golang.org/pkg/expvar "bytes, encoding/json, fmt, log, math, net/http, os, runtime, sort, strconv, sync, sync/atomic")

[net{52}](https://golang.org/pkg/net "bufio, bytes, compress/gzip, container/list, context, crypto/hmac, crypto/md5, crypto/rand, crypto/tls, crypto/x509, encoding/base64, encoding/binary, encoding/gob, encoding/json, errors, flag, fmt, html/template, internal/nettrace, internal/poll, internal/singleflight, io, io/ioutil, log, math, math/rand, mime, mime/multipart, os, os/exec, path, path/filepath, reflect, regexp, runtime, runtime/pprof, runtime/trace, sort, strconv, strings, sync, sync/atomic, syscall, time, unicode, unicode/utf8, unsafe, vendor/golang_org/x/net/http2/hpack, vendor/golang_org/x/net/idna, vendor/golang_org/x/net/lex/httplex, vendor/golang_org/x/net/proxy, vendor/golang_org/x/net/route")

[cmd{72}](https://golang.org/pkg/cmd "bufio, bytes, compress/gzip, container/heap, container/list, crypto/md5, crypto/sha1, crypto/sha256, crypto/tls, debug/dwarf, debug/elf, debug/gosym, debug/macho, debug/pe, debug/plan9obj, encoding/base64, encoding/binary, encoding/hex, encoding/json, encoding/xml, errors, flag, fmt, go/ast, go/build, go/constant, go/doc, go/format, go/importer, go/parser, go/printer, go/scanner, go/token, go/types, hash, html, html/template, internal/singleflight, internal/trace, io, io/ioutil, log, math, math/big, math/rand, net, net/http, net/url, os, os/exec, os/signal, path, path/filepath, reflect, regexp, runtime, runtime/debug, runtime/pprof, runtime/trace, sort, strconv, strings, sync, syscall, testing, text/scanner, text/tabwriter, text/template, time, unicode, unicode/utf8, unsafe")


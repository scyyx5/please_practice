subinclude("///go//build_defs:go")

go_library(
    name = "uuid",
    srcs = [
        "dce.go",
        "doc.go",
        "hash.go",
        "marshal.go",
        "node.go",
        "node_net.go",
        "null.go",
        "sql.go",
        "time.go",
        "util.go",
        "uuid.go",
        "version1.go",
        "version4.go",
    ],
    _module = "github.com/google/uuid@v1.3.0",
    _subrepo = "third_party/go/github.com_google_uuid",
    cover = False,
    labels = [
        "go_replace_directive",
        "go_module_path:github.com/google/uuid",
        "go_module:github.com/google/uuid@v1.3.0",
    ],
    visibility = ["PUBLIC"],
)

filegroup(
    name = "installs",
    exported_deps = [],
    visibility = ["PUBLIC"],
)

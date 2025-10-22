subinclude("///go//build_defs:go")

go_library(
    name = "testify",
    srcs = ["doc.go"],
    _module = "github.com/stretchr/testify@v1.11.1",
    _subrepo = "third_party/go/github.com_stretchr_testify",
    cover = False,
    labels = [
        "go_module_path:github.com/stretchr/testify",
        "go_module:github.com/stretchr/testify@v1.11.1",
    ],
    visibility = ["PUBLIC"],
)

filegroup(
    name = "installs",
    exported_deps = [],
    visibility = ["PUBLIC"],
)

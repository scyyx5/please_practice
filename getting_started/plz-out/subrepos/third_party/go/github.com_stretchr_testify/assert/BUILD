subinclude("///go//build_defs:go")

go_library(
    name = "assert",
    srcs = [
        "assertion_compare.go",
        "assertion_format.go",
        "assertion_forward.go",
        "assertion_order.go",
        "assertions.go",
        "doc.go",
        "errors.go",
        "forward_assertions.go",
        "http_assertions.go",
    ],
    _module = "github.com/stretchr/testify@v1.11.1",
    _subrepo = "third_party/go/github.com_stretchr_testify",
    cover = False,
    labels = [
        "go_module_path:github.com/stretchr/testify",
        "go_module:github.com/stretchr/testify@v1.11.1",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "///third_party/go/github.com_davecgh_go-spew//spew",
        "///third_party/go/github.com_pmezard_go-difflib//difflib",
        "//assert/yaml",
    ],
)

subinclude("///go//build_defs:go")

go_library(
    name = "yaml.v3",
    srcs = [
        "apic.go",
        "decode.go",
        "emitterc.go",
        "encode.go",
        "parserc.go",
        "readerc.go",
        "resolve.go",
        "scannerc.go",
        "sorter.go",
        "writerc.go",
        "yaml.go",
        "yamlh.go",
        "yamlprivateh.go",
    ],
    _module = "gopkg.in/yaml.v3@v3.0.1",
    _subrepo = "third_party/go/gopkg.in_yaml.v3",
    cover = False,
    labels = [
        "go_module_path:gopkg.in/yaml.v3",
        "go_module:gopkg.in/yaml.v3@v3.0.1",
    ],
    visibility = ["PUBLIC"],
)

filegroup(
    name = "installs",
    exported_deps = [],
    visibility = ["PUBLIC"],
)

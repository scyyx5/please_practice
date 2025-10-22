go_library(
    name = "greetings",
    srcs = ["greetings.go"],
    visibility = ["//src/..."],
)

go_test(
    name = "greetings_test",
    srcs = ["greetings_test.go"],
    external = True,
    deps = [
        ":greetings",
        "///third_party/go/github.com_stretchr_testify//assert",
    ],
)

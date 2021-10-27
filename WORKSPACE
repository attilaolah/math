# Declares that this directory is the root of a Bazel workspace.
# See https://docs.bazel.build/versions/master/build-ref.html#workspace.
workspace(
    # How this workspace would be referenced with absolute labels from another workspace
    name = "math",
)

load("//:package.bzl", "github_archive")

github_archive(
    name = "bazel_federation",
    repo = "bazelbuild/bazel-federation",
    sha256 = "506dfbfd74ade486ac077113f48d16835fdf6e343e1d4741552b450cfc2efb53",
    url = "releases/download/{version}/bazel_federation-{version}.tar.gz",
    version = "0.0.1",
)

github_archive(
    name = "rules_python",
    repo = "bazelbuild/rules_python",
    url = "releases/download/{version}/rules_python-{version}.tar.gz",
    version="0.4.0",
    sha256="954aa89b491be4a083304a2cb838019c8b8c3720a7abb9c4cb81ac7a24230cea",
)

load(
    "@bazel_federation//:repositories.bzl",
    "bazel_gazelle",
    "bazel_skylib",
    "rules_go",
)

bazel_gazelle()

bazel_skylib()

load("@bazel_federation//setup:bazel_skylib.bzl", "bazel_skylib_setup")

bazel_skylib_setup()

rules_go()

load("@bazel_federation//setup:rules_go.bzl", "rules_go_setup")

rules_go_setup()

github_archive(
    name = "com_github_atlassian_bazel_tools",
    repo = "atlassian/bazel-tools",
    sha256 = "6a991df7a79db78229cbabded60c98641400f31fc88244847b519fbb904fc360",
    strip_prefix = "bazel-tools-{version}",
    url = "archive/{version}.zip",
    version = "1056bf1d619b432063841df24b9eca86eb716527",
)

load("@com_github_atlassian_bazel_tools//golangcilint:deps.bzl", "golangcilint_dependencies")

golangcilint_dependencies()

# gazelle:repo bazel_gazelle

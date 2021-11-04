# Declares that this directory is the root of a Bazel workspace.
# See https://docs.bazel.build/versions/master/build-ref.html#workspace.
workspace(
    # How this workspace would be referenced with absolute labels from another workspace
    name = "math",
)

load("//:package.bzl", "github_archive")

github_archive(
    name = "rules_python",
    repo = "bazelbuild/rules_python",
    url = "releases/download/{version}/rules_python-{version}.tar.gz",
    version="0.4.0",
    sha256="954aa89b491be4a083304a2cb838019c8b8c3720a7abb9c4cb81ac7a24230cea",
)

github_archive(
    name = "io_bazel_rules_go",
    repo = "bazelbuild/rules_go",
    url = "releases/download/v{version}/rules_go-v{version}.zip",
    version="0.29.0",
    sha256 = "2b1641428dff9018f9e85c0384f03ec6c10660d935b750e3fa1492a281a53b0f",
)

github_archive(
    name = "bazel_gazelle",
    repo = "bazelbuild/bazel-gazelle",
    url = "releases/download/v{version}/bazel-gazelle-v{version}.tar.gz",
    version="0.24.0",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
)

github_archive(
    name = "bazel_skylib",
    repo = "bazelbuild/bazel-skylib",
    url = "releases/download/{version}/bazel-skylib-{version}.tar.gz",
    version = "1.1.1",
    sha256 = "c6966ec828da198c5d9adbaa94c05e3a1c7f21bd012a0b29ba8ddbccb2c93b0d",
)

github_archive(
    name = "com_github_atlassian_bazel_tools",
    repo = "atlassian/bazel-tools",
    sha256 = "6a991df7a79db78229cbabded60c98641400f31fc88244847b519fbb904fc360",
    strip_prefix = "bazel-tools-{version}",
    url = "archive/{version}.zip",
    version = "1056bf1d619b432063841df24b9eca86eb716527",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")
load("@com_github_atlassian_bazel_tools//golangcilint:deps.bzl", "golangcilint_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@rules_python//python:pip.bzl", "pip_install")

go_rules_dependencies()

go_register_toolchains(version = "1.17.1")

gazelle_dependencies()

bazel_skylib_workspace()

golangcilint_dependencies()

pip_install(
    name = "third_party",
    requirements = "//:requirements.txt",
)

# gazelle:repo bazel_gazelle

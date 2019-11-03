# Declares that this directory is the root of a Bazel workspace.
# See https://docs.bazel.build/versions/master/build-ref.html#workspace.
workspace(
    # How this workspace would be referenced with absolute labels from another workspace
    name = "math",
)

load("//:package.bzl", "register_repositories")

register_repositories()

load("//:deps.bzl", "install_dependencies")

install_dependencies()

# gazelle:repo bazel_gazelle

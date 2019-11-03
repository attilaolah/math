load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

def install_dependencies():
    """Install Go dependencies.

    If the user wants to get a different version of these, they can just fetch
    it from their WORKSPACE before calling this function, or not call this
    function at all.
    """

    # Register Go dependencies.
    go_rules_dependencies()
    go_register_toolchains()

    # Register Gazelle dependencies.
    gazelle_dependencies()

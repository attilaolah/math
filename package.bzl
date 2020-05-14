load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Version information.
# This is here to make it easier to update this file.
VERSIONS = {
    "atlassian/bazel-tools": "1056bf1d619b432063841df24b9eca86eb716527",
    "bazelbuild/bazel-gazelle": "v0.19.1",
    "bazelbuild/bazel-skylib": "1.0.2",
    "bazelbuild/rules_go": "v0.20.2",
}
SHA256_SUMS = {
    "atlassian/bazel-tools": "6a991df7a79db78229cbabded60c98641400f31fc88244847b519fbb904fc360",
    "bazelbuild/bazel-gazelle": "86c6d481b3f7aedc1d60c1c211c6f76da282ae197c3b3160f54bd3a8f847896f",
    "bazelbuild/bazel-skylib": "e5d90f0ec952883d56747b7604e2a15ee36e288bb556c3d0ed33e818a4d971f2",
    "bazelbuild/rules_go": "b9aa86ec08a292b97ec4591cf578e020b35f98e12173bbd4a921f84f583aebd9",
}

def register_repositories():
    """Fetch transitive dependencies.

    If the user wants to get a different version of these, they can just fetch
    it from their WORKSPACE before calling this function, or not call this
    function at all.
    """

    # Skylib is a dependency of our own .bzl files.
    _github_archive(
        name = "bazel_skylib",
        repo = "bazelbuild/bazel-skylib",
        url = "archive/{version}.tar.gz",
        strip_prefix = "bazel-skylib-{version}",
    )

    # Rules for Go libraries.
    _github_archive(
        name = "io_bazel_rules_go",
        repo = "bazelbuild/rules_go",
        url = "releases/download/{version}/rules_go-{version}.tar.gz",
    )
    _github_archive(
        name = "bazel_gazelle",
        repo = "bazelbuild/bazel-gazelle",
        url = "releases/download/{version}/bazel-gazelle-{version}.tar.gz",
    )

    # Go Linter.
    _github_archive(
        name = "com_github_atlassian_bazel_tools",
        repo = "atlassian/bazel-tools",
        url = "archive/{version}.zip",
        strip_prefix = "bazel-tools-{version}",
    )

def _github_archive(name, repo, url, strip_prefix = None):
    """Add a GitHub repository, if not already present."""
    if name in native.existing_rules():
        return

    version = VERSIONS[repo]

    if strip_prefix:
        strip_prefix = strip_prefix.format(version = version)

    http_archive(
        name = name,
        sha256 = SHA256_SUMS[repo],
        urls = [
            "https://github.com/" + repo + "/" + url.format(version = version),
        ],
        strip_prefix = strip_prefix,
    )

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Version information.
# This is here to make it easier to update this file.
VERSIONS = {
    "bazelbuild/bazel-gazelle": "v0.19.0",
    "bazelbuild/bazel-skylib": "1.0.2",
    "bazelbuild/rules_go": "v0.20.1",
}
SHA256_SUMS = {
    "bazelbuild/bazel-gazelle": "41bff2a0b32b02f20c227d234aa25ef3783998e5453f7eade929704dcff7cd4b",
    "bazelbuild/bazel-skylib": "e5d90f0ec952883d56747b7604e2a15ee36e288bb556c3d0ed33e818a4d971f2",
    "bazelbuild/rules_go": "842ec0e6b4fbfdd3de6150b61af92901eeb73681fd4d185746644c338f51d4c0",
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

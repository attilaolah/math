load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def github_archive(name, repo, url, version, sha256, strip_prefix = None):
    """Add a GitHub repository, if not already present."""
    if name in native.existing_rules():
        return

    if strip_prefix:
        strip_prefix = strip_prefix.format(version = version)

    http_archive(
        name = name,
        sha256 = sha256,
        urls = [
            "https://github.com/" + repo + "/" + url.format(version = version),
        ],
        strip_prefix = strip_prefix,
    )

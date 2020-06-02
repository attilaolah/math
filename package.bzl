"""Common repository rules."""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def github_archive(name, repo, url, version, sha256, strip_prefix = None):
    """Add a GitHub repository, if not already present.

    Placeholders are substituted based on passed-in values, e.g. {version} is
    replaced with the passed-in version.

    Args:
      name: Workspace name.
      repo: GitHub repo, e.g. attilaolah/math.
      url: Relative URL for the archive to download. May contain placeholders.
      version: Value for the placeholder in the `url` and `strip_prefix`params.
      sha256: SHA256 to be passed on to http_archive().
      strip_prefix: Passed on to http_archive(). May contain substitutions.
    """
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

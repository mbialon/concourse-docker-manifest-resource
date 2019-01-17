# Concourse Docker Manifest Resource

[![Travis CI](https://img.shields.io/travis/mbialon/concourse-docker-manifest-resource.svg?style=for-the-badge)](https://travis-ci.org/mbialon/concourse-docker-manifest-resource)

Creates [Docker](https://docker.io/) manifests.

## Source Configuration

* `repository`: *Required.* The name of the repository, e.g. `mbialon/concourse-docker-manifest-resource`.

* `tag`: The tag. Can be overwritten by the tag file.

* `username`: The username to user when authentidcating.

* `password`: The password to use when authentidcating.

## Behavior

### `out`: Push the manifest.

Create, annotate, and push the manifest. The resulting version is the manifest's digest.

#### Parameters

* `tag_file`: manifest tag file

* `manifests`: an array of:

    * `arch`: architecture

    * `os`: operating system

    * `tag_file`: a tag file

## Example

```yaml
resource_types:
- name: docker-manifest
  type: docker-image
  source:
    repository: mbialon/concourse-docker-manifest-resource

resources:
- name: image-manifest
  type: docker-manifest
  source:
    repository: mbialon/image
    username: ((docker.username))
    password: ((docker.password))

jobs:
- name: build manifest
  plan:
  - put: image-manifest
      params:
        tag_file: version/version
        manifests:
        - arch: amd64
          os: linux
          tag_file: build/tag
```

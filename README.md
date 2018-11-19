# Concourse Docker Manifest Resource

Creates [Docker](https://docker.io/) manifests.

## Source Configuration

* `repository`: *Required.* The name of the repository, e.g. `mbialon/concourse-docker-manifest-resource`.

* `tag`: *Required.* The tag.

* `username`: The username to user when authentidcating.

* `password`: The password to use when authentidcating.

## Behavior

### `out`: Push the manifest.

Create, annotate, and push the manifest. The resulting version is the manifest's digest.

#### Parameters

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
    tag: latest
    username: ((docker.username))
    password: ((docker.password))

jobs:
- name: build manifest
  plan:
  - put: image-manifest
      params:
        manifests:
        - arch: amd64
          os: linux
          tag_file: build/tag
```

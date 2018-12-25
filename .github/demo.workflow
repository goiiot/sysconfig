workflow "Build Demo" {
  on = "push"
  resolves = ["Docker Publish"]
}

action "Docker Login" {
  uses = "actions/docker/login@master"
  secrets = ["DOCKER_USERNAME", "DOCKER_PASSWORD"]
}

action "Docker Build" {
  uses = "actions/docker/cli@master"
  args = "build -t goiiot/sysconfig:demo -f demo.dockerfile ."
}

action "Docker Publish" {
  uses = "actions/docker/cli@master"
  needs = ["Docker Login", "Docker Build"]
  args = "push goiiot/sysconfig:demo"
}

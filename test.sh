#!/bin/bash

# test that world.yaml implements world.cue
echo "Testing world.yaml"
cue vet data/world/world.yaml schema/types.cue -d '#World'

# test person directory
echo "Testing data/person"
cue vet data/person/*.yaml schema/types.cue -d '#Person'

# test site directory
echo "Testing data/site"
cue vet data/site/*.yaml schema/types.cue -d '#Site'
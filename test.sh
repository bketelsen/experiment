#!/bin/bash

# test that world.yaml implements world.cue
echo "Testing world.yaml"
cue vet data/world.yaml schema/types.cue -d '#World'

# test that brian.yaml implements person.cue
echo "Testing brian.yaml"
cue vet data/person/brian.yaml schema/types.cue -d '#Person'
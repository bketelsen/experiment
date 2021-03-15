all: world person

world:
	@echo Testing world/*.yaml	
	@cue vet data/world/*.yaml schema/types.cue -d '#World'

person:
	@echo Testing person/*.yaml
	@cue vet data/person/*.yaml schema/types.cue -d '#Person'
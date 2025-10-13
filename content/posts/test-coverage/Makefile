PLANTUML := "plantuml/plantuml:1.2025"

build:
	mkdir -p ./img
	docker run --rm -v ${PWD}:/data $(PLANTUML) -verbose -o /data/ /data/*.puml
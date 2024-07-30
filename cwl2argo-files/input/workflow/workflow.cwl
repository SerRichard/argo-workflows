cwlVersion: v1.2
class: Workflow

inputs:
  OGmessage: string

steps:
  step1:
    id: step1id
    requirements:
      - class: DockerRequirement
        dockerPull: docker_image_name:tag
        dockerLoad: true
    run: step1.cwl
    in:
      message: OGmessage
    out: [output]

  step2:
    id: step2id
    run: step2.cwl
    doc:
      - This is documented
    in:
      message:
        source: step1/output
    hints:
      joke:
        knock: knock
      who:
        is: there
    label: some-label
    requirements:
      - class: DockerRequirement
        dockerPull: docker_image_name:tag
        dockerLoad: true
    scatter:
      - scatter1
      - scatter2
      - scatter3
    scatterMethod: dotproduct
    out: [output1]

  step3:
    run: step3.cwl
    doc: Not documented
    hints:
      - Not a good hint
    in:
      message:
        id: argumentfor3
        source: step2/output1
    scatter: scatterme
    requirements:
      - class: DockerRequirement
        dockerPull: docker_image_name:tag
        dockerLoad: true

outputs:
  final_output:
    type: string

hints:
  joke:
    knock: knock
  who:
    is: there

cwlVersion: v1.2
class: CommandLineTool
baseCommand: echo

inputs:
  message:
    type: string
    inputBinding:
      position: 1
      prefix: "--message"

outputs:
  output_file:
    type: File
    outputBinding:
      glob: output_step2.txt

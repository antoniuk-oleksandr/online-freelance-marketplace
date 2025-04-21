<script lang="ts">
  import { handleAttachFilesInputChange } from '../../handlers'

  type AttachFilesButtonInputProps = {
    inputRef: HTMLInputElement | undefined
    addFiles: (files: FileList) => void
    accept?: 'all' | 'image'
  }

  let { inputRef = $bindable(), addFiles, accept = 'all' }: AttachFilesButtonInputProps = $props()
  const acceptValue = accept === 'all' ? '*' : '.png, .svg, .jpg, .jpeg, .gif'
</script>

<input
  bind:this={inputRef}
  onchange={(e) => handleAttachFilesInputChange(e, addFiles)}
  multiple
  accept={acceptValue}
  type="file"
  hidden
/>

<script lang="ts">
  import { handleAttachFilesInputChange } from '../../handlers'

  type AttachFilesButtonInputProps = {
    inputRef: HTMLInputElement | undefined
    addFiles: (files: FileList) => void
  }

  let { inputRef = $bindable(), addFiles }: AttachFilesButtonInputProps = $props()
</script>

<input
  bind:this={inputRef}
  onchange={(e) => handleAttachFilesInputChange(e, addFiles)}
  multiple
  type="file"
  hidden
/>

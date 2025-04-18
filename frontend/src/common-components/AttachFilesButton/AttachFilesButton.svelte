<script lang="ts">
  import AttachFilesButtonElement from './components/AttachFilesButtonElement/AttachFilesButtonElement.svelte'
  import AttachFilesButtonInput from './components/AttachFilesButtonInput/AttachFilesButtonInput.svelte'

  type AttachFilesButtonProps = {
    addFiles: (files: FileList) => void
    styles?: string
  }

  const { addFiles, ...rest }: AttachFilesButtonProps = $props()
  let inputRef = $state<HTMLInputElement | undefined>()
</script>

<AttachFilesButtonInput bind:inputRef {addFiles} />
<AttachFilesButtonElement {...rest} {inputRef} />

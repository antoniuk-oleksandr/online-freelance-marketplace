<script lang="ts">
  import AttachFilesButtonElement from './components/AttachFilesButtonElement/AttachFilesButtonElement.svelte'
  import AttachFilesButtonInput from './components/AttachFilesButtonInput/AttachFilesButtonInput.svelte'

  type AttachFilesButtonProps = {
    addFiles: (files: FileList) => void
    styles?: string
    text?: string
    icon?: string
    accept?: 'all' | 'image'
  }

  const { addFiles, ...rest }: AttachFilesButtonProps = $props()
  let inputRef = $state<HTMLInputElement | undefined>()
</script>

<AttachFilesButtonInput {...rest} bind:inputRef {addFiles} />
<AttachFilesButtonElement {...rest} {inputRef} />

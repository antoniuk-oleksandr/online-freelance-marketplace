<script lang="ts">
  import AttachFilesButtonElement from './components/AttachFilesButtonElement/AttachFilesButtonElement.svelte'
  import AttachFilesButtonInput from './components/AttachFilesButtonInput/AttachFilesButtonInput.svelte'

  type AttachFilesButtonProps = {
    addFiles: (files: FileList) => void
  }

  const { addFiles }: AttachFilesButtonProps = $props()
  let inputRef = $state<HTMLInputElement | undefined>()
</script>

<AttachFilesButtonInput bind:inputRef {addFiles} />
<AttachFilesButtonElement {inputRef} />

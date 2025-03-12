<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import Iconify from '@iconify/svelte'
  import { handleAttachFilesButtonClick } from '../../handlers'

  type AttachFilesButtonElementProps = {
    inputRef: HTMLInputElement | undefined
  }

  const { inputRef }: AttachFilesButtonElementProps = $props()
</script>

<Button
  color="grey"
  styles="!w-full md:!w-48 !gap-3"
  type="button"
  clickAction={() => handleAttachFilesButtonClick(inputRef)}
>
  <Iconify width={24} height={24} icon="hugeicons:attachment-02" />
  <span>Attach Files</span>
</Button>

<script module lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import Icon from '@iconify/svelte'
  import {
    handleChatFileDrop,
    handleChatFileAttachDrag,
    handleChatFileChange,
  } from '../../handlers'

  let fileInput = $state<HTMLInputElement | undefined>()
  let files = $state<File[] | undefined>()
  const setFiles = (value: File[]) => (files = value)

  let dargValue = $state<boolean>(false)
  const setDragValue = (value: boolean) => (dargValue = value)

  export { OrderByIdChatTabAttachModelContent }
</script>

{#snippet OrderByIdChatTabAttachModelContent()}
  <!-- svelte-ignore a11y_no_static_element_interactions -->
  <div class="p-6">
    <div
      ondragover={(e) => handleChatFileAttachDrag(e, setDragValue, true)}
      ondragleave={(e) => handleChatFileAttachDrag(e, setDragValue, false)}
      ondrop={(e) => handleChatFileDrop(e, setFiles)}
      class="{dargValue
        ? 'border-cyan-500'
        : 'border-light-palette-divider dark:border-dark-palette-divider'} flex flex-col gap-6 justify-center items-center text-center py-12 border-2 border-dashed rounded-xl"
    >
      <div class="size-12">
        <Icon icon="hugeicons:cloud-saving-done-02" width="48" height="48" />
      </div>
      <div class="flex flex-col gap-1">
        <p class="font-semibold">Choose a file or drag and drop it here</p>
        <p class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
          Any formats up to 50MB
        </p>
      </div>
      <Button clickAction={() => fileInput?.click()} type="button" color="outline"
        >Browse Files</Button
      >
      <input
        onchange={(e) => handleChatFileChange(e, setFiles)}
        hidden
        bind:this={fileInput}
        type="file"
      />
    </div>
  </div>
{/snippet}

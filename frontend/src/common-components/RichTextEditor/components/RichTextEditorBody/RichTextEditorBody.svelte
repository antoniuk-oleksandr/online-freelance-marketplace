<script lang="ts">
  import type { Editor } from '@tiptap/core'

  type RichTextEditorProps = {
    editor: Editor | undefined
    element: HTMLDivElement | undefined
  }

  let { editor, element = $bindable() }: RichTextEditorProps = $props()
</script>

<!-- svelte-ignore element_invalid_self_closing_tag -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  onclick={() => editor?.commands.focus()}
  bind:this={element}
  class="text-light-palette-text-primary dark:text-dark-palette-text-primary min-h-48 cursor-text first:!mt-0 !w-full !max-w-full outline-none p-6"
/>

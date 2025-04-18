<script lang="ts">
  import type { Editor } from '@tiptap/core'
  import RichTextEditorButton from '../RichTextEditorButton/RichTextEditorButton.svelte'
  import RichTextEditorHeaderLayout from './RichTextEditorHeaderLayout.svelte'
  import { getInitialRichTextEditorHeaderItems, refreshItemsValueStore } from '../../helpers'
  import type { RichTextEditorItem } from '@/types/RichTextEditorItem'

  type RichTextEditorHeaderProps = {
    editor: Editor
    itemsValueStore: Record<string, { disabled: boolean; active: boolean }> | undefined
  }

  const { editor, itemsValueStore }: RichTextEditorHeaderProps = $props()

  const items = $state<Record<string, RichTextEditorItem>>(
    getInitialRichTextEditorHeaderItems(editor),
  )
</script>

<RichTextEditorHeaderLayout>
  {#if itemsValueStore && Object.keys(itemsValueStore).length > 0}
    {#each Object.entries(items) as [key, value]}
      <RichTextEditorButton
        {...value}
        onClick={() => value.toggleFunc()}
        disabled={itemsValueStore[key] === undefined ? undefined : itemsValueStore[key].disabled}
        active={itemsValueStore[key] === undefined ? undefined : itemsValueStore[key].active}
      />
    {/each}
  {/if}
</RichTextEditorHeaderLayout>

<script lang="ts">
  import { Editor } from '@tiptap/core'
  import { onDestroy, onMount } from 'svelte'
  import RichTextEditorHeader from './components/RichTextEditorHeader/RichTextEditorHeader.svelte'
  import { refreshItemsValueStore } from './helpers'
  import RichTextEditorLayout from './RichTextEditorLayout.svelte'
  import RichTextEditorBody from './components/RichTextEditorBody/RichTextEditorBody.svelte'
  import { handleRichTextEditorDestroy, handleRichTextEditorMount } from './handlers'
  import RichTextEditorFooter from './components/RichTextEditorFooter/RichTextEditorFooter.svelte'
  import LabeledInputError from '../LabeledInput/components/LabeledInputError/LabeledInputError.svelte'

  type RichTextEditorProps = {
    limit: number
    onChange: (value: string) => void
    content?: string
    error?: string | null
  }

  const { limit, onChange, error, content = '' }: RichTextEditorProps = $props()

  let element = $state<undefined | HTMLDivElement>()
  let editor = $state<undefined | Editor>()
  let itemsValueStore = $state<Record<string, { disabled: boolean; active: boolean }> | undefined>()
  let count = $state<number>(0)

  const handler = () => {
    if (!editor) return
    itemsValueStore = refreshItemsValueStore(editor)
    count = editor.storage.characterCount.characters()
    onChange(editor.getHTML())
  }

  onMount(() => (editor = handleRichTextEditorMount(element, handler, limit, content)))
  onDestroy(() => handleRichTextEditorDestroy(editor, handler))
</script>

<RichTextEditorLayout {error}>
  {#if editor}
    <RichTextEditorHeader {itemsValueStore} {editor} />
  {/if}
  <RichTextEditorBody {editor} bind:element />
  <RichTextEditorFooter {error} {count} {limit} />
</RichTextEditorLayout>
<LabeledInputError {error} />

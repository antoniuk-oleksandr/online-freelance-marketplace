import { Editor } from "@tiptap/core"
import Color from "@tiptap/extension-color"
import StarterKit from "@tiptap/starter-kit"
import ListItem from '@tiptap/extension-list-item'
import TextStyle from '@tiptap/extension-text-style'
import CharacterCount from '@tiptap/extension-character-count'

export const handleRichTextEditorMount = (
  element: HTMLDivElement | undefined,
  handler: () => void,
  limit: number,
  content: string
) => {
  const editor = new Editor({
    element: element,
    extensions: [
      Color.configure({ types: [TextStyle.name, ListItem.name] }),
      // @ts-ignore
      TextStyle.configure({ types: [ListItem.name] }),
      StarterKit,
      CharacterCount.configure({ limit }),
    ],
    content,
  })
  editor.on('transaction', handler)

  return editor
}

export const handleRichTextEditorDestroy = (
  editor: Editor | undefined,
  handler: () => void,
) => {
  if (!editor) return
  editor.off('transaction', handler)
  editor.destroy()
}

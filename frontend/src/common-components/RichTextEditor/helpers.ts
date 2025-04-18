import type { RichTextEditorItem } from "@/types/RichTextEditorItem";
import type { RichTextEditorValueStore } from "@/types/RichTextEditorValueStore";
import type { Editor } from "@tiptap/core";

export const getInitialRichTextEditorHeaderItems = (
  editor: Editor,
): Record<string, RichTextEditorItem> => {
  const arrBefore = {
    bold: {
      icon: 'mdi:format-bold',
      tooltipText: 'Bold',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().toggleBold().run(),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleBold().run(),
      active: editor.isActive('bold'),
    },
    italic: {
      icon: 'mdi:format-italic',
      tooltipText: 'Italic',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().toggleItalic().run(),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleItalic().run(),
      active: editor.isActive('italic'),
    },
    strike: {
      icon: 'mdi:format-strikethrough',
      tooltipText: 'Strikethrough',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().toggleStrike().run(),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleStrike().run(),
      active: editor.isActive('strike'),
    },
    underline: {
      icon: 'mdi:code-tags',
      tooltipText: 'Inline code',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().toggleCode().run(),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleCode().run(),
      active: editor.isActive('code'),
    },
    clearMarks: {
      icon: 'mdi:format-clear',
      tooltipText: 'Clear marks',
      toggleFunc: () => editor.chain().focus().unsetAllMarks().run(),
      disabled: !editor.can().chain().focus().unsetAllMarks().run(),
    },
    clearNodes: {
      icon: 'mdi:eraser',
      tooltipText: 'Clear nodes',
      toggleFunc: () => editor.chain().focus().clearNodes().run(),
    },
    paragraph: {
      icon: 'mdi:format-paragraph',
      tooltipText: 'Paragraph',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().setParagraph().run(),
      //@ts-ignore
      disabled: !editor.can().chain().focus().setParagraph().run(),
      active: editor.isActive('paragraph'),
    },
  };

  for (let i = 1; i <= 6; i++) {
      //@ts-ignore
    arrBefore[`heading${i}`] = {
      icon: `mdi:format-header-${i}`,
      tooltipText: `Heading ${i}`,
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().toggleHeading({ level: i }).run(),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleHeading({ level: i }).run(),
      active: editor.isActive('heading', { level: i }),
    };
  }

  const arrAfter = {
    bulletList: {
      icon: 'mdi:format-list-bulleted',
      tooltipText: 'Bullet list',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().toggleBulletList().run(),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleBulletList().run(),
      active: editor.isActive('bulletList'),
    },
    orderedList: {
      icon: 'mdi:format-list-numbered',
      tooltipText: 'Ordered list',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().toggleOrderedList().run(),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleOrderedList().run(),
      active: editor.isActive('orderedList'),
    },
    codeBlock: {
      icon: 'mdi:code-braces',
      tooltipText: 'Code block',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().toggleCodeBlock().run(),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleCodeBlock().run(),
      active: editor.isActive('codeBlock'),
    },
    blockquote: {
      icon: 'mdi:format-quote-close',
      tooltipText: 'Blockquote',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().toggleBlockquote().run(),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleBlockquote().run(),
      active: editor.isActive('blockquote'),
    },
    horizontalRule: {
      icon: 'mdi:minus',
      tooltipText: 'Horizontal rule',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().setHorizontalRule().run(),
    },
    hardBreak: {
      icon: 'mdi:keyboard-return',
      tooltipText: 'Hard break',
      //@ts-ignore
      toggleFunc: () => editor.chain().focus().setHardBreak().run(),
    },
  };

  //@ts-ignore
  return { ...arrBefore, ...arrAfter };
};

export const refreshItemsValueStore = (editor: Editor) => {
  const items: Record<string, { disabled: boolean; active: boolean }> = {
    bold: {
      active: editor.isActive('bold'),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleBold().run(),
    },
    italic: {
      active: editor.isActive('italic'),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleItalic().run(),
    },
    strike: {
      active: editor.isActive('strike'),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleStrike().run(),
    },
    underline: {
      active: editor.isActive('code'),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleCode().run(),
    },
    clearMarks: {
      active: false,
      disabled: !editor.can().chain().focus().unsetAllMarks().run(),
    },
    paragraph: {
      active: editor.isActive('paragraph'),
      //@ts-ignore
      disabled: !editor.can().chain().focus().setParagraph().run(),
    },
    bulletList: {
      active: editor.isActive('bulletList'),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleBulletList().run(),
    },
    orderedList: {
      active: editor.isActive('orderedList'),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleOrderedList().run(),
    },
    codeBlock: {
      active: editor.isActive('codeBlock'),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleCodeBlock().run(),
    },
    blockquote: {
      active: editor.isActive('blockquote'),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleBlockquote().run(),
    },
  };

  // Dynamically add heading1 to heading6
  for (let i = 1; i <= 6; i++) {
    items[`heading${i}`] = {
      active: editor.isActive('heading', { level: i }),
      //@ts-ignore
      disabled: !editor.can().chain().focus().toggleHeading({ level: i }).run(),
    };
  }

  return items;
}


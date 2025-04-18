export type RichTextEditorItem = {
  icon: string,
  tooltipText: string,
  toggleFunc: () => void,
  disabled: boolean | undefined,
  active: boolean | undefined,
}

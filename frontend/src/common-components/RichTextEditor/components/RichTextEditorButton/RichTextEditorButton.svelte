<script lang="ts">
  import Tooltip from '@/common-components/Tooltip/Tooltip.svelte'
  import Icon from '@iconify/svelte'

  type RichTextEditorButtonProps = {
    onClick: () => void
    tooltipText?: string
    icon?: string
    active?: boolean
    disabled?: boolean
  }

  const {
    onClick,
    active = false,
    disabled = false,
    tooltipText = 'none',
    icon = 'mdi:block-helper',
  }: RichTextEditorButtonProps = $props()
</script>

<Tooltip text={tooltipText}>
  <button
    type="button"
    class="{active ? '!bg-cyan-500 text-white' : ''} grid place-items-center w-11 h-7 bg-light-palette-action-hover rounded-md dark:bg-dark-palette-action-hover whitespace-nowrap px-3 disabled:bg-light-palette-action-disabled dark:disabled:bg-dark-palette-action-disabled"
    {disabled}
    onclick={onClick}
  >
    <Icon width={20} height={20} {icon}/>
  </button>
</Tooltip>

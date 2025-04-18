<script lang="ts">
  import type { DropdownItem } from '@/types/DropdownItem'
  import type { LayoutProps } from '@/types/LayoutProps'

  type DropdownMenuItemElementLayoutPropsProps = DropdownItem &
    LayoutProps & {
      setShown: (value: boolean) => void
      clickAction: () => void
      customColor?: string
    }

  const {
    setShown,
    clickAction,
    children,
    customColor,
    closeDropdown = true,
  }: DropdownMenuItemElementLayoutPropsProps = $props()
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<div
  onclick={() => {
    closeDropdown && setShown(false)
    clickAction()
  }}
  class="{customColor
    ? customColor
    : ' text-light-palette-text-primary dark:text-dark-palette-text-primary'} justify-between cursor-pointer select-none text-start text-base flex items-center hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover px-3 rounded-[4px] duration-200 ease-out py-1"
>
  {@render children()}
</div>

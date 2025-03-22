<script lang="ts">
  import type { DropdownItem } from '@/types/DropdownItem'
  import type { LayoutProps } from '@/types/LayoutProps'
  import DropdownMenu from './components/DropdownMenu/DropdownMenu.svelte'
  import DropdownTrigger from './components/DropdownTrigger/DropdownTrigger.svelte'
  import DropdownLayout from './DropdownLayout.svelte'

  type DropdownProps = LayoutProps & {
    items: DropdownItem[]
    modalHeaderTitle: string
    shownValue?: boolean
    layoutStyles?: string
    menuWidth?: string
    positionX?: 'left' | 'right'
    positionY?: 'top' | 'bottom'
  }

  let shown = $state(false)
  const setShown = (value: boolean) => (shown = value)

  let triggerRef = $state<HTMLDivElement | undefined>()

  let { children, shownValue = $bindable(), ...rest }: DropdownProps = $props()

  $effect(() => {
    if (shownValue !== undefined && shownValue !== shown) shownValue = shown
  })
</script>

<DropdownLayout {...rest}>
  <DropdownTrigger bind:triggerRef bind:shown>
    {@render children()}
  </DropdownTrigger>
  {#if shown}
    <DropdownMenu {setShown} {triggerRef} {...rest} />
  {/if}
</DropdownLayout>

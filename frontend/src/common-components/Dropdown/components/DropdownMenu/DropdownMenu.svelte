<script lang="ts">
  import type { DropdownItem } from '@/types/DropdownItem'
  import DropdownMenuLayout from './DropdownMenuLayout.svelte'
  import DropdownMenuItemElementList from '../DropdownMenuItemElementList/DropdownMenuItemElementList.svelte'
  import { onMount } from 'svelte'
  import DropdownModalBackdrop from '../DropdownModalBackdrop/DropdownModalBackdrop.svelte'
  import DropdownModalHeader from '../DropdownModal/components/DropdownModalHeader/DropdownModalHeader.svelte'

  type DropdownMenuProps = {
    items: DropdownItem[]
    setShown: (value: boolean) => void
    triggerRef: HTMLDivElement | undefined
    modalHeaderTitle: string
    menuWidth?: string
    positionX?: 'left' | 'right'
    positionY?: 'top' | 'bottom'
    menuStyles?: string
  }

  let menuRef = $state<HTMLDivElement | undefined>()
  let { setShown, triggerRef, ...rest }: DropdownMenuProps = $props()

  const handleClick = (e: Event) => {
    const containsTrigger = triggerRef && triggerRef.contains(e.target as Node)
    const containsMenu = menuRef && menuRef.contains(e.target as Node)
    if (containsMenu || containsTrigger) return

    setShown(false)
  }

  onMount(() => {
    if (window.innerWidth < 1024) document.body.style.overflow = 'hidden'
    document.addEventListener('click', handleClick)

    return () => {
      document.removeEventListener('click', handleClick)
      document.body.style.overflow = 'auto'
    }
  })
</script>

<DropdownMenuLayout bind:menuRef {...rest}>
  <DropdownModalHeader {...rest} {setShown} />
  <DropdownMenuItemElementList {setShown} {...rest} />
</DropdownMenuLayout>
<DropdownModalBackdrop />

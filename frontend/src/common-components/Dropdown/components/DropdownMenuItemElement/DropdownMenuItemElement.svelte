<script lang="ts">
  import type { DropdownItem } from '@/types/DropdownItem'
  import DropdownMenuItemElementIcon from './components/DropdownMenuItemElementIcon/DropdownMenuItemElementIcon.svelte'
  import DropdownMenuItemElementLayout from './DropdownMenuItemElementLayout.svelte'
  import DropdownMenuItemElementText from './components/DropdownMenuItemElementText/DropdownMenuItemElementText.svelte'
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import DropdownMenuItemElementBadge from './components/DropdownMenuItemElementBadge/DropdownMenuItemElementBadge.svelte'
  import DropdownMenuItemIconTextLayout from './components/DropdownMenuItemIconText/DropdownMenuItemIconTextLayout.svelte'

  type DropdownMenuItemElementProps = DropdownItem & {
    setShown: (value: boolean) => void
  }

  let props: DropdownMenuItemElementProps = $props()
</script>

{#if props.dividerBefore}
  <DividerElement />
{/if}
<DropdownMenuItemElementLayout {...props}>
  <DropdownMenuItemIconTextLayout>
    <DropdownMenuItemElementIcon {...props} />
    <DropdownMenuItemElementText {...props} />
  </DropdownMenuItemIconTextLayout>
  <DropdownMenuItemElementBadge {...props} />
</DropdownMenuItemElementLayout>
{#if props.dividerAfter}
  <DividerElement />
{/if}

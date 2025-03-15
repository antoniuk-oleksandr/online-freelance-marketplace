<script lang="ts">
  import Iconify from '@iconify/svelte'

  type DropdownMenuItemElementIconProps = {
    icon?: string
  }

  const { icon }: DropdownMenuItemElementIconProps = $props()
</script>

{#if icon}
  <Iconify {icon} />
{/if}

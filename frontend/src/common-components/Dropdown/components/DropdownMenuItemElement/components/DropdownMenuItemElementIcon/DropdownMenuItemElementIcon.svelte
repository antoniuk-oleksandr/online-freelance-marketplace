<script lang="ts">
  import Iconify from '@iconify/svelte'

  type DropdownMenuItemElementIconProps = {
    icon?: string
    iconColor?: string
  }

  const { icon, iconColor }: DropdownMenuItemElementIconProps = $props()
</script>

{#if icon}
  <Iconify class={iconColor ? 'text-' + iconColor : ''} {icon} />
{/if}

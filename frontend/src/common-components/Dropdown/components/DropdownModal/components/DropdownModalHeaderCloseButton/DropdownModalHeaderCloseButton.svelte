<script lang="ts">
  import Icon from '@iconify/svelte'

  type DropdownModalHeaderCloseButtonProps = {
    setShown?: (value: boolean) => void
  }

  const { setShown }: DropdownModalHeaderCloseButtonProps = $props()
  const handleModalCloseButtonClick = () => setShown && setShown(false)
</script>

<button onclick={handleModalCloseButtonClick} class="overflow-hidden">
  <Icon
    class="rotate-45 hover:text-light-palette-action-hover hover:dark:text-dark-palette-action-hover duration-200 ease-out"
    width={24}
    height={24}
    icon="hugeicons:add-01"
  />
</button>

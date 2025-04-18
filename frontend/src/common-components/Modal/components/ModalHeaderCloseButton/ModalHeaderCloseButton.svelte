<script lang="ts">
  import Icon from '@iconify/svelte'
  import { handleModalCloseButtonClick } from '@/common-components/Modal/handlers'
</script>

<button onclick={handleModalCloseButtonClick} class="overflow-hidden">
  <Icon
    class="rotate-45 hover:text-light-palette-text-secondary hover:dark:text-dark-palette-text-secondary duration-200 ease-out"
    width={24}
    height={24}
    icon="hugeicons:add-01"
  />
</button>

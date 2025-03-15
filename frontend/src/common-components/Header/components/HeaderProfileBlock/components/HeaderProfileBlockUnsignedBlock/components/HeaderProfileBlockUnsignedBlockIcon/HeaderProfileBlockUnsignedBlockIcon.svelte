<script lang="ts">
  import Icon from '@iconify/svelte'

  type HeaderProfileBlockUnsignedBlockIconProps = {
    darkMode: boolean | null
  }

  const { darkMode }: HeaderProfileBlockUnsignedBlockIconProps = $props()
</script>

<Icon
  class="size-6 text-light-palette-text-primary dark:text-dark-palette-text-primary"
  icon={`hugeicons:${!darkMode ? 'sun-03' : 'moon-02'}`}
/>

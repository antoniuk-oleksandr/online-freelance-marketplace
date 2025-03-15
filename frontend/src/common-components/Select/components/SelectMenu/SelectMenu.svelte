<script lang="ts">
  import SelectMenuLayout from '@/common-components/Select/components/SelectMenu/SelectMenuLayout.svelte'
  import type { SelectMenuProps } from '@/types/SelectMenuProps'
  import SelectMenuContent from '@/common-components/Select/components/SelectMenuContent/SelectMenuContent.svelte'

  const props: SelectMenuProps = $props()
</script>

{#if props.showExitAnimation}
  <SelectMenuLayout {...props}>
    <SelectMenuContent {...props} />
  </SelectMenuLayout>
{/if}

<script lang="ts">
  import type { TabType } from '@/types/TabType'
  import TabElement from '../TabElement/TabElement.svelte'
  import TabElementListLayout from './TabElementListLayout.svelte'

  type TabElementListProps = {
    tabs: TabType[]
    tabIndex: number
    setTabIndex: (index: number) => void
    elementStyles?: string
  }

  const { tabs, ...rest }: TabElementListProps = $props()
</script>

<TabElementListLayout>
  {#each tabs as item, index}
    <TabElement {index} {...item} {...rest} />
  {/each}
</TabElementListLayout>

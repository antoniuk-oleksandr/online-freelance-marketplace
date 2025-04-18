<script lang="ts">
  import type { TabType } from '@/types/TabType'
  import TabElement from '../TabElement/TabElement.svelte'
  import TabElementListLayout from './TabElementListLayout.svelte'

  type TabElementListProps = {
    tabs: TabType[]
    activeTabIndex: number
    setActiveTabIndex: (index: number) => void
  }

  const { tabs, ...rest }: TabElementListProps = $props()
</script>

<TabElementListLayout>
  {#each tabs as item, index}
    <TabElement {index} {...item} {...rest} />
  {/each}
</TabElementListLayout>

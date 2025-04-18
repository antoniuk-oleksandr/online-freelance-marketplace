<script lang="ts">
  import type { TabType } from '@/types/TabType'
  import TabElementList from './components/TabElementList/TabElementList.svelte'
  import TabsLayout from './TabsLayout.svelte'
    import type { MyProfileOrderByIdData } from '@/types/MyProfileOrderByIdData'

  type TabsProps = {
    tabs: TabType[]
    initialActiveTabIndex?: number
    tabChangeAction?: (index: number) => void
    tabComponentsData: MyProfileOrderByIdData
  }

  const { tabs, initialActiveTabIndex = 0, tabChangeAction, tabComponentsData }: TabsProps = $props()

  let activeTabIndex = $state(initialActiveTabIndex)
  const setActiveTabIndex = (index: number) => {
    activeTabIndex = index
    tabChangeAction && tabChangeAction(index)
  }
</script>

<TabsLayout>
  <TabElementList {tabs} {activeTabIndex} {setActiveTabIndex} />
  {@render tabs[activeTabIndex].component(tabComponentsData[activeTabIndex])}
</TabsLayout>

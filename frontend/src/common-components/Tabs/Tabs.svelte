<script lang="ts">
  import type { TabType } from '@/types/TabType'
  import TabElementList from './components/TabElementList/TabElementList.svelte'
  import TabsLayout from './TabsLayout.svelte'
  import type { MyProfileOrderByIdData } from '@/types/MyProfileOrderByIdData'

  type TabsProps = {
    tabs: TabType[]
    tabIndex: number
    setTabIndex: (index: number) => void
    tabComponentsData: MyProfileOrderByIdData
    orderId: number
    elementStyles?: string
  }

  const {
    tabs,
    tabIndex,
    setTabIndex,
    tabComponentsData,
    orderId,
    elementStyles,
  }: TabsProps = $props()
</script>

<TabsLayout>
  <TabElementList {elementStyles} {tabs} {tabIndex} {setTabIndex} />
  {@render tabs[tabIndex].component(tabComponentsData[tabIndex], orderId)}
</TabsLayout>

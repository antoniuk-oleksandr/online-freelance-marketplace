<script lang="ts">
  import TabElementButton from './components/TabElementButton/TabElementButton.svelte'
  import TabElementUnderline from './components/TabElementUnderline/TabElementUnderline.svelte'
  import TabElementLayout from './TabElementLayout.svelte'

  type TabElementProps = {
    title: string
    index: number
    tabIndex: number
    setTabIndex: (index: number) => void
    icon?: string
    elementStyles?: string
  }

  const props: TabElementProps = $props()
</script>

<TabElementLayout>
  <TabElementButton {...props} />
  <TabElementUnderline {...props} />
</TabElementLayout>

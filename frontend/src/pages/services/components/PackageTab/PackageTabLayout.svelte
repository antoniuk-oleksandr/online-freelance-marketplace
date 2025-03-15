<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import { themeStore } from '@/common-stores/theme-storage'

  type PackageTabLayoutProps = LayoutProps & {
    length: number
    index: number
    setSelectedTab: (value: number) => void
    selectedTab: number
  }

  let unselectedColor = $state<string | null>(null)
  themeStore.subscribe((value) => {
    unselectedColor = value === true ? 'rgba(255, 255, 255, 0.12)' : 'rgba(0, 0, 0, 0.12)'
  })

  const tabProps: PackageTabLayoutProps = $props()

  const { children, length, index, setSelectedTab } = tabProps

  let selected = $state<boolean>(false)
  $effect(() => {
    selected = index === tabProps.selectedTab
  })

  const width = length === 1 ? 'w-full' : length === 2 ? 'w-1/2' : 'w-1/3'

  const selectedColor = 'rgb(6, 182, 212)'
</script>

{#if unselectedColor !== null}
  <button
    onclick={() => setSelectedTab(index)}
    style="border-color: {selected ? selectedColor : unselectedColor}"
    class="{width} cursor-pointer hover:bg-transparent !p-4 !text-base !font-semibold text-center border-b"
  >
    {@render children()}
  </button>
{/if}

<script lang="ts">
  import Icon from '@iconify/svelte'

  type TabElementButtonProps = {
    title: string
    index: number
    activeTabIndex: number
    setActiveTabIndex: (index: number) => void
    icon?: string
  }

  const { title, index, icon, activeTabIndex, setActiveTabIndex }: TabElementButtonProps = $props()
</script>

<div class="{activeTabIndex === index ? 'text-cyan-500' : ''} durantion-300 ease-out">
  {#if icon !== undefined}
    <button class="md:hidden px-2.5" onclick={() => setActiveTabIndex(index)}>
      <Icon {icon} width="24" height="24" />
    </button>
  {/if}
  <button class={icon ? 'hidden md:block' : ''} onclick={() => setActiveTabIndex(index)}>
    {title}
  </button>
</div>

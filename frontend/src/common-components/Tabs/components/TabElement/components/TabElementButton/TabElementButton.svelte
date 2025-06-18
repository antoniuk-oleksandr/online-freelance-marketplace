<script lang="ts">
  import Icon from '@iconify/svelte'

  type TabElementButtonProps = {
    title: string
    index: number
    tabIndex: number
    setTabIndex: (index: number) => void
    icon?: string
    elementStyles?: string
  }

  const {  elementStyles, title, index, icon, tabIndex, setTabIndex }: TabElementButtonProps = $props()
</script>

<div class="{tabIndex === index ? 'text-cyan-500' : ''} durantion-300 ease-out">
  {#if icon !== undefined}
    <button class="{elementStyles} md:hidden px-2.5 text-2xl" onclick={() => setTabIndex(index)}>
      <Icon {icon}  />
    </button>
  {/if}
  <button class={icon ? 'hidden md:block' : ''} onclick={() => setTabIndex(index)}>
    {title}
  </button>
</div>

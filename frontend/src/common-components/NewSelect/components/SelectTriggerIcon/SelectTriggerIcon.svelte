<script lang="ts">
  import Icon from '@iconify/svelte'

  type SelectTriggerIconProps = {
    iconLoaded: boolean
    shownValue: boolean
    triggerIcon?: string
  }

  let { iconLoaded = $bindable(), triggerIcon, shownValue }: SelectTriggerIconProps = $props()
</script>

{#if triggerIcon}
  <Icon width="24" height="24" icon={triggerIcon} on:load={() => (iconLoaded = true)} />
{:else}
  <div class="transition-transform duration-300 ease-out" class:rotate-180={shownValue}>
    <Icon width="24" height="24" icon="hugeicons:arrow-up-01" />
  </div>
{/if}
